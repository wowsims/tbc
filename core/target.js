import { MobType } from '/tbc/core/proto/common.js';
import { SpellSchool } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Target as TargetProto } from '/tbc/core/proto/common.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import * as Mechanics from '/tbc/core/constants/mechanics.js';
import { TypedEvent } from './typed_event.js';
// Manages all the settings for a single Target.
export class Target {
    constructor(sim) {
        this.id = 0;
        this.name = '';
        this.level = Mechanics.BOSS_LEVEL;
        this.mobType = MobType.MobTypeDemon;
        this.tankIndex = 0;
        this.stats = new Stats();
        this.swingSpeed = 0;
        this.minBaseDamage = 0;
        this.dualWield = false;
        this.dualWieldPenalty = false;
        this.canCrush = true;
        this.suppressDodge = false;
        this.parryHaste = true;
        this.spellSchool = SpellSchool.SpellSchoolPhysical;
        this.idChangeEmitter = new TypedEvent();
        this.nameChangeEmitter = new TypedEvent();
        this.levelChangeEmitter = new TypedEvent();
        this.mobTypeChangeEmitter = new TypedEvent();
        this.propChangeEmitter = new TypedEvent();
        this.statsChangeEmitter = new TypedEvent();
        // Emits when any of the above emitters emit.
        this.changeEmitter = new TypedEvent();
        this.sim = sim;
        [
            this.idChangeEmitter,
            this.nameChangeEmitter,
            this.levelChangeEmitter,
            this.mobTypeChangeEmitter,
            this.propChangeEmitter,
            this.statsChangeEmitter,
        ].forEach(emitter => emitter.on(eventID => this.changeEmitter.emit(eventID)));
        this.changeEmitter.on(eventID => this.sim.encounter?.changeEmitter.emit(eventID));
    }
    getId() {
        return this.id;
    }
    setId(eventID, newId) {
        if (newId == this.id)
            return;
        this.id = newId;
        this.idChangeEmitter.emit(eventID);
    }
    getName() {
        return this.name;
    }
    setName(eventID, newName) {
        if (newName == this.name)
            return;
        this.name = newName;
        this.nameChangeEmitter.emit(eventID);
    }
    getLevel() {
        return this.level;
    }
    setLevel(eventID, newLevel) {
        if (newLevel == this.level)
            return;
        this.level = newLevel;
        this.levelChangeEmitter.emit(eventID);
    }
    getMobType() {
        return this.mobType;
    }
    setMobType(eventID, newMobType) {
        if (newMobType == this.mobType)
            return;
        this.mobType = newMobType;
        this.mobTypeChangeEmitter.emit(eventID);
    }
    getTankIndex() {
        return this.tankIndex;
    }
    setTankIndex(eventID, newTankIndex) {
        if (newTankIndex == this.tankIndex)
            return;
        this.tankIndex = newTankIndex;
        this.propChangeEmitter.emit(eventID);
    }
    getSwingSpeed() {
        return this.swingSpeed;
    }
    setSwingSpeed(eventID, newSwingSpeed) {
        if (newSwingSpeed == this.swingSpeed)
            return;
        this.swingSpeed = newSwingSpeed;
        this.propChangeEmitter.emit(eventID);
    }
    getMinBaseDamage() {
        return this.minBaseDamage;
    }
    setMinBaseDamage(eventID, newMinBaseDamage) {
        if (newMinBaseDamage == this.minBaseDamage)
            return;
        this.minBaseDamage = newMinBaseDamage;
        this.propChangeEmitter.emit(eventID);
    }
    getDualWield() {
        return this.dualWield;
    }
    setDualWield(eventID, newDualWield) {
        if (newDualWield == this.dualWield)
            return;
        this.dualWield = newDualWield;
        this.propChangeEmitter.emit(eventID);
    }
    getDualWieldPenalty() {
        return this.dualWieldPenalty;
    }
    setDualWieldPenalty(eventID, newDualWieldPenalty) {
        if (newDualWieldPenalty == this.dualWieldPenalty)
            return;
        this.dualWieldPenalty = newDualWieldPenalty;
        this.propChangeEmitter.emit(eventID);
    }
    getCanCrush() {
        return this.canCrush;
    }
    setCanCrush(eventID, newCanCrush) {
        if (newCanCrush == this.canCrush)
            return;
        this.canCrush = newCanCrush;
        this.propChangeEmitter.emit(eventID);
    }
    getSuppressDodge() {
        return this.suppressDodge;
    }
    setSuppressDodge(eventID, newSuppressDodge) {
        if (newSuppressDodge == this.suppressDodge)
            return;
        this.suppressDodge = newSuppressDodge;
        this.propChangeEmitter.emit(eventID);
    }
    getParryHaste() {
        return this.parryHaste;
    }
    setParryHaste(eventID, newParryHaste) {
        if (newParryHaste == this.parryHaste)
            return;
        this.parryHaste = newParryHaste;
        this.propChangeEmitter.emit(eventID);
    }
    getSpellSchool() {
        return this.spellSchool;
    }
    setSpellSchool(eventID, newSpellSchool) {
        if (newSpellSchool == this.spellSchool)
            return;
        this.spellSchool = newSpellSchool;
        this.propChangeEmitter.emit(eventID);
    }
    getStats() {
        return this.stats;
    }
    setStats(eventID, newStats) {
        if (newStats.equals(this.stats))
            return;
        this.stats = newStats;
        this.statsChangeEmitter.emit(eventID);
    }
    matchesPreset(preset) {
        return TargetProto.equals(this.toProto(), preset.target);
    }
    applyPreset(eventID, preset) {
        this.fromProto(eventID, preset.target || TargetProto.create());
    }
    toProto() {
        return TargetProto.create({
            id: this.getId(),
            name: this.getName(),
            level: this.getLevel(),
            mobType: this.getMobType(),
            tankIndex: this.getTankIndex(),
            swingSpeed: this.getSwingSpeed(),
            minBaseDamage: this.getMinBaseDamage(),
            dualWield: this.getDualWield(),
            dualWieldPenalty: this.getDualWieldPenalty(),
            canCrush: this.getCanCrush(),
            suppressDodge: this.getSuppressDodge(),
            parryHaste: this.getParryHaste(),
            spellSchool: this.getSpellSchool(),
            stats: this.stats.asArray(),
        });
    }
    fromProto(eventID, proto) {
        TypedEvent.freezeAllAndDo(() => {
            let stats = new Stats(proto.stats);
            if (proto.armor) {
                stats = stats.withStat(Stat.StatArmor, proto.armor);
            }
            this.setId(eventID, proto.id);
            this.setName(eventID, proto.name);
            this.setLevel(eventID, proto.level);
            this.setMobType(eventID, proto.mobType);
            this.setTankIndex(eventID, proto.tankIndex);
            this.setSwingSpeed(eventID, proto.swingSpeed);
            this.setMinBaseDamage(eventID, proto.minBaseDamage);
            this.setDualWield(eventID, proto.dualWield);
            this.setDualWieldPenalty(eventID, proto.dualWieldPenalty);
            this.setCanCrush(eventID, proto.canCrush);
            this.setSuppressDodge(eventID, proto.suppressDodge);
            this.setParryHaste(eventID, proto.parryHaste);
            this.setSpellSchool(eventID, proto.spellSchool);
            this.setStats(eventID, stats);
        });
    }
    clone(eventID) {
        const newTarget = new Target(this.sim);
        newTarget.fromProto(eventID, this.toProto());
        return newTarget;
    }
    static defaultProto() {
        return TargetProto.create({
            level: Mechanics.BOSS_LEVEL,
            mobType: MobType.MobTypeDemon,
            tankIndex: 0,
            swingSpeed: 2,
            minBaseDamage: 10000,
            dualWield: false,
            dualWieldPenalty: false,
            canCrush: true,
            suppressDodge: false,
            parryHaste: true,
            spellSchool: SpellSchool.SpellSchoolPhysical,
            stats: Stats.fromMap({
                [Stat.StatArmor]: 7683,
                [Stat.StatBlockValue]: 54,
                [Stat.StatAttackPower]: 320,
            }).asArray(),
        });
    }
    static fromDefaults(eventID, sim) {
        const target = new Target(sim);
        target.fromProto(eventID, Target.defaultProto());
        return target;
    }
}
