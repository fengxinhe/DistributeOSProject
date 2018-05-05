package main

import (
    "golang.org/x/net/websocket"
    "net/http"
    "net/rpc"
    "flag"
    "servermanager"
    //"server"
    "fmt"
)

func main() {
    port := flag.String("port", ":8000", "http service address")
    htdocs := flag.String("../client", "../client", "http dir")
    flag.Parse()
    servermanager.SM.StartServer()
    fmt.Println("sratttttt")
    rpc.Register(servermanager.SS)
    go servermanager.H.Run()
    http.Handle("/jsonrpc", websocket.Handler(servermanager.JsonrpcHandler))
    http.Handle("/notify", websocket.Handler(servermanager.NotifyHandler))
    http.Handle("/push", websocket.Handler(servermanager.PushHandler))
    http.Handle("/", http.FileServer(http.Dir(*htdocs)))
    err := http.ListenAndServe(*port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }

}
