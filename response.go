package gowaves

import (
	"time"
)

type ActivationStatusResponse struct {
	Features []struct {
		ActivationHeight int    `json:"activationHeight"`
		BlockchainStatus string `json:"blockchainStatus"`
		ID               int    `json:"id"`
		NodeStatus       string `json:"nodeStatus"`
	} `json:"features"`
	Height          int `json:"height"`
	NextCheck       int `json:"nextCheck"`
	VotingInterval  int `json:"votingInterval"`
	VotingThreshold int `json:"votingThreshold"`
}

type NodeVersionResponse struct {
	Version string `json:"version"`
}

type NodeStatusResponse struct {
	BlockchainHeight int       `json:"blockchainHeight"`
	StateHeight      int       `json:"stateHeight"`
	UpdatedTimestamp int64     `json:"updatedTimestamp"`
	UpdatedDate      time.Time `json:"updatedDate"`
}

type NodeStopResponse struct {
	Stopped bool `json:"stopped"`
}

type AddressesResponse []string

type AddressValidateResponse struct {
	Address string `json:"address"`
	Valid   bool   `json:"valid"`
}

type AddressesBalanceResponse struct {
	Address       string `json:"address"`
	Balance       int    `json:"balance"`
	Confirmations int    `json:"confirmations"`
}

type AssetsTransferResponse struct {
	Amount          int         `json:"amount"`
	AssetID         string      `json:"assetId"`
	Attachment      string      `json:"attachment"`
	Fee             int         `json:"fee"`
	FeeAsset        interface{} `json:"feeAsset"`
	ID              string      `json:"id"`
	Recipient       string      `json:"recipient"`
	Sender          string      `json:"sender"`
	SenderPublicKey string      `json:"senderPublicKey"`
	Signature       string      `json:"signature"`
	Timestamp       int         `json:"timestamp"`
	Type            int         `json:"type"`
}

type AssetsMassTransferResponse struct {
	AssetID         string   `json:"assetId"`
	Attachment      string   `json:"attachment"`
	Fee             int      `json:"fee"`
	ID              string   `json:"id"`
	Proofs          []string `json:"proofs"`
	Sender          string   `json:"sender"`
	SenderPublicKey string   `json:"senderPublicKey"`
	Timestamp       int      `json:"timestamp"`
	TotalAmount     int      `json:"totalAmount"`
	TransferCount   int      `json:"transferCount"`
	Transfers       []struct {
		Amount    int    `json:"amount"`
		Recipient string `json:"recipient"`
	} `json:"transfers"`
	Type    int `json:"type"`
	Version int `json:"version"`
}

type AssetsOrderResponse struct {
	Version          int    `json:"version"`
	ID               string `json:"id"`
	Sender           string `json:"sender"`
	SenderPublicKey  string `json:"senderPublicKey"`
	MatcherPublicKey string `json:"matcherPublicKey"`
	AssetPair        struct {
		AmountAsset string      `json:"amountAsset"`
		PriceAsset  interface{} `json:"priceAsset"`
	} `json:"assetPair"`
	OrderType         string      `json:"orderType"`
	Amount            int         `json:"amount"`
	Price             int         `json:"price"`
	Timestamp         int64       `json:"timestamp"`
	Expiration        int64       `json:"expiration"`
	MatcherFee        int         `json:"matcherFee"`
	Signature         string      `json:"signature"`
	Proofs            []string    `json:"proofs"`
	MatcherFeeAssetID interface{} `json:"matcherFeeAssetId"`
}

type BlocksAtResponse struct {
	Version      int    `json:"version"`
	Timestamp    int64  `json:"timestamp"`
	Reference    string `json:"reference"`
	NxtConsensus struct {
		BaseTarget          int    `json:"base-target"`
		GenerationSignature string `json:"generation-signature"`
	} `json:"nxt-consensus"`
	Generator        string `json:"generator"`
	Signature        string `json:"signature"`
	Blocksize        int    `json:"blocksize"`
	TransactionCount int    `json:"transactionCount"`
	Fee              int    `json:"fee"`
	Transactions     []struct {
		Type            int    `json:"type"`
		ID              string `json:"id"`
		Sender          string `json:"sender"`
		SenderPublicKey string `json:"senderPublicKey"`
		Fee             int    `json:"fee"`
		Timestamp       int64  `json:"timestamp"`
		Signature       string `json:"signature"`
		Recipient       string `json:"recipient"`
		Amount          int64  `json:"amount"`
	} `json:"transactions"`
	Height int `json:"height"`
}

