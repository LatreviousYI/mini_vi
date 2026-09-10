/*
 * @Author       : lvyitao 
 * @Date         : 2024-06-07 16:41:03
 * @LastEditTime: 2026-09-10 13:35:08
 */
package systemConfig

import (
	// "log"

	// "src/common"
	// "src/utils"

	// "github.com/go-co-op/gocron/v2"
	// v1 "src/modules/systemConfig/v1"

	"github.com/gofiber/fiber/v3"
)

func Init(factoryRouter fiber.Router) {
	// 接口

	// factoryRouter.Post("/upload/images", v1.UploadImageWithOperate)
	// factoryRouter.Get("/get/images", v1.GetImg)

	// 定时任务
	// _, err := common.GoCronTab.NewJob(
	// 	gocron.CronJob("0 0 9 ? * MON", true),
	// 	gocron.NewTask(utils.TimeCleanlog),
	// )
	// if err != nil {
	// 	log.Println(err)
	// }
}
