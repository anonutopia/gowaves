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

type WavesNodeClient struct {
	Host   string
	Port   uint
	ApiKey string
}

func (w *WavesNodeClient) DoRequest(uri string, method string, body interface{}, resp interface{}) error {
	cl := http.Client{}
	var url string

	if w.Port != 80 {
		url = fmt.Sprintf("%s:%s%s", w.Host, strconv.Itoa(int(w.Port)), uri)
	} else {
		url = fmt.Sprintf("%s%s", w.Host, uri)
	}

	log.Println(url)

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

	req.Header.Set("X-API-Key", w.ApiKey)
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
			log.Printf("[WavesNodeClient.DoRequest] Error, body: %s", string(body))
			return errors.New(string(body))
		}
		json.Unmarshal(body, resp)
		// log.Println(string(body))
	} else {
		return err
	}

	return nil
}

func (w *WavesNodeClient) ActivationStatus() (*ActivationStatusResponse, error) {
	asr := &ActivationStatusResponse{}
	err := w.DoRequest("/activation/status", http.MethodGet, nil, asr)
	return asr, err
}

func (w *WavesNodeClient) NodeVersion() (*NodeVersionResponse, error) {
	nvr := &NodeVersionResponse{}
	err := w.DoRequest("/node/version", http.MethodGet, nil, nvr)
	return nvr, err
}

func (w *WavesNodeClient) NodeStop() (*NodeStopResponse, error) {
	nvr := &NodeStopResponse{}
	err := w.DoRequest("/node/stop", http.MethodPost, nil, nvr)
	return nvr, err
}

func (w *WavesNodeClient) NodeStatus() (*NodeStatusResponse, error) {
	nsr := &NodeStatusResponse{}
	err := w.DoRequest("/node/status", http.MethodGet, nil, nsr)
	return nsr, err
}

func (w *WavesNodeClient) AddressValidate(address string) (*AddressValidateResponse, error) {
	avr := &AddressValidateResponse{}
	err := w.DoRequest(fmt.Sprintf("/addresses/validate/%s", address), http.MethodGet, nil, avr)
	return avr, err
}

func (w *WavesNodeClient) AddressesBalance(address string) (*AddressesBalanceResponse, error) {
	abr := &AddressesBalanceResponse{}
	err := w.DoRequest(fmt.Sprintf("/addresses/balance/%s", address), http.MethodGet, nil, abr)
	return abr, err
}

func (w *WavesNodeClient) AssetsTransfer(atr *AssetsTransferRequest) (*AssetsTransferResponse, error) {
	atres := &AssetsTransferResponse{}
	err := w.DoRequest("/assets/transfer", http.MethodPost, atr, atres)
	return atres, err
}

func (w *WavesNodeClient) AssetsMassTransfer(amtr *AssetsMassTransferRequest) (*AssetsMassTransferResponse, error) {
	amtres := &AssetsMassTransferResponse{}
	err := w.DoRequest("/assets/masstransfer", http.MethodPost, amtr, amtres)
	return amtres, err
}

func (w *WavesNodeClient) AssetsOrder(aor *AssetsOrderRequest) (*AssetsOrderResponse, error) {
	aors := &AssetsOrderResponse{}
	err := w.DoRequest("/assets/order", http.MethodPost, aor, aors)
	return aors, err
}

func (w *WavesNodeClient) BlocksAt(n uint) (*BlocksAtResponse, error) {
	bar := &BlocksAtResponse{}
	err := w.DoRequest(fmt.Sprintf("/blocks/at/%d", n), http.MethodGet, nil, bar)
	return bar, err
}

func (w *WavesNodeClient) BlocksHeight() (*BlocksHeightResponse, error) {
	bhr := &BlocksHeightResponse{}
	err := w.DoRequest("/blocks/height", http.MethodGet, nil, bhr)
	return bhr, err
}

func (w *WavesNodeClient) TransactionsAddressLimit(address string, limit uint) ([][]TransactionsAddressLimitResponse, error) {
	var talr [][]TransactionsAddressLimitResponse
	err := w.DoRequest(fmt.Sprintf("/transactions/address/%s/limit/%d", address, limit), http.MethodGet, nil, &talr)
	return talr, err
}

func (w *WavesNodeClient) AssetsBalance(address string, assetId string) (*AssetsBalanceResponse, error) {
	abr := &AssetsBalanceResponse{}
	err := w.DoRequest(fmt.Sprintf("/assets/balance/%s/%s", address, assetId), http.MethodGet, nil, abr)
	return abr, err
}

func (w *WavesNodeClient) AssetsBalanceDistribution(assetId string, height int, limit int, after string) (*AssetsBalanceDistributionResponse, error) {
	abdr := &AssetsBalanceDistributionResponse{}
	url := fmt.Sprintf("/assets/%s/distribution/%s/limit/%s", assetId, strconv.Itoa(height), strconv.Itoa(limit))
	if len(after) > 0 {
		url = fmt.Sprintf("%s?after=%s", url, after)
	}
	err := w.DoRequest(url, http.MethodGet, nil, abdr)
	return abdr, err
}
