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
		if _, err := os.Stat(envFile); os.IsNotExist(err) {
			return fmt.Errorf("%s not found", envFile)
		}

		secretsPassword, err := GetSecretsPassword()
		if err != nil {
			return fmt.Errorf("failed to get secrets password")
		}

		if err := crypto.FileDecrypter(encryptedEnvFile, envFile, []byte(secretsPassword)); err != nil {
			return fmt.Errorf("failed to decrypt %s", encryptedEnvFile)
		}
	}

	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("failed to load %s", envFile)
	}

	return nil
}
