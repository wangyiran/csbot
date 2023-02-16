package service

import (
	"fmt"
	"strconv"
	"time"
)

const (
	hotSellLimit     = 20
	hotPurchaseLimit = 5
)

type itemMeta struct {
	ItemName string
	C5Gid    string
	BuffGid  string
	UuGid    string
	IgxeGid  string
}
type HotItemInfo struct {
	ItemName             string  `gorm:"Column:item_name"`
	C5Gid                string  `gorm:"Column:c5_gid"`
	C5SellPrice          float32 `gorm:"Column:c5_sell"`
	C5PurchasePrice      float32 `gorm:"Column:c5_purchase"`
	C5LastWeekAvgPrice   float32 `gorm:"Column:c5_lastweek_avgprice"`
	BuffGid              string  `gorm:"Column:buff_gid"`
	BuffSellPrice        float32 `gorm:"Column:buff_sell"`
	BuffPurchasePrice    float32 `gorm:"Column:buff_purchase"`
	BuffLastWeekAvgPrice float32 `gorm:"Column:buff_lastweek_avgprice"`
	UuGid                string  `gorm:"Column:uu_gid"`
	UuSellPrice          float32 `gorm:"Column:uu_sell"`
	UuPurchasePrice      float32 `gorm:"Column:uu_purchase"`
	UuLastWeekAvgPrice   float32 `gorm:"Column:uu_lastweek_avgprice"`
	IgxeGid              string  `gorm:"Column:igxe_gid"`
	IgxeSellPrice        float32 `gorm:"Column:igxe_sell"`
	IgxePurchasePrice    float32 `gorm:"Column:igxe_purchase"`
	IgxeLastWeekAvgPrice float32 `gorm:"Column:igxe_lastweek_avgprice"`
}

func (s *Service) FilterHotItems(itemMetas []ItemInfo) (err error) {
	hotItems := make([]HotItemInfo, 0)
	for index, itemMeta := range itemMetas {
		fmt.Println(index)
		gid := itemMeta.C5Gid
		time.Sleep(time.Millisecond * 10)
		sellCnt, sl, err := s.GetC5SellData(gid)
		if err != nil {
			fmt.Printf("GetC5SellData err , err:(%s)\n", err.Error())
		}
		time.Sleep(time.Millisecond * 50)

		purchaseCnt, pl, err := s.GetC5PurchaseData(gid)
		if err != nil {
			fmt.Printf("GetC5PurchaseData err , err:(%s)\n", err.Error())
		}

		if sellCnt >= hotSellLimit && purchaseCnt >= hotPurchaseLimit {
			row := HotItemInfo{
				ItemName:        itemMeta.ItemName,
				C5Gid:           itemMeta.C5Gid,
				UuGid:           itemMeta.UuGid,
				BuffGid:         itemMeta.BuffGid,
				IgxeGid:         itemMeta.IgxeGid,
				C5SellPrice:     str2float(sl[0].SellPrice),
				C5PurchasePrice: pl[0].PurchasePrice,
				UuSellPrice:     float32(itemMeta.UuPrice),
			}
			hotItems = append(hotItems, row)
		}
		//get all info
		//hotItemInfo, err := s.GetHotItemsCompleteInfo(itemMeta)
		// 	if err != nil {
		// 		fmt.Printf("GetHotItemsCompleteInfo err , err:(%s)\n", err.Error())
		// 	}
		// 	err = s.InsertHotItemInfo(hotItemInfo)
		// 	if err != nil {
		// 		fmt.Printf("InsertHotItemInfo err , err:(%s)\n", err.Error())
		// 	}
	}
	s.InsertHotItemList(hotItems)

	return err
}

func str2float(s string) float32 {
	num64, _ := strconv.ParseFloat(s, 32)
	num := float32(num64)
	return num
}

func (s *Service) GetHotItemsCompleteInfo(itemMeta itemMeta) ([]HotItemInfo, error) {
	return nil, nil
}

func (s *Service) InsertHotItemInfo(hotItemInfo []HotItemInfo) error {
	return nil
}
