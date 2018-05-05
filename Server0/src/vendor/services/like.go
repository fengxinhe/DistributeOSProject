package services

import(
    //"errors"
    //"strconv"
    "fmt"
    "sync"
)

type LikeInfo struct{
    Id  int
    likemutex  sync.Mutex
}
type Likes struct {
    Num int
    Id  int
}
var Like = new(LikeInfo)

func (like *LikeInfo) LikeHandler(args *Command, reply *ReplyMessage) error {
    fmt.Println("LikeHandler")
    like.likemutex.Lock()
    defer like.likemutex.Unlock()
    Node.LastApplied++
    if args.Num==1 {
        LikeDB[args.Id]+=1
        reply.Val=LikeDB[args.Id]
    }else if args.Num==-1{
        LikeDB[args.Id]-=1
        reply.Val=LikeDB[args.Id]
    }else {
        fmt.Println("like error")
        return nil
    }
     return nil
}
