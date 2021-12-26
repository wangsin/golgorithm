package skiplist

/*
	Implement paper "Skip Lists: A Probabilistic Alternative to Balanced Trees" in Go.
*/

const (
	DefaultMaxLevel    uint8 = 32
	DefaultProbability       = 1 / 4
)

type element struct {
	Key, Value                  interface{}
	After, Before, Below, Above *element
}

type list struct {
	MaxLevel, CurrentHeight uint8
	Count                   int64
	Probability             float32
	HeadNode                *element
}

func NewSkipList() *list {
	return &list{
		MaxLevel:      DefaultMaxLevel,
		Probability:   DefaultProbability,
		CurrentHeight: 1,
		HeadNode:      new(element),
		Count:         0,
	}
}
