package modules

import (
	"src/modules/systemConfig"

	"github.com/gofiber/fiber/v3"
)

func RouterInit(factoryApp fiber.Router) {
	v1 := factoryApp.Group("/v1")
	systemConfig.Init(v1)
}
