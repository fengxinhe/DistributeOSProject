package main

import (
    "golang.org/x/net/websocket"
    "net/http"
    "net/rpc"
    "flag"
    "services"
    "server"
    "fmt"
)

func main() {
    port := flag.String("port", ":8010", "http service address")
    htdocs := flag.String("../client", "../client", "http dir")
    flag.Parse()
    server.SM.StartServer()
    rpc.Register(services.User)
    fmt.Println("sratttttt")
    rpc.Register(services.Blog)
    rpc.Register(services.Like)
    rpc.Register(server.SS)
    rpc.Register(services.Follow)
    go services.H.Run()
    http.Handle("/jsonrpc", websocket.Handler(services.JsonrpcHandler))
    http.Handle("/notify", websocket.Handler(services.NotifyHandler))
    http.Handle("/push", websocket.Handler(services.PushHandler))
    http.Handle("/", http.FileServer(http.Dir(*htdocs)))
    err := http.ListenAndServe(*port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }

}
