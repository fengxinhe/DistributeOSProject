package view

import (
    "net/http"
    "html/template"
    "net/url"
    "log"
    "fmt"
    "path/filepath"
    "github.com/oxtoacart/bpool"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool
var mainTmpl = `{{define "main"}} {{template "base" .}} {{end}}`
var viewInfo View

func init(){
    //templates = template.Must(template.ParseGlob("../templates/*")) //Template caching

    if templates == nil {
        templates = make(map[string]*template.Template)
    }
    templatesDir := "/home/firebug/go/src/gobird/templates/"
    //templatesDir := viewInfo.TemplateDir
    layoutFiles, err := filepath.Glob(templatesDir + "layouts/*.tmpl")
    if err != nil {
        log.Fatal(err)
    }

    includeFiles, err := filepath.Glob(templatesDir + "*.tmpl")
    if err != nil {
        log.Fatal(err)
    }
    mainTemplate := template.New("main")
    mainTemplate, err = mainTemplate.Parse(mainTmpl)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range includeFiles {
        fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
    }
    bufpool = bpool.NewBufferPool(64)
    fmt.Println(templates)
}

type View struct {
    BaseURI string
    StaticURL string
    TemplateDir string
    Name string
    Data map[string]interface{}
    request *http.Request
}

func Repopulate(list []string, src url.Values, dst map[string]interface{}) {
	for _, v := range list {
		dst[v] = src.Get(v)
	}
}

func Configure(vi View) {
    viewInfo = vi
}

func New(req *http.Request) *View {
    v := &View{}
    v.Data = make(map[string]interface{})

    v.request = req
    v.Data["Static"] = viewInfo.StaticURL
    return v
}


func (v *View)RenderTemplate(w http.ResponseWriter){

    //err := templates.ExecuteTemplate(w, v.Name+".html", v.Data)
    tmpl, ok := templates[v.Name+".tmpl"]
    if !ok {
        fmt.Errorf("The template does not exist.")
    }

	// if err != nil {
	// 	http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
    // }
    buf := bufpool.Get()
    defer bufpool.Put(buf)
    tmpl.Execute(buf,v.Data)
    buf.WriteTo(w)
}
