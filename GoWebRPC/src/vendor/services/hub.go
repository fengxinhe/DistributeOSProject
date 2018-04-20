package services

import "fmt"
type hub struct {
    // Registered connections.
    connections map[*connection]bool

    // Inbound messages from the connections.
    broadcast chan string

    // Register requests from the connections.
    register chan *connection

    // Unregister requests from connections.
    unregister chan *connection
}

var H = hub{
    broadcast: make(chan string),
    register: make(chan *connection),
    unregister: make(chan *connection),
    connections: make(map[*connection]bool),
}

func (h *hub) Run() {
    for {
        select {
        case c := <-h.register:
            h.connections[c] = true
        case c := <-h.unregister:
            delete(h.connections, c)
            close(c.send)
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
