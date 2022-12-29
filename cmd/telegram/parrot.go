package telegram

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type ParrotStruct struct {
	ChatId  int    `json:"chat_id"`
	Text    string `json:"text"`
	Sticker string `json:"sticker"`
}

func (r ResultStruct) ParrotVoice(baseUrl *string, token *string) {
	fmt.Println(r.Msg)
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	var (
		p     ParrotStruct
		route string
	)

	if r.Msg.Sticker != nil {
		p = ParrotStruct{ChatId: r.Msg.From.UsrID, Sticker: r.Msg.Sticker.Sticker}
		route = *baseUrl + "/" + *token + "/sendSticker"

	} else {
		p = ParrotStruct{ChatId: r.Msg.From.UsrID, Text: r.Msg.Txt}
		route = *baseUrl + "/" + *token + "/sendMessage"

	}
	body, _ := json.Marshal(p)
	req.SetRequestURI(route)
	req.SetBody(body)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")

	err := fasthttp.Do(req, resp)

	fmt.Printf("Message: %v has been send to %v\n", p.Text, p.ChatId)

	if err != nil {
		println("error")
	}
}
