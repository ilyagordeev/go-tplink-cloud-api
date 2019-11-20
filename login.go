package tplinkcloudapi

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
)

var Token, DeviceId string
var Deices map[string]string

type LoginRequest struct {
	Method string `json:"method"`
	Params `json:"params"`
}

type Params struct {
	AppType       string `json:"appType"`
	CloudUserName string `json:"cloudUserName"`
	CloudPassword string `json:"cloudPassword"`
	TerminalUUID  string `json:"terminalUUID"`
}

type DeviceListRequest struct {
	Method       string `json:"method"`
	ParamsDevice `json:"params"`
}

type ParamsDevice struct {
	AppName      string `json:"appName"`
	TerminalUUID string `json:"termID"`
	AppVer       string `json:"appVer"`
	Ospf         string `json:"ospf"`
	NetType      string `json:"netType"`
	Locale       string `json:"locale"`
	Token        string `json:"token"`
}

func Login(cloudUserName string, cloudPassword string) {
	u, _ := uuid.NewV4()
	loginJSON := &LoginRequest{
		Method: "login",
		Params: Params{
			AppType:       "Kasa_Android",
			CloudUserName: cloudUserName,
			CloudPassword: cloudPassword,
			TerminalUUID:  u.String(),
		},
	}
	body, _ := json.Marshal(loginJSON)
	respBody := Request(body)
	var dat map[string]interface{}
	if err := json.Unmarshal(respBody, &dat); err != nil {
		panic(err)
	}
	Token = dat["result"].(map[string]interface{})["token"].(string)

	deviceList := &DeviceListRequest{
		Method: "getDeviceList",
		ParamsDevice: ParamsDevice{
			AppName:      "Kasa_Android",
			TerminalUUID: u.String(),
			AppVer:       "1.4.4.607",
			Ospf:         "Android+6.0.1",
			NetType:      "wifi",
			Locale:       "es_ES",
			Token:        Token,
		},
	}
	body, _ = json.Marshal(deviceList)
	respBody = Request(body)

	if err := json.Unmarshal(respBody, &dat); err != nil {
		panic(err)
	}

	respDevMap := dat["result"].(map[string]interface{})["deviceList"].([]interface{})

	if len(respDevMap) == 1 {
		alias, _ := respDevMap[0].(map[string]interface{})["alias"].(string)
		deviceModel, _ := respDevMap[0].(map[string]interface{})["deviceModel"].(string)
		fmt.Println("You have 1 device:", alias, deviceModel)
	} else {
		for idx := range respDevMap {
			if alias, ok := respDevMap[idx].(map[string]interface{})["alias"].(string); ok {
				deviceId, _ := respDevMap[idx].(map[string]interface{})["deviceId"].(string)
				Deices[alias] = deviceId
			}
		}
	}
	DeviceId = respDevMap[0].(map[string]interface{})["deviceId"].(string)

}
