package main

import (
	"fmt"
	"peano"
)

func main() {
	contexts := peano.CreateContexts(5)
	game := contexts.Game()

	// System registration
	systems := peano.CreateSystemPool()
	systems.Add(&Translate{})
	systems.Add(&ReactiveTranslate{})

	// Create entity
	player := game.CreateEntity()

	// Add component
	player.AddPosition(10, 30)
	player.AddDirection(0, 0)
	player.AddSpeed(5)

	// Remove component
	player.RemoveSpeed()

	// Replace component
	player.ReplacePosition(30, 10)
	fmt.Printf("X = %f,Y = %f", player.GetPosition().X, player.GetPosition().Y)

	// On or Off component
	player.OffDirection()
	player.OnDirection()

	// Destroy entity
	//player.Destroy()

	// GameLoop
	systems.Init(contexts)
	systems.Execute()
	systems.Clean()
	systems.Exit(contexts)
}

type Translate struct {
	group peano.Group
}

func (s *Translate) Initer(contexts peano.Contexts) {
	game := contexts.Game()
	s.group = game.Group(peano.NewMatcher().AllOf(peano.Position))
}

func (s *Translate) Executer() {
	for _, e := range s.group.GetEntities() {
		pos := e.GetPosition()
		e.ReplacePosition(pos.X+10, pos.X+10)
	}
	fmt.Println("执行Tanslate的Executer")
}

type ReactiveTranslate struct {
}

func (s *ReactiveTranslate) Trigger(contexts peano.Contexts) peano.Collector {
	game := contexts.Game()
	return game.Collector(peano.NewMatcher().AllOf(peano.Position)).OnUpdate().OnAdd()
}

func (s *ReactiveTranslate) Filter(entity *peano.Entity) bool {
	return entity.Has(peano.Position)
}

func (s *ReactiveTranslate) Executer(entities []*peano.Entity) {
	for _, e := range entities {
		pos := e.GetPosition()
		e.ReplacePosition(pos.X+10, pos.X+10)
	}
}
