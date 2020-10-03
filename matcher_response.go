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

type OrderbookResponse struct {
	Success bool `json:"success"`
	Message struct {
		Version          int    `json:"version"`
		ID               string `json:"id"`
		Sender           string `json:"sender"`
		SenderPublicKey  string `json:"senderPublicKey"`
		MatcherPublicKey string `json:"matcherPublicKey"`
		AssetPair        struct {
			AmountAsset string `json:"amountAsset"`
			PriceAsset  string `json:"priceAsset"`
		} `json:"assetPair"`
		OrderType         string      `json:"orderType"`
		Amount            int         `json:"amount"`
		Price             int         `json:"price"`
		Timestamp         int64       `json:"timestamp"`
		Expiration        int64       `json:"expiration"`
		MatcherFee        int         `json:"matcherFee"`
		MatcherFeeAssetID interface{} `json:"matcherFeeAssetId"`
		Signature         string      `json:"signature"`
		Proofs            []string    `json:"proofs"`
	} `json:"message"`
	Status string `json:"status"`
}
