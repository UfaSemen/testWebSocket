package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/UfaSemen/testWebSocket/pkg/data"
	"github.com/gorilla/websocket"
)

//interface for number generators
type Generator interface {
	Generate() int
}

//function starts exection of the websocket client, configPath is a path to config, gen is a generator of numbers
func StartClient(configPath string, gen Generator) {
	conf, err := readConfig(configPath)
	if err != nil {
		log.Fatal("decode of config file:", err)
	}
	limiter := time.NewTicker(time.Duration(conf.RateLimit) * time.Millisecond)
	defer limiter.Stop()
	var wg sync.WaitGroup
	for i := 0; i < conf.ConnectionNum; i++ {
		wg.Add(1)
		go connectionExecution(conf, gen, limiter, &wg)
	}
	wg.Wait()
}

func connectionExecution(conf Config, gen Generator, limiter *time.Ticker, wg *sync.WaitGroup) {
	defer wg.Done()
	u := url.URL{Scheme: "ws", Host: conf.ServerAddress + ":" + strconv.Itoa(conf.ServerPort), Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	var n1, n2 int
	var req data.Request
	var resp data.Response
	var res data.Result
	for i := 0; i < conf.MessageNumPerCon; i++ {
		<-limiter.C
		n1 = gen.Generate()
		n2 = gen.Generate()
		req = data.Request{
			Number1: n1,
			Number2: n2,
		}
		err = c.WriteJSON(req)
		if err != nil {
			log.Println("write:", err)
			return
		}

		err = c.ReadJSON(&resp)
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %v, %v", resp.Sum, resp.Product)

		res = data.Result{
			Req:  req,
			Resp: resp,
		}
		file, err := json.Marshal(res)
		if err != nil {
			log.Println("json marshal:", err)
			return
		}
		err = ioutil.WriteFile(conf.OutputPath+"/"+strconv.Itoa(req.Number1)+","+strconv.Itoa(req.Number2)+".json", file, 0644)
		if err != nil {
			log.Println("write to file:", err)
			return
		}
	}
	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
		return
	}
}
