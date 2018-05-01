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

func (f *FollowInfo)FollowHandler(args *Following, reply *int) error{
    f.followmutex.Lock()
    defer f.followmutex.Unlock()
    fmt.Printf("follow action: %d\n", args.Action)
    if list, ok := FollowDB[args.InterestId]; ok {
        list[args.UserId]=args.Action
    }else{
        FollowDB[args.InterestId]=&[5]int{0}
        FollowDB[args.InterestId][args.UserId]=args.Action
    }
    *reply=1
    return nil
}

func GetFriends(name string) [5]int{
    if list, ok := FollowDB[name]; ok {
        return *list
    }
    return [5]int{0}
}
