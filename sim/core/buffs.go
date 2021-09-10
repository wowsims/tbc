package core

type Buffs struct {
	// Raid buffs
	ArcaneInt                bool
	GiftOfTheWild            bool
	BlessingOfKings          bool
	ImprovedBlessingOfWisdom bool
	ImprovedDivineSpirit     bool

	// Party Buffs
	Moonkin             bool
	MoonkinRavenGoddess bool   // adds 20 spell crit to moonkin aura
	SpriestDPS          uint16 // adds Mp5 ~ 25% (dps*5%*5sec = 25%)
	EyeOfNight          bool   // Eye of night bonus from party member (not you)
	TwilightOwl         bool   // from party member

	// Target Debuff
	JudgementOfWisdom         bool
	ImprovedSealOfTheCrusader bool
	Misery                    bool
}

type RaceBonusType byte

// These values are used directly in the dropdown in index.html
const (
	RaceBonusTypeNone RaceBonusType = iota
	RaceBonusTypeBloodElf
	RaceBonusTypeDraenei
	RaceBonusTypeDwarf
	RaceBonusTypeGnome
	RaceBonusTypeHuman
	RaceBonusTypeNightElf
	RaceBonusTypeOrc
	RaceBonusTypeTauren
	RaceBonusTypeTroll10
	RaceBonusTypeTroll30
	RaceBonusTypeUndead
)

type Consumes struct {
	// Buffs
	BrilliantWizardOil       bool
	MajorMageblood           bool
	FlaskOfBlindingLight     bool
	FlaskOfMightyRestoration bool
	BlackendBasilisk         bool

	// Used in rotations
	DestructionPotion bool
	SuperManaPotion   bool
	DarkRune          bool
	DrumsOfBattle     bool
}

func (c Consumes) AddStats(s Stats) Stats {
	if c.BrilliantWizardOil {
		s[StatSpellCrit] += 14
		s[StatSpellPower] += 36
	}
	if c.MajorMageblood {
		s[StatMP5] += 16.0
	}
	if c.FlaskOfBlindingLight {
		s[StatSpellPower] += 80
	}
	if c.FlaskOfMightyRestoration {
		s[StatMP5] += 25
	}
	if c.BlackendBasilisk {
		s[StatSpellPower] += 23
	}
	return s
}
