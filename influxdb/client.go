package influxdb

import (
	"fmt"
	"net/url"

	client "github.com/influxdata/influxdb1-client"
)

func ClientWrite(db string, pts []client.Point) error {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		return err
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		return err
	}
	bps := client.BatchPoints{
		Points:   pts,
		Database: db,
	}

	_, err = con.Write(bps)
	return err
}
