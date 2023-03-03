package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	uuSellDataUrl     = "https://api.youpin898.com/api/homepage/v2/es/commodity/GetCsGoPagedList"
	uuPurchaseDataUrl = "https://api.youpin898.com/api/youpin/commodity/purchase/find"
)

type UuSellReq struct {
	Gid            string `json:"templateId"`
	PageSize       int    `json:"pageSize"`
	PageIndex      int    `json:"pageIndex"`
	SortType       int    `json:"sortType"`
	ListSortType   int    `json:"listSortType"`
	ListType       int    `json:"listType"`
	StickersIsSort bool   `json:"stickersIsSort"`
}

type UuPurchaseReq struct {
	Gid       string `json:"templateId"`
	PageSize  int    `json:"pageSize"`
	PageIndex int    `json:"pageIndex"`
}

type UuSellResp struct {
	Data UuSellData `json:"Data"`
}

type UuPurchaseResp struct {
	Data UuPurchaseData `json:"data"`
}

type UuSellData struct {
	List []UuSellGoodInfo `json:"CommodityList"`
}

type UuPurchaseData struct {
	List []UuPurchaseGoodInfo `json:"response"`
}

type UuSellGoodInfo struct {
	Gid       string `json:"TemplateId"`
	Gname     string `json:"CommodityName"`
	SellPrice string `json:"Price"`
}

type UuPurchaseGoodInfo struct {
	Gid           string  `json:"templateId"`
	Gname         string  `json:"commodityName"`
	PurchasePrice float32 `json:"unitPrice"`
}

func (s *Service) GetUuLowestSellPrice(gid string) (float32, error) {
	infos, err := s.GetUuSellData(gid)
	if len(infos) == 0 {
		return 1, err
	}
	num64, err := strconv.ParseFloat(infos[0].SellPrice, 32)
	num := float32(num64)
	return num, err
}

func (s *Service) GetUuLowestPurchasePrice(gid string) (float32, error) {
	infos, err := s.GetUuPurchaseData(gid)
	if len(infos) == 0 {
		return 1, err
	}
	return infos[0].PurchasePrice / 100, err
}

func (s *Service) GetUuSellData(gid string) ([]UuSellGoodInfo, error) {
	url := uuSellDataUrl

	body := UuSellReq{
		Gid:            gid,
		PageSize:       10,
		PageIndex:      1,
		SortType:       1,
		ListSortType:   1,
		ListType:       10,
		StickersIsSort: false,
	}
	bodyData, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GetUuSellData err , err:(%s)\n", err.Error())
	}

	defer resp.Body.Close()
	reader := resp.Body
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("GetUuSellData unmarshel body err , err:(%s)\n", err.Error())
		return nil, nil
	}
	sellResp := UuSellResp{}
	_ = json.Unmarshal(b, &sellResp)
	return sellResp.Data.List, err
}

func (s *Service) GetUuPurchaseData(gid string) ([]UuPurchaseGoodInfo, error) {
	url := uuPurchaseDataUrl

	body := UuPurchaseReq{
		Gid:       gid,
		PageSize:  10,
		PageIndex: 1,
	}
	bodyData, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GetUuSellData err , err:(%s)\n", err.Error())
	}

	defer resp.Body.Close()
	reader := resp.Body
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("GetUuSellData unmarshel body err , err:(%s)\n", err.Error())
		return nil, nil
	}
	sellResp := UuPurchaseResp{}
	_ = json.Unmarshal(b, &sellResp)
	return sellResp.Data.List, err
}
