package service

import "strconv"

func (s Service) GetAllGoodsInfoFromPlatform() {
	buffMP := make(map[string]BuffItem, 0)
	uuMP := make(map[string]UUItem, 0)

	c5Items := s.GetAllC5Goods()
	uuItems := s.GetAllUUGoods()
	buffItems := s.GetAllBuffGoods()

	for i := 0; i < len(uuItems); i++ {
		uuMP[uuItems[i].Gname] = *uuItems[i]
	}

	for i := 0; i < len(buffItems); i++ {
		buffMP[buffItems[i].Gname] = buffItems[i]
	}

	itemsInfo := make([]ItemInfo, 0)
	for index := 0; index < len(c5Items); index++ {
		_, okuu := uuMP[c5Items[index].Gname]
		_, okbuff := buffMP[c5Items[index].Gname]
		if !okbuff && !okuu {
			continue
		}
		c5Price, _ := strconv.ParseFloat(c5Items[index].CnyPrice, 64)
		uuPrice, _ := strconv.ParseFloat(uuMP[c5Items[index].Gname].CnyPrice, 64)
		buffPrice, _ := strconv.ParseFloat(buffMP[c5Items[index].Gname].CnyPrice, 64)

		itemsInfo = append(itemsInfo, ItemInfo{
			ItemName:  c5Items[index].Gname,
			C5Gid:     c5Items[index].Gid,
			UuGid:     strconv.Itoa(uuMP[c5Items[index].Gname].Gid),
			BuffGid:   strconv.Itoa(buffMP[c5Items[index].Gname].Gid),
			C5Price:   c5Price,
			UuPrice:   uuPrice,
			BuffPrice: buffPrice,
		})
	}

	_ = s.InsertItemList(itemsInfo)
	return
}
