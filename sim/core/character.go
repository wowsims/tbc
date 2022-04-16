package core

import (
	"fmt"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Character is a data structure to hold all the shared values that all
// class logic shares.
// All players have stats, equipment, auras, etc
type Character struct {
	Unit

	Name  string // Different from Label, needed for returned results.
	Race  proto.Race
	Class proto.Class

	// Current gear.
	Equip items.Equipment

	// Pets owned by this Character.
	Pets []PetAgent

	rageBar
	energyBar

	// Consumables this Character will be using.
	Consumes proto.Consumes

	// Base stats for this Character.
	baseStats stats.Stats

	// Cast speed without any temporary effects.
	initialCastSpeed float64

	// Melee swing speed without any temporary effects.
	initialMeleeSwingSpeed float64

	// Ranged swing speed without any temporary effects.
	initialRangedSwingSpeed float64

	// Provides stat dependency management behavior.
	stats.StatDependencyManager

	// Provides major cooldown management behavior.
	majorCooldownManager

	// Up reference to this Character's Party.
	Party *Party

	// This character's index within its party [0-4].
	PartyIndex int

	// Used for applying the effects of hardcast / channeled spells at a later time.
	// By definition there can be only 1 hardcast spell being cast at any moment.
	Hardcast Hardcast
	//Hardcast PendingAction

	// AutoAttacks is the manager for auto attack swings.
	// Must be enabled to use "EnableAutoAttacks()"
	AutoAttacks AutoAttacks

	// Total amount of remaining additional mana expected for the current sim iteration,
	// beyond this Character's mana pool. This should include mana potions / runes /
	// innervates / etc.
	ExpectedBonusMana float64

	// Hack for ensuring we don't apply windfury totem aura if there's already
	// a MH imbue.
	// TODO: Figure out a cleaner way to do this.
	HasMHWeaponImbue bool

	// GCD-related PendingActions for this character.
	gcdAction      *PendingAction
	hardcastAction *PendingAction

	// Fields related to waiting for certain events to happen.
	waitingForMana float64
	waitStartTime  time.Duration

	// Cached mana return values per tick.
	manaTickWhileCasting    float64
	manaTickWhileNotCasting float64
}

func NewCharacter(party *Party, partyIndex int, player proto.Player) Character {
	character := Character{
		Unit: Unit{
			Type:        PlayerUnit,
			Index:       int32(party.Index*5 + partyIndex),
			Level:       CharacterLevel,
			auraTracker: newAuraTracker(),
			PseudoStats: stats.NewPseudoStats(),
			Metrics:     NewCharacterMetrics(),
		},

		Name:  player.Name,
		Race:  player.Race,
		Class: player.Class,
		Equip: items.ProtoToEquipment(*player.Equipment),

		Party:      party,
		PartyIndex: partyIndex,

		majorCooldownManager: newMajorCooldownManager(player.Cooldowns),
	}

	character.Label = fmt.Sprintf("%s (#%d)", character.Name, character.Index+1)

	if player.Consumes != nil {
		character.Consumes = *player.Consumes
	}

	character.baseStats = BaseStats[BaseStatsKey{Race: character.Race, Class: character.Class}]

	bonusStats := stats.Stats{}
	if player.BonusStats != nil {
		copy(bonusStats[:], player.BonusStats[:])
	}

	character.AddStats(character.baseStats)
	character.AddStats(bonusStats)
	character.addUniversalStatDependencies()

	return character
}

func (character *Character) addUniversalStatDependencies() {
	character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Stamina,
		ModifiedStat: stats.Health,
		Modifier: func(stamina float64, health float64) float64 {
			return health + stamina*10
		},
	})
	character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.Armor,
		Modifier: func(agility float64, armor float64) float64 {
			return armor + agility*2
		},
	})
}

func (character *Character) applyAllEffects(agent Agent, raidBuffs proto.RaidBuffs, partyBuffs proto.PartyBuffs, individualBuffs proto.IndividualBuffs) {
	agent.ApplyTalents()
	applyRaceEffects(agent)

	character.AddStats(character.Equip.Stats())
	character.applyItemEffects(agent)
	character.applyItemSetBonusEffects(agent)

	applyBuffEffects(agent, raidBuffs, partyBuffs, individualBuffs)
	applyConsumeEffects(agent, raidBuffs, partyBuffs)

	for _, petAgent := range character.Pets {
		applyPetBuffEffects(petAgent, raidBuffs, partyBuffs, individualBuffs)
	}
}

