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
	Type            int         `json:"type"`
	ID              string      `json:"id"`
	Sender          string      `json:"sender"`
	SenderPublicKey string      `json:"senderPublicKey"`
	Fee             int         `json:"fee"`
	Timestamp       int64       `json:"timestamp"`
	Signature       string      `json:"signature,omitempty"`
	Version         int         `json:"version"`
	Recipient       string      `json:"recipient,omitempty"`
	AssetID         string      `json:"assetId"`
	FeeAssetID      interface{} `json:"feeAssetId,omitempty"`
	FeeAsset        interface{} `json:"feeAsset,omitempty"`
	Amount          int         `json:"amount,omitempty"`
	Attachment      string      `json:"attachment"`
	Height          int         `json:"height"`
	Proofs          []string    `json:"proofs,omitempty"`
	TransferCount   int         `json:"transferCount,omitempty"`
	TotalAmount     int         `json:"totalAmount,omitempty"`
	Transfers       []struct {
		Recipient string `json:"recipient"`
		Amount    int    `json:"amount"`
	} `json:"transfers,omitempty"`
}

type AssetsBalanceResponse struct {
	Address string `json:"address"`
	AssetID string `json:"assetId"`
	Balance int    `json:"balance"`
}
