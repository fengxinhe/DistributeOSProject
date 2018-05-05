package services

import (
    "sync"
    "fmt"
)

type FollowInfo struct{
    Id int
    followmutex sync.Mutex
}

type Following struct{
    UserId int
    InterestId string
    Action  int
}
var Follow=new(FollowInfo)
//var followinglist=[]string{}
// var FollowDB = map[string]*[5]int{
//     "aaa":{1,0,0,0,0},
//     "bbb":{0,1,0,0,0},
// }
//var followmutex = &sync.Mutex{}

func (f *FollowInfo)FollowHandler(args *Command, reply *ReplyMessage) error{
    f.followmutex.Lock()
    defer f.followmutex.Unlock()
    fmt.Printf("follow action: %d\n", args.Action)
    if list, ok := FollowDB[args.InterestId]; ok {
        list[args.UserId]=args.Action
    }else{
        FollowDB[args.InterestId]=&[5]int{0}
        FollowDB[args.InterestId][args.UserId]=args.Action
    }
    Node.LastApplied++
    reply.Val=1
    return nil
}
