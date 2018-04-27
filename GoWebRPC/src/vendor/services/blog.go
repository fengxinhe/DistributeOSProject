package services

import(
    //"errors"
    "strconv"
    "sync"
    "sync/atomic"
    "fmt"
)

type BlogInfo struct{
    Id  uint64

}
type Blogs struct {
    Author  string
    Content string
//    Heat    int
}

var blogmutex = &sync.Mutex{}
var Blog = new(BlogInfo)
var BlogDB = make(map[int]string)
func (b *BlogInfo) AddBlog(args *Blogs, reply *int) error {
     blogmutex.Lock()
     fmt.Println("send")
     BlogDB[int(b.Id)]=args.Content
     LikeDB[int(b.Id)]=0
     *reply=int(b.Id)
     msg:="addblog"+" "+strconv.Itoa(int(b.Id))+" "+args.Author+" "+args.Content
     friends:=GetFriends(args.Author)
     smsg:=Message{friends:friends, data: msg,}
    // H.broadcast <- msg
     H.message <- smsg
     //b.Id++
     atomic.AddUint64(&b.Id, 1)
     blogmutex.Unlock()
     return nil
}
