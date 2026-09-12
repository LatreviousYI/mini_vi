/*
 * @Author       : lvyitao
 * @Date         : 2024-06-07 16:41:03
 * @LastEditTime: 2026-09-12 10:41:16
 */
package systemConfig

import (

	v1 "src/modules/systemConfig/v1"

	"github.com/gofiber/fiber/v3"
)

func Init(factoryRouter fiber.Router) {
	// 接口

	factoryRouter.Get("/ts", v1.Ts)

}
