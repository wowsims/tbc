import { Class } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { BuffBot as BuffBotProto } from '/tbc/core/proto/ui.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { Sim } from '/tbc/core/sim.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { BuffBotSettings, buffBotPresets } from './presets.js';

export const NO_ASSIGNMENT = -1;

// Represents a buff bot in a raid.
export class BuffBot {
	settings: BuffBotSettings;
	spec: Spec = 0;
	name: string = '';

	private raidIndex: number = NO_ASSIGNMENT;
	private innervateAssignment: RaidTarget = emptyRaidTarget();
	private powerInfusionAssignment: RaidTarget = emptyRaidTarget();

  readonly raidIndexChangeEmitter = new TypedEvent<void>();
  readonly innervateAssignmentChangeEmitter = new TypedEvent<void>();
  readonly powerInfusionAssignmentChangeEmitter = new TypedEvent<void>();
  readonly changeEmitter = new TypedEvent<void>();

	private readonly sim: Sim;

	constructor(id: string, sim: Sim) {
		const settings = buffBotPresets.find(preset => preset.buffBotId == id);
		if (!settings) {
			throw new Error('No buff bot config with id \'' + id + '\'!');
		}

		this.sim = sim;
		this.settings = settings;
		this.updateSettings();

		[
      this.raidIndexChangeEmitter,
      this.innervateAssignmentChangeEmitter,
      this.powerInfusionAssignmentChangeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));

		this.changeEmitter.on(() => sim.raid.getParty(this.getPartyIndex()).changeEmitter.emit());
	}

	private updateSettings() {
		this.spec = this.settings.spec;
		this.name = this.settings.name;
	}

	getLabel(): string {
		return `${this.name} (#${this.getRaidIndex() + 1})`;
	}

	getClass(): Class {
		return specToClass[this.settings.spec];
	}

	getClassColor(): string {
		return classColors[this.getClass()];
	}

	getRaidIndex(): number {
		return this.raidIndex;
	}
	setRaidIndex(newRaidIndex: number) {
		if (newRaidIndex != this.raidIndex) {
			this.raidIndex = newRaidIndex;
			this.raidIndexChangeEmitter.emit();
			this.sim.raid.compChangeEmitter.emit();
		}
	}

	getPartyIndex(): number {
		return Math.floor(this.getRaidIndex() / 5);
	}

	getInnervateAssignment(): RaidTarget {
		// Defensive copy.
		return RaidTarget.clone(this.innervateAssignment);
	}
	setInnervateAssignment(newInnervateAssignment: RaidTarget) {
		if (RaidTarget.equals(newInnervateAssignment, this.innervateAssignment))
			return;

		// Defensive copy.
		this.innervateAssignment = RaidTarget.clone(newInnervateAssignment);
		this.innervateAssignmentChangeEmitter.emit();
	}

	getPowerInfusionAssignment(): RaidTarget {
		// Defensive copy.
		return RaidTarget.clone(this.powerInfusionAssignment);
	}
	setPowerInfusionAssignment(newPowerInfusionAssignment: RaidTarget) {
		if (RaidTarget.equals(newPowerInfusionAssignment, this.powerInfusionAssignment))
			return;

		// Defensive copy.
		this.powerInfusionAssignment = RaidTarget.clone(newPowerInfusionAssignment);
		this.powerInfusionAssignmentChangeEmitter.emit();
	}

	toProto(): BuffBotProto {
		return BuffBotProto.create({
			id: this.settings.buffBotId,
			raidIndex: this.getRaidIndex(),
			innervateAssignment: this.getInnervateAssignment(),
			powerInfusionAssignment: this.getPowerInfusionAssignment(),
		});
	}

	fromProto(proto: BuffBotProto) {
		const settings = buffBotPresets.find(preset => preset.buffBotId == proto.id);
		if (!settings) {
			throw new Error('No buff bot config with id \'' + proto.id + '\'!');
		}
		this.settings = settings;
		this.updateSettings();
		this.setRaidIndex(proto.raidIndex);
		this.setInnervateAssignment(proto.innervateAssignment || emptyRaidTarget());
		this.setPowerInfusionAssignment(proto.powerInfusionAssignment || emptyRaidTarget());
	}

	clone(): BuffBot {
		const newBuffBot = new BuffBot(this.settings.buffBotId, this.sim);
		newBuffBot.fromProto(this.toProto());
		return newBuffBot;
	}
}
