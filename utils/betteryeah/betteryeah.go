package betteryeah

import (
	"log"
	"src/utils"

	"resty.dev/v3"
)

type GenImagesBodyInputs struct {
	Message string `json:"message"`
	Img     string `json:"img"`
}

type GenImagesBody struct {
	Inputs GenImagesBodyInputs `json:"inputs"`
}

type GenImagesResult struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		TaskId    string `json:"task_id"`
		Status    string `json:"status"`
		RunResult any    `json:"run_result"`
		Message   any    `json:"message"`
	} `json:"data"`
	RequestID string `json:"request_id"`
	NowTime   int    `json:"now_time"`
}

func GenImages(messgae, url string) (string, error) {
	img := url
	headers := map[string]string{"Content-Type": "application/json", "Access-key": utils.ConfigValue.AccessKey, "Workspace-Id": utils.ConfigValue.WorkspaceId}
	body := GenImagesBody{Inputs: GenImagesBodyInputs{Message: messgae, Img: img}}
	client := resty.New() // 创建一个restry客户端
	var genRe GenImagesResult
	res, err := client.R().SetHeaders(headers).SetResult(&genRe).SetBody(body).Post("https://ai-api.betteryeah.com/v1/public_api/rest_api/2af03e6c48b24234b0783640af4cd093/execute_flow")
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(genRe)
	if res.StatusCode() != 200 {
		log.Println(err)
		return "", err
	}
	if genRe.Code != 200 {
		log.Println(genRe)
		return "", err
	}
	if genRe.Data.Status != "SUCCEEDED" {
		log.Println(genRe.Data.Message)
		return "", err
	}
	message := genRe.Data.RunResult.(map[string]any)["message"].(string)
	return message, nil
}
