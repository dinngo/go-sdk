package dotenv

import (
	"fmt"
	"os"

	"github.com/dinngo/go-sdk/crypto"
	"github.com/dinngo/go-sdk/utils"
	"github.com/joho/godotenv"
)

func LoadByStage() error {
	envFile := ".env"
	if stage := utils.GetNullableEnv("STAGE"); stage != nil {
		envFile = fmt.Sprintf("%s.%s", envFile, *stage)
	}

	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		encryptedEnvFile := envFile + ".enc"
		if _, err := os.Stat(encryptedEnvFile); os.IsNotExist(err) {
			return fmt.Errorf("%s not found", encryptedEnvFile)
		}

		secretsPassword, err := GetSecretsPassword()
		if err != nil {
			return fmt.Errorf("failed to get secrets password: %s", err.Error())
		}

		if err := crypto.FileDecrypter(encryptedEnvFile, envFile, []byte(secretsPassword)); err != nil {
			return fmt.Errorf("failed to decrypt %s (%s): %s", encryptedEnvFile, secretsPassword, err.Error())
		}
	}

	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("failed to load %s: %s", envFile, err.Error())
	}

	return nil
}
