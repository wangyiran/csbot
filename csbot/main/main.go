package main

import (
	"csbot/csbot/internal/service"
)

func main() {
	//fmt.Println("11")
	s := service.New()
	//s.GetAllGoodsInfoFromPlatform()
	s.Analyse()

}
