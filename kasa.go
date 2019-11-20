package tplinkcloudapi

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	Login(os.Getenv("TPLINK_USERNAME"), os.Getenv("TPLINK_PASSWORD"))

	response := PassthroughRequest("{\"system\":{\"set_relay_state\":{\"state\":1}}}")
	fmt.Printf(response)

	for {
		go getRealtime()
		time.Sleep(3 * time.Second)
	}

}

func getRealtime() {
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
