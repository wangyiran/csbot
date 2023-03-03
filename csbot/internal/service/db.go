package service

import (
	"fmt"
)

type ItemInfo struct {
	ID        int64   `gorm:"PRIMARY_KEY;Column:id"`
	ItemName  string  `gorm:"Column:item_name"`
	C5Gid     string  `gorm:"Column:c5_gid"`
	BuffGid   string  `gorm:"Column:buff_gid"`
	UuGid     string  `gorm:"Column:uu_gid"`
	IgxeGid   string  `gorm:"Column:igxe_gid"`
	C5Price   float64 `gorm:"Column:c5_price"`
	BuffPrice float64 `gorm:"Column:buff_price"`
	UuPrice   float64 `gorm:"Column:uu_price"`
	IgxePrice float64 `gorm:"Column:igxe_price"`
	IsHot     int32   `gorm:"Colume:is_hot"`
}

func (s *Service) Test() {
	fmt.Println("!")
	tx := s.db.Begin()
	row := &ItemInfo{}
	err := tx.Table("items_list").Where("1=1").Find(row).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)
}

func (s *Service) InsertItemList(infos []ItemInfo) error {
	fmt.Printf("ready to insert item info , count:(%d)\n", len(infos))
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	if err := tx.Table("items_list").Create(infos).Error; err != nil {
		fmt.Printf("insert item info err, err:(%s)\n", err.Error())
		return err
	}
	fmt.Printf("insert item info success\n")
	tx.Commit()
	return nil
}

func (s *Service) InsertHotItemList(infos []HotItemInfo) error {
	fmt.Printf("ready to insert hot item info , count:(%d)\n", len(infos))
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	if err := tx.Table("hot_items").Create(infos).Error; err != nil {
		fmt.Printf("insert hot item info err, err:(%s)\n", err.Error())
		return err
	}
	fmt.Printf("insert hot item info success\n")
	tx.Commit()
	return nil
}

func (s *Service) GetCompleteItemListFromDB() (infos []ItemInfo) {
	var rows []ItemInfo
	err := s.db.Table("items_list").Where("c5_price > 20 and c5_price < 2000").Find(&rows).Error
	if err != nil || len(rows) == 0 {
		fmt.Printf("get items list err:(%s)\n", err.Error())
		return rows
	}
	return rows
}

func (s *Service) GetHotItemListFromDB() (infos []HotItemInfo) {
	var rows []HotItemInfo
	err := s.db.Table("hot_items").Where("c5_sell > 20 and c5_sell < 3000").Find(&rows).Error
	if err != nil || len(rows) == 0 {
		fmt.Printf("get items list err:(%s)\n", err.Error())
		return rows
	}
	return rows
}

func (s *Service) DeleteOldHotItemsList() error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Rollback()

	if err := tx.Exec("DELETE FROM hot_items").Error; err != nil {
		fmt.Printf("DeleteOldHotItemsList err, err:(%s)\n", err.Error())
		return err
	}
	fmt.Printf("DeleteOldHotItemsList success\n")
	tx.Commit()
	return nil
}
