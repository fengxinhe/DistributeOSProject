package controller

import (
    "fmt"
    //"io/ioutil"
    "net/http"
    "view"
    "model"
    //"regexp"
    //"errors"
    mux "github.com/julienschmidt/httprouter"
    "controller/session"
    "github.com/gorilla/schema"
)

func IndexGet(w http.ResponseWriter, r *http.Request,_ mux.Params) {
    v := view.New(r)
    sess := session.Instance(r)
    v.Name = "index"
    v.Data["Blogs"] = model.GetBlogs()
    if sess.Values["authenticated"]==1{
        v.Data["Username"] = sess.Values["username"]
    }else{
        v.Data["Username"] = "guest"
    }
    v.RenderTemplate(w)
    return
}
type likeBook struct{
    Id  string
    Like  int
}
func LikeHandler(w http.ResponseWriter, r *http.Request,_ mux.Params) {
    //sess :=session.Instance(r)
    err:=r.ParseForm()
    if err != nil {
        fmt.Println("ajax data error")
    }
    var lb likeBook
    var decoder = schema.NewDecoder()
    err = decoder.Decode(&lb, r.PostForm)
    if err!=nil{
        fmt.Println("decode fail!")
    }
    fmt.Println(lb)
    model.ModifyHeat(lb.Id,lb.Like)
}
