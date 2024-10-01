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

func PutErrorMetric(namesapce string, service string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := cloudwatch.NewFromConfig(cfg)

	input := &cloudwatch.PutMetricDataInput{
		Namespace: utils.Pointer(namesapce),
		MetricData: []types.MetricDatum{
			{
				MetricName: utils.Pointer("Error"),
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
