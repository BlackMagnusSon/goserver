package telegram

import (
	"fmt"
	"time"
)

func TelegramParrot(url *string, token *string) error {
	r := requestStruct{Offset: -1}
	resultArr := r.GetUpdates(url, token)

	for _, val := range resultArr.Result {
		r.Offset = val.Id
	}

	for {
		resultArr = r.GetUpdates(url, token)
		for _, val := range resultArr.Result {
			// fmt.Println(i)
			if val.Id <= r.Offset || val.Msg == nil {
				// fmt.Printf("%v - %v\n", val.Id, r.Offset)
				fmt.Println("Have no Updates")
				continue
			}
			if r.Offset < int(val.Id) {
				r.Offset = int(val.Id)
			}

			go val.ParrotVoice(url, token)
		}
		time.Sleep(1 * time.Second)
	}
}
