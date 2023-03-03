package service

import (
	"csbot/csbot/utils"
	"fmt"
	"os"
	"strings"
	"time"
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

func (s *Service) DailyAnalyseHotItems() {
	HotItemInfos := s.GetHotItemListFromDB()
	HotItemInfos, err := s.UpdateHotItemsInfo(HotItemInfos)
	if err != nil {
		fmt.Printf("dailyAnalyseHotItems UpdateHotItemsInfo err , err:(%s)\n", err.Error)
		return
	}
	filterHot(HotItemInfos)
	filterC5HotPurchase(HotItemInfos)
	filterUuHotPurchase(HotItemInfos)
}

func filterHot(infos []HotItemInfo) {
	path := "analyse_hot.txt"
	_exist, _err := HasDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
		_ = os.Remove(path)
	}

	txt, _ := os.OpenFile(`analyse_hot.txt`, os.O_APPEND|os.O_CREATE, 0666)
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

		minPrice := utils.MinPrice(float64(infos[i].C5SellPrice), float64(infos[i].UuSellPrice))
		maxPrice := utils.MaxPrice(float64(infos[i].C5SellPrice), float64(infos[i].UuSellPrice))
		// minPrice := utils.MinPrice(infos[i].C5Price, utils.MinPrice(infos[i].UuPrice, infos[i].BuffPrice))
		// maxPrice := utils.MaxPrice(infos[i].C5Price, utils.MaxPrice(infos[i].UuPrice, infos[i].BuffPrice))
		cha := (maxPrice - minPrice) / minPrice
		if cha < 0.06 || cha > 0.2 {
			continue
		}
		str := fmt.Sprintf("%s:%v:%v:%v\n", strings.Replace(infos[i].ItemName, "\n", "", -1), float32(cha), infos[i].C5SellPrice, infos[i].UuSellPrice)
		_, _ = txt.Write([]byte(str))

	}
	return
}

func filterC5HotPurchase(infos []HotItemInfo) {
	path := "analyse_c5hot_purchase.txt"
	_exist, _err := HasDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
		_ = os.Remove(path)
	}

	txt, _ := os.OpenFile(`analyse_c5hot_purchase.txt`, os.O_APPEND|os.O_CREATE, 0666)
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

		// minPrice := utils.MinPrice(infos[i].C5Price, utils.MinPrice(infos[i].UuPrice, infos[i].BuffPrice))
		// maxPrice := utils.MaxPrice(infos[i].C5Price, utils.MaxPrice(infos[i].UuPrice, infos[i].BuffPrice))
		cha := (float64(infos[i].C5SellPrice) - float64(infos[i].C5PurchasePrice)) / float64(infos[i].C5PurchasePrice)
		if cha < 0.08 || cha > 0.2 {
			continue
		}
		str := fmt.Sprintf("%s:%v:%v:%v\n", strings.Replace(infos[i].ItemName, "\n", "", -1), float32(cha), infos[i].C5SellPrice, infos[i].C5PurchasePrice)
		_, _ = txt.Write([]byte(str))
	}
	return
}

func filterUuHotPurchase(infos []HotItemInfo) {
	path := "analyse_uuhot_purchase.txt"
	_exist, _err := HasDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
		_ = os.Remove(path)
	}

	txt, _ := os.OpenFile(`analyse_uuhot_purchase.txt`, os.O_APPEND|os.O_CREATE, 0666)
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

		// minPrice := utils.MinPrice(infos[i].C5Price, utils.MinPrice(infos[i].UuPrice, infos[i].BuffPrice))
		// maxPrice := utils.MaxPrice(infos[i].C5Price, utils.MaxPrice(infos[i].UuPrice, infos[i].BuffPrice))
		cha := (float64(infos[i].UuSellPrice) - float64(infos[i].UuPurchasePrice)) / float64(infos[i].UuPurchasePrice)
		if cha < 0.08 || cha > 0.2 {
			continue
		}
		str := fmt.Sprintf("%s:%v:%v:%v\n", strings.Replace(infos[i].ItemName, "\n", "", -1), float32(cha), infos[i].UuSellPrice, infos[i].UuPurchasePrice)
		_, _ = txt.Write([]byte(str))
	}
	return
}

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func (s *Service) UpdateHotItemsInfo(infos []HotItemInfo) (res []HotItemInfo, err error) {
	res = make([]HotItemInfo, 0)
	for index, info := range infos {
		c5SellPirce, _ := s.GetC5LowestSellPrice(info.C5Gid)
		c5PurchasePirce, _ := s.GetC5LowestPurchasePrice(info.C5Gid)
		uuSellPirce, _ := s.GetUuLowestSellPrice(info.UuGid)
		uuPurchasePirce, _ := s.GetUuLowestPurchasePrice(info.UuGid)
		time.Sleep(time.Millisecond * 100)
		row := HotItemInfo{
			ItemName:        info.ItemName,
			C5Gid:           info.C5Gid,
			UuGid:           info.UuGid,
			BuffGid:         info.BuffGid,
			IgxeGid:         info.IgxeGid,
			C5SellPrice:     c5SellPirce,
			C5PurchasePrice: c5PurchasePirce,
			UuSellPrice:     uuSellPirce,
			UuPurchasePrice: uuPurchasePirce,
		}
		res = append(res, row)
		fmt.Println(index)
	}
	err = s.DeleteOldHotItemsList()
	err = s.InsertHotItemList(res)
	return res, err
}