type BlocksHeightResponse struct {
	Height int `json:"height"`
}

type TransactionsAddressLimitResponse struct {
	Type            int      `json:"type"`
	ID              string   `json:"id"`
	Sender          string   `json:"sender"`
	SenderPublicKey string   `json:"senderPublicKey"`
	Fee             int      `json:"fee"`
	Timestamp       int64    `json:"timestamp"`
	Proofs          []string `json:"proofs,omitempty"`
	Version         int      `json:"version"`
	AssetID         string   `json:"assetId"`
	Attachment      string   `json:"attachment"`
	TransferCount   int      `json:"transferCount,omitempty"`
	Total           int      `json:"total,omitempty"`
	Transfers       []struct {
		Recipient string `json:"recipient"`
		Amount    int    `json:"amount"`
	} `json:"transfers,omitempty"`
	Height     int         `json:"height"`
	Signature  string      `json:"signature,omitempty"`
	Recipient  string      `json:"recipient,omitempty"`
	FeeAssetID interface{} `json:"feeAssetId,omitempty"`
	FeeAsset   interface{} `json:"feeAsset,omitempty"`
	Amount     int         `json:"amount,omitempty"`
	Order1     struct {
		Version          int    `json:"version"`
		ID               string `json:"id"`
		Sender           string `json:"sender"`
		SenderPublicKey  string `json:"senderPublicKey"`
		MatcherPublicKey string `json:"matcherPublicKey"`
		AssetPair        struct {
			AmountAsset string      `json:"amountAsset"`
			PriceAsset  interface{} `json:"priceAsset"`
		} `json:"assetPair"`
		OrderType         string      `json:"orderType"`
		Amount            int         `json:"amount"`
		Price             int         `json:"price"`
		Timestamp         int64       `json:"timestamp"`
		Expiration        int64       `json:"expiration"`
		MatcherFee        int         `json:"matcherFee"`
		Signature         string      `json:"signature"`
		Proofs            []string    `json:"proofs"`
		MatcherFeeAssetID interface{} `json:"matcherFeeAssetId"`
	} `json:"order1"`
	Order2 struct {
		Version          int    `json:"version"`
		ID               string `json:"id"`
		Sender           string `json:"sender"`
		SenderPublicKey  string `json:"senderPublicKey"`
		MatcherPublicKey string `json:"matcherPublicKey"`
		AssetPair        struct {
			AmountAsset string      `json:"amountAsset"`
			PriceAsset  interface{} `json:"priceAsset"`
		} `json:"assetPair"`
		OrderType         string      `json:"orderType"`
		Amount            int64       `json:"amount"`
		Price             int         `json:"price"`
		Timestamp         int64       `json:"timestamp"`
		Expiration        int64       `json:"expiration"`
		MatcherFee        int         `json:"matcherFee"`
		Signature         string      `json:"signature"`
		Proofs            []string    `json:"proofs"`
		MatcherFeeAssetID interface{} `json:"matcherFeeAssetId"`
	} `json:"order2"`
	Price             int    `json:"price"`
	BuyMatcherFee     int    `json:"buyMatcherFee"`
	SellMatcherFee    int    `json:"sellMatcherFee"`
	ApplicationStatus string `json:"applicationStatus"`
}

type AssetsBalanceResponse struct {
	Address string `json:"address"`
	AssetID string `json:"assetId"`
	Balance int64  `json:"balance"`
}

type AssetsBalanceDistributionResponse struct {
	HasNext  bool           `json:"hasNext"`
	LastItem string         `json:"lastItem"`
	Items    map[string]int `json:"items,omitempty"`
}

type DataResponse []struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type UtilsHashSecureResponse struct {
	Message string `json:"message"`
	Hash    string `json:"hash"`
}
