package services

import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    // "encoding/json"
)

type connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // Buffered channel of outbound messages.
    send chan string
}
type SocketMsg struct{
    Id  int
    WS		*websocket.Conn
}
func (c *connection) reader() {
    for {
        var msg string
        fmt.Println("wwwrrrrr")
        err := websocket.Message.Receive(c.ws, &msg)
        if err != nil {
            break
        }
        H.broadcast <- msg
    }
    c.ws.Close()
}

func (c *connection) writer() {
    for message := range c.send {
        fmt.Println("wwwwwwwwwwww")
        fmt.Println(message)

        err := websocket.Message.Send(c.ws, message)
        if err != nil{
            log.Print("socket send error:", err)
            break
        }
    }
    c.ws.Close()
}

func seperateWriter(ws *websocket.Conn,msg string) {
    //for {
        fmt.Println("wwwwwwwwwwww")
        fmt.Println(msg)
        err := websocket.Message.Send(ws, msg)
        if err != nil{
            log.Print("socket send error:", err)
            ws.Close()
        }
    //}

}

func AddSocketConnection(ws *websocket.Conn, id int){
    sm := SocketMsg{Id:id, WS:ws}
    H.socketregister <- sm
}
func NotifyHandler(ws *websocket.Conn) {
    c := &connection{send: make(chan string, 100), ws: ws}
    fmt.Println("NotifyHandler")
    H.register <- c
    defer func() { H.unregister <- c }()
    go c.writer()
    c.reader()
}
