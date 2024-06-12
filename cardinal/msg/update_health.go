package msg

type UpdatePlayerHealthMsg struct {
	Nickname string
	Health   int
}

type UpdatePlayerHealthMsgReply struct {
}
