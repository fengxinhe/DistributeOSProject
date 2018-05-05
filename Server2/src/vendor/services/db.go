package services

import "golang.org/x/net/websocket"

type Command struct {
    ServiceName    string
    Method      string
    Username    string
    Psd         int

    Author  string
    Content string

    UserId int
    InterestId string
    Action  int

    Num int
    Id  int

    A int
    B int

    DBid int
}


type ReplyMessage struct{
    Val  int
    List    []string
    Friends [5]int
}
var FollowDB = map[string]*[5]int{
    "aaa":{1,0,0,0,0},
    "bbb":{0,1,0,0,0},
}
///////////////////////////////
var UserList = []string{"aaa","bbb"}
var UserDB = map[string]int{
    "aaa": 111,
    "bbb": 222,
}
var UserStatus = map[string]int{
    "aaa": 0,
    "bbb": 0,
}
//////////////////////////////////
var BlogDB = make(map[int]string)
/////////////////////////////////////////
var LikeDB = make(map[int]int)
//////////////////////////////////////////
var LogDB=make(map[int]string)
/////

type connection struct {
    ws *websocket.Conn
    send chan string
}
type SocketMsg struct{
    Id  int
    WS		*websocket.Conn
}
