package dotenv

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/dinngo/go-sdk/utils"
)

func GetSecretsPassword() (string, error) {
	if secretsPassword := utils.GetNullableEnv("SECRETS_PASSWORD"); secretsPassword != nil {
		return *secretsPassword, nil
	}

	name, region := utils.GetNullableEnv("SECRETS_PASSWORD_PS_NAME"), utils.GetNullableEnv("SECRETS_PASSWORD_PS_REGION")
	if name == nil || region == nil {
		return "", errors.New("secrets password not found")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}
	cfg.Region = *region
	client := ssm.NewFromConfig(cfg)

	input := &ssm.GetParameterInput{
		Name:           name,
		WithDecryption: aws.Bool(true),
	}
	resp, err := client.GetParameter(context.TODO(), input)
	if err != nil {
		return "", err
	}

	return *resp.Parameter.Value, nil
}
