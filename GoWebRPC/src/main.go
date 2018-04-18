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
    htdocs := flag.String("../htdocs", "../htdocs", "http dir")
    flag.Parse()

    rpc.Register(services.User)
    //rpc.Register(new(services.Register))
    rpc.Register(new(services.Arith))
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
