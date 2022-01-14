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
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { getTalentTreeIcon } from '/tbc/core/proto_utils/utils.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { bucket } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';

import {
	AuraUptimeLog,
	DamageDealtLog,
	DpsLog,
	Entity,
	MajorCooldownUsedLog,
	ManaChangedLogGroup,
	SimLog,
} from './logs_parser.js';

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
	readonly logs: Array<SimLog>;

	private constructor(request: RaidSimRequest, result: RaidSimResult, raidMetrics: RaidMetrics, encounterMetrics: EncounterMetrics, logs: Array<SimLog>) {
		this.request = request;
		this.result = result;
		this.raidMetrics = raidMetrics;
		this.encounterMetrics = encounterMetrics;
		this.logs = logs;
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
		return ActionMetrics.joinById(this.getPlayers(filter).map(player => player.getPlayerAndPetActions()).flat());
	}

	getSpellMetrics(filter: SimResultFilter): Array<ActionMetrics> {
		return this.getActionMetrics(filter).filter(e => e.hitAttempts != 0 && !e.isMeleeAction)
	}

	getMeleeMetrics(filter: SimResultFilter): Array<ActionMetrics> {
		return this.getActionMetrics(filter).filter(e => e.hitAttempts != 0 && e.isMeleeAction);
	}

	getBuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.joinById(this.getPlayers(filter).map(player => player.auras).flat());
	}

	getDebuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.joinById(this.getTargets(filter).map(target => target.auras).flat());
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
		const logs = await SimLog.parseAll(result);

		const raidPromise = RaidMetrics.makeNew(iterations, duration, request.raid!, result.raidMetrics!, logs);
		const encounterPromise = EncounterMetrics.makeNew(iterations, duration, request.encounter!, result.encounterMetrics!, logs);

		const raidMetrics = await raidPromise;
		const encounterMetrics = await encounterPromise;

		return new SimResult(request, result, raidMetrics, encounterMetrics, logs);
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

	static async makeNew(iterations: number, duration: number, raid: RaidProto, metrics: RaidMetricsProto, logs: Array<SimLog>): Promise<RaidMetrics> {
		const numParties = Math.min(raid.parties.length, metrics.parties.length);
		
		const parties = await Promise.all(
				[...new Array(numParties).keys()]
						.map(i => PartyMetrics.makeNew(
								iterations,
								duration,
								raid.parties[i],
								metrics.parties[i],
								i,
								logs)));

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

	static async makeNew(iterations: number, duration: number, party: PartyProto, metrics: PartyMetricsProto, partyIndex: number, logs: Array<SimLog>): Promise<PartyMetrics> {
		const numPlayers = Math.min(party.players.length, metrics.players.length);
		const players = await Promise.all(
				[...new Array(numPlayers).keys()]
						.filter(i => party.players[i].class != Class.ClassUnknown)
						.map(i => PlayerMetrics.makeNew(
								iterations,
								duration,
								party.players[i],
								metrics.players[i],
								partyIndex * 5 + i,
								false,
								logs)));

		return new PartyMetrics(party, metrics, partyIndex, players);
	}
}

export class PlayerMetrics {
	// If this Player is a pet, player is the owner.
	private readonly player: PlayerProto;
	private readonly metrics: PlayerMetricsProto;

	readonly raidIndex: number;
	readonly name: string;
	readonly spec: Spec;
	readonly isPet: boolean;
	readonly iconUrl: string;
	readonly classColor: string;
	readonly dps: DpsMetricsProto;
	readonly actions: Array<ActionMetrics>;
	readonly auras: Array<AuraMetrics>;
	readonly pets: Array<PlayerMetrics>;
	private readonly iterations: number;
	private readonly duration: number;

	readonly logs: Array<SimLog>;
	readonly damageDealtLogs: Array<DamageDealtLog>;
	readonly manaChangedLogs: Array<ManaChangedLogGroup>;
	readonly dpsLogs: Array<DpsLog>;
	readonly auraUptimeLogs: Array<AuraUptimeLog>;
	readonly majorCooldownLogs: Array<MajorCooldownUsedLog>;

	// Aura uptime logs, filtered to include only auras that correspond to a
	// major cooldown.
	readonly majorCooldownAuraUptimeLogs: Array<AuraUptimeLog>;