// Apply effects from all equipped items.
func (character *Character) applyItemEffects(agent Agent) {
	for slot, eq := range character.Equip {
		if applyItemEffect, ok := itemEffects[eq.ID]; ok {
			applyItemEffect(agent)
		}

		for _, g := range eq.Gems {
			if applyGemEffect, ok := itemEffects[g.ID]; ok {
				applyGemEffect(agent)
			}
		}

		// TODO: should we use eq.Enchant.EffectID because some enchants use a spellID instead of itemID?
		if applyEnchantEffect, ok := itemEffects[eq.Enchant.ID]; ok {
			applyEnchantEffect(agent)
		}

		if applyWeaponEffect, ok := weaponEffects[eq.Enchant.ID]; ok {
			applyWeaponEffect(agent, proto.ItemSlot(slot))
		}
	}
}

func (character *Character) AddPet(pet PetAgent) {
	if character.Unit.finalized {
		panic("Pets must be added before finalization!")
	}

	character.Pets = append(character.Pets, pet)
}

func (character *Character) AddStats(stat stats.Stats) {
	character.stats = character.stats.Add(stat)

	if len(character.Pets) > 0 {
		for _, petAgent := range character.Pets {
			petAgent.GetPet().addOwnerStats(stat)
		}
	}
}
func (character *Character) AddStatsDynamic(sim *Simulation, stat stats.Stats) {
	directStats := stat
	directStats[stats.Mana] = 0 // TODO: Mana needs special treatment

	if directStats[stats.MeleeHaste] != 0 {
		character.AddMeleeHaste(sim, directStats[stats.MeleeHaste])
		directStats[stats.MeleeHaste] = 0
	}

	character.stats = character.stats.Add(directStats)

	if len(character.Pets) > 0 {
		for _, petAgent := range character.Pets {
			petAgent.GetPet().addOwnerStats(stat)
		}
	}
}
func (character *Character) AddStat(stat stats.Stat, amount float64) {
	if character.Unit.finalized {
		if stat == stats.Mana {
			panic("Use AddMana!")
		}
		if stat == stats.MeleeHaste {
			panic("Use AddMeleeHaste!")
		}
	}

	character.stats[stat] += amount

	if stat == stats.MP5 || stat == stats.Intellect || stat == stats.Spirit {
		character.UpdateManaRegenRates()
	}

	if len(character.Pets) > 0 {
		for _, petAgent := range character.Pets {
			petAgent.GetPet().addOwnerStat(stat, amount)
		}
	}
}

func (character *Character) AddMeleeHaste(sim *Simulation, amount float64) {
	if amount > 0 {
		mod := 1 + (amount / (HasteRatingPerHastePercent * 100))
		character.AutoAttacks.ModifySwingTime(sim, mod)
	} else {
		mod := 1 / (1 + (-amount / (HasteRatingPerHastePercent * 100)))
		character.AutoAttacks.ModifySwingTime(sim, mod)
	}
	character.stats[stats.MeleeHaste] += amount

	// Could add melee haste to pets too, but not aware of any pets that scale with
	// owner's melee haste.
}

// MultiplyMeleeSpeed will alter the attack speed multiplier and change swing speed of all autoattack swings in progress.
func (character *Character) MultiplyMeleeSpeed(sim *Simulation, amount float64) {
	character.PseudoStats.MeleeSpeedMultiplier *= amount
	character.AutoAttacks.ModifySwingTime(sim, amount)
}

func (character *Character) MultiplyRangedSpeed(sim *Simulation, amount float64) {
	character.PseudoStats.RangedSpeedMultiplier *= amount
}

