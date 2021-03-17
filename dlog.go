package dllog

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	goWS "github.com/sacOO7/gowebsocket"
)

type ParamBody struct {
	State  string      `json:"state"`
	Data   interface{} `json:"data"`
	Config Config      `json:"config"`
	At     string      `json:"at"`
}
type IO struct {
	*goWS.Socket
	Config
}
type Config struct {
	Mode    string `json:"mode"`
	Channel string `json:"channel"`
}

func recovery() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}
func Init(url string, config Config, callback func(string)) (socket IO) {
	defer recovery()
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
	return IO{
		&s,
		config,
	}
}

func (socket IO) Send(state string, data interface{}) {
	defer recovery()
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02-15-04-05")
	b := ParamBody{
		State:  state,
		Data:   data,
		Config: socket.Config,
		At:     currentDate,
	}
	body, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	socket.SendText(string(body))
}
