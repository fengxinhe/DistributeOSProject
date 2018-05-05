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

func findUserID(name string) int{
        for i, val := range UserList{
            if val==name{
                return i
            }
        }
        return -1
}
func (u *UserInfo) Register (args *Command, id *ReplyMessage) error {
    u.usermutex.Lock()
    defer u.usermutex.Unlock()
    Node.LastApplied++
    if -1 != findUserID(args.Username){
        id.Val=-1
        //usermutex.Unlock()
        return nil
    }
    UserList=append(UserList, args.Username)
    UserDB[args.Username]=args.Psd
    UserStatus[args.Username]=0
    var arr [5]int
    arr[findUserID(args.Username)]=1
    FollowDB[args.Username]=&arr
    id.Val=1
    // msg:="register"+" "+ args.Username
    // H.broadcast <- msg
    return nil
}

func (u *UserInfo) GetMember(args *Command, reply *ReplyMessage) error{
    Node.LastApplied++
    tmp:=append([]string(nil),UserList...)
    tmp[args.Psd]="me"
    reply.List=tmp
    return nil
}
func (u *UserInfo) Signin (args *Command, id *ReplyMessage) error {
    //*id = u.Id
    u.usermutex.Lock()
    defer u.usermutex.Unlock()
    Node.LastApplied++
    fmt.Println("signin....")
    fmt.Println(args.Username)
    fmt.Println(UserDB)
    //fmt.Println(args.DBid)
    if args.Psd != UserDB[args.Username]{
        fmt.Println("error user psd")
        id.Val=-1

    }else{
        userid:=findUserID(args.Username)
        //fmt.Println(UserStatus[args.Username])
        if UserStatus[args.Username]==0{
            id.Val=userid
            UserStatus[args.Username]=1
        }else{
            id.Val=-1
        }
    }
    //fmt.Printf("signin handler%d\n", userid)
    return nil
}
func (u *UserInfo) Signout (args *Command, success *ReplyMessage) error {
    u.usermutex.Lock()
    defer u.usermutex.Unlock()
    Node.LastApplied++
    userid:=findUserID(args.Username)
    //fmt.Println(userid)
    if UserStatus[args.Username]==1{
        success.Val=userid
        fmt.Printf("signout handler%d\n", userid)
        //u.Mutex[userid].Unlock()
        UserStatus[args.Username]=0
        //H.socketunregister <- userid
    }else{
        success.Val=-1
    }
    return nil

}
