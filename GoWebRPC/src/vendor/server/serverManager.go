package server

import(
    "time"
    //"os"
    "services"
    "log"
    "net/rpc"
    "net"
    "net/http"
    "fmt"
    "sync"
)

type RequestMsg struct{
    ServiceName     string
}

type Command struct {
    Type    string
    Username    string
    Psd         int

    Author  string
    Content string

    UserId int
    InterestId string
    Action  int

    Num int
    Id  int

    A int
    B int
}

type Server struct{

    name    string
    host    string
    port   string
    rpcclient *rpc.Client

    votedFor    string
    log     *Log
    leader  string

    stopped  chan  bool
    electionTimeout   time.Duration
    heartbeatInterval  time.Duration

}
type ServerManager struct{
    servers   map[string]*Server


}
type Log struct {
    // ApplyFunc   func(*LogEntry, Command) (interface{}, error)
    // file        *os.File
    // path        string
    // entries     []*LogEntry
     commitIndex uint64
    // mutex       sync.RWMutex
    // startIndex  uint64
    // startTerm   uint64
    initialized bool
}

func (s *Server)NewServer(){
    l,e:=net.Listen("tcp","localhost:"+s.port)
    if e != nil{
		log.Fatal("listen error:", e)
	}
    fmt.Println("new server")
	go http.Serve(l, nil)
}
func (s *Server)DialServer(){
    s.rpcclient,err:=rpc.DialHTTP("tcp", "localhost:"+s.port)
    if err != nil {
        log.Fatal("dialing:", err)
    }
    wg.Done()
}
var serverportmap = map[string]string{
    "node0":"6000",
    "node1":"6001",
    "node2":"6002",
    "node3":"6003",
    "node4":"6004",
}
var SS = new(Command)
var SM= ServerManager{
    servers: make(map[string]*Server),
}
var wg sync.WaitGroup

func (sm *ServerManager)StartServer(){
    rpc.Register(new(services.Arith))
    // rpc.Register(services.User)
    // rpc.Register(services.Blog)
    // rpc.Register(services.Like)
    // rpc.Register(services.Follow)
    rpc.HandleHTTP()
    for name, port := range serverportmap{
        wg.Add(1)
        go func(sName string, sPort string){
            server:=&Server{name: sName, port:sPort,}
            fmt.Println(sName)
            sm.servers[sName]=server
            server.NewServer()
            wg.Done()
        }(name, port)
    }
    wg.Wait()
    fmt.Println("server init ok")
    for _,server := range sm.servers{
        wg.Add(1)
        go server.DialServer()

    }
    wg.Wait()
}

func (s *ServerManager) DispatchCommand(method string, args *interface{},reply *int){
    var dpwg sync.WaitGroup

    for _,server := range sm.servers{
        dpwg.Add(1)
        go func(method string, args *interface{},reply *int){
            err:=server.rpcclient.Call(method,args,&reply)
            if err != nil {
                log.Fatal("arith error:", err)
            }
            dpwg.Done()
        }(method, args, reply)
    }
    wg.Wait()
}
func (s *Command) RequestHandler(args *Command, reply *int) error {
    fmt.Println("request handler")
    if args.Type=="Arith"{
        args := &services.Args{7,8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
        err2 = client2.Call("Arith.Multiply", args, &reply)
        err2 = client3.Call("Arith.Multiply", args, &reply)


	//	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

    }
    *reply=1000
    return nil
}
