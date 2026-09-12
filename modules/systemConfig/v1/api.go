/*
 * @Author       : lvyitao 
 * @Date         : 2024-06-05 11:25:20
 * @LastEditTime: 2026-09-12 10:40:59
 */
package v1

import (

	"src/utils"

	"github.com/gofiber/fiber/v3"
)


func Ts(c fiber.Ctx) error {
	return utils.SuccessResp(c, "ok")
}

