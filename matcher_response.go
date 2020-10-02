package gowaves

type OrderbookStatusResponse struct {
	Success    bool   `json:"success"`
	Ask        int    `json:"ask"`
	BidAmount  int64  `json:"bidAmount"`
	Bid        int    `json:"bid"`
	LastAmount int64  `json:"lastAmount"`
	AskAmount  int64  `json:"askAmount"`
	LastSide   string `json:"lastSide"`
	Status     string `json:"status"`
	LastPrice  int    `json:"lastPrice"`
}
