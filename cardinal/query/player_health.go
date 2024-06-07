package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "GravitationioShard/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerHealthRequest struct {
	Nickname string
}

type PlayerHealthResponse struct {
	HP int
}

func PlayerHealth(world cardinal.WorldContext, req *PlayerHealthRequest) (*PlayerHealthResponse, error) {
	fmt.Println("PlayerHealth")
	var playerHealth *comp.Health
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Health](), filter.Component[comp.Balance]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			fmt.Println("Name:", player.Nickname)
			fmt.Println("Req:", req.Nickname)
			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
				playerHealth, err = cardinal.GetComponent[comp.Health](world, id)
				fmt.Println("HP:", playerHealth.HP)
				if err != nil {

					return false
				}
				fmt.Println("Error:", err)
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return nil, searchErr
	}
	if err != nil {
		return nil, err
	}

	if playerHealth == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerHealthResponse{HP: playerHealth.HP}, nil
}
