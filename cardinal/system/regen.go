package system

import (
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "e2e_test/component"
)

// RegenSystem replenishes the player's HP at every tick.
// This provides an example of a system that doesn't rely on a transaction to update a component.
func RegenSystem(world cardinal.WorldContext) error {
	return cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Health]())).
		Each(world, func(id types.EntityID) bool {
			health, err := cardinal.GetComponent[comp.Health](world, id)
			if err != nil {
				return true
			}
			health.HP++
			if err := cardinal.SetComponent[comp.Health](world, id, health); err != nil {
				return true
			}
			return true
		})
}
