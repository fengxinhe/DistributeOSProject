package controller

import (
    //"fmt"
    //"io/ioutil"
    "net/http"
    "view"
    "model"
    //"regexp"
    //"errors"
    mux "github.com/julienschmidt/httprouter"
    "controller/session"
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
