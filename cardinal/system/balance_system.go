package system

// update balance of player

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"

	comp "GravitationioShard/component"
	"GravitationioShard/msg"
)

// UpdatePlayerBalanceSystem update balance of player
func UpdatePlayerBalanceSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.UpdatePlayerBalanceMsg, msg.UpdatePlayerBalanceMsgReply](
		world,
		func(update cardinal.TxData[msg.UpdatePlayerBalanceMsg]) (msg.UpdatePlayerBalanceMsgReply, error) {
			playerId, playerBalance, err := queryTargetPlayerBalance(world, update.Msg.Nickname)
			if err != nil {
				return msg.UpdatePlayerBalanceMsgReply{}, fmt.Errorf("failed to update balance: %w", err)
			}
			fmt.Println("Coins before: ", playerBalance.Coins)
			playerBalance.Coins = update.Msg.Balance
			fmt.Println("Coins after: ", playerBalance.Coins)
			if err := cardinal.SetComponent[comp.Balance](world, playerId, playerBalance); err != nil {
				return msg.UpdatePlayerBalanceMsgReply{}, fmt.Errorf("failed to update balance: %w", err)
			}

			return msg.UpdatePlayerBalanceMsgReply{}, nil
		})
}
