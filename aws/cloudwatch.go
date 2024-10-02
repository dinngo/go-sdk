package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/dinngo/go-sdk/utils"
)

func putMetric(name, namesapce, service string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := cloudwatch.NewFromConfig(cfg)

	input := &cloudwatch.PutMetricDataInput{
		Namespace: utils.Pointer(namesapce),
		MetricData: []types.MetricDatum{
			{
				MetricName: utils.Pointer(name),
				Timestamp:  utils.Pointer(time.Now().UTC()),
				Unit:       types.StandardUnitCount,
				Value:      aws.Float64(1),
				Dimensions: []types.Dimension{
					{
						Name:  utils.Pointer("Service"),
						Value: utils.Pointer(service),
					},
				},
			},
		},
	}
	if _, err := client.PutMetricData(context.TODO(), input); err != nil {
		return err
	}

	return nil
}

func PutErrorMetric(namesapce, service string) error {
	return putMetric("Error", namesapce, service)
}

func PutHealthyMetric(namesapce, service string) error {
	return putMetric("Healthy", namesapce, service)
}

func MonitorHealthy(namesapce, service string) {
	go func() {
		for {
			PutHealthyMetric(namesapce, service)
			time.Sleep(time.Minute)
		}
	}()
}
