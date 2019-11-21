package tplinkcloudapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

func GetRealtime() {
	responseData := PassthroughRequest("{\"emeter\":{\"get_realtime\":null}}")
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(responseData), &dat); err != nil {
		fmt.Println("Bad json:", err)
	}
	getRealtime := dat["emeter"].(map[string]interface{})["get_realtime"].(map[string]interface{})

	power := math.Round(getRealtime["power_mw"].(float64) / 1000)
	voltage := math.Round(getRealtime["voltage_mv"].(float64) / 1000)
	amperage := getRealtime["current_ma"].(float64)
	totalWh := getRealtime["total_wh"].(float64)

	fmt.Println("Power:", power, " | Voltage: ", voltage, " | Amperage: ", amperage, " | Total w/h: ", totalWh)
}

func SwitchOn() error {
	return checkError(PassthroughRequest("{\"system\":{\"set_relay_state\":{\"state\":1}}}"))
}

func SwitchOf() error {
	return checkError(PassthroughRequest("{\"system\":{\"set_relay_state\":{\"state\":0}}}"))
}

func checkError(response string) error {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(response), &dat); err != nil {
		return err
	}
	errCode := dat["system"].(map[string]interface{})["set_relay_state"].(map[string]interface{})["err_code"].(float64)
	if errCode != 0 {
		return errors.New("Failed")
	}
	return nil
}