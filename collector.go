package peano

type Collector interface {
	Entities() []*Entity

	OnAdd() Collector
	//OnRemove() Collector
	OnUpdate() Collector

	Clear()
}

type collector struct {
	group    Group
	entities []*Entity
}

func newCollector(group Group) Collector {
	return &collector{group: group}
}

func (c *collector) Entities() []*Entity {
	return c.entities
}

func (c *collector) OnAdd() Collector {
	c.group.OnEntityAdded(c.addEntity)
	return c
}

func (c *collector) OnUpdate() Collector {
	c.group.OnEntityUpdate(c.addEntity)
	return c
}

func (c *collector) Clear() {
	c.entities = c.entities[:0]
}

func (c *collector) addEntity(g Group, e *Entity, tp int, cm Component) {
	c.entities = append(c.entities, e)
}
