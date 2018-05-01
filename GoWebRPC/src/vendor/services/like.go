package services

import(
    //"errors"
    "strconv"
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

func (like *LikeInfo) LikeHandler(args *Likes, reply *int) error {
    fmt.Println("LikeHandler")
    like.likemutex.Lock()
    defer like.likemutex.Unlock()
    if args.Num==1 {
        LikeDB[args.Id]+=1
        *reply=LikeDB[args.Id]
    }else if args.Num==-1{
        LikeDB[args.Id]-=1
        *reply=LikeDB[args.Id]
    }else {
        fmt.Println("like error")
        return nil
    }
     msg:="modifylike"+" "+strconv.Itoa(args.Id)+" "+strconv.Itoa(LikeDB[args.Id])
     H.broadcast <- msg
     return nil
}
