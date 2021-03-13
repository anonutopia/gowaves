package gowaves

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type WavesMatcherClient struct {
	Host   string
	Port   uint
	ApiKey string
}

func (w *WavesMatcherClient) DoRequest(uri string, method string, body interface{}, resp interface{}) error {
	cl := http.Client{}

	url := fmt.Sprintf("%s:%s%s", w.Host, strconv.Itoa(int(w.Port)), uri)

	var req *http.Request
	var err error

	if body != nil {
		b := new(bytes.Buffer)
		err = json.NewEncoder(b).Encode(body)
		if err != nil {
			return err
		}
		req, err = http.NewRequest(method, url, b)
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
	}

	if len(w.ApiKey) > 0 {
		req.Header.Set("X-API-Key", w.ApiKey)
	}
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	res, err := cl.Do(req)

	if err == nil {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return errors.New(string(body))
		}
		json.Unmarshal(body, resp)

		if WNC.Debug {
			log.Println(string(body))
		}
	} else {
		return err
	}

	return nil
}

func (w *WavesMatcherClient) OrderbookStatus(amountAssetID string, priceAssetID string) (*OrderbookStatusResponse, error) {
	osr := &OrderbookStatusResponse{}
	err := w.DoRequest(fmt.Sprintf("/matcher/orderbook/%s/%s/status", amountAssetID, priceAssetID), http.MethodGet, nil, osr)
	return osr, err
}

func (w *WavesMatcherClient) Orderbook(aor *AssetsOrderResponse) (*OrderbookResponse, error) {
	ors := &OrderbookResponse{}
	err := w.DoRequest("/matcher/orderbook", http.MethodPost, aor, ors)
	return ors, err
}

func (w *WavesMatcherClient) OrderbookMarket(aor *AssetsOrderResponse) (*OrderbookResponse, error) {
	ors := &OrderbookResponse{}
	err := w.DoRequest("/matcher/orderbook/market", http.MethodPost, aor, ors)
	return ors, err
}

func (w *WavesMatcherClient) OrderbookCancel(amountAssetID string, priceAssetID string, ocr *OrderbookCancelRequest) (*OrderbookCancelResponse, error) {
	ocrs := &OrderbookCancelResponse{}
	err := w.DoRequest(fmt.Sprintf("/matcher/orderbook/%s/%s/cancel", amountAssetID, priceAssetID), http.MethodPost, ocr, ocrs)
	return ocrs, err
}

func (w *WavesMatcherClient) OrderbookPair(amountAssetID string, priceAssetID string, depth int) (*OrderbookPairResponse, error) {
	opr := &OrderbookPairResponse{}
	err := w.DoRequest(fmt.Sprintf("/matcher/orderbook/%s/%s?depth=%d", amountAssetID, priceAssetID, depth), http.MethodGet, nil, opr)
	return opr, err
}
