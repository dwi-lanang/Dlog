package dllog

import (
	"encoding/json"
	"fmt"
	"time"

	goWS "github.com/sacOO7/gowebsocket"
)

type ParamBody struct {
	State   string      `json:"state"`
	Data    interface{} `json:"data"`
	Options interface{} `json:"options"`
	At      string
}
type Socket struct {
	*goWS.Socket
	Options
}
type Options struct {
	Mode, Src string
}

func Dlog(url, options Options, callback func(string)) (socket Socket) {
	s := goWS.New(fmt.Sprint("wss://", url))

	s.OnConnectError = func(err error, socket goWS.Socket) {
		callback("Error :" + err.Error())
	}

	s.OnConnected = func(socket goWS.Socket) {
		callback("Connected")
	}

	s.OnTextMessage = func(message string, socket goWS.Socket) {
		callback(message)
	}

	s.OnPingReceived = func(data string, socket goWS.Socket) {
		callback(data)
	}

	s.OnPongReceived = func(data string, socket goWS.Socket) {
		callback(data)
	}

	s.OnDisconnected = func(err error, socket goWS.Socket) {
		callback("Disconnected")
		socket.Connect()
		return
	}

	s.Connect()
	return Socket{
		&s,
		options,
	}
}

func (socket Socket) Send(state string, data interface{}) {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02-15-04-05")
	b := ParamBody{
		State:   state,
		Data:    data,
		Options: socket.Options,
		At:      currentDate,
	}
	body, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	socket.SendText(string(body))
}
