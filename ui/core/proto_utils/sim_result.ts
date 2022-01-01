import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { AuraMetrics as AuraMetricsProto } from '/tbc/core/proto/api.js';
import { DpsMetrics as DpsMetricsProto } from '/tbc/core/proto/api.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterMetrics as EncounterMetricsProto } from '/tbc/core/proto/api.js';
import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { PartyMetrics as PartyMetricsProto } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { PlayerMetrics as PlayerMetricsProto } from '/tbc/core/proto/api.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidMetrics as RaidMetricsProto } from '/tbc/core/proto/api.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { TargetMetrics as TargetMetricsProto } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { ActionId } from '/tbc/core/resources.js';
import { getIconUrl } from '/tbc/core/resources.js';
import { getName } from '/tbc/core/resources.js';
import { actionIdToString } from '/tbc/core/resources.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { getTalentTreeIcon } from '/tbc/core/proto_utils/utils.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { bucket } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';

export interface SimResultFilter {
	// Raid index of the player to display, or null for all players.
	player?: number | null;

	// Target index of the target to display, or null for all targets.
	target?: number | null;
}

// Holds all the data from a simulation call, and provides helper functions
// for parsing it.
export class SimResult {
	readonly request: RaidSimRequest;
	readonly result: RaidSimResult;

	readonly raidMetrics: RaidMetrics;
	readonly encounterMetrics: EncounterMetrics;

	private constructor(request: RaidSimRequest, result: RaidSimResult, raidMetrics: RaidMetrics, encounterMetrics: EncounterMetrics) {
		this.request = request;
		this.result = result;
		this.raidMetrics = raidMetrics;
		this.encounterMetrics = encounterMetrics;
	}

	getPlayers(filter?: SimResultFilter): Array<PlayerMetrics> {
		if (filter?.player || filter?.player === 0) {
			const player = this.getPlayerWithRaidIndex(filter.player);
			return player ? [player] : [];
		} else {
			return this.raidMetrics.parties.map(party => party.players).flat();
		}
	}

	// Returns the first player, regardless of which party / raid slot its in.
	getFirstPlayer(): PlayerMetrics | null {
		return this.getPlayers()[0] || null;
	}

	getPlayerWithRaidIndex(raidIndex: number): PlayerMetrics | null {
		return this.getPlayers().find(player => player.raidIndex == raidIndex) || null;
	}

	getTargets(filter?: SimResultFilter): Array<TargetMetrics> {
		if (filter?.target || filter?.target === 0) {
			const target = this.getTargetWithIndex(filter.target);
			return target ? [target] : [];
		} else {
			return this.encounterMetrics.targets.slice();
		}
	}

	getTargetWithIndex(index: number): TargetMetrics | null {
		return this.getTargets().find(target => target.index == index) || null;
	}

	getDamageMetrics(filter: SimResultFilter): DpsMetricsProto {
		if (filter.player || filter.player === 0) {
			return this.getPlayerWithRaidIndex(filter.player)?.dps || DpsMetricsProto.create();
		}

		return this.raidMetrics.dps;
	}

	getActionMetrics(filter: SimResultFilter): Array<ActionMetrics> {
		return ActionMetrics.join(this.getPlayers(filter).map(player => player.actions).flat());
	}

	getSpellMetrics(filter: SimResultFilter): Array<ActionMetrics> {
		return this.getActionMetrics(filter).filter(e => e.hits + e.misses != 0);
	}

	getBuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.join(this.getPlayers(filter).map(player => player.auras).flat());
	}

	getDebuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.join(this.getTargets(filter).map(target => target.auras).flat());
	}

	getLogs(): Array<string> {
		return this.result.logs.split('\n');
	}

	toJson(): any {
		return {
			'request': RaidSimRequest.toJson(this.request),
			'result': RaidSimResult.toJson(this.result),
		};
	}

	static async fromJson(obj: any): Promise<SimResult> {
		const request = RaidSimRequest.fromJson(obj['request']);
		const result = RaidSimResult.fromJson(obj['result']);
		return SimResult.makeNew(request, result);
	}

	static async makeNew(request: RaidSimRequest, result: RaidSimResult): Promise<SimResult> {
		const iterations = request.simOptions?.iterations || 1;
		const duration = request.encounter?.duration || 1;

		const raidPromise = RaidMetrics.makeNew(iterations, duration, request.raid!, result.raidMetrics!);
		const encounterPromise = EncounterMetrics.makeNew(iterations, duration, request.encounter!, result.encounterMetrics!);

		const raidMetrics = await raidPromise;
		const encounterMetrics = await encounterPromise;

		return new SimResult(request, result, raidMetrics, encounterMetrics);
	}
}

export class RaidMetrics {
	private readonly raid: RaidProto;
	private readonly metrics: RaidMetricsProto;

