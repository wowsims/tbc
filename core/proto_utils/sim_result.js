import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { AuraMetrics as AuraMetricsProto } from '/tbc/core/proto/api.js';
import { DpsMetrics as DpsMetricsProto } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { getIconUrl } from '/tbc/core/resources.js';
import { getName } from '/tbc/core/resources.js';
import { actionIdToString } from '/tbc/core/resources.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { getTalentTreeIcon } from '/tbc/core/proto_utils/utils.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { bucket } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';
// Holds all the data from a simulation call, and provides helper functions
// for parsing it.
export class SimResult {
    constructor(request, result, raidMetrics, encounterMetrics) {
        this.request = request;
        this.result = result;
        this.raidMetrics = raidMetrics;
        this.encounterMetrics = encounterMetrics;
    }
    getPlayers(filter) {
        if (filter?.player || filter?.player === 0) {
            const player = this.getPlayerWithRaidIndex(filter.player);
            return player ? [player] : [];
        }
        else {
            return this.raidMetrics.parties.map(party => party.players).flat();
        }
    }
    // Returns the first player, regardless of which party / raid slot its in.
    getFirstPlayer() {
        return this.getPlayers()[0] || null;
    }
    getPlayerWithRaidIndex(raidIndex) {
        return this.getPlayers().find(player => player.raidIndex == raidIndex) || null;
    }
    getTargets(filter) {
        if (filter?.target || filter?.target === 0) {
            const target = this.getTargetWithIndex(filter.target);
            return target ? [target] : [];
        }
        else {
            return this.encounterMetrics.targets.slice();
        }
    }
    getTargetWithIndex(index) {
        return this.getTargets().find(target => target.index == index) || null;
    }
    getDamageMetrics(filter) {
        if (filter.player || filter.player === 0) {
            return this.getPlayerWithRaidIndex(filter.player)?.dps || DpsMetricsProto.create();
        }
        return this.raidMetrics.dps;
    }
    getActionMetrics(filter) {
        return ActionMetrics.join(this.getPlayers(filter).map(player => player.actions).flat());
    }
    getSpellMetrics(filter) {
        return this.getActionMetrics(filter).filter(e => e.hits + e.misses != 0);
    }
    getBuffMetrics(filter) {
        return AuraMetrics.join(this.getPlayers(filter).map(player => player.auras).flat());
    }
    getDebuffMetrics(filter) {
        return AuraMetrics.join(this.getTargets(filter).map(target => target.auras).flat());
    }
    getLogs() {
        return this.result.logs.split('\n');
    }
    toJson() {
        return {
            'request': RaidSimRequest.toJson(this.request),
            'result': RaidSimResult.toJson(this.result),
        };
    }
    static async fromJson(obj) {
        const request = RaidSimRequest.fromJson(obj['request']);
        const result = RaidSimResult.fromJson(obj['result']);
        return SimResult.makeNew(request, result);
    }
    static async makeNew(request, result) {
        const iterations = request.simOptions?.iterations || 1;
        const duration = request.encounter?.duration || 1;
        const raidPromise = RaidMetrics.makeNew(iterations, duration, request.raid, result.raidMetrics);
        const encounterPromise = EncounterMetrics.makeNew(iterations, duration, request.encounter, result.encounterMetrics);
        const raidMetrics = await raidPromise;
        const encounterMetrics = await encounterPromise;
        return new SimResult(request, result, raidMetrics, encounterMetrics);
    }
}
export class RaidMetrics {
    constructor(raid, metrics, parties) {
        this.raid = raid;
        this.metrics = metrics;
        this.dps = this.metrics.dps;
        this.parties = parties;
    }
    static async makeNew(iterations, duration, raid, metrics) {
        const numParties = Math.min(raid.parties.length, metrics.parties.length);
        const parties = await Promise.all([...new Array(numParties).keys()]
            .map(i => PartyMetrics.makeNew(iterations, duration, raid.parties[i], metrics.parties[i], i)));
        return new RaidMetrics(raid, metrics, parties);
    }
}
export class PartyMetrics {
    constructor(party, metrics, partyIndex, players) {
        this.party = party;
        this.metrics = metrics;
        this.partyIndex = partyIndex;
        this.dps = this.metrics.dps;
        this.players = players;
    }
    static async makeNew(iterations, duration, party, metrics, partyIndex) {
        const numPlayers = Math.min(party.players.length, metrics.players.length);
        const players = await Promise.all([...new Array(numPlayers).keys()]
            .filter(i => party.players[i].class != Class.ClassUnknown)
            .map(i => PlayerMetrics.makeNew(iterations, duration, party.players[i], metrics.players[i], partyIndex * 5 + i)));
        return new PartyMetrics(party, metrics, partyIndex, players);
    }
}
export class PlayerMetrics {
    constructor(player, metrics, raidIndex, actions, auras, iterations, duration) {
        this.player = player;
        this.metrics = metrics;
        this.raidIndex = raidIndex;
        this.name = player.name;
        this.spec = playerToSpec(player);
        this.iconUrl = getTalentTreeIcon(this.spec, player.talentsString);
        this.classColor = classColors[specToClass[this.spec]];
        this.dps = this.metrics.dps;
        this.actions = actions;
        this.auras = auras;
        this.iterations = iterations;
        this.duration = duration;
    }
    get label() {
        return `${this.name} (#${this.raidIndex + 1})`;
    }
    get secondsOomAvg() {
        return this.metrics.secondsOomAvg;
    }
    get totalDamage() {
        return this.dps.avg * this.duration;
    }
    static async makeNew(iterations, duration, player, metrics, raidIndex) {
        const actionsPromise = Promise.all(metrics.actions.map(actionMetrics => ActionMetrics.makeNew(iterations, duration, actionMetrics)));
        const aurasPromise = Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(iterations, duration, auraMetrics)));
        const actions = await actionsPromise;
        const auras = await aurasPromise;
        return new PlayerMetrics(player, metrics, raidIndex, actions, auras, iterations, duration);
    }
}
export class EncounterMetrics {
    constructor(encounter, metrics, targets) {
        this.encounter = encounter;
        this.metrics = metrics;
        this.targets = targets;
    }
    static async makeNew(iterations, duration, encounter, metrics) {
        const numTargets = Math.min(encounter.targets.length, metrics.targets.length);
        const targets = await Promise.all([...new Array(numTargets).keys()]
            .map(i => TargetMetrics.makeNew(iterations, duration, encounter.targets[i], metrics.targets[i], i)));
        return new EncounterMetrics(encounter, metrics, targets);
    }
    get durationSeconds() {
        return this.encounter.duration;
    }
}
export class TargetMetrics {
    constructor(target, metrics, index, auras) {
        this.target = target;
        this.metrics = metrics;
        this.index = index;
        this.auras = auras;
    }
    static async makeNew(iterations, duration, target, metrics, index) {
        const auras = await Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(iterations, duration, auraMetrics)));
        return new TargetMetrics(target, metrics, index, auras);
    }
}
export class AuraMetrics {
    constructor(actionId, name, iconUrl, iterations, duration, data) {
        this.actionId = actionId;
        this.name = name;
        this.iconUrl = iconUrl;
        this.iterations = iterations;
        this.duration = duration;
        this.data = data;
    }
    get uptimePercent() {
        return this.data.uptimeSecondsAvg / this.duration * 100;
    }
    static async makeNew(iterations, duration, auraMetrics) {
        const actionId = {
            id: {
                spellId: auraMetrics.id,
            },
            tag: 0,
        };
        const namePromise = getName(actionId.id);
        const iconPromise = getIconUrl(actionId.id);
        const name = await namePromise;
        const iconUrl = await iconPromise;
        return new AuraMetrics(actionId, name, iconUrl, iterations, duration, auraMetrics);
    }
    // Merges aura metrics that have the same name/ID, adding their stats together.
    static join(auras) {
        const joinedById = bucket(auras, aura => actionIdToString(aura.actionId));
        return Object.values(joinedById).map(aurasToJoin => {
            const firstAura = aurasToJoin[0];
            return new AuraMetrics(firstAura.actionId, firstAura.name, firstAura.iconUrl, firstAura.iterations, firstAura.duration, AuraMetricsProto.create({
                uptimeSecondsAvg: Math.max(...aurasToJoin.map(a => a.data.uptimeSecondsAvg)),
            }));
        });
    }
}
;
// Manages the metrics for a single player action (e.g. Lightning Bolt).
export class ActionMetrics {
    constructor(actionId, name, iconUrl, iterations, duration, data) {
        this.actionId = actionId;
        this.name = name;
        this.iconUrl = iconUrl;
        this.iterations = iterations;
        this.duration = duration;
        this.data = data;
    }
    get damage() {
        return this.data.damage;
    }
    get dps() {
        return this.data.damage / this.iterations / this.duration;
    }
    get casts() {
        return this.data.casts / this.iterations;
    }
    get castsPerMinute() {
        return this.data.casts / this.iterations / (this.duration / 60);
    }
    get avgCast() {
        return this.data.damage / this.data.casts;
    }
    get hits() {
        return this.data.hits / this.iterations;
    }
    get avgHit() {
        return this.data.damage / this.data.hits;
    }
    get critPercent() {
        return (this.data.crits / this.data.hits) * 100;
    }
    get misses() {
        return this.data.misses / this.iterations;
    }
    get missPercent() {
        return (this.data.misses / (this.data.hits + this.data.misses)) * 100;
    }
    static async makeNew(iterations, duration, actionMetrics) {
        let actionId = null;
        if (actionMetrics.id.rawId.oneofKind == 'spellId') {
            actionId = {
                id: {
                    spellId: actionMetrics.id.rawId.spellId,
                },
                tag: actionMetrics.id.tag,
            };
        }
        else if (actionMetrics.id.rawId.oneofKind == 'itemId') {
            actionId = {
                id: {
                    itemId: actionMetrics.id.rawId.itemId,
                },
                tag: actionMetrics.id.tag,
            };
        }
        else if (actionMetrics.id.rawId.oneofKind == 'otherId') {
            actionId = {
                id: {
                    otherId: actionMetrics.id.rawId.otherId,
                },
                tag: actionMetrics.id.tag,
            };
        }
        else {
            throw new Error('Invalid action metric with no ID');
        }
        const namePromise = getName(actionId.id);
        const iconPromise = getIconUrl(actionId.id);
        let name = await namePromise;
        if (actionId.tag != 0) {
            if (name == "Mind Flay") { // for now we can just check the name and use special tagging rules.
                if (actionId.tag == 1) {
                    name += ' (1 Tick)';
                }
                else if (actionId.tag == 2) {
                    name += ' (2 Tick)';
                }
                else if (actionId.tag == 3) {
                    name += ' (3 Tick)';
                }
            }
            else {
                if (actionId.tag == 1) {
                    name += ' (LO)';
                }
                else if (actionId.tag == 10) {
                    name += ' (Auto)';
                }
                else if (actionId.tag == 11) {
                    name += ' (Offhand Auto)';
                }
                else {
                    name += ' (??)';
                }
            }
        }
        const iconUrl = await iconPromise;
        return new ActionMetrics(actionId, name, iconUrl, iterations, duration, actionMetrics);
    }
    // Merges action metrics that have the same name/ID, adding their stats together.
    static join(actions) {
        const joinedById = bucket(actions, action => actionIdToString(action.actionId));
        return Object.values(joinedById).map(actionsToJoin => {
            const firstAction = actionsToJoin[0];
            return new ActionMetrics(firstAction.actionId, firstAction.name, firstAction.iconUrl, firstAction.iterations, firstAction.duration, ActionMetricsProto.create({
                casts: sum(actionsToJoin.map(a => a.data.casts)),
                hits: sum(actionsToJoin.map(a => a.data.hits)),
                crits: sum(actionsToJoin.map(a => a.data.crits)),
                misses: sum(actionsToJoin.map(a => a.data.misses)),
                damage: sum(actionsToJoin.map(a => a.data.damage)),
            }));
        });
    }
}
