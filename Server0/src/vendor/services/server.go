package services

import(
    "net/rpc"
   "time"
    "net"
    "log"
    "fmt"
    //"net/http"
    "net/rpc/jsonrpc"
)

type Server struct{

    Name    string
    ServerId       int
    Host    string
    Port   string
//    Rpcclient *rpc.Client
    VotedFor    string
    Log     [100]string
    CurrentTerm     int
    LastApplied     int
    Leader  int

    Stopped  chan  bool
    ElectionTimeout  chan  time.Duration
    HeartbeatInterval  time.Duration
}

var Node =Server{Name: "node-0", Port:"3000",ServerId:0,LastApplied:0,}

func (s *Server)NewServer(){
    rpc.Register(new(HeartBeat))
    rpc.Register(User)
    rpc.Register(Blog)
    rpc.Register(Like)
    rpc.Register(Follow)
    //rpc.HandleHTTP()
    l,e:=net.Listen("tcp","localhost:3000")
    if e != nil{
		log.Fatal("listen error:", e)
	}
    fmt.Println("Node-0 running...")
	//go http.Serve(l, nil)
    for {
       conn, err := l.Accept()
       if err != nil {
           continue
       }
       jsonrpc.ServeConn(conn)
       fmt.Println("Node-0 listening...")
   }
}
