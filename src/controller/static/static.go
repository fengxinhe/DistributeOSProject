package static

import(
    "net/http"
//    "strings"
    "fmt"
    "io"
    "time"
)

type StaticInfo struct {
    STATIC_URL string
    STATIC_ROOT string
}

var staticinfo StaticInfo

func Configure(s StaticInfo) {
    staticinfo = s
}

func GetInfo() string{
    return staticinfo.STATIC_ROOT
}

func Static(w http.ResponseWriter, req *http.Request) {
    static_file := req.URL.Path[len(staticinfo.STATIC_URL):]
    if len(static_file) != 0 {
        f, err := http.Dir(staticinfo.STATIC_ROOT).Open(static_file)
        if err == nil {
            content := io.ReadSeeker(f)
            http.ServeContent(w, req, static_file, time.Now(), content)
            return
        }
    }
    http.NotFound(w, req)
}


func GETStaticURL() string {
    return staticinfo.STATIC_URL
}

func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not Found 404")
}
