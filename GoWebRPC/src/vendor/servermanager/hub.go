package servermanager

import "fmt"
import "golang.org/x/net/websocket"

type hub struct {
    // Registered connections.
    connections map[*connection]bool
    connectionMap map[int]*websocket.Conn

    // Inbound messages from the connections.
    broadcast chan string
    message chan Message

    // Register requests from the connections.
    register chan *connection
    socketregister chan SocketMsg
    // Unregister requests from connections.
    unregister chan *connection
    socketunregister chan int
}
type Message struct{
    friends [5]int
    data string
}

var H = hub{
    broadcast: make(chan string),
    register: make(chan *connection),
    unregister: make(chan *connection),
    connections: make(map[*connection]bool),
    connectionMap: make(map[int]*websocket.Conn),
    message: make(chan Message),
    socketregister: make(chan SocketMsg),
    socketunregister: make(chan int),
}

func (h *hub) Run() {
    for {
        select {
        case c := <-h.register:
            h.connections[c] = true
        case c := <-h.unregister:
            delete(h.connections, c)
            close(c.send)
        case id:=<-h.socketunregister:
            delete(h.connectionMap,id)
        case c:= <-h.socketregister:
            h.connectionMap[c.Id] = c.WS

        case m:=<-h.message:
            for i:=0;i<5;i++ {
                if m.friends[i]==1{
                    if ws,ok := h.connectionMap[i]; ok{
                        seperateWriter(ws,m.data)
                    }
                }
            }
        case m := <-h.broadcast:
            //fmt.Println(m)
            for c := range h.connections {
                fmt.Println("hub write...")
                select {
                case c.send <- m:
                default:
                    delete(h.connections, c)
                    close(c.send)
                    go c.ws.Close()
                }
            }
        }
    }
}
