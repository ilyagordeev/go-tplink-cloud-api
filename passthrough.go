package tplinkcloudapi

import "encoding/json"

type Passthrough struct {
	Method       string `json:"method"`
	ParamsRequst `json:"params"`
}

type ParamsRequst struct {
	DeviceId    string `json:"deviceId"`
	RequestData string `json:"requestData"`
}

func PassthroughRequest(requestData string) string {
	passthrough := &Passthrough{
		Method: "passthrough",
		ParamsRequst: ParamsRequst{
			DeviceId:    DeviceId,
			RequestData: requestData,
		},
	}
	body, _ := json.Marshal(passthrough)
	respBody := Request(body)
	var dat map[string]interface{}
	if err := json.Unmarshal(respBody, &dat); err != nil {
		panic(err)
	}
	return dat["result"].(map[string]interface{})["responseData"].(string)
}
