package druid

type DruidForm uint8

const (
	Humanoid = 1 << iota
	Bear
	Cat
	Moonkin
)

func (druid *Druid) FormMatches(other DruidForm) bool {
	return (druid.Form & other) != 0
}

// We currently don't model shapeshifting directly. If we add it, add code here similar to sim/warrior/stances.go.
