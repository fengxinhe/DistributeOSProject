package services

import(
    //"errors"
    "strconv"
)

type BlogInfo struct{
    Id  int

}
type Blogs struct {
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
     msg:="addblog"+" "+strconv.Itoa(b.Id)+" "+args.Content+" "+strconv.Itoa(args.Heat)
     H.broadcast <- msg
     b.Id++
     return nil
}
