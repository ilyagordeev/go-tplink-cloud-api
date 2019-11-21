package main

import (
	"fmt"
	"github.com/ilyagordeev/go-tplink-cloud-api"
	"os"
	"time"
)

func main() {

	tplinkcloudapi.Login(os.Getenv("TPLINK_USERNAME"), os.Getenv("TPLINK_PASSWORD"))

	fmt.Println(tplinkcloudapi.SwitchOn())

	for {
		go tplinkcloudapi.GetRealtime()
		time.Sleep(3 * time.Second)
	}
}

