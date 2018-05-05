package servermanager

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

    CurrentTerm int

}

type DBMsg struct{
    FollowDB  map[string]*[5]int
    UserList  []string
    UserDB  map[string]int
    UserStatus  map[string]int
    BlogDB  map[int]string
    LikeDB  map[int]int
}
type DBQueryMsg struct{
    Name string
}

type ServerReply struct{
    Ack bool
    ServerId int
    ServerName string
    LastCommit int
}
type HeartMessage struct{
    LeaderId  int
}
type ReplyMessage struct{
    Val  int
    List    []string
    Friends [5]int
}
type Log struct {
    // ApplyFunc   func(*LogEntry, Command) (interface{}, error)
    // file        *os.File
    // path        string
    // entries     []*LogEntry
     commitIndex uint64
    // mutex       sync.RWMutex
    // startIndex  uint64
    // startTerm   uint64
    initialized bool
}
//
// var NDB = [5]NodeDB{
//     NodeDB{
//     FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
//     UserList :[]string{"aaa","bbb"},
//     UserDB :map[string]int{"aaa": 111,"bbb": 222,},
//     UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
//     BlogDB :make(map[int]string),
//     LikeDB :make(map[int]int),
//     },
//     NodeDB{
//         FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
//         UserList :[]string{"aaa","bbb"},
//         UserDB :map[string]int{"aaa": 111,"bbb": 222,},
//         UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
//         BlogDB :make(map[int]string),
//         LikeDB :make(map[int]int),
//     },
//     NodeDB{
//         FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
//         UserList :[]string{"aaa","bbb"},
//         UserDB :map[string]int{"aaa": 111,"bbb": 222,},
//         UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
//         BlogDB :make(map[int]string),
//         LikeDB :make(map[int]int),
//     },
//     NodeDB{
//         FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
//         UserList :[]string{"aaa","bbb"},
//         UserDB :map[string]int{"aaa": 111,"bbb": 222,},
//         UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
//         BlogDB :make(map[int]string),
//         LikeDB :make(map[int]int),
//     },
//     NodeDB{
//         FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
//         UserList :[]string{"aaa","bbb"},
//         UserDB :map[string]int{"aaa": 111,"bbb": 222,},
//         UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
//         BlogDB :make(map[int]string),
//         LikeDB :make(map[int]int),
//     },}

// var FollowDB = map[string]*[5]int{
//     "aaa":{1,0,0,0,0},
//     "bbb":{0,1,0,0,0},
// }
///////////////////////////////
// var UserList = []string{"aaa","bbb"}
// var UserDB = map[string]int{
//     "aaa": 111,
//     "bbb": 222,
// }
// var UserStatus = map[string]int{
//     "aaa": 0,
//     "bbb": 0,
// }
////////////////////////////////////
// var BlogDB = make(map[int]string)
// /////////////////////////////////////////
// var LikeDB = make(map[int]int)
//////////////////////////////////////////
type connection struct {
    ws *websocket.Conn
    send chan string
}
type SocketMsg struct{
    Id  int
    WS		*websocket.Conn
}
