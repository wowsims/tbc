import { RaidTarget } from '/tbc/core/proto/common.js';
import { BuffBot as BuffBotProto } from '/tbc/core/proto/ui.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { buffBotPresets } from './presets.js';
export const NO_ASSIGNMENT = -1;
// Represents a buff bot in a raid.
export class BuffBot {
    constructor(id, sim) {
        this.spec = 0;
        this.name = '';
        this.raidIndex = NO_ASSIGNMENT;
        this.innervateAssignment = emptyRaidTarget();
        this.powerInfusionAssignment = emptyRaidTarget();
        this.raidIndexChangeEmitter = new TypedEvent();
        this.innervateAssignmentChangeEmitter = new TypedEvent();
        this.powerInfusionAssignmentChangeEmitter = new TypedEvent();
        this.changeEmitter = new TypedEvent();
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
    updateSettings() {
        this.spec = this.settings.spec;
        this.name = this.settings.name;
    }
    getLabel() {
        return `${this.name} (#${this.getRaidIndex() + 1})`;
    }
    getClass() {
        return specToClass[this.settings.spec];
    }
    getClassColor() {
        return classColors[this.getClass()];
    }
    getRaidIndex() {
        return this.raidIndex;
    }
    setRaidIndex(newRaidIndex) {
        if (newRaidIndex != this.raidIndex) {
            this.raidIndex = newRaidIndex;
            this.raidIndexChangeEmitter.emit();
            this.sim.raid.compChangeEmitter.emit();
        }
    }
    getPartyIndex() {
        return Math.floor(this.getRaidIndex() / 5);
    }
    getInnervateAssignment() {
        // Defensive copy.
        return RaidTarget.clone(this.innervateAssignment);
    }
    setInnervateAssignment(newInnervateAssignment) {
        if (RaidTarget.equals(newInnervateAssignment, this.innervateAssignment))
            return;
        // Defensive copy.
        this.innervateAssignment = RaidTarget.clone(newInnervateAssignment);
        this.innervateAssignmentChangeEmitter.emit();
    }
    getPowerInfusionAssignment() {
        // Defensive copy.
        return RaidTarget.clone(this.powerInfusionAssignment);
    }
    setPowerInfusionAssignment(newPowerInfusionAssignment) {
        if (RaidTarget.equals(newPowerInfusionAssignment, this.powerInfusionAssignment))
            return;
        // Defensive copy.
        this.powerInfusionAssignment = RaidTarget.clone(newPowerInfusionAssignment);
        this.powerInfusionAssignmentChangeEmitter.emit();
    }
    toProto() {
        return BuffBotProto.create({
            id: this.settings.buffBotId,
            raidIndex: this.getRaidIndex(),
            innervateAssignment: this.getInnervateAssignment(),
            powerInfusionAssignment: this.getPowerInfusionAssignment(),
        });
    }
    fromProto(proto) {
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
    clone() {
        const newBuffBot = new BuffBot(this.settings.buffBotId, this.sim);
        newBuffBot.fromProto(this.toProto());
        return newBuffBot;
    }
}
