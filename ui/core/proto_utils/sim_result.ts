import { ActionMetrics as ActionMetricsProto } from '/tbc/core/proto/api.js';
import { AuraMetrics as AuraMetricsProto } from '/tbc/core/proto/api.js';
import { DistributionMetrics as DistributionMetricsProto } from '/tbc/core/proto/api.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EncounterMetrics as EncounterMetricsProto } from '/tbc/core/proto/api.js';
import { Party as PartyProto } from '/tbc/core/proto/api.js';
import { PartyMetrics as PartyMetricsProto } from '/tbc/core/proto/api.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { PlayerMetrics as PlayerMetricsProto } from '/tbc/core/proto/api.js';
import { Raid as RaidProto } from '/tbc/core/proto/api.js';
import { RaidMetrics as RaidMetricsProto } from '/tbc/core/proto/api.js';
import { ResourceMetrics as ResourceMetricsProto, ResourceType } from '/tbc/core/proto/api.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { TargetMetrics as TargetMetricsProto } from '/tbc/core/proto/api.js';
import { RaidSimRequest, RaidSimResult } from '/tbc/core/proto/api.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { SimRun } from '/tbc/core/proto/ui.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { getTalentTreeIcon } from '/tbc/core/proto_utils/utils.js';
import { playerToSpec } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { bucket } from '/tbc/core/utils.js';
import { sum } from '/tbc/core/utils.js';

import {
	AuraUptimeLog,
	CastLog,
	DamageDealtLog,
	DpsLog,
	Entity,
	MajorCooldownUsedLog,
	ResourceChangedLogGroup,
	SimLog,
	ThreatLogGroup,
} from './logs_parser.js';

export interface SimResultFilter {
	// Raid index of the player to display, or null for all players.
	player?: number | null;

	// Target index of the target to display, or null for all targets.
	target?: number | null;
}

class SimResultData {
	readonly request: RaidSimRequest;
	readonly result: RaidSimResult;

	constructor(request: RaidSimRequest, result: RaidSimResult) {
		this.request = request;
		this.result = result;
	}

	get iterations() {
		return this.request.simOptions?.iterations || 1;
	}

	get duration() {
		return this.request.encounter?.duration || 1;
	}

	get firstIterationDuration() {
		return this.result.firstIterationDuration || 1;
	}
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

	getDamageMetrics(filter: SimResultFilter): DistributionMetricsProto {
		if (filter.player || filter.player === 0) {
			return this.getPlayerWithRaidIndex(filter.player)?.dps || DistributionMetricsProto.create();
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

	getResourceMetrics(filter: SimResultFilter, resourceType: ResourceType): Array<ResourceMetrics> {
		return ResourceMetrics.joinById(this.getPlayers(filter).map(player => player.resources.filter(resource => resource.type == resourceType)).flat());
	}

	getBuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.joinById(this.getPlayers(filter).map(player => player.auras).flat());
	}

	getDebuffMetrics(filter: SimResultFilter): Array<AuraMetrics> {
		return AuraMetrics.joinById(this.getTargets(filter).map(target => target.auras).flat());
	}

	toProto(): SimRun {
		return SimRun.create({
			request: this.request,
			result: this.result,
		});
	}

	static async fromProto(proto: SimRun): Promise<SimResult> {
		return SimResult.makeNew(proto.request || RaidSimRequest.create(), proto.result || RaidSimResult.create());
	}

	static async makeNew(request: RaidSimRequest, result: RaidSimResult): Promise<SimResult> {
		const resultData = new SimResultData(request, result);
		const logs = await SimLog.parseAll(result);

		const raidPromise = RaidMetrics.makeNew(resultData, request.raid!, result.raidMetrics!, logs);
		const encounterPromise = EncounterMetrics.makeNew(resultData, request.encounter!, result.encounterMetrics!, logs);

		const raidMetrics = await raidPromise;
		const encounterMetrics = await encounterPromise;

		return new SimResult(request, result, raidMetrics, encounterMetrics, logs);
	}
}

export class RaidMetrics {
	private readonly raid: RaidProto;
	private readonly metrics: RaidMetricsProto;

	readonly dps: DistributionMetricsProto;
	readonly parties: Array<PartyMetrics>;

