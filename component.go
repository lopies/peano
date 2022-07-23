package peano

type Component any

type componentPool []Component

func (s *componentPool) Push(e Component) {
	*s = append(*s, e)
}

func (s *componentPool) Pop() (Component, bool) {
	lenght := len(*s)
	if lenght > 0 {
		last := lenght - 1
		entity := (*s)[last]
		*s = (*s)[:last]
		return entity, true
	}
	return nil, false
}
