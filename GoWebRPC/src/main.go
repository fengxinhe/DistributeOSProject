package main

import (
    "golang.org/x/net/websocket"
    "net/http"
    "net/rpc"
    "flag"
    "services"
)

func main() {
    port := flag.String("port", ":8010", "http service address")
    htdocs := flag.String("../client", "../client", "http dir")
    flag.Parse()

    rpc.Register(services.User)
    rpc.Register(services.Blog)
    rpc.Register(services.Like)
    rpc.Register(new(services.Arith))
    go services.H.Run()
    go http.Handle("/jsonrpc", websocket.Handler(services.JsonrpcHandler))
    go http.Handle("/notify", websocket.Handler(services.NotifyHandler))
    go http.Handle("/push", websocket.Handler(services.PushHandler))
    http.Handle("/", http.FileServer(http.Dir(*htdocs)))
    err := http.ListenAndServe(*port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }

}
