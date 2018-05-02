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
// var LikeDB = make(map[int]int)

func (like *LikeInfo) LikeHandler(args *Command, reply *int) error {
    fmt.Println("LikeHandler")
    like.likemutex.Lock()
    defer like.likemutex.Unlock()
    if args.Num==1 {
        NDB[args.DBid].LikeDB[args.Id]+=1
        *reply=NDB[args.DBid].LikeDB[args.Id]
    }else if args.Num==-1{
        NDB[args.DBid].LikeDB[args.Id]-=1
        *reply=NDB[args.DBid].LikeDB[args.Id]
    }else {
        fmt.Println("like error")
        return nil
    }
     // msg:="modifylike"+" "+strconv.Itoa(args.Id)+" "+strconv.Itoa(LikeDB[args.Id])
     // H.broadcast <- msg
     return nil
}
