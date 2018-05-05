package main

import(
//     "net/rpc"
// //    "time"
//     "net"
//     "log"
//     "fmt"
    //"net/http"
    "services"
//    "net/rpc/jsonrpc"
)

func main(){
    //server:=&Server{Name: "node-0", Port:"3000",ServerId:0,}
    services.Node.NewServer()

}
