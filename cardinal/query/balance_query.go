package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "GravitationioShard/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerGetBalanceRequest struct {
	Nickname string
}

type PlayerGetBalanceResponse struct {
	Balance [4]int
}

func PlayerGetBalance(world cardinal.WorldContext, req *PlayerGetBalanceRequest) (*PlayerGetBalanceResponse, error) {
	var playerBalance *comp.Balance
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Health](), filter.Component[comp.Balance]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
				playerBalance, err = cardinal.GetComponent[comp.Balance](world, id)
				if err != nil {
					return false
				}
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

	if playerBalance == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}
	fmt.Println("[GET] Coins: ", playerBalance.Coins)
	return &PlayerGetBalanceResponse{Balance: playerBalance.Coins}, nil
}
