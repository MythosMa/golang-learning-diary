package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var upgrader = websocket.FastHTTPUpgrader{
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool { return true },
}

// websocket
func webSocketHandler(ctx *fasthttp.RequestCtx) {
	upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		defer conn.Close()
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("读取消息错误:", err)
				return
			}
			responseMessage := "服务器收到了你的消息：" + string(message)
			log.Printf("收到消息: %s", message)

			err = conn.WriteMessage(messageType, []byte(responseMessage))
			if err != nil {
				log.Println("写入消息错误:", err)
				return
			}
		}
	})
}

// http
func httpHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprint(ctx, "Hello, World!")
	case "/ws":
		webSocketHandler(ctx)
	default:
		ctx.Error("404 Not found", fasthttp.StatusNotFound)
	}

}

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		httpHandler(ctx)
	}

	fmt.Println("Server started on http://localhost:9000")

	if err := fasthttp.ListenAndServe(":9000", requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
