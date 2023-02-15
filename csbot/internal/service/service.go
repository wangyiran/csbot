package service

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	db       *gorm.DB
	linkOnce sync.Once
}

func New() (s *Service) {
	s = &Service{}
	s.linkDb()
	return s
}

func (s *Service) linkDb() {
	s.linkOnce.Do(func() {
		dsn := "csbot:wyr1998@tcp(101.42.19.193:3306)/csbot?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		s.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		fmt.Println("connect db ok")
	})
}
