package controller
import (
    "fmt"
    //"os"
    //"io"
    "io/ioutil"
    "net/http"
    "view"
    "model"
    "log"
     mux "github.com/julienschmidt/httprouter"
     "strconv"
     "time"
     "controller/session"
)

func CreateBlogGet(w http.ResponseWriter, r *http.Request, _ mux.Params) {
    v := view.New(r)
    sess := session.Instance(r)
    if sess.Values["authenticated"]==1{
        v.Data["Username"] = sess.Values["username"]
        v.Name = "write"
        v.RenderTemplate(w)
    }else{
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

}

func CreateBlogPost(w http.ResponseWriter, r *http.Request,_ mux.Params) {

    //title:=getProject(name)
    fmt.Println("post")
    //check the title if exist

    sess := session.Instance(r)

    r.ParseMultipartForm(32 << 20)
   var blog model.Blog
   blog.BlogID = strconv.FormatInt(time.Now().UnixNano(),10)
   blog.BlogTitle = r.FormValue("blog_title")
   blog.BlogAuthor=sess.Values["username"].(string)
   blog.BlogHeat = 0

   var temp=r.FormValue("summernotecode")
   fmt.Println(temp)
  // var fn=blog.BlogID+"_"+blog.BlogAuthor
   blog.BlogContent=temp


   if err := model.CreateBlog(blog); err != nil {
       log.Println(err)
       sess.Save(r,w)

       //respondWithError(w, http.StatusInternalServerError, err.Error())
       return
   } else {
       sess.Save(r,w)
       http.Redirect(w, r, "/", http.StatusFound)
       return
   }


    fmt.Println("ok")
}

func readfile(path string) string{
    data,_ := ioutil.ReadFile(path)
    return string(data[:])
}
func saveFile(fn string, content string) string {
    data := []byte(content)
    path := "/home/firebug/go/src/gobird/src/db/blog/"+fn
    ioutil.WriteFile(path, data,0600)
    return path
}
