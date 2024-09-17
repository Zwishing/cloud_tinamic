package main

import (
	vector "cloud_tinamic/kitex_gen/service/vector/vectorservice"
	"log"
)

func main() {
	svr := vector.NewServer(new(VectorServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
