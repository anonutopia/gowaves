package gowaves

type AssetsTransferRequest struct {
	Amount     int    `json:"amount"`
	AssetID    string `json:"assetId,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	Fee        int    `json:"fee"`
	FeeAssetID string `json:"feeAssetId,omitempty"`
	Recipient  string `json:"recipient"`
	Sender     string `json:"sender"`
	Timestamp  int    `json:"timestamp,omitempty"`
}

type AssetsMassTransferRequest struct {
	AssetID    string `json:"assetId"`
	Attachment string `json:"attachment"`
	Fee        int    `json:"fee"`
	Sender     string `json:"sender"`
	Timestamp  int    `json:"timestamp"`
	Transfers  []struct {
		Amount    int    `json:"amount"`
		Recipient string `json:"recipient"`
	} `json:"transfers"`
	Version int `json:"version"`
}
