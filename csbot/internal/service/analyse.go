package service

import (
	"csbot/csbot/utils"
	"fmt"
	"os"
	"strings"
)

func (s *Service) Analyse() {
	ItemInfos := s.GetCompleteItemListFromDB()

	filterV1(ItemInfos)

}

//0.06<差价比<0.2
func filterV1(infos []ItemInfo) {
	txt, _ := os.OpenFile(`analyse.txt`, os.O_APPEND|os.O_CREATE, 0666)
	defer func(txt *os.File) {
		err := txt.Close()
		if err != nil {
			panic(err)
		}
	}(txt)

	fmt.Println(len(infos))
	mp := make(map[string]int)
	for i := 0; i < len(infos); i++ {
		if mp[infos[i].ItemName] > 0 {
			continue
		}
		mp[infos[i].ItemName]++
		if strings.Contains(infos[i].ItemName, "战痕累累") || strings.Contains(infos[i].ItemName, "破损不堪") || strings.Contains(infos[i].ItemName, "印花") || strings.Contains(infos[i].ItemName, "刀") || strings.Contains(infos[i].ItemName, "匕首") {
			continue
		}
		minPrice := utils.MinPrice(infos[i].C5Price, infos[i].UuPrice)
		maxPrice := utils.MaxPrice(infos[i].C5Price, infos[i].UuPrice)
		// minPrice := utils.MinPrice(infos[i].C5Price, utils.MinPrice(infos[i].UuPrice, infos[i].BuffPrice))
		// maxPrice := utils.MaxPrice(infos[i].C5Price, utils.MaxPrice(infos[i].UuPrice, infos[i].BuffPrice))
		cha := (maxPrice - minPrice) / minPrice
		if cha < 0.06 || cha > 0.2 {
			continue
		}
		str := fmt.Sprintf("%s:%v:%v:%v:%v\n", infos[i].ItemName, cha, infos[i].C5Price, infos[i].UuPrice, infos[i].BuffPrice)
		_, _ = txt.Write([]byte(str))
	}
	return
}
