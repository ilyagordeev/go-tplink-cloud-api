# Golang bindings for the TP-Link smart devices API
Init version

This build works only with hs100/110

##Installation
 `go get github.com/ilyagordeev/go-tplink-cloud-api`
 
 ##Example

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

	    tplinkcloudapi.GetRealtime()

    }
