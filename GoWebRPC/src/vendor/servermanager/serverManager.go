package servermanager

import(
    "time"
    //"os"
    //"log"
    "net/rpc"
    "net/rpc/jsonrpc"
///    "net"
//    "net/http"
    "fmt"
    "sync"
    "strconv"
//    "strings"
    //"sync/atomic"

)



type ServerManager struct{
    //servers   map[string]*Server
    LeaderName string
    Leaderid    int
    RpcServer   map[string]*rpc.Client

    MapLock sync.Mutex
}

var serverportmap = map[string]string{
    "node-0":"3000",
    "node-1":"3001",
    "node-2":"3002",
}
var SS = new(Command)
var SM= ServerManager{
    RpcServer: make(map[string]*rpc.Client),
}
var wg sync.WaitGroup
func (s *ServerManager)DialServer(name string, port string){
    fmt.Println("Dialling")
    sv, err :=jsonrpc.Dial("tcp", "localhost:"+port)//+s.Port)
    if err != nil {
         //log.Fatal("dialing:", err)
         fmt.Println("dial error")
     }else{
         s.MapLock.Lock()
         defer s.MapLock.Unlock()
         s.RpcServer[name]=sv
    }
    wg.Done()
}
func (sm *ServerManager)StartServer(){

    sm.LeaderName="node-1"
    sm.Leaderid=1
    fmt.Println("server init ok")
    for name, port := range serverportmap{
        wg.Add(1)
        go sm.DialServer(name,port)
    }

    wg.Wait()
    fmt.Println("heartbeat")
    for name, _ := range sm.RpcServer{
        go sm.HeartBeat(name)
    }
}

func (sm *ServerManager) DispatchCommand(service string, args *Command)(ReplyMessage,error){

    var wrongsignal int
    HeartBeatMsg:=&HeartMessage{sm.Leaderid}
    var serverreply ServerReply
    //fmt.Println("send heartbeat to "+s.Name)
    var dpwg sync.WaitGroup

    for name, _ := range sm.RpcServer{
        dpwg.Add(1)
        go func(name string, wrong *int) error{
            err:=sm.RpcServer[name].Call("HeartBeat.HeartBeat",HeartBeatMsg,&serverreply)
            if err!=nil || !serverreply.Ack{
                *wrong++
                if(name==sm.LeaderName){
                sm.LeaderElection()
                    fmt.Println("Leader fail.....")
                    *wrong=3
                    return err
                }
            }
            if *wrong>1{
                fmt.Printf("err---------%v\n",err)
                return  err
            }
            dpwg.Done()
            return nil
        }(name,&wrongsignal)

    }
    dpwg.Wait()

    unsuccess:=0
    var reply ReplyMessage
    var err error
    for name, sv := range sm.RpcServer {
        dpwg.Add(1)
        go func(sname string,server *rpc.Client,cmd *Command,wrong *int) error{
            var tmpreply ReplyMessage
            if sname==sm.LeaderName {
                err=server.Call(service, cmd,&reply)
            }else{
                err=server.Call(service, cmd,&tmpreply)
            }
            if err!=nil{
                reply.Val=-1
                *wrong++
                if(sname==sm.LeaderName){
                    fmt.Println("Leader fail.....,wait......")
                    sm.LeaderElection()
                }
            }
            if *wrong>1{
                return err
            }
            dpwg.Done()
            return nil
        }(name, sv,args,&unsuccess)
    }
    dpwg.Wait()

    //return replys[sm.leaderid]
    //fmt.Println(replys)
    return reply,err
}
func (sm *ServerManager)LeaderElection(){
    cmt:=0
    count:=0
    var wg sync.WaitGroup
    for name, _ := range sm.RpcServer {
        wg.Add(1)
        go func(cnt *int, cindex *int, name string){
            HeartBeatMsg:=&HeartMessage{-1}
            var Reply ServerReply
            err := sm.RpcServer[name].Call("HeartBeat.Voting",HeartBeatMsg,&Reply)
            if err!=nil{
                fmt.Println(name+" down!")
                *cnt++
                if *cnt>1{
                    return
                }
            }else{
                if Reply.LastCommit>=*cindex {
                    sm.Leaderid=Reply.ServerId
                    sm.LeaderName=name
                    *cindex=Reply.LastCommit
                }
            }
            wg.Done()
        }(&count,&cmt, name)
    }
    wg.Wait()
    fmt.Println("choosing....->"+sm.LeaderName)
    count=0
    // broadcast every node who is the leader
    for name, _ := range sm.RpcServer {
        wg.Add(1)
        go func(cnt *int, name string){
            HeartBeatMsg:=&HeartMessage{sm.Leaderid}
            var Reply ServerReply
            err := sm.RpcServer[name].Call("HeartBeat.SetLeader",HeartBeatMsg,&Reply)
            if err!=nil{
                fmt.Println(name+" down!")
                *cnt++
                if *cnt>1 {
                    return
                }
            }
            wg.Done()
        }(&count,name)
    }
    wg.Wait()

}

func (s *ServerManager)HeartBeat(name string){
    //i:=0
    for {
        time.Sleep(3*time.Second)
        //var reply bool
        HeartBeatMsg:=&HeartMessage{s.Leaderid}
        var Reply ServerReply
        //fmt.Println("send heartbeat to "+s.Name)
        err:=s.RpcServer[name].Call("HeartBeat.HeartBeat",HeartBeatMsg,&Reply)
        //fmt.Println(Reply)
        if err!=nil{
            //log.Fatal("call error: ",err)
            s.ReDialling(name,serverportmap[name])
            if name == s.LeaderName{
                fmt.Println("start leader election......")
                s.LeaderElection()
            }
            fmt.Println(err)
        }
        if Reply.Ack{
            fmt.Println(name +"....... normal!")
        }

    }
}
func (s *ServerManager)ReDialling(name string, port string){
    fmt.Println("Redialling")
    sv, err :=jsonrpc.Dial("tcp", "localhost:"+port)//+s.Port)
    if err != nil {
         fmt.Println("dial error")
     }else{
         s.MapLock.Lock()
         defer s.MapLock.Unlock()
         s.RpcServer[name]=sv
         s.DBRecovery(name)
    }
}

func (s *ServerManager)DBRecovery(name string){

    var db DBMsg
    DBBeatMsg:=&HeartMessage{s.Leaderid}
    var successid int
    s.RpcServer[s.LeaderName].Call("HeartBeat.GetDB",DBBeatMsg,&db)
    s.RpcServer[name].Call("HeartBeat.DBRecovery",&db,&successid)
    fmt.Printf("DB recovering successfully---%d\n",successid)

}
func (c *Command) RequestHandler(args *Command, reply *ReplyMessage) error {
    fmt.Println("request handler")
    //fmt.Println(args)
    //var reply int
    service := args.ServiceName
    var err error
    *reply,err=SM.DispatchCommand(service, args)
    if err!=nil{
        reply.Val=-1
        return err
    }
    //Websocket channel
    switch method:=args.Method; method{
    case "Signout":
        H.socketunregister <- reply.Val
    case "AddBlog":
        bmsg:="addblog"+" "+strconv.Itoa(reply.Val)+" "+args.Author+" "+args.Content
        friends:=reply.Friends
        smsg:=Message{friends:friends, data: bmsg,}
        H.message <- smsg
    case "LikeHandler":
        lmsg:="modifylike"+" "+strconv.Itoa(args.Id)+" "+strconv.Itoa(reply.Val)
        H.broadcast <- lmsg
    case "Register":
        rmsg:="register"+" "+ args.Username
        H.broadcast <- rmsg
    }

    return nil
}
