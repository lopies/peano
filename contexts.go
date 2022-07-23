package peano

const (
	Game = iota

	//NEXT2
)

type Contexts []interface{}

func SetContexts(ebs ...interface{}) Contexts {
	var contexts Contexts
	for _, v := range ebs {
		contexts = append(contexts, v)
	}
	return contexts
}

func (c *Contexts) Add(element interface{}) {
	*c = append(*c, element)
}

func (c *Contexts) Get(id int) interface{} {
	return (*c)[id]
}

func CreateContexts(gameComponentTotal int) Contexts {
	return SetContexts(
		CreateEntityBase(gameComponentTotal),
		//NEXT2
	)
}

func (c *Contexts) Game() EntityBase {
	return (*c)[Game].(EntityBase)
}
