/*
 * @Author       : lvyitao 
 * @Date         : 2024-06-05 11:25:20
 * @LastEditTime: 2026-09-10 14:37:23
 */
package v1

import (
	"fmt"
	"log"
	"path"
	"path/filepath"

	"src/utils"
	"src/utils/betteryeah"

	"github.com/gofiber/fiber/v3"
	uuid "github.com/satori/go.uuid"
)

func UploadImageWithOperate(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return utils.ErrResponse(c, 400, err.Error(), "")
	}
	basePath := utils.GetExcutePath()
	imgSavePath := filepath.Join(basePath, "data")
	fileExt := path.Ext(file.Filename)
	u4 := uuid.NewV4().String()
	uuidFileName := fmt.Sprintf("%s%s", u4, fileExt)
	imgPath := filepath.Join(imgSavePath, uuidFileName)
	err = c.SaveFile(file, imgPath)
	if err != nil {
		return utils.ErrResponse(c, 400, err.Error(), "")
	}
	// imgUrl := "https://ai-hrbp-prod.oss-cn-shanghai.aliyuncs.com/temporary/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20251226164452.jpg?x-oss-credential=LTAI5tHNPheRf61CQRGLxoN8%2F20260120%2Fcn-shanghai%2Foss%2Faliyun_v4_request&x-oss-date=20260120T075708Z&x-oss-expires=3600&x-oss-signature-version=OSS4-HMAC-SHA256&x-oss-signature=abae57a5a275867930910817945171917419a8e6de19a0aca6df503aa8845e83"
	config := utils.ConfigValue
	imgUrl := fmt.Sprintf("http://106.15.194.206:%d/get/images?uuidFileName=%s", config.Port, uuidFileName)
	url, err := betteryeah.GenImages("根据人像脸型帮我生成好打理的,适合的发型,最少生成6中,放在一张图片内,类似6宫格", imgUrl)
	if err != nil {
		log.Println(err)
		return utils.ErrResponse(c, 500, err.Error(), "")
	}
	return utils.SuccessResp(c, url)
}

func GetImg(c fiber.Ctx) error {
	uuidFileName := c.Query("uuidFileName", "")
	if uuidFileName == "" {
		return utils.ErrResponse(c, 400, "参数错误", "")
	}
	basePath := utils.GetExcutePath()
	imgSavePath := filepath.Join(basePath, "data")
	imgPath := filepath.Join(imgSavePath, uuidFileName)
	return c.SendFile(imgPath)
}
