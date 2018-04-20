package services

import(
    //"errors"
    "strconv"
)

type BlogInfo struct{
    Id  int

}
type Blogs struct {
    Author  string
    Content string
    Heat    int
}

var Blog = new(BlogInfo)
var BlogDB = make(map[int]string)
var HeatDB = make(map[int]int)
func (b *BlogInfo) AddBlog(args *Blogs, reply *int) error {
     BlogDB[b.Id]=args.Content
     HeatDB[b.Id]=args.Heat
     *reply=b.Id
     msg:="addblog"+" "+strconv.Itoa(b.Id)+" "+args.Author+" "+args.Content+" "+strconv.Itoa(args.Heat)
    // mmsg := Message{
    //     Method: "addblog",
    //     BlogID: b.Id,
    //     Content: args.Content,
    //     Like : args.Heat,
    // }
     H.broadcast <- msg
     b.Id++
     return nil
}
