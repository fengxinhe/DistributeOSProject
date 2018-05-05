package services

import(
    //"errors"
    //"strconv"
    "sync"
    "sync/atomic"
    "fmt"
)

type BlogInfo struct{
    Id  uint64
    blogmutex   sync.Mutex
}
type Blogs struct {
    Author  string
    Content string
}

var Blog = new(BlogInfo)
func (b *BlogInfo) AddBlog(args *Command, reply *ReplyMessage) error {
     b.blogmutex.Lock()
     defer b.blogmutex.Unlock()
     fmt.Println("send")
     BlogDB[int(b.Id)]=args.Content
     LikeDB[int(b.Id)]=0
     reply.Val=int(b.Id)
     reply.Friends=GetFriends(args.Author)
     Node.LastApplied++
     atomic.AddUint64(&b.Id, 1)
     return nil
}

func GetFriends(name string) [5]int{
    if list, ok := FollowDB[name]; ok {
        Node.LastApplied++
        return *list
    }
    return [5]int{0}
}
