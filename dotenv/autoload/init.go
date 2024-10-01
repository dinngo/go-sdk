package autoload

import (
	"github.com/dinngo/go-sdk/dotenv"
)

func init() {
	if err := dotenv.LoadByStage(); err != nil {
		panic(err)
	}
}
