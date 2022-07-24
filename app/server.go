package main

import (
	"fmt"
	"peano"
)

func main() {
	contexts := peano.CreateContexts(5)
	game := contexts.Game()

	players := make([]*peano.Entity, 0, 10)
	for i := 0; i < 10; i++ {
		player := game.CreateEntity()
		player.AddPosition(0, 0)
		player.AddDirection(1000, 1000)
		players = append(players, player)
	}

	// System registration
	systems := peano.CreateSystemPool()
	systems.Add(&InputSystem{players: players})
	systems.Add(&DirectionReactiveTranslate{})
	systems.Add(&PositionReactiveTranslate{})

	// GameLoop
	systems.Init(contexts)
	systems.Execute()
	systems.Clean()
	systems.Exit(contexts)
}

type InputSystem struct {
	players []*peano.Entity
}

func (s *InputSystem) Initer(contexts peano.Contexts) {
	fmt.Println("input system init")
}

func (s *InputSystem) Executer() {
	for _, player := range s.players {
		pos := player.GetPosition()
		player.ReplacePosition(pos.X+1, pos.X+1)

		direction := player.GetDirection()
		player.ReplaceDirection(direction.Y+1, direction.X+1)
	}
}

type PositionReactiveTranslate struct {
}

func (s *PositionReactiveTranslate) Trigger(contexts peano.Contexts) peano.Collector {
	game := contexts.Game()
	return game.Collector(peano.NewMatcher().AllOf(peano.Position)).OnUpdate().OnAdd()
}

func (s *PositionReactiveTranslate) Filter(entity *peano.Entity) bool {
	return entity.Has(peano.Position)
}

func (s *PositionReactiveTranslate) Executer(entities []*peano.Entity) {
	for _, e := range entities {
		pos := e.GetPosition()
		fmt.Printf("PositionReactiveTranslate listen player pos change ,ID=%d,X = %f,Y=%f\n", e.ID(), pos.X, pos.Y)
	}
}

type DirectionReactiveTranslate struct {
}

func (s *DirectionReactiveTranslate) Trigger(contexts peano.Contexts) peano.Collector {
	game := contexts.Game()
	return game.Collector(peano.NewMatcher().AllOf(peano.Direction)).OnUpdate().OnAdd()
}

func (s *DirectionReactiveTranslate) Filter(entity *peano.Entity) bool {
	return entity.Has(peano.Direction)
}

func (s *DirectionReactiveTranslate) Executer(entities []*peano.Entity) {
	for _, e := range entities {
		d := e.GetDirection()
		fmt.Printf("DirectionReactiveTranslate listen player direction change ,ID=%d,d.x=%f,d.y=%f\n", e.ID(), d.X, d.Y)
	}
}
