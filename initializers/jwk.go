package initializers

import "github.com/vaults-dev/vaults-backend/utils"

func GenerateJwk() {
	err := utils.GenerateJwk()
	if err != nil {
		panic(err.Error())
	}
}