// Helper for when both MultiplyMeleeSpeed and MultiplyRangedSpeed are needed.
func (character *Character) MultiplyAttackSpeed(sim *Simulation, amount float64) {
	character.PseudoStats.MeleeSpeedMultiplier *= amount
	character.PseudoStats.RangedSpeedMultiplier *= amount
	character.AutoAttacks.ModifySwingTime(sim, amount)
}

func (character *Character) GetInitialStat(stat stats.Stat) float64 {
	return character.initialStats[stat]
}
func (character *Character) GetBaseStats() stats.Stats {
	return character.baseStats
}
func (character *Character) GetStats() stats.Stats {
	return character.stats
}
func (character *Character) GetStat(stat stats.Stat) float64 {
	return character.stats[stat]
}

// Returns whether the indicates stat is currently modified by a temporary bonus.
func (character *Character) HasTemporaryBonusForStat(stat stats.Stat) bool {
	return character.initialStats[stat] != character.stats[stat]
}

// Returns if spell casting has any temporary increases active.
func (character *Character) HasTemporarySpellCastSpeedIncrease() bool {
	return character.HasTemporaryBonusForStat(stats.SpellHaste) ||
		character.PseudoStats.CastSpeedMultiplier != 1
}

// Returns if melee swings have any temporary increases active.
func (character *Character) HasTemporaryMeleeSwingSpeedIncrease() bool {
	return character.SwingSpeed() != character.initialMeleeSwingSpeed
}

// Returns if ranged swings have any temporary increases active.
func (character *Character) HasTemporaryRangedSwingSpeedIncrease() bool {
	return character.RangedSwingSpeed() != character.initialRangedSwingSpeed
}

func (character *Character) InitialCastSpeed() float64 {
	return character.initialCastSpeed
}

func (character *Character) SpellGCD() time.Duration {
	return MaxDuration(GCDMin, time.Duration(float64(GCDDefault)/character.CastSpeed()))
}

func (character *Character) CastSpeed() float64 {
	return character.PseudoStats.CastSpeedMultiplier * (1 + (character.stats[stats.SpellHaste] / (HasteRatingPerHastePercent * 100)))
}

func (character *Character) ApplyCastSpeed(dur time.Duration) time.Duration {
	return time.Duration(float64(dur) / character.CastSpeed())
}

func (character *Character) SwingSpeed() float64 {
	return character.PseudoStats.MeleeSpeedMultiplier * (1 + (character.stats[stats.MeleeHaste] / (HasteRatingPerHastePercent * 100)))
}

func (character *Character) RangedSwingSpeed() float64 {
	return character.PseudoStats.RangedSpeedMultiplier * (1 + (character.stats[stats.MeleeHaste] / (HasteRatingPerHastePercent * 100)))
}

// Returns the crit multiplier for a spell.
// https://web.archive.org/web/20081014064638/http://elitistjerks.com/f31/t12595-relentless_earthstorm_diamond_-_melee_only/p4/
// https://github.com/TheGroxEmpire/TBC_DPS_Warrior_Sim/issues/30
func (character *Character) calculateCritMultiplier(normalCritDamage float64, primaryModifiers float64, secondaryModifiers float64) float64 {
	if character.HasMetaGemEquipped(34220) || character.HasMetaGemEquipped(32409) { // CSD and RED
		primaryModifiers *= 1.03
	}
	return 1.0 + (normalCritDamage*primaryModifiers-1.0)*(1.0+secondaryModifiers)
}
func (character *Character) SpellCritMultiplier(primaryModifiers float64, secondaryModifiers float64) float64 {
	return character.calculateCritMultiplier(1.5, primaryModifiers, secondaryModifiers)
}
func (character *Character) MeleeCritMultiplier(primaryModifiers float64, secondaryModifiers float64) float64 {
	return character.calculateCritMultiplier(2.0, primaryModifiers, secondaryModifiers)
}
func (character *Character) DefaultSpellCritMultiplier() float64 {
	return character.SpellCritMultiplier(1, 0)
}
func (character *Character) DefaultMeleeCritMultiplier() float64 {
	return character.MeleeCritMultiplier(1, 0)
}

