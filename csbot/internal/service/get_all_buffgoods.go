package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	bufflistUrlFmt = "https://buff.163.com/api/market/goods?game=csgo&page_num=%d&use_suggestion=0&_=1676041947483&page_size=100"
)

//接口返回体
type BuffListResp struct {
	ListData BuffListData `json:"data"`
}

//接口返回体-data(饰品总数、页数、当前页饰品list)
type BuffListData struct {
	Page      int        `json:"page_num"`
	TotalPage int        `json:"total_page"`
	List      []BuffItem `json:"items"`
}

//Buff饰品信息
type BuffItem struct {
	Gid      int    `json:"id"`
	Gname    string `json:"name"`
	CnyPrice string `json:"sell_min_price"`
}

func (s *Service) GetAllBuffGoods() []BuffItem {

	BuffItems := make([]BuffItem, 0, 20000)
	return BuffItems
	_, pages, err := s.GetBuffItemsByPage(1)
	if err != nil {
		fmt.Printf("get Buff items list pages err , err:(%s)\n", err.Error())
	}
	fmt.Println(pages)

	rand.Seed(time.Now().UnixNano())
	for page := 1; page <= pages; page++ {
		items, _, err := s.GetBuffItemsByPage(page)
		fmt.Printf("%d-", page)
		slptm := 3000 + rand.Intn(1001)

		time.Sleep(time.Duration(slptm) * time.Millisecond)
		if err != nil {
			fmt.Printf("get Buff items list err , page:(%d) , err:(%s)\n", page, err.Error())
			return BuffItems
		}
		BuffItems = append(BuffItems, items...)
		//fmt.Println(BuffItems)

		fmt.Printf("%d\n", len(BuffItems))

	}
	fmt.Println("ok")
	fmt.Println(len(BuffItems))
	// txt, _ := os.OpenFile(`buff.txt`, os.O_APPEND|os.O_CREATE, 0666)
	// defer func(txt *os.File) {
	// 	err := txt.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }(txt)
	// for _, i := range BuffItems {
	// 	str := fmt.Sprintf("%+v\n", i)
	// 	_, _ = txt.Write([]byte(str))

	// }
	return BuffItems

}

func (s *Service) GetBuffItemsByPage(page int) (items []BuffItem, maxPage int, err error) {
	url := fmt.Sprintf(bufflistUrlFmt, page)
	//fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("new c5 request err")
		return
	}

	req.Header.Set("Cookie", "Device-Id=8PkRksQnIAb0XLPFt4ky; csrf_token=ImIwZjNmYWY2MTBhY2ZhYzQ5YWIxNzA2MTU2MmQwNzhiOWJjMzgxODYi.Fsj7nQ.nHNuwSwkiu2H2Dir2oK5kmIpcQ8; session=1-Zd1kcgq_V4f7q8sqkxCCNy9kmpKWeu-3eCDWitQh-mEk2045606301")
	req.Header.Set("Host", "<calculated when request is sent>")
	req.Header.Set("cookie", "nts_mail_user=18918068385@163.com:-1:1; _ntes_nnid=518a376ceaf0f225468095e4d1dca0eb,1645287208147; _ntes_nuid=518a376ceaf0f225468095e4d1dca0eb; NTES_P_UTID=NXekXtocqvKhOjJ5r7LzCuhSMN4thHXx|1669310007; Device-Id=SbqtApSjyweqcR7ohWUw; Locale-Supported=zh-Hans; game=csgo; NTES_YD_SESS=8r5Dt5Lr4xpeHaAeGBW5WOOWZHgpPl19hoJWULu_SSCKQkujQqfOaEtgNW4m2DjT761fDcuCNJkhRXyWF2W_Ng2m.6aZnSqI3GhOVZeKwJHhjyve.DZLdUIa7faVz8Pf88fjXseISaX7XkbtqQ4YIMlqbEriVSdC34CABlmLj5WNNZlKZxPEHenPxOQ6BJzuTtPLfN5ZFNxoRihVnurRewmy8dq5elplgVtIs.AGjGPdm; S_INFO=1676124147|0|0&60##|18918068385; P_INFO=18918068385|1676124147|1|netease_buff|00&99|null&null&null#shh&null#10#0|&0|null|18918068385; remember_me=U1092871877|kOnhSBbSKBJ9cgyoIXIBj7cwmTu75Va1; session=1-sQlgMttIXLjUL_CJ4Am3QH7_QA39_IvbjXG-kSch27NZ2045606301; csrf_token=Ijc1YTM1ZThlMzBiYWI3MWVhZDYzODlmZTQzM2ZlNWZlMzg4OTk5NGEi.Fsk73A.4SIT0GYVN_q23nQSY48-aOjJ8l4")

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
	//fmt.Printf("%v\n", string(b))
	resp := &BuffListResp{}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		fmt.Printf("unmarshel resp err:%s", string(b))
		return
	}
	//fmt.Println(resp.ListData.List)

	return resp.ListData.List, resp.ListData.TotalPage, nil
}
