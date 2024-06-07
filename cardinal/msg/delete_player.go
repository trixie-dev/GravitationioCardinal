package msg

type DeletePlayerMsg struct {
	Nickname string `json:"nickname"`
}

type DeletePlayerResult struct {
	Success bool `json:"success"`
}