func (character *Character) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}
func (character *Character) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	if character.Race == proto.Race_RaceDraenei {
		class := character.Class
		if class == proto.Class_ClassHunter ||
			class == proto.Class_ClassPaladin ||
			class == proto.Class_ClassWarrior {
			partyBuffs.DraeneiRacialMelee = true
		} else if class == proto.Class_ClassMage ||
			class == proto.Class_ClassPriest ||
			class == proto.Class_ClassShaman {
			partyBuffs.DraeneiRacialCaster = true
		}
	}

	if character.Consumes.Drums > 0 {
		partyBuffs.Drums = character.Consumes.Drums
	}

	if character.Equip[items.ItemSlotMainHand].ID == ItemIDAtieshMage {
		partyBuffs.AtieshMage += 1
	}
	if character.Equip[items.ItemSlotMainHand].ID == ItemIDAtieshWarlock {
		partyBuffs.AtieshWarlock += 1
	}

	if character.Equip[items.ItemSlotNeck].ID == ItemIDBraidedEterniumChain {
		partyBuffs.BraidedEterniumChain = true
	}
	if character.Equip[items.ItemSlotNeck].ID == ItemIDChainOfTheTwilightOwl {
		partyBuffs.ChainOfTheTwilightOwl = true
	}
	if character.Equip[items.ItemSlotNeck].ID == ItemIDEyeOfTheNight {
		partyBuffs.EyeOfTheNight = true
	}
	if character.Equip[items.ItemSlotNeck].ID == ItemIDJadePendantOfBlasting {
		partyBuffs.JadePendantOfBlasting = true
	}
}

func (character *Character) Finalize(raid *Raid) {
	if character.Unit.finalized {
		return
	}

	character.StatDependencyManager.Finalize()
	character.stats = character.ApplyStatDependencies(character.stats)

	character.Unit.finalize()

	// All stats added up to this point are part of the 'initial' stats.
	character.initialCastSpeed = character.CastSpeed()
	character.initialMeleeSwingSpeed = character.SwingSpeed()
	character.initialRangedSwingSpeed = character.RangedSwingSpeed()

	character.majorCooldownManager.finalize(character)

	for _, petAgent := range character.Pets {
		petAgent.GetPet().Finalize(raid)
	}
}

func (character *Character) init(sim *Simulation, agent Agent) {
	character.Unit.init(sim)
	agent.Init(sim)
}