	readonly dps: DpsMetricsProto;
	readonly parties: Array<PartyMetrics>;

	private constructor(raid: RaidProto, metrics: RaidMetricsProto, parties: Array<PartyMetrics>) {
		this.raid = raid;
		this.metrics = metrics;
		this.dps = this.metrics.dps!;
		this.parties = parties;
	}

	static async makeNew(iterations: number, duration: number, raid: RaidProto, metrics: RaidMetricsProto): Promise<RaidMetrics> {
		const numParties = Math.min(raid.parties.length, metrics.parties.length);
		
		const parties = await Promise.all(
				[...new Array(numParties).keys()]
						.map(i => PartyMetrics.makeNew(
								iterations,
								duration,
								raid.parties[i],
								metrics.parties[i],
								i)));

		return new RaidMetrics(raid, metrics, parties);
	}
}

export class PartyMetrics {
	private readonly party: PartyProto;
	private readonly metrics: PartyMetricsProto;

	readonly partyIndex: number;
	readonly dps: DpsMetricsProto;
	readonly players: Array<PlayerMetrics>;

	private constructor(party: PartyProto, metrics: PartyMetricsProto, partyIndex: number, players: Array<PlayerMetrics>) {
		this.party = party;
		this.metrics = metrics;
		this.partyIndex = partyIndex;
		this.dps = this.metrics.dps!;
		this.players = players;
	}

	static async makeNew(iterations: number, duration: number, party: PartyProto, metrics: PartyMetricsProto, partyIndex: number): Promise<PartyMetrics> {
		const numPlayers = Math.min(party.players.length, metrics.players.length);
		const players = await Promise.all(
				[...new Array(numPlayers).keys()]
						.filter(i => party.players[i].class != Class.ClassUnknown)
						.map(i => PlayerMetrics.makeNew(
								iterations,
								duration,
								party.players[i],
								metrics.players[i],
								partyIndex * 5 + i)));

		return new PartyMetrics(party, metrics, partyIndex, players);
	}
}

export class PlayerMetrics {
	private readonly player: PlayerProto;
	private readonly metrics: PlayerMetricsProto;

	readonly raidIndex: number;
	readonly name: string;
	readonly spec: Spec;
	readonly iconUrl: string;
	readonly classColor: string;
	readonly dps: DpsMetricsProto;
	readonly actions: Array<ActionMetrics>;
	readonly auras: Array<AuraMetrics>;
	private readonly iterations: number;
	private readonly duration: number;

	private constructor(player: PlayerProto, metrics: PlayerMetricsProto, raidIndex: number, actions: Array<ActionMetrics>, auras: Array<AuraMetrics>, iterations: number, duration: number) {
		this.player = player;
		this.metrics = metrics;

		this.raidIndex = raidIndex;
		this.name = player.name;
		this.spec = playerToSpec(player);
		this.iconUrl = getTalentTreeIcon(this.spec, player.talentsString);
		this.classColor = classColors[specToClass[this.spec]];
		this.dps = this.metrics.dps!;
		this.actions = actions;
		this.auras = auras;
		this.iterations = iterations;
		this.duration = duration;
	}

	get label() {
		return `${this.name} (#${this.raidIndex + 1})`;
	}

	get secondsOomAvg() {
		return this.metrics.secondsOomAvg
	}

	get totalDamage() {
		return this.dps.avg * this.duration;
	}

	static async makeNew(iterations: number, duration: number, player: PlayerProto, metrics: PlayerMetricsProto, raidIndex: number): Promise<PlayerMetrics> {
		const actionsPromise = Promise.all(metrics.actions.map(actionMetrics => ActionMetrics.makeNew(iterations, duration, actionMetrics)));
		const aurasPromise = Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(iterations, duration, auraMetrics)));

		const actions = await actionsPromise;
		const auras = await aurasPromise;
		return new PlayerMetrics(player, metrics, raidIndex, actions, auras, iterations, duration);
	}
}

export class EncounterMetrics {
	private readonly encounter: EncounterProto;
	private readonly metrics: EncounterMetricsProto;

	readonly targets: Array<TargetMetrics>;

	private constructor(encounter: EncounterProto, metrics: EncounterMetricsProto, targets: Array<TargetMetrics>) {
		this.encounter = encounter;
		this.metrics = metrics;
		this.targets = targets;
	}

	static async makeNew(iterations: number, duration: number, encounter: EncounterProto, metrics: EncounterMetricsProto): Promise<EncounterMetrics> {
		const numTargets = Math.min(encounter.targets.length, metrics.targets.length);
		const targets = await Promise.all(
				[...new Array(numTargets).keys()]
						.map(i => TargetMetrics.makeNew(
								iterations,
								duration,
								encounter.targets[i],
								metrics.targets[i],
								i)));

		return new EncounterMetrics(encounter, metrics, targets);
	}

