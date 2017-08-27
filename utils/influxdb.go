package utils

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"

	"node/common"
)

func init() {
	//conf := common.GetSettings()
	return
}

func GetInfluxDBWriteClient() (client.Client, error) {
	conf := common.GetSettings()
	addr := conf.Getv("INFLUXDB_HOST") + ":" + conf.Getv("INFLUXDB_WRITE_PORT")
	return getInfluxDBClient(addr)
}

func GetInfluxDBREADClient() (client.Client, error) {
	conf := common.GetSettings()
	addr := conf.Getv("INFLUXDB_HOST") + ":" + conf.Getv("INFLUXDB_WRITE_PORT")
	return getInfluxDBClient(addr)
}

func getInfluxDBClient(addr string) (client.Client, error) {
	client, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: addr,
	})
	return client, err
}

func newBatchPoints() (client.BatchPoints, error) {
	conf := common.GetSettings()
	database := conf.Getv("INFLUXDB_DATABASE")
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: database,
	})
	return bp, err
}

func WriteData(cli client.Client, measurement string, tags map[string]string, fields map[string]interface{}, t time.Time) (bool, error) {
	// Create a new point batch
	batchpoints, err := newBatchPoints()
	if err != nil {
		return false, err
	}
	// Create a point and add to batch
	point, err := client.NewPoint(measurement, tags, fields, t)
	if err != nil {
		return false, err
	}
	batchpoints.AddPoint(point)
	// Write the batch
	if err := cli.Write(batchpoints); err != nil {
		return false, err
	}
	return true, nil
}
