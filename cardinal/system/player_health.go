package system

import (
	comp "GravitationioShard/component"
	"GravitationioShard/msg"
	"fmt"
	"pkg.world.dev/world-engine/cardinal"
)

func UpdatePlayerHealthSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.UpdatePlayerHealthMsg, msg.UpdatePlayerHealthMsgReply](
		world,
		func(update cardinal.TxData[msg.UpdatePlayerHealthMsg]) (msg.UpdatePlayerHealthMsgReply, error) {
			playerId, playerHealth, err := queryTargetPlayerHealth(world, update.Msg.Nickname)
			if err != nil {
				return msg.UpdatePlayerHealthMsgReply{}, fmt.Errorf("failed to update health: %w", err)
			}
			fmt.Println("Health before: ", playerHealth.HP)
			playerHealth.HP = update.Msg.Health
			fmt.Println("Health after: ", playerHealth.HP)
			if err := cardinal.SetComponent[comp.Health](world, playerId, playerHealth); err != nil {
				return msg.UpdatePlayerHealthMsgReply{}, fmt.Errorf("failed to update health: %w", err)
			}

			return msg.UpdatePlayerHealthMsgReply{}, nil
		})
}