func (character *Character) reset(sim *Simulation, agent Agent) {
	character.majorCooldownManager.reset(sim)
	character.Unit.reset(sim)

	character.ExpectedBonusMana = 0
	character.UpdateManaRegenRates()

	character.energyBar.reset(sim)
	character.rageBar.reset(sim)

	character.AutoAttacks.reset(sim)

	for _, petAgent := range character.Pets {
		petAgent.GetPet().reset(sim, petAgent)
	}

	if character.gcdAction != nil {
		sim.pendingActionPool.Put(character.gcdAction)
	}
	character.gcdAction = character.newGCDAction(sim, agent)

	agent.Reset(sim)
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (character *Character) advance(sim *Simulation, elapsedTime time.Duration) {
	character.Unit.advance(sim, elapsedTime)

	if character.Hardcast.Expires != 0 && character.Hardcast.Expires <= sim.CurrentTime {
		character.Hardcast.Expires = 0
		character.Hardcast.OnExpire(sim)
	}

	if len(character.Pets) > 0 {
		for _, petAgent := range character.Pets {
			petAgent.GetPet().advance(sim, elapsedTime)
		}
	}
}

func (character *Character) HasTrinketEquipped(itemID int32) bool {
	return character.Equip[items.ItemSlotTrinket1].ID == itemID ||
		character.Equip[items.ItemSlotTrinket2].ID == itemID
}

func (character *Character) HasMetaGemEquipped(gemID int32) bool {
	for _, gem := range character.Equip[items.ItemSlotHead].Gems {
		if gem.ID == gemID {
			return true
		}
	}
	return false
}

// Returns the MH weapon if one is equipped, and null otherwise.
func (character *Character) GetMHWeapon() *items.Item {
	weapon := &character.Equip[proto.ItemSlot_ItemSlotMainHand]
	if weapon.ID == 0 {
		return nil
	} else {
		return weapon
	}
}
func (character *Character) HasMHWeapon() bool {
	return character.GetMHWeapon() != nil
}

// Returns the OH weapon if one is equipped, and null otherwise. Note that
// shields / Held-in-off-hand items are NOT counted as weapons in this function.
func (character *Character) GetOHWeapon() *items.Item {
	weapon := &character.Equip[proto.ItemSlot_ItemSlotOffHand]
	if weapon.ID == 0 ||
		weapon.WeaponType == proto.WeaponType_WeaponTypeShield ||
		weapon.WeaponType == proto.WeaponType_WeaponTypeOffHand {
		return nil
	} else {
		return weapon
	}
}
func (character *Character) HasOHWeapon() bool {
	return character.GetOHWeapon() != nil
}

// Returns the ranged weapon if one is equipped, and null otherwise.
func (character *Character) GetRangedWeapon() *items.Item {
	weapon := &character.Equip[proto.ItemSlot_ItemSlotRanged]
	if weapon.ID == 0 ||
		weapon.RangedWeaponType == proto.RangedWeaponType_RangedWeaponTypeIdol ||
		weapon.RangedWeaponType == proto.RangedWeaponType_RangedWeaponTypeLibram ||
		weapon.RangedWeaponType == proto.RangedWeaponType_RangedWeaponTypeTotem {
		return nil
	} else {
		return weapon
	}
}
func (character *Character) HasRangedWeapon() bool {
	return character.GetRangedWeapon() != nil
}

// Returns the hands that the item is equipped in, as (MH, OH).
func (character *Character) GetWeaponHands(itemID int32) (bool, bool) {
	mh := false
	oh := false
	if weapon := character.GetMHWeapon(); weapon != nil && weapon.ID == itemID {
		mh = true
	}
	if weapon := character.GetOHWeapon(); weapon != nil && weapon.ID == itemID {
		oh = true
	}
	return mh, oh
}

func (character *Character) doneIteration(sim *Simulation) {
	// Need to do pets first so we can add their results to the owners.
	if len(character.Pets) > 0 {
		for _, petAgent := range character.Pets {
			pet := petAgent.GetPet()
			pet.doneIteration(sim)
			character.Metrics.AddFinalPetMetrics(&pet.Metrics)
		}
	}

	character.Hardcast = Hardcast{}
	character.doneIterationGCD(sim.Duration)

	character.Unit.doneIteration(sim)
}

func (character *Character) GetStatsProto() *proto.PlayerStats {
	gearStats := character.Equip.Stats()
	finalStats := character.GetStats()

	return &proto.PlayerStats{
		BaseStats:  character.baseStats[:],
		GearStats:  gearStats[:],
		FinalStats: finalStats[:],

		Sets:      character.GetActiveSetBonusNames(),
		Cooldowns: character.GetMajorCooldownIDs(),
	}
}

func (character *Character) GetMetricsProto(numIterations int32) *proto.PlayerMetrics {
	metrics := character.Metrics.ToProto(numIterations)
	metrics.Name = character.Name
	metrics.Auras = character.auraTracker.GetMetricsProto(numIterations)

	metrics.Pets = []*proto.PlayerMetrics{}
	for _, petAgent := range character.Pets {
		metrics.Pets = append(metrics.Pets, petAgent.GetPet().GetMetricsProto(numIterations))
	}

	return metrics
}

type BaseStatsKey struct {
	Race  proto.Race
	Class proto.Class
}

var BaseStats = map[BaseStatsKey]stats.Stats{}

// To calculate base stats, get a naked level 70 of the race/class you want, ideally without any talents to mess up base stats.
//  Basic stats are as-shown (str/agi/stm/int/spirit)

// Base Spell Crit is calculated by
//   1. Take as-shown value (troll shaman have 3.5%)
//   2. Calculate the bonus from int (for troll shaman that would be 104/78.1=1.331% crit)
//   3. Subtract as-shown from int bouns (3.5-1.331=2.169)
//   4. 2.169*22.08 (rating per crit percent) = 47.89 crit rating.

// Base mana can be looked up here: https://wowwiki-archive.fandom.com/wiki/Base_mana

// I assume a similar processes can be applied for other stats.
