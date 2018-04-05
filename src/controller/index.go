package controller

import (
    //"fmt"
    //"io/ioutil"
    "net/http"
    "../view"
    "../model"
    //"regexp"
    //"errors"
    mux "github.com/julienschmidt/httprouter"

)

func IndexGet(w http.ResponseWriter, r *http.Request,_ mux.Params) {
    v := view.New(r)
    v.Name = "index"
    //v.Data["Title"] = "Xinhe Feng"
    //v.Data["Classes"] = model.GetClasses()
    v.Data["Blogs"] = model.GetBlogs()
    v.RenderTemplate(w)
    return
}
