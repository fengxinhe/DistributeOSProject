package services

import (
    "net/rpc"
    "sync"
    "fmt"
)

type UserInfo struct {
    Id int
    Client [32]*rpc.Client
    Mutex  [32]sync.Mutex
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
var UserList = []string{"aaa","bbb"}
var UserDB = map[string]int{
    "aaa": 111,
    "bbb": 222,
}
var UserStatus = map[string]int{
    "aaa": 0,
    "bbb": 0,
}
var User = new(UserInfo)

func findUserID(name string) int{
        for i, val := range UserList{
            if val==name{
                return i
            }
        }
        return -1
}
func (u *UserInfo) Register (args *Users, id *int) error {
    if -1 != findUserID(args.Username){
        *id=-1
        return nil
    }
    UserList=append(UserList, args.Username)
    UserDB[args.Username]=args.Psd
    UserStatus[args.Username]=0
    *id=1
    return nil
}
func (u *UserInfo) Signup (args *Users, id *int) error {
    //*id = u.Id

    fmt.Println(args.Psd)
    if args.Psd != UserDB[args.Username]{
        fmt.Println("error user psd")
        *id=-1
        return nil
    }

        userid:=findUserID(args.Username)
        fmt.Println(userid)
        if UserStatus[args.Username]==0{
            *id=userid
            UserStatus[args.Username]=1
            //u.Mutex[*id].Lock()
        }else{
            *id=-1
        }

    // if *id == 32 {
    //     *id = 0
    //     u.Id = 0
    // } else {
    //     u.Id++
    // }

    //u.Mutex[*id].Lock()
    return nil
}
