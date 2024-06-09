package system

import (
	"fmt"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"pkg.world.dev/world-engine/cardinal"

	comp "GravitationioShard/component"
	"GravitationioShard/msg"
)

const (
	InitialHP = 100
)

// PlayerSpawnerSystem spawns players based on `CreatePlayer` transactions.
// This provides an example of a system that creates a new entity.
func PlayerSpawnerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreatePlayerMsg, msg.CreatePlayerResult](
		world,
		func(create cardinal.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerResult, error) {
			// Check if the player already exists
			exists, err := playerExists(world, create.Msg.Nickname)
			fmt.Println("Player exists: ", exists)
			if err != nil {
				return msg.CreatePlayerResult{}, fmt.Errorf("error checking if player exists: %w", err)
			}
			if exists {
				return msg.CreatePlayerResult{Success: false}, nil
			}
			// Create the player entity
			id, err := cardinal.Create(world,
				comp.Player{Nickname: create.Msg.Nickname},
				comp.Health{HP: InitialHP},
				comp.Balance{Coins: [4]int{0, 0, 0, 0}},
			)
			fmt.Println("Created Player ID: ", id)
			if err != nil {
				return msg.CreatePlayerResult{}, fmt.Errorf("error creating player: %w", err)
			}

			err = world.EmitEvent(map[string]any{
				"event": "new_player",
				"id":    id,
			})
			if err != nil {
				return msg.CreatePlayerResult{}, err
			}
			return msg.CreatePlayerResult{Success: true}, nil
		})
}

// checks if a player with the given nickname already exists.
func playerExists(world cardinal.WorldContext, nickname string) (bool, error) {
	var exists bool
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Health](), filter.Component[comp.Balance]())).Each(world,
		func(id types.EntityID) bool {
			player, err := cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			if player.Nickname == nickname {
				exists = true
				return false
			}

			return true
		})
	if searchErr != nil {
		return false, searchErr
	}
	return exists, nil
}
