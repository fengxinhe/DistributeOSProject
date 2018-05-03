package services

import(
    "time"
    //"os"
    "log"
    "net/rpc"
    "net"
    "net/http"
    "fmt"
    "sync"
    "strconv"
    "strings"
    //"sync/atomic"

)

type Command struct {
    ServiceName    string
    Method      string
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

    DBid int
}

type Server struct{

    Name    string
    ServerId       int
    Host    string
    Port   string
    Rpcclient *rpc.Client

    votedFor    string
    log     [100]string
    commitIndex   uint64
    leader  string

    stopped  chan  bool
    electionTimeout   time.Duration
    heartbeatInterval  time.Duration

}
type ServerManager struct{
    servers   map[string]*Server
    leaderName string
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
    l,e:=net.Listen("tcp","localhost:"+s.Port)
    if e != nil{
		log.Fatal("listen error:", e)
	}
    fmt.Println("new server")
	go http.Serve(l, nil)
}
func (s *Server)DialServer(){
    var err error
    s.Rpcclient, err =rpc.DialHTTP("tcp", "localhost:"+s.Port)
    if err != nil {
         log.Fatal("dialing:", err)
     }
    wg.Done()
}
var serverportmap = map[string]string{
    "node-0":"3000",
    "node-1":"3001",
    "node-2":"3002",
    "node-3":"3003",
    "node-4":"3004",
}
var SS = new(Command)
var SM= ServerManager{
    servers: make(map[string]*Server),
}
var wg sync.WaitGroup

func (sm *ServerManager)StartServer(){
    rpc.Register(new(HeartBeat))
    rpc.Register(User)
    rpc.Register(Blog)
    rpc.Register(Like)
    rpc.Register(Follow)
    rpc.HandleHTTP()
    var mapmutex = &sync.Mutex{}
    for name, port := range serverportmap{
        wg.Add(1)
        go func(sName string, sPort string){
            server:=&Server{Name: sName, Port:sPort,}
            server.ServerId,_=strconv.Atoi(strings.Split(sName,"-")[1])
            //fmt.Println(server.id)
            mapmutex.Lock()
            sm.servers[sName]=server
            mapmutex.Unlock()
            server.NewServer()
            wg.Done()
        }(name, port)
    }
    wg.Wait()
    sm.leaderName="node-0"
    fmt.Println("server init ok")
    for _,server := range sm.servers{
        wg.Add(1)
        go server.DialServer()

    }
    wg.Wait()
    for _,server := range sm.servers{
        go server.HeartBeat()
    }
}

func (s *ServerManager) DispatchCommand(service string, args *Command) int{
    var dpwg sync.WaitGroup
    replys := make(map[string]int, 5)
    var mutex = &sync.Mutex{}
    for node,server := range s.servers{
        dpwg.Add(1)
        go func(node string, cmd *Command,server *Server){
            var reply int
            mutex.Lock()
            cmd.DBid=server.ServerId
            err:=server.Rpcclient.Call(service,&cmd,&reply)
            //fmt.Println(cmd)
            fmt.Println(reply)
            replys[node]=reply
            mutex.Unlock()
            if err != nil {
                log.Fatal("rpc error:", err)
            }
            dpwg.Done()
        }(node, args, server)
    }
    dpwg.Wait()
    return replys[s.leaderName]
    //fmt.Println(replys)
}

// func ResponseHandler(replys map){
//     var mistake int
//     var res int
//     var lostserver []string
//     num:=0
//     same:=make(map[int]int)
//     for name, _ := range serverportmap{
//         if reply, ok:=replys[name];ok{
//             if reply==-1{
//                 atomic.AddUint64(&mistake, 1)
//
//             }
//             else{
//                 atomic.AddUint64(&num, 1)
//                 same[reply]=1
//             }
//         }else{
//             atomic.AddUint64(&mistake, 1)
//             lostserver=append(lostserver,name)
//         }
//     }
//     if num >= 3{
//         return same[]
//     }
//
// }

type ServerReply struct{
    Ack bool

}
func (s *Server)HeartBeat(){
    //i:=0
    for {
        time.Sleep(1*time.Second)
        var reply bool
        HeartBeatMsg:=HeartMessage{s.ServerId}
        fmt.Println("send heartbeat to "+s.Name)
        err:=s.Rpcclient.Call("HeartBeat.HeartBeat",HeartBeatMsg,&reply)
        fmt.Println(reply)
        if err!=nil{
            log.Fatal("call error: ",err)
        }
        if reply{
            fmt.Println(s.Name+" normal!")
        }

    }
}
func DBRecovery(){

}
func (c *Command) RequestHandler(args *Command, reply *int) error {
    fmt.Println("request handler")
    //var reply int
    service := args.ServiceName
    *reply=SM.DispatchCommand(service, args)
    switch method:=args.Method; method{
    case "Signout":
        H.socketunregister <- *reply
    case "AddBlog":
        bmsg:="addblog"+" "+strconv.Itoa(*reply)+" "+args.Author+" "+args.Content
        friends:=GetFriends(0,args.Author)
        smsg:=Message{friends:friends, data: bmsg,}
        H.message <- smsg
    case "LikeHandler":
        fmt.Println(*reply)
        lmsg:="modifylike"+" "+strconv.Itoa(args.Id)+" "+strconv.Itoa(*reply)
        H.broadcast <- lmsg
    case "Register":
        rmsg:="register"+" "+ args.Username
        H.broadcast <- rmsg
    }

    return nil
}
