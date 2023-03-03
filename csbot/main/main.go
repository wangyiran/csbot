package main

import (
	"csbot/csbot/internal/service"
	"fmt"
)

func main() {
	//fmt.Println("11")
	s, cf, err := service.New()
	defer cf()
	if err != nil {
		fmt.Printf("new service err , err:(%s)", err.Error())
	}
	//s.GetAllGoodsInfoFromPlatform()
	//s.Analyse()
	//ItemInfos := s.GetCompleteItemListFromDB()
	//s.FilterHotItems(ItemInfos)
	//s.AnalyseHotItems()

	s.DailyAnalyseHotItems()

}
