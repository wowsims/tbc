package core

type cache struct {
	castPool []*DirectCastAction
}

// NewCast returns a cast from the pool, also fills the pool if there
// are no casts available in the pool.
func (p *cache) NewCast() *DirectCastAction {
	poolSize := len(p.castPool)

	if poolSize <= 0 {
		p.fillCasts()
		poolSize = len(p.castPool)
	}

	c := p.castPool[poolSize-1]
	p.castPool = p.castPool[:poolSize-1]
	return c
}

// fillCasts pre-allocates cast structs for use in simulation.
func (p *cache) fillCasts() {
	newCasts := make([]DirectCastAction, 1000)
	for i := range newCasts {
		p.castPool = append(p.castPool, &newCasts[i])
	}
}

// ReturnCasts returns a slice of casts back to the pool for reuse.
//  the casts are also zero'd
func (p *cache) ReturnCasts(casts []*DirectCastAction) {
	//for _, v := range casts {
	//	v.Spell = nil
	//	v.Caster = nil
	//	v.Tag = 0
	//	v.CastTime = 0
	//	v.ManaCost = 0
	//	v.BonusSpellPower = 0
	//	v.BonusHit = 0
	//	v.BonusCrit = 0
	//	v.CritDamageMultipier = 0
	//	v.DidHit = false
	//	v.DidCrit = false
	//	v.DidDmg = 0
	//	v.Effect = nil
	//	v.DoItNow = nil
	//}

	p.castPool = append(p.castPool, casts...)
}
