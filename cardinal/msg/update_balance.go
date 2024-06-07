package msg

import (
	"GravitationioShard/component"
	"fmt"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
)

type UpdatePlayerBalanceMsg struct {
	Nickname string
	Balance  [4]int
}

type UpdatePlayerBalanceMsgReply struct {
}

func UpdatePlayerBalance(world cardinal.WorldContext, req *UpdatePlayerBalanceMsg) (*UpdatePlayerBalanceMsgReply, error) {
	var playerBalance *component.Balance
	var err error
	var playerId types.EntityID
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[component.Player](), filter.Component[component.Balance]())).
		Each(world, func(id types.EntityID) bool {
			var player *component.Player
			player, err = cardinal.GetComponent[component.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
				playerBalance, err = cardinal.GetComponent[component.Balance](world, id)
				if err != nil {
					return false
				}
				playerId = id
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
	fmt.Println("Coins before: ", playerBalance.Coins)
	playerBalance.Coins = req.Balance
	fmt.Println("Coins after: ", playerBalance.Coins)

	err = cardinal.SetComponent[component.Balance](world, playerId, playerBalance)
	if err != nil {
		return nil, err
	}
	fmt.Println("Coins after set: ", playerBalance.Coins)

	return &UpdatePlayerBalanceMsgReply{}, nil
}
