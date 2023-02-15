package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	c5listUrlFmt = "https://www.c5game.com/napi/trade/steamtrade/sga/item-search/v2/list?limit=%d&appId=730&page=%d&orderBy=0&changePrice=5000&minPrice=%d&maxPrice=%d&quality=normal&rarity="
)

//接口返回体
type C5ListResp struct {
	Success  bool       `json:"success"`
	ListData C5ListData `json:"data"`
}

//接口返回体-data(饰品总数、页数、当前页饰品list)
type C5ListData struct {
	Pages int       `json:"pages"`
	List  []*C5Item `json:"list"`
}

//C5饰品信息
type C5Item struct {
	Gid      string `json:"itemId"`
	Gname    string `json:"itemName"`
	CnyPrice string `json:"cnyPrice"`
}

func (s *Service) GetAllC5Goods() []*C5Item {

	C5Items := make([]*C5Item, 0, 10000)

	_, pages, err := s.GetC5ItemsByPage(1)
	if err != nil {
		fmt.Printf("get c5 items list pages err , err:(%s)\n", err.Error())
	}
	for page := 1; page <= pages; page++ {
		items, _, err := s.GetC5ItemsByPage(page)
		if err != nil {
			fmt.Printf("get c5 items list err , page:(%d) , err:(%s)\n", page, err.Error())
			return C5Items
		}
		C5Items = append(C5Items, items...)
	}
	fmt.Println(len(C5Items))
	return C5Items

}

func (s *Service) GetC5ItemsByPage(page int) (items []*C5Item, maxPage int, err error) {
	limit := 200
	minPrice := 10
	maxPrice := 5000
	url := fmt.Sprintf(c5listUrlFmt, limit, page, minPrice, maxPrice)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("new c5 request err")
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("do c5 request err")
		return
	}
	defer res.Body.Close()

	reader := res.Body
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("reader err")
		return
	}

	resp := &C5ListResp{}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		fmt.Println("unmarshel resp err")
		return
	}
	return resp.ListData.List, resp.ListData.Pages, nil
}
