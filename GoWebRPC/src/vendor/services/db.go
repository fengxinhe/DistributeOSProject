package services

import "golang.org/x/net/websocket"

type NodeDB struct{
    FollowDB  map[string]*[5]int
    UserList  []string
    UserDB  map[string]int
    UserStatus  map[string]int
    BlogDB  map[int]string
    LikeDB  map[int]int
}

var NDB = [5]NodeDB{
    NodeDB{
    FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
    UserList :[]string{"aaa","bbb"},
    UserDB :map[string]int{"aaa": 111,"bbb": 222,},
    UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
    BlogDB :make(map[int]string),
    LikeDB :make(map[int]int),
    },
    NodeDB{
        FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
        UserList :[]string{"aaa","bbb"},
        UserDB :map[string]int{"aaa": 111,"bbb": 222,},
        UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
        BlogDB :make(map[int]string),
        LikeDB :make(map[int]int),
    },
    NodeDB{
        FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
        UserList :[]string{"aaa","bbb"},
        UserDB :map[string]int{"aaa": 111,"bbb": 222,},
        UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
        BlogDB :make(map[int]string),
        LikeDB :make(map[int]int),
    },
    NodeDB{
        FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
        UserList :[]string{"aaa","bbb"},
        UserDB :map[string]int{"aaa": 111,"bbb": 222,},
        UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
        BlogDB :make(map[int]string),
        LikeDB :make(map[int]int),
    },
    NodeDB{
        FollowDB :map[string]*[5]int{"aaa":{1,0,0,0,0},"bbb":{0,1,0,0,0},},
        UserList :[]string{"aaa","bbb"},
        UserDB :map[string]int{"aaa": 111,"bbb": 222,},
        UserStatus :map[string]int{"aaa": 0,"bbb": 0,},
        BlogDB :make(map[int]string),
        LikeDB :make(map[int]int),
    },}

var FollowDB = map[string]*[5]int{
    "aaa":{1,0,0,0,0},
    "bbb":{0,1,0,0,0},
}
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
var BlogDB = make(map[int]string)
/////////////////////////////////////////
var LikeDB = make(map[int]int)
//////////////////////////////////////////
type connection struct {
    ws *websocket.Conn
    send chan string
}
type SocketMsg struct{
    Id  int
    WS		*websocket.Conn
}
