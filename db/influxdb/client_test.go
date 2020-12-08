package influxdb

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	client "github.com/influxdata/influxdb1-client"
)

func TestClientWrite(t *testing.T) {
	pts := make([]client.Point, 1)
	pts[0] = client.Point{
		Measurement: "payment_records",
		Tags: map[string]string{
			"type":    "cc",
			"channel": "CTwo",
		},
		Fields: map[string]interface{}{
			"success":  1,
			"rsp_time": 12.3434,
		},
	}

	if err := ClientWrite("mydb", pts); err != nil {
		t.Errorf("client write test err : %s", err.Error())
	}
}

func TestClientWrite2(t *testing.T) {
	for {
		WriteOneHundred()
		time.Sleep(time.Second * 2)
	}

}

func WriteOneHundred() {
	payTypes := []string{"cc", "fpx"}
	channels := map[string][]string{
		"cc":  {"COne", "CTwo", "CThree"},
		"fpx": {"IBank", "ISB", "SOC"},
	}

	pts := make([]client.Point, 100)
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		payType := payTypes[rand.Intn(2)]
		channel := channels[payType][rand.Intn(3)]
		pts[i] = client.Point{
			Measurement: "payment_records",
			Tags: map[string]string{
				"type":    payType,
				"channel": channel,
			},
			Fields: map[string]interface{}{
				"success":  rand.Intn(2),
				"rsp_time": rand.Float32() * 20,
			},
		}
	}

	if err := ClientWrite("mydb", pts); err != nil {
		fmt.Printf("client write test err : %s", err.Error())
	}
}
