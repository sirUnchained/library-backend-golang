package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"message"`
	Ok     bool        `json:"ok"`
}

func Send(isOk bool, msg interface{}, status int, w http.ResponseWriter) {
	sendingResponse := &response{}

	sendingResponse.Ok = isOk
	sendingResponse.Msg = msg
	sendingResponse.Status = status

	sendingJson, err := json.Marshal(&sendingResponse)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status)
	w.Write(sendingJson)
}