	get durationSeconds() {
		return this.encounter.duration;
	}
}

export class TargetMetrics {
	private readonly target: TargetProto;
	private readonly metrics: TargetMetricsProto;

	readonly index: number;
	readonly auras: Array<AuraMetrics>;

	private constructor(target: TargetProto, metrics: TargetMetricsProto, index: number, auras: Array<AuraMetrics>) {
		this.target = target;
		this.metrics = metrics;

		this.index = index;
		this.auras = auras;
	}

	static async makeNew(iterations: number, duration: number, target: TargetProto, metrics: TargetMetricsProto, index: number): Promise<TargetMetrics> {
		const auras = await Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(iterations, duration, auraMetrics)));
		return new TargetMetrics(target, metrics, index, auras);
	}
}

export class AuraMetrics {
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: AuraMetricsProto;

	private constructor(actionId: ActionId, name: string, iconUrl: string, iterations: number, duration: number, data: AuraMetricsProto) {
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

	static async makeNew(iterations: number, duration: number, auraMetrics: AuraMetricsProto): Promise<AuraMetrics> {
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
	static join(auras: Array<AuraMetrics>): Array<AuraMetrics> {
		const joinedById = bucket(auras, aura => actionIdToString(aura.actionId));

		return Object.values(joinedById).map(aurasToJoin => {
			const firstAura = aurasToJoin[0];
			return new AuraMetrics(
					firstAura.actionId, firstAura.name, firstAura.iconUrl, firstAura.iterations, firstAura.duration,
					AuraMetricsProto.create({
						uptimeSecondsAvg: Math.max(...aurasToJoin.map(a => a.data.uptimeSecondsAvg)),
					}));
		});
	}
};

// Manages the metrics for a single player action (e.g. Lightning Bolt).
export class ActionMetrics {
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: ActionMetricsProto;

	private constructor(actionId: ActionId, name: string, iconUrl: string, iterations: number, duration: number, data: ActionMetricsProto) {
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

	static async makeNew(iterations: number, duration: number, actionMetrics: ActionMetricsProto): Promise<ActionMetrics> {
		let actionId: ActionId | null = null;
		if (actionMetrics.id!.rawId.oneofKind == 'spellId') {
			actionId = {
				id: {
					spellId: actionMetrics.id!.rawId.spellId,
				},
				tag: actionMetrics.id!.tag,
			};
		} else if (actionMetrics.id!.rawId.oneofKind == 'itemId') {
			actionId = {
				id: {
					itemId: actionMetrics.id!.rawId.itemId,
				},
				tag: actionMetrics.id!.tag,
			};
		} else if (actionMetrics.id!.rawId.oneofKind == 'otherId') {
			actionId = {
				id: {
					otherId: actionMetrics.id!.rawId.otherId,
				},
				tag: actionMetrics.id!.tag,
			};
		} else {
			throw new Error('Invalid action metric with no ID');
		}

		const namePromise = getName(actionId.id);
		const iconPromise = getIconUrl(actionId.id);

		let name = await namePromise;
		if (actionId.tag != 0) {
			if (name == "Mind Flay") { // for now we can just check the name and use special tagging rules.
				if (actionId.tag == 1) {
					name += ' (1 Tick)';
				} else if (actionId.tag == 2) {
					name += ' (2 Tick)';
				} else if (actionId.tag == 3) {
					name += ' (3 Tick)';
				}
			} else if (name == 'Fireball') {
				name += ' (DoT)';
			} else if (name == 'Pyroblast') {
				name += ' (DoT)';
			} else {
				if (actionId.tag == 1) {
				 	name += ' (LO)';
				} else if (actionId.tag == 10) {
					name += ' (Auto)';
				} else if (actionId.tag == 11) {
					name += ' (Offhand Auto)';
				} else {
					name += ' (??)';
				}
			} 
		} 

		const iconUrl = await iconPromise;

		return new ActionMetrics(actionId, name, iconUrl, iterations, duration, actionMetrics);
	}

	// Merges action metrics that have the same name/ID, adding their stats together.
	static join(actions: Array<ActionMetrics>): Array<ActionMetrics> {
		const joinedById = bucket(actions, action => actionIdToString(action.actionId));

		return Object.values(joinedById).map(actionsToJoin => {
			const firstAction = actionsToJoin[0];
			return new ActionMetrics(
					firstAction.actionId, firstAction.name, firstAction.iconUrl, firstAction.iterations, firstAction.duration,
					ActionMetricsProto.create({
						casts: sum(actionsToJoin.map(a => a.data.casts)),
						hits: sum(actionsToJoin.map(a => a.data.hits)),
						crits: sum(actionsToJoin.map(a => a.data.crits)),
						misses: sum(actionsToJoin.map(a => a.data.misses)),
						damage: sum(actionsToJoin.map(a => a.data.damage)),
					}));
		});
	}
}