	private constructor(
			player: PlayerProto,
			isPet: boolean,
			metrics: PlayerMetricsProto,
			raidIndex: number,
			actions: Array<ActionMetrics>,
			auras: Array<AuraMetrics>,
			pets: Array<PlayerMetrics>,
			logs: Array<SimLog>,
			iterations: number,
			duration: number) {
		this.player = player;
		this.metrics = metrics;

		this.raidIndex = raidIndex;
		this.name = player.name;
		this.spec = playerToSpec(player);
		this.isPet = isPet;
		this.iconUrl = getTalentTreeIcon(this.spec, player.talentsString);
		this.classColor = classColors[specToClass[this.spec]];
		this.dps = this.metrics.dps!;
		this.actions = actions;
		this.auras = auras;
		this.pets = pets;
		this.logs = logs;
		this.iterations = iterations;
		this.duration = duration;

		this.damageDealtLogs = this.logs.filter((log): log is DamageDealtLog => log.isDamageDealt());
		this.dpsLogs = DpsLog.fromLogs(this.damageDealtLogs);

		this.auraUptimeLogs = AuraUptimeLog.fromLogs(this.logs, new Entity(this.name, '', this.raidIndex, false, this.isPet));
		this.majorCooldownLogs = this.logs.filter((log): log is MajorCooldownUsedLog => log.isMajorCooldownUsed());

		this.manaChangedLogs = ManaChangedLogGroup.fromLogs(this.logs);
		AuraUptimeLog.populateActiveAuras(this.dpsLogs, this.auraUptimeLogs);
		AuraUptimeLog.populateActiveAuras(this.manaChangedLogs, this.auraUptimeLogs);

		this.majorCooldownAuraUptimeLogs = this.auraUptimeLogs.filter(auraLog => this.majorCooldownLogs.find(mcdLog => mcdLog.cooldownId.equals(auraLog.aura)));
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

	getPlayerAndPetActions(): Array<ActionMetrics> {
		return this.actions.concat(this.pets.map(pet => pet.getPlayerAndPetActions()).flat());
	}

	static async makeNew(iterations: number, duration: number, player: PlayerProto, metrics: PlayerMetricsProto, raidIndex: number, isPet: boolean, logs: Array<SimLog>): Promise<PlayerMetrics> {
		const playerLogs = logs.filter(log => log.source && (!log.source.isTarget && (isPet == log.source.isPet) && log.source.index == raidIndex));

		const actionsPromise = Promise.all(metrics.actions.map(actionMetrics => ActionMetrics.makeNew(iterations, duration, actionMetrics, raidIndex)));
		const aurasPromise = Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(iterations, duration, auraMetrics, raidIndex)));
		const petsPromise = Promise.all(metrics.pets.map(petMetrics => PlayerMetrics.makeNew(iterations, duration, player, petMetrics, raidIndex, true, playerLogs)));

		const actions = await actionsPromise;
		const auras = await aurasPromise;
		const pets = await petsPromise;
		return new PlayerMetrics(player, isPet, metrics, raidIndex, actions, auras, pets, playerLogs, iterations, duration);
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

	static async makeNew(iterations: number, duration: number, encounter: EncounterProto, metrics: EncounterMetricsProto, logs: Array<SimLog>): Promise<EncounterMetrics> {
		const numTargets = Math.min(encounter.targets.length, metrics.targets.length);
		const targets = await Promise.all(
				[...new Array(numTargets).keys()]
						.map(i => TargetMetrics.makeNew(
								iterations,
								duration,
								encounter.targets[i],
								metrics.targets[i],
								i,
								logs)));

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
	readonly logs: Array<SimLog>;

	private constructor(target: TargetProto, metrics: TargetMetricsProto, index: number, auras: Array<AuraMetrics>, logs: Array<SimLog>) {
		this.target = target;
		this.metrics = metrics;

		this.index = index;
		this.auras = auras;
		this.logs = logs;
	}

	static async makeNew(iterations: number, duration: number, target: TargetProto, metrics: TargetMetricsProto, index: number, logs: Array<SimLog>): Promise<TargetMetrics> {
		const targetLogs = logs.filter(log => log.source && (log.source.isTarget && log.source.index == index));
		const auras = await Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(iterations, duration, auraMetrics)));
		return new TargetMetrics(target, metrics, index, auras, targetLogs);
	}
}

export class AuraMetrics {
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: AuraMetricsProto;

	private constructor(actionId: ActionId, iterations: number, duration: number, data: AuraMetricsProto) {
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.iterations = iterations;
		this.duration = duration;
		this.data = data;
	}

	get uptimePercent() {
		return this.data.uptimeSecondsAvg / this.duration * 100;
	}

	static async makeNew(iterations: number, duration: number, auraMetrics: AuraMetricsProto, playerIndex?: number): Promise<AuraMetrics> {
		const actionId = await ActionId.fromProto(auraMetrics.id!).fill(playerIndex);
		return new AuraMetrics(actionId, iterations, duration, auraMetrics);
	}

