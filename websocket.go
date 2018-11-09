package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"fmt"
//	"time"
	"github.com/gorilla/websocket"
	"encoding/json"
)

// var addr = flag.String("addr", "ropsten.infura.io", "http service address")
var addr = flag.String("addr", "rinkeby.infura.io", "http service address")



func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	//done := make(chan struct{})


	// if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
	// 	return
	// }

	// w, err := c.conn.NextWriter(websocket.TextMessage)
	// if err != nil {
	// 	return
	// }
	writeerr := c.WriteMessage(websocket.TextMessage,[]byte(`{"jsonrpc":"2.0","method":"eth_newFilter","params":[{"address":"0x4e83362442b8d1bec281594cea3050c8eb01311c"}],"id":73}`))

	if writeerr != nil {
		fmt.Println ("faield to write something")
	}


	msgtype,msgtext,readerr := c.ReadMessage()
	if readerr != nil{
		fmt.Println("hello read faield ")
		// break
	}
	var EventInfo EthNewFilterRespose
	EventInfoerr := json.Unmarshal(msgtext,&EventInfo)

	if EventInfoerr != nil{
		fmt.Println("hello faileld ")
	}
	fmt.Println("msg type is ",msgtype,"msg text is ",msgtext)
	fmt.Printf("hell th result is  %+v\n",EventInfo)
	

	var ethnewfilter  EthGetFilterReq
	ethnewfilter.Jsonrpc = "2.0"
	ethnewfilter.Method = "eth_getFilterLogs"
	ethnewfilter.Params = append(ethnewfilter.Params,EventInfo.Result)
	ethnewfilter.Id = EventInfo.Id

	hello,_ := json.Marshal(ethnewfilter)

	getchangeerr := c.WriteMessage(websocket.TextMessage,hello)

	if getchangeerr != nil {
		fmt.Println("hello world ")
	}

	_,msgtextres,readerrres := c.ReadMessage()
	if readerrres != nil{
		fmt.Println("hello read faield ")
		// break
	}
	var EventInfores EthGetFilterResponse
	EventInfoerrres := json.Unmarshal(msgtextres,&EventInfores)
	if EventInfoerrres != nil{
		fmt.Println("failed again")
	}

	fmt.Printf("the result is ....%+v\n",EventInfores)

	
	// for{
	// 	msgtype,msgtext,readerr := c.ReadMessage()
	// 	if readerr != nil{
	// 		fmt.Println("hello read faield ")
	// 		break
	// 	}
	// 	var EventInfo EthNewFilterRespose
	// 	EventInfoerr := json.Unmarshal(msgtext,&EventInfo)

	// 	if EventInfoerr != nil{
	// 		fmt.Println("hello faileld ")
	// 	}
	// 	fmt.Println("msg type is ",msgtype,"msg text is ",msgtext)
	// 	break
	// }
	// go func() {
	// 	defer close(done)
	// 	for {
	// 		_, message, err := c.ReadMessage()
	// 		if err != nil {
	// 			log.Println("read:", err)
	// 			return
	// 		}
	// 		log.Printf("recv: %s", message)
	// 	}
	// }()

	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case t := <-ticker.C:
	// 		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
	// 		if err != nil {
	// 			log.Println("write:", err)
	// 			return
	// 		}
	// 	case <-interrupt:
	// 		log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err11 := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err11 != nil {
				log.Println("write close:", err11)
				return
			}
			// select {
			// case <-done:
			// case <-time.After(time.Second):
			// }
			// returns
		// }
	// }
}





//1.eth_newfilter

//{"jsonrpc":"2.0","method":"eth_newFilter","params":[{"topics": ["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"]}],"id":73}
type EthNewFilterReq struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []AddressType 	`json:"params"`
	Id      int    `json:"id"`
}

type EthGetFilterReq struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []string 	`json:"params"`
	Id      int    `json:"id"`
}

type EthGetFilterResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  interface{} `json:"result"`
}



type AddressType struct {
	Address string `json:"address"`
}

type ParamsInfo struct {
	Address string `json:"address"`
	FromBlock string `json:"fromBlock"`
	ToBlock string `json:"toBlock"`
	Topics  string `json:"topics"`
}

// {
//     "jsonrpc":"2.0",
//     "id":73,
//     "result":"0x7db09f66a25e197d995d3895278b731"
// }

type EthNewFilterRespose struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}