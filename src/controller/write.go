package controller
import (
    "fmt"
    //"os"
    //"io"
    "io/ioutil"
    "net/http"
    "../view"
    "../model"
    "log"
    //"regexp"
    //"errors"
    //"gopkg.in/mgo.v2/bson"

    //"image"
    //"image/jpeg"
    //"github.com/disintegration/imaging"
    //"github.com/gorilla/schema"
    //"strings"
    // "strconv"
     mux "github.com/julienschmidt/httprouter"
     "strconv"
     "time"

)

func CreateBlogGet(w http.ResponseWriter, r *http.Request, _ mux.Params) {
    v := view.New(r)
    v.Name = "write"
    //view.Repopulate([]string{"class_surface_img","class_title", "class_summary", "class_content","first_tag","secondtag"}, r.Form, v.Data)
    v.RenderTemplate(w)
}

// type Blog struct {
//     BlogID      string
//     BlogAuthor  string
//     BlogTitle   string
//     BlogContent string
//     BlogHeat    int
// }

func CreateBlogPost(w http.ResponseWriter, r *http.Request,_ mux.Params) {

    //title:=getProject(name)
    fmt.Println("post")
    //check the title if exist


    r.ParseMultipartForm(32 << 20)
   var blog model.Blog
   blog.BlogID = strconv.FormatInt(time.Now().UnixNano(),10)
   blog.BlogTitle = r.FormValue("blog_title")
   blog.BlogAuthor="aa"
   blog.BlogHeat = 10

   var temp=r.FormValue("summernotecode")
   fmt.Println(temp)
  // var fn=blog.BlogID+"_"+blog.BlogAuthor
   blog.BlogContent=temp


   if err := model.CreateBlog(blog); err != nil {
       log.Println(err)
       //respondWithError(w, http.StatusInternalServerError, err.Error())
       return
   } else {
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
    path := "/home/firebug/go/gobird/src/db/blog/"+fn
    ioutil.WriteFile(path, data,0600)
    return path
}