	private constructor(raid: RaidProto, metrics: RaidMetricsProto, parties: Array<PartyMetrics>) {
		this.raid = raid;
		this.metrics = metrics;
		this.dps = this.metrics.dps!;
		this.parties = parties;
	}

	static async makeNew(resultData: SimResultData, raid: RaidProto, metrics: RaidMetricsProto, logs: Array<SimLog>): Promise<RaidMetrics> {
		const numParties = Math.min(raid.parties.length, metrics.parties.length);

		const parties = await Promise.all(
			[...new Array(numParties).keys()]
				.map(i => PartyMetrics.makeNew(
					resultData,
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
	readonly dps: DistributionMetricsProto;
	readonly players: Array<PlayerMetrics>;

	private constructor(party: PartyProto, metrics: PartyMetricsProto, partyIndex: number, players: Array<PlayerMetrics>) {
		this.party = party;
		this.metrics = metrics;
		this.partyIndex = partyIndex;
		this.dps = this.metrics.dps!;
		this.players = players;
	}

	static async makeNew(resultData: SimResultData, party: PartyProto, metrics: PartyMetricsProto, partyIndex: number, logs: Array<SimLog>): Promise<PartyMetrics> {
		const numPlayers = Math.min(party.players.length, metrics.players.length);
		const players = await Promise.all(
			[...new Array(numPlayers).keys()]
				.filter(i => party.players[i].class != Class.ClassUnknown)
				.map(i => PlayerMetrics.makeNew(
					resultData,
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
	readonly petActionId: ActionId | null;
	readonly iconUrl: string;
	readonly classColor: string;
	readonly dps: DistributionMetricsProto;
	readonly tps: DistributionMetricsProto;
	readonly actions: Array<ActionMetrics>;
	readonly auras: Array<AuraMetrics>;
	readonly resources: Array<ResourceMetrics>;
	readonly pets: Array<PlayerMetrics>;
	private readonly iterations: number;
	private readonly duration: number;

	readonly logs: Array<SimLog>;
	readonly damageDealtLogs: Array<DamageDealtLog>;
	readonly groupedResourceLogs: Record<ResourceType, Array<ResourceChangedLogGroup>>;
	readonly dpsLogs: Array<DpsLog>;
	readonly auraUptimeLogs: Array<AuraUptimeLog>;
	readonly majorCooldownLogs: Array<MajorCooldownUsedLog>;
	readonly castLogs: Array<CastLog>;
	readonly threatLogs: Array<ThreatLogGroup>;

	// Aura uptime logs, filtered to include only auras that correspond to a
	// major cooldown.
	readonly majorCooldownAuraUptimeLogs: Array<AuraUptimeLog>;

	private constructor(
		player: PlayerProto,
		petActionId: ActionId | null,
		metrics: PlayerMetricsProto,
		raidIndex: number,
		actions: Array<ActionMetrics>,
		auras: Array<AuraMetrics>,
		resources: Array<ResourceMetrics>,
		pets: Array<PlayerMetrics>,
		logs: Array<SimLog>,
		resultData: SimResultData) {
		this.player = player;
		this.metrics = metrics;

		this.raidIndex = raidIndex;
		this.name = metrics.name;
		this.spec = playerToSpec(player);
		this.petActionId = petActionId;
		this.iconUrl = getTalentTreeIcon(this.spec, player.talentsString);
		this.classColor = classColors[specToClass[this.spec]];
		this.dps = this.metrics.dps!;
		this.tps = this.metrics.threat!;
		this.actions = actions;
		this.auras = auras;
		this.resources = resources;
		this.pets = pets;
		this.logs = logs;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;

		this.damageDealtLogs = this.logs.filter((log): log is DamageDealtLog => log.isDamageDealt());
		this.dpsLogs = DpsLog.fromLogs(this.damageDealtLogs);
		this.castLogs = CastLog.fromLogs(this.logs);
		this.threatLogs = ThreatLogGroup.fromLogs(this.logs);

		this.auraUptimeLogs = AuraUptimeLog.fromLogs(this.logs, new Entity(this.name, '', this.raidIndex, false, this.isPet), resultData.firstIterationDuration);
		this.majorCooldownLogs = this.logs.filter((log): log is MajorCooldownUsedLog => log.isMajorCooldownUsed());

		this.groupedResourceLogs = ResourceChangedLogGroup.fromLogs(this.logs);
		AuraUptimeLog.populateActiveAuras(this.dpsLogs, this.auraUptimeLogs);
		AuraUptimeLog.populateActiveAuras(this.groupedResourceLogs[ResourceType.ResourceTypeMana], this.auraUptimeLogs);

		this.majorCooldownAuraUptimeLogs = this.auraUptimeLogs.filter(auraLog => this.majorCooldownLogs.find(mcdLog => mcdLog.actionId!.equals(auraLog.actionId!)));
	}

	get label() {
		return `${this.name} (#${this.raidIndex + 1})`;
	}

	get isPet() {
		return this.petActionId != null;
	}

	get maxThreat() {
		return this.threatLogs[this.threatLogs.length - 1]?.threatAfter || 0;
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

	getMeleeActions(): Array<ActionMetrics> {
		return this.actions.filter(e => e.hitAttempts != 0 && e.isMeleeAction);
	}

	getSpellActions(): Array<ActionMetrics> {
		return this.actions.filter(e => e.hitAttempts != 0 && !e.isMeleeAction)
	}

	getResourceMetrics(resourceType: ResourceType): Array<ResourceMetrics> {
		return this.resources.filter(resource => resource.type == resourceType);
	}

	static async makeNew(resultData: SimResultData, player: PlayerProto, metrics: PlayerMetricsProto, raidIndex: number, isPet: boolean, logs: Array<SimLog>): Promise<PlayerMetrics> {
		const playerLogs = logs.filter(log => log.source && (!log.source.isTarget && (isPet == log.source.isPet) && log.source.index == raidIndex));

		const actionsPromise = Promise.all(metrics.actions.map(actionMetrics => ActionMetrics.makeNew(null, resultData, actionMetrics, raidIndex)));
		const aurasPromise = Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(null, resultData, auraMetrics, raidIndex)));
		const resourcesPromise = Promise.all(metrics.resources.map(resourceMetrics => ResourceMetrics.makeNew(null, resultData, resourceMetrics, raidIndex)));
		const petsPromise = Promise.all(metrics.pets.map(petMetrics => PlayerMetrics.makeNew(resultData, player, petMetrics, raidIndex, true, playerLogs)));

		let petIdPromise: Promise<ActionId | null> = Promise.resolve(null);
		if (isPet) {
			petIdPromise = ActionId.fromPetName(metrics.name).fill(raidIndex);
		}

		const actions = await actionsPromise;
		const auras = await aurasPromise;
		const resources = await resourcesPromise;
		const pets = await petsPromise;
		const petActionId = await petIdPromise;

		const playerMetrics = new PlayerMetrics(player, petActionId, metrics, raidIndex, actions, auras, resources, pets, playerLogs, resultData);
		actions.forEach(action => action.player = playerMetrics);
		auras.forEach(aura => aura.player = playerMetrics);
		resources.forEach(resource => resource.player = playerMetrics);
		return playerMetrics;
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

	static async makeNew(resultData: SimResultData, encounter: EncounterProto, metrics: EncounterMetricsProto, logs: Array<SimLog>): Promise<EncounterMetrics> {
		const numTargets = Math.min(encounter.targets.length, metrics.targets.length);
		const targets = await Promise.all(
			[...new Array(numTargets).keys()]
				.map(i => TargetMetrics.makeNew(
					resultData,
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
	readonly auraUptimeLogs: Array<AuraUptimeLog>;

	private constructor(target: TargetProto, metrics: TargetMetricsProto, index: number, auras: Array<AuraMetrics>, logs: Array<SimLog>, resultData: SimResultData) {
		this.target = target;
		this.metrics = metrics;

		this.index = index;
		this.auras = auras;
		this.logs = logs;

		this.auraUptimeLogs = AuraUptimeLog.fromLogs(this.logs, new Entity('Target ' + (this.index + 1), '', this.index, true, false), resultData.firstIterationDuration);
	}

	static async makeNew(resultData: SimResultData, target: TargetProto, metrics: TargetMetricsProto, index: number, logs: Array<SimLog>): Promise<TargetMetrics> {
		const targetLogs = logs.filter(log => log.source && (log.source.isTarget && log.source.index == index));
		const auras = await Promise.all(metrics.auras.map(auraMetrics => AuraMetrics.makeNew(null, resultData, auraMetrics)));
		return new TargetMetrics(target, metrics, index, auras, targetLogs, resultData);
	}
}

export class AuraMetrics {
	player: PlayerMetrics | null;
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	private readonly resultData: SimResultData;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: AuraMetricsProto;

	private constructor(player: PlayerMetrics | null, actionId: ActionId, data: AuraMetricsProto, resultData: SimResultData) {
		this.player = player;
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.data = data;
		this.resultData = resultData;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;
	}

	get uptimePercent() {
		return this.data.uptimeSecondsAvg / this.duration * 100;
	}

	static async makeNew(player: PlayerMetrics | null, resultData: SimResultData, auraMetrics: AuraMetricsProto, playerIndex?: number): Promise<AuraMetrics> {
		const actionId = await ActionId.fromProto(auraMetrics.id!).fill(playerIndex);
		return new AuraMetrics(player, actionId, auraMetrics, resultData);
	}

	// Merges an array of metrics into a single metrics.
	static merge(auras: Array<AuraMetrics>, removeTag?: boolean, actionIdOverride?: ActionId): AuraMetrics {
		const firstAura = auras[0];
		const player = auras.every(aura => aura.player == firstAura.player) ? firstAura.player : null;
		let actionId = actionIdOverride || firstAura.actionId;
		if (removeTag) {
			actionId = actionId.withoutTag();
		}
		return new AuraMetrics(
			player,
			actionId,
			AuraMetricsProto.create({
				uptimeSecondsAvg: Math.max(...auras.map(a => a.data.uptimeSecondsAvg)),
			}),
			firstAura.resultData);
	}

	// Groups similar metrics, i.e. metrics with the same item/spell/other ID but
	// different tags, and returns them as separate arrays.
	static groupById(auras: Array<AuraMetrics>, useTag?: boolean): Array<Array<AuraMetrics>> {
		if (useTag) {
			return Object.values(bucket(auras, aura => aura.actionId.toString()));
		} else {
			return Object.values(bucket(auras, aura => aura.actionId.toStringIgnoringTag()));
		}
	}

	// Merges aura metrics that have the same name/ID, adding their stats together.
	static joinById(auras: Array<AuraMetrics>, useTag?: boolean): Array<AuraMetrics> {
		return AuraMetrics.groupById(auras, useTag).map(aurasToJoin => AuraMetrics.merge(aurasToJoin));
	}
};

export class ResourceMetrics {
	player: PlayerMetrics | null;
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	readonly type: ResourceType;
	private readonly resultData: SimResultData;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: ResourceMetricsProto;

	private constructor(player: PlayerMetrics | null, actionId: ActionId, data: ResourceMetricsProto, resultData: SimResultData) {
		this.player = player;
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.type = data.type;
		this.resultData = resultData;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;
		this.data = data;
	}

	get events() {
		return this.data.events / this.iterations;
	}

	get gain() {
		return this.data.gain / this.iterations;
	}

	get gainPerSecond() {
		return this.data.gain / this.iterations / this.duration;
	}

	get avgGain() {
		return this.data.gain / this.data.events;
	}

	get wastedGain() {
		return (this.data.gain - this.data.actualGain) / this.iterations;
	}

	static async makeNew(player: PlayerMetrics | null, resultData: SimResultData, resourceMetrics: ResourceMetricsProto, playerIndex?: number): Promise<ResourceMetrics> {
		const actionId = await ActionId.fromProto(resourceMetrics.id!).fill(playerIndex);
		return new ResourceMetrics(player, actionId, resourceMetrics, resultData);
	}

	// Merges an array of metrics into a single metrics.
	static merge(resources: Array<ResourceMetrics>, removeTag?: boolean, actionIdOverride?: ActionId): ResourceMetrics {
		const firstResource = resources[0];
		const player = resources.every(resource => resource.player == firstResource.player) ? firstResource.player : null;
		let actionId = actionIdOverride || firstResource.actionId;
		if (removeTag) {
			actionId = actionId.withoutTag();
		}
		return new ResourceMetrics(
			player,
			actionId,
			ResourceMetricsProto.create({
				events: sum(resources.map(a => a.data.events)),
				gain: sum(resources.map(a => a.data.gain)),
				actualGain: sum(resources.map(a => a.data.actualGain)),
			}),
			firstResource.resultData);
	}

	// Groups similar metrics, i.e. metrics with the same item/spell/other ID but
	// different tags, and returns them as separate arrays.
	static groupById(resources: Array<ResourceMetrics>, useTag?: boolean): Array<Array<ResourceMetrics>> {
		if (useTag) {
			return Object.values(bucket(resources, resource => resource.actionId.toString()));
		} else {
			return Object.values(bucket(resources, resource => resource.actionId.toStringIgnoringTag()));
		}
	}

	// Merges resource metrics that have the same name/ID, adding their stats together.
	static joinById(resources: Array<ResourceMetrics>, useTag?: boolean): Array<ResourceMetrics> {
		return ResourceMetrics.groupById(resources, useTag).map(resourcesToJoin => ResourceMetrics.merge(resourcesToJoin));
	}
};

// Manages the metrics for a single player action (e.g. Lightning Bolt).
export class ActionMetrics {
	player: PlayerMetrics | null;
	readonly actionId: ActionId;
	readonly name: string;
	readonly iconUrl: string;
	private readonly resultData: SimResultData;
	private readonly iterations: number;
	private readonly duration: number;
	private readonly data: ActionMetricsProto;

	private constructor(player: PlayerMetrics | null, actionId: ActionId, data: ActionMetricsProto, resultData: SimResultData) {
		this.player = player;
		this.actionId = actionId;
		this.name = actionId.name;
		this.iconUrl = actionId.iconUrl;
		this.resultData = resultData;
		this.iterations = resultData.iterations;
		this.duration = resultData.duration;
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

	get tps() {
		return this.data.threat / this.iterations / this.duration;
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

	get avgCastThreat() {
		return this.data.threat / this.data.casts;
	}

	get hits() {
		return this.data.hits / this.iterations;
	}

	private get landedHitsRaw() {
		if (this.data.isMelee) {
			return this.data.hits + this.data.crits + this.data.blocks + this.data.glances;
		} else {
			return this.data.hits;
		}
	}
	get landedHits() {
		return this.landedHitsRaw / this.iterations;
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
		return this.data.damage / this.landedHitsRaw;
	}

	get avgHitThreat() {
		return this.data.threat / this.landedHitsRaw;
	}

	get critPercent() {
		return (this.data.crits / this.hitAttempts) * 100;
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

	static async makeNew(player: PlayerMetrics | null, resultData: SimResultData, actionMetrics: ActionMetricsProto, playerIndex?: number): Promise<ActionMetrics> {
		const actionId = await ActionId.fromProto(actionMetrics.id!).fill(playerIndex);
		return new ActionMetrics(player, actionId, actionMetrics, resultData);
	}

	// Merges an array of metrics into a single metric.
	static merge(actions: Array<ActionMetrics>, removeTag?: boolean, actionIdOverride?: ActionId): ActionMetrics {
		const firstAction = actions[0];
		const player = actions.every(action => action.player == firstAction.player) ? firstAction.player : null;
		let actionId = actionIdOverride || firstAction.actionId;
		if (removeTag) {
			actionId = actionId.withoutTag();
		}
		return new ActionMetrics(
			player,
			actionId,
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
				threat: sum(actions.map(a => a.data.threat)),
			}),
			firstAction.resultData);
	}

	// Groups similar metrics, i.e. metrics with the same item/spell/other ID but
	// different tags, and returns them as separate arrays.
	static groupById(actions: Array<ActionMetrics>, useTag?: boolean): Array<Array<ActionMetrics>> {
		if (useTag) {
			return Object.values(bucket(actions, action => action.actionId.toString()));
		} else {
			return Object.values(bucket(actions, action => action.actionId.toStringIgnoringTag()));
		}
	}

	// Merges action metrics that have the same name/ID, adding their stats together.
	static joinById(actions: Array<ActionMetrics>, useTag?: boolean): Array<ActionMetrics> {
		return ActionMetrics.groupById(actions, useTag).map(actionsToJoin => ActionMetrics.merge(actionsToJoin));
	}
}
