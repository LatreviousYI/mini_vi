/*
 * @Author       : lvyitao
 * @Date         : 2024-06-05 11:20:39
 * @LastEditTime: 2026-09-10 13:13:31
 */
package factory

import (
	"context"
	"path/filepath"
	"sync"

	"src/modules"
	"src/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
)

var FiberApp *fiber.App

var GlobalWG sync.WaitGroup

var GlobalCtx context.Context

var GlobalCancel context.CancelFunc

func index(c fiber.Ctx) error {
	basePath := utils.GetExcutePath()
	return c.SendFile(filepath.Join(basePath, "www", "index.html"))
}

func Init() {
	// 加载配置
	utils.GetConfig()
	FiberApp = fiber.New()
	FiberApp.Use(logger.New(logger.Config{
		// 常用占位符：${time} ${ip} ${method} ${path} ${status} ${latency} ${ua} ${referer}
		Format:     "[${time}] ${ip} ${method} ${path} ${status} ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
	}))
	FiberApp.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))
	GlobalCtx, GlobalCancel = context.WithCancel(context.Background())
	FiberApp.Get("/healthy", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})
	basepath := utils.GetExcutePath()
	staticPath := filepath.Join(basepath, "www", "static")
	FiberApp.Use("/static", static.New(staticPath))
	FiberApp.Get("/", index)
	api := FiberApp.Group("api")
	modules.RouterInit(api)
}
