package model

import(
    "fmt"
    "io/ioutil"
    "encoding/json"
    "log"
)

type Blog struct {
    BlogID      string
    BlogAuthor  string
    BlogTitle   string
    BlogContent string
    BlogHeat    int
}

func CreateBlog(blog Blog) error{

    path:= "/home/firebug/go/gobird/src/db/blog/"+blog.BlogID+"-"+blog.BlogTitle+".json"
    data, err := json.Marshal(blog)
    if err != nil {
        log.Fatal(err)
        return err
    }
    ioutil.WriteFile(path, data,0600)
    fmt.Println("create_bb")
    return nil
}

func BlogUnmarshal(data []byte) *Blog{
    var blog Blog
    err := json.Unmarshal(data, &blog)
    if err != nil {
        log.Fatal(err)
    }
    return &blog
}

func GetBlogs() (*[]Blog){
    files, _ := ioutil.ReadDir("/home/firebug/go/gobird/src/db/blog/")
    var blogs []Blog

    for _, file := range files {

        fmt.Println("indexxxxxxx")
        data, _:= ioutil.ReadFile("/home/firebug/go/gobird/src/db/blog/"+file.Name())
        blogs=append(blogs,*BlogUnmarshal(data))
    }
    return &blogs
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
