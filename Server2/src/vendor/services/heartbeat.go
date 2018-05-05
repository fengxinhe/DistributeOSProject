package services

import(
 //  "time"
//    "net"
//    "log"
    "fmt"
    //"net/http"
)

type HeartMessage struct{
    LeaderId  int
}
type ServerReply struct{
    Ack bool
    ServerId int
    ServerName string
    LastCommit int

}
type HeartBeat int
func (hb *HeartBeat)HeartBeat(args *HeartMessage, reply *ServerReply) error{
    fmt.Printf("leader id ->%d\n",args.LeaderId)
    reply.Ack=true
    reply.ServerId=0
    //Server.ElectionTimeout <- 5e8
    return nil
}

func (hb *HeartBeat) Voting(args *HeartMessage, reply *ServerReply) error{
    fmt.Println("Voting......")
    reply.ServerId=Node.ServerId
    reply.LastCommit=Node.LastApplied
    return nil
}

func (hb *HeartBeat) SetLeader(args *HeartMessage, reply *ServerReply) error{
    fmt.Println("Set Leader......")
    fmt.Println(args.LeaderId)
    Node.Leader=args.LeaderId
    return nil
}


func (hb*HeartBeat) GetDB(args *HeartMessage, reply *DBMsg) error{

    reply.FollowDB=FollowDB
    reply.UserList=UserList
    reply.UserDB=UserDB
    reply.UserStatus=UserStatus
    reply.BlogDB=BlogDB
    reply.LikeDB=LikeDB
    return nil
}

func (hb*HeartBeat) DBRecovery(args *DBMsg, reply *int) error{
    FollowDB=args.FollowDB
    UserList=args.UserList
    UserDB=args.UserDB
    UserStatus=args.UserStatus
    BlogDB=args.BlogDB
    LikeDB=args.LikeDB
    *reply=Node.ServerId
    return nil
}
