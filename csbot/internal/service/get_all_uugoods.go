package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	uulistUrlFmt = "https://api.youpin898.com/api/homepage/es/template/GetCsGoPagedList"
)

type uuListRequest struct {
	GameId         string `json:"gameId"`
	ListSortType   string `json:"listSortType"`
	ListType       string `json:"listType"`
	MaxPrice       int    `json:"maxPrice"`
	MinPrice       int    `json:"minPrice"`
	PageIndex      int    `json:"pageIndex"`
	PageSize       int    `json:"pageSize"`
	SortType       string `json:"sortType"`
	StickersIsSort bool   `json:"stickersIsSort"`
}

//接口返回体
type UUListResp struct {
	ListData []*UUItem `json:"Data"`
}

//接口返回体-data(饰品总数、页数、当前页饰品list)
type UUListData struct {
	List []*UUItem `json:"commodityTemplateList"`
}

//C5饰品信息
type UUItem struct {
	Gid      int    `json:"Id"`
	Gname    string `json:"CommodityName"`
	CnyPrice string `json:"Price"`
}

func (s *Service) GetAllUUGoods() []*UUItem {

	UUItems := make([]*UUItem, 0, 10000)

	for page := 0; ; page++ {
		items, _, err := s.GetUUItemsByPage(page)
		//fmt.Println(page)
		if len(items) == 0 && err != nil {
			break
		}
		UUItems = append(UUItems, items...)

	}
	fmt.Println(len(UUItems))
	// txt, _ := os.OpenFile(`uu.txt`, os.O_APPEND|os.O_CREATE, 0666)
	// defer func(txt *os.File) {
	// 	err := txt.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }(txt)
	// for _, i := range UUItems {
	// 	str := fmt.Sprintf("%+v\n", i)
	// 	_, _ = txt.Write([]byte(str))

	// }
	return UUItems
}

func (s *Service) GetUUItemsByPage(page int) (items []*UUItem, maxPage int, err error) {
	url := uulistUrlFmt

	body := uuListRequest{
		GameId:         "730",
		ListSortType:   "2",
		ListType:       "10",
		MaxPrice:       5000,
		MinPrice:       10,
		PageIndex:      page,
		PageSize:       200,
		SortType:       "0",
		StickersIsSort: false,
	}
	bodyData, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyData))
	if err != nil {
		fmt.Println("new uu request err")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("do uu request err")
		return
	}
	defer res.Body.Close()

	reader := res.Body
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("reader err")
		return
	}
	resp := &UUListResp{}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		fmt.Println("unmarshel resp err")
		fmt.Println(err)
		return
	}
	return resp.ListData, 0, nil
}
