package gowaves

import (
	"bytes"
	"encoding/json"
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

	url := fmt.Sprintf("http://%s:%s%s", w.Host, strconv.Itoa(int(w.Port)), uri)

	var req *http.Request
	var err error

	if body != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(body)
		req, err = http.NewRequest(method, url, b)
	} else {
		req, err = http.NewRequest(method, url, nil)
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
		}
		json.Unmarshal(body, resp)
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
