package core

type cache struct {
	castPool []*Cast
}

// NewCast returns a cast from the pool, also fills the pool if there
// are no casts available in the pool.
func (p *cache) NewCast() *Cast {
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
	newCasts := make([]Cast, 1000)
	for i := range newCasts {
		p.castPool = append(p.castPool, &newCasts[i])
	}
}

// ReturnCasts returns a slice of casts back to the pool for reuse.
//  the casts are also zero'd
func (p *cache) ReturnCasts(casts []*Cast) {
	for _, v := range casts {
		v.Spell = nil
		v.CastTime = 0
		v.ManaCost = 0
		v.Dmg = 0
		v.Hit = 0
		v.Crit = 0
		v.CritBonus = 0
		v.DidHit = false
		v.DidCrit = false
		v.DidDmg = 0
		v.Effect = nil
		v.DoItNow = nil
	}

	p.castPool = append(p.castPool, casts...)
}
