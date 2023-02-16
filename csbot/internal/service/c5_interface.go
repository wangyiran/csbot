package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	c5SellDataUrl     = "https://www.c5game.com/napi/trade/steamtrade/sga/sell/v3/list?itemId=%s&page=1&limit=20"
	c5PurchaseDataUrl = "https://www.c5game.com/napi/trade/steamtrade/sga/purchase/v2/list?itemId=%s&delivery=&page=1"
)

type C5SellResp struct {
	Data C5SellData `json:"data"`
}

type C5PurchaseResp struct {
	Data C5PurchaseData `json:"data"`
}

type C5SellData struct {
	Count int              `json:"total"`
	List  []C5SellGoodInfo `json:"list"`
}

type C5PurchaseData struct {
	Count int                  `json:"total"`
	List  []C5PurchaseGoodInfo `json:"list"`
}

type C5SellGoodInfo struct {
	Gid       string `json:"itemId"`
	Gname     string `json:"itemName"`
	SellPrice string `json:"cnyPrice"`
}

type C5PurchaseGoodInfo struct {
	Gid           string  `json:"itemId"`
	Gname         string  `json:"itemName"`
	PurchasePrice float32 `json:"cnyPrice"`
}

func (s *Service) GetC5SellData(gid string) (int, []C5SellGoodInfo, error) {
	url := fmt.Sprintf(c5SellDataUrl, gid)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GetC5SellData err , err:(%s)\n", err.Error())
	}

	defer resp.Body.Close()
	reader := resp.Body
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("GetC5SellData unmarshel body err , err:(%s)\n", err.Error())
		return 0, nil, nil
	}
	sellResp := C5SellResp{}
	_ = json.Unmarshal(b, &sellResp)
	return sellResp.Data.Count, sellResp.Data.List, err
}

func (s *Service) GetC5PurchaseData(gid string) (int, []C5PurchaseGoodInfo, error) {
	url := fmt.Sprintf(c5PurchaseDataUrl, gid)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GetC5PurchaseData err , err:(%s)\n", err.Error())
	}

	defer resp.Body.Close()
	reader := resp.Body
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("GetC5PurchaseData unmarshel body err , err:(%s)\n", err.Error())
		return 0, nil, nil
	}
	purchaseResp := C5PurchaseResp{}
	_ = json.Unmarshal(b, &purchaseResp)
	return purchaseResp.Data.Count, purchaseResp.Data.List, err
}
