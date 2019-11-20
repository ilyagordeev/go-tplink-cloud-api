package tplinkcloudapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Request(body []byte) []byte {
	token := "?token=" + Token
	if len(Token) == 0 {
		token = ""
	}
	resp, err := http.Post("https://wap.tplinkcloud.com/"+token, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Print("Some error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody
}
