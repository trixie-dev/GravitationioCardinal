package component

type Balance struct {
	Coins [4]int
}

func (Balance) Name() string {
	return "Balance"
}
