package services

import (
    "net/rpc"
    "sync"
    "fmt"
)

type UserInfo struct {
    Id int
    Client [5]*rpc.Client
    Mutex  [5]sync.Mutex
    usermutex  sync.Mutex
    //SocketClient  [32]*websocket.Conn
}
type Users struct {
    Username    string
    Psd         int
}
type LoginResponse struct {
    Id        int
    Username  string
    Success   int
}
// var UserList = []string{"aaa","bbb"}
// var UserDB = map[string]int{
//     "aaa": 111,
//     "bbb": 222,
// }
// var UserStatus = map[string]int{
//     "aaa": 0,
//     "bbb": 0,
// }
var User = new(UserInfo)
//var usermutex = &sync.Mutex{}

func findUserID(dbid int,name string) int{
        for i, val := range NDB[dbid].UserList{
            if val==name{
                return i
            }
        }
        return -1
}
func (u *UserInfo) Register (args *Command, id *int) error {
    u.usermutex.Lock()
    defer u.usermutex.Unlock()
    if -1 != findUserID(args.DBid,args.Username){
        *id=-1
        //usermutex.Unlock()
        return nil
    }
    NDB[args.DBid].UserList=append(NDB[args.DBid].UserList, args.Username)
    NDB[args.DBid].UserDB[args.Username]=args.Psd
    NDB[args.DBid].UserStatus[args.Username]=0
    var arr [5]int
    arr[findUserID(args.DBid,args.Username)]=1
    NDB[args.DBid].FollowDB[args.Username]=&arr
    *id=1
    // msg:="register"+" "+ args.Username
    // H.broadcast <- msg
    return nil
}

func (u *UserInfo) GetMember(args *Command, reply *[]string) error{
    tmp:=append([]string(nil), NDB[args.DBid].UserList...)
    tmp[args.Psd]="me"
    *reply=tmp
    return nil
}
func (u *UserInfo) Signin (args *Command, id *int) error {
    //*id = u.Id
    u.usermutex.Lock()
    defer u.usermutex.Unlock()
    fmt.Println("signin....")
    //fmt.Println(args.DBid)
    if args.Psd != NDB[args.DBid].UserDB[args.Username]{
        fmt.Println("error user psd")
        *id=-1

    }else{
        userid:=findUserID(args.DBid,args.Username)
        //fmt.Println(UserStatus[args.Username])
        if NDB[args.DBid].UserStatus[args.Username]==0{
            *id=userid
            NDB[args.DBid].UserStatus[args.Username]=1
        }else{
            *id=-1
        }
    }
    //fmt.Printf("signin handler%d\n", userid)
    return nil
}
func (u *UserInfo) Signout (args *Command, success *int) error {
    u.usermutex.Lock()
    defer u.usermutex.Unlock()
    userid:=findUserID(0,args.Username)
    //fmt.Println(userid)
    if NDB[args.DBid].UserStatus[args.Username]==1{
        *success=userid
        fmt.Printf("signout handler%d\n", userid)
        //u.Mutex[userid].Unlock()
        NDB[args.DBid].UserStatus[args.Username]=0
        //H.socketunregister <- userid
    }else{
        *success=-1
    }
    return nil

}
