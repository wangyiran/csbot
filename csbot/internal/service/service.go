package service

import (
	"fmt"
	"sync"

	"github.com/robfig/cron"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	db       *gorm.DB
	linkOnce sync.Once
	cron     *cron.Cron
}

func New() (s *Service, cf func(), err error) {
	s = &Service{
		cron: cron.New(),
	}
	err = s.linkDb()
	if err != nil {
		return s, nil, err
	}
	err = s.addCronFunc()
	if err != nil {
		return s, nil, err
	}
	s.cron.Start()
	cf = s.close
	return s, cf, err
}

func (s *Service) linkDb() (err error) {
	s.linkOnce.Do(func() {
		dsn := "csbot:wyr1998@tcp(101.42.19.193:3306)/csbot?charset=utf8mb4&parseTime=True&loc=Local"
		s.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		fmt.Println("connect db ok")
	})
	return
}

func (s *Service) addCronFunc() (err error) {
	//每周 拉全量数据,update一次hotItem
	//每天 拉两次hotItem数据，并分析给出结论
	err = s.cron.AddFunc("0 0 9,21 * * ?", s.DailyAnalyseHotItems)
	if err != nil {
		return err
	}
	//err = s.cron.ADdFunc("0 0 22 ? * 1 *", s.weeklyFilterCompleteItems)
	return err
}

func (s *Service) close() {
	s.cron.Stop()
}
