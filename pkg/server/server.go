package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/UfaSemen/testWebSocket/pkg/data"
	"github.com/gorilla/websocket"
)

//interface for final numbers calculations
type Calculator interface {
	Sum(n1, n2 int) int
	Product(n1, n2 int) int
}

type handlerContext struct {
	calc Calculator
}

func (hctx handlerContext) handler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		var req data.Request
		err = c.ReadJSON(&req)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, 1000) {
				log.Println("read:", err)
			}
			break
		}
		log.Printf("recv: %v, %v", req.Number1, req.Number2)
		resp := data.Response{
			Sum:     hctx.calc.Sum(req.Number1, req.Number2),
			Product: hctx.calc.Product(req.Number1, req.Number2),
		}
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

//function starts exection of the websocket server, configPath is a path to config, calc is a final numbers calculator
func StartServer(configPath string, calc Calculator) {
	conf, err := readConfig(configPath)
	if err != nil {
		log.Fatal("decode of config file:", err)
	}
	hctx := handlerContext{calc: calc}
	http.HandleFunc("/", hctx.handler)

	err = http.ListenAndServe("localhost:"+strconv.Itoa(conf.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
