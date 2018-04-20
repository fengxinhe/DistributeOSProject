package services

import(
    //"errors"
    "strconv"
)

type LikeInfo struct{
    Id  int

}
type Likes struct {
    Num int
    Id  int
}

var Like = new(LikeInfo)
var LikeDB = make(map[int]string)
func (like *LikeInfo) LikeHandler(args *Likes, reply *int) error {

     msg:="likeChange"+" "+strconv.Itoa(like.Id)+" "+strconv.Itoa(args.Num)+" "+strconv.Itoa(args.Id)
    
     H.broadcast <- msg
     like.Id++
     return nil
}
