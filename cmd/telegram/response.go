package telegram

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type telegramResponse struct {
	Status bool           `json:"ok"`
	Result []ResultStruct `json:"result"`
}

type ResultStruct struct {
	Id  int        `json:"update_id"`
	Msg *msgStruct `json:"message"`
}

type msgStruct struct {
	Txt     string         `json:"text"`
	Sticker *stickerStruct `json:"sticker"`
	From    *fromStruct
}

type stickerStruct struct {
	Sticker string `json:"file_id"`
}

type fromStruct struct {
	UsrID int    `json:"id"`
	Usr   string `json:"username"`
}

type requestStruct struct {
	Offset int `json:"offset"`
}

func (r *requestStruct) GetUpdates(baseUrl *string, token *string) *telegramResponse {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	var response telegramResponse

	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	route := *baseUrl + "/" + *token + "/getUpdates"

	req.SetRequestURI(route)

	err := fasthttp.Do(req, resp)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		panic(err)
	}
	return &response
}
