package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"

	comp "GravitationioShard/component"
)

// SpawnDefaultPlayersSystem creates 10 players with nicknames "default-[0-9]". This System is registered as an
// Init system, meaning it will be executed exactly one time on tick 0.
func SpawnDefaultPlayersSystem(world cardinal.WorldContext) error {
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("default-%d", i)
		_, err := cardinal.Create(world,
			comp.Player{Nickname: name},
			comp.Health{HP: InitialHP},
			comp.Balance{Coins: [4]int{0, 0, 0, 0}},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
