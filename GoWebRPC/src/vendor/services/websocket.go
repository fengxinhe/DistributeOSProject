package services

import (
    "golang.org/x/net/websocket"
    "net/rpc/jsonrpc"
    "log"
    "fmt"
)

func JsonrpcHandler(ws *websocket.Conn) {
     jsonrpc.ServeConn(ws)
}

func PushHandler(ws *websocket.Conn) {
    var id int
    c := jsonrpc.NewClient(ws)
    //fmt.Println("websocket")

    err := c.Call("User.Getpsd", nil, &id)
    if err != nil {
        log.Print("User.Getid error:", err)
        return
    }
    //fmt.Println(c)
    User.Client[id] = c;
    User.Mutex[id].Lock()
    fmt.Printf("push handler%d\n", id)
    NotifyHandler(ws)

}

func PopHandler(ws *websocket.Conn) {

}