	// Merges an array of metrics into a single metrics.
	static merge(auras: Array<AuraMetrics>): AuraMetrics {
		const firstAura = auras[0];
		return new AuraMetrics(
				firstAura.actionId, firstAura.iterations, firstAura.duration,
				AuraMetricsProto.create({
					uptimeSecondsAvg: Math.max(...auras.map(a => a.data.uptimeSecondsAvg)),
				}));
	}

	// Merges aura metrics that have the same name/ID, adding their stats together.
	static joinById(auras: Array<AuraMetrics>): Array<AuraMetrics> {
		const joinedById = bucket(auras, aura => aura.actionId.toString());
		return Object.values(joinedById).map(aurasToJoin => AuraMetrics.merge(aurasToJoin));
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

	private constructor(actionId: ActionId, iterations: number, duration: number, data: ActionMetricsProto) {
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.iterations = iterations;
		this.duration = duration;
		this.data = data;
	}

	get isMeleeAction() {
		return this.data.isMelee;
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

	get landedHits() {
		if (this.data.isMelee) {
			return this.data.hits + this.data.crits + this.data.blocks + this.data.glances;
		} else {
			return this.data.hits;
		}
	}

	get hitAttempts() {
		if (this.data.isMelee) {
			return this.data.misses
					+ this.data.dodges
					+ this.data.parries
					+ this.data.blocks
					+ this.data.glances
					+ this.data.crits
					+ this.data.hits;
		} else {
			return this.data.hits + this.data.misses;
		}
	}

	get avgHit() {
		return this.data.damage / this.landedHits;
	}

	get critPercent() {
		return (this.data.crits / this.landedHits) * 100;
	}

	get misses() {
		return this.data.misses / this.iterations;
	}

	get missPercent() {
		return (this.data.misses / this.hitAttempts) * 100;
	}

	get dodges() {
		return this.data.dodges / this.iterations;
	}

	get dodgePercent() {
		return (this.data.dodges / this.hitAttempts) * 100;
	}

	get parries() {
		return this.data.parries / this.iterations;
	}

	get parryPercent() {
		return (this.data.parries / this.hitAttempts) * 100;
	}

	get blocks() {
		return this.data.blocks / this.iterations;
	}

	get blockPercent() {
		return (this.data.blocks / this.hitAttempts) * 100;
	}

	get glances() {
		return this.data.glances / this.iterations;
	}

	get glancePercent() {
		return (this.data.glances / this.hitAttempts) * 100;
	}

	static async makeNew(iterations: number, duration: number, actionMetrics: ActionMetricsProto, playerIndex?: number): Promise<ActionMetrics> {
		const actionId = await ActionId.fromProto(actionMetrics.id!).fill(playerIndex);
		return new ActionMetrics(actionId, iterations, duration, actionMetrics);
	}

	// Merges an array of metrics into a single metric.
	static merge(actions: Array<ActionMetrics>, removeTag?: boolean): ActionMetrics {
		const firstAction = actions[0];
		let actionId = firstAction.actionId;
		if (removeTag) {
			actionId = actionId.withoutTag();
		}
		return new ActionMetrics(
				actionId, firstAction.iterations, firstAction.duration,
				ActionMetricsProto.create({
					isMelee: firstAction.isMeleeAction,
					casts: sum(actions.map(a => a.data.casts)),
					hits: sum(actions.map(a => a.data.hits)),
					crits: sum(actions.map(a => a.data.crits)),
					misses: sum(actions.map(a => a.data.misses)),
					dodges: sum(actions.map(a => a.data.dodges)),
					parries: sum(actions.map(a => a.data.parries)),
					blocks: sum(actions.map(a => a.data.blocks)),
					glances: sum(actions.map(a => a.data.glances)),
					damage: sum(actions.map(a => a.data.damage)),
				}));
	}

	// Merges action metrics that have the same name/ID, adding their stats together.
	static joinById(actions: Array<ActionMetrics>): Array<ActionMetrics> {
		const joinedById = bucket(actions, action => action.actionId.toString());
		return Object.values(joinedById).map(actionsToJoin => ActionMetrics.merge(actionsToJoin));
	}

	// Groups similar metrics, i.e. metrics with the same item/spell/other ID but
	// different tags, and returns them as separate arrays.
	static groupById(actions: Array<ActionMetrics>): Array<Array<ActionMetrics>> {
		return Object.values(bucket(actions, action => action.actionId.toStringIgnoringTag()));
	}
}
