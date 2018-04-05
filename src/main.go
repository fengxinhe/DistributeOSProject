package main
import(

    "fmt"
    "os"
    "encoding/json"
    "./controller/server"
    "./controller/static"
    "./view"
    "./controller/route"
    //"./public/database"
    "./public/jsonconfig"
)

type configuration struct {

	Server    server.Server   `json:"Server"`
//	Session   session.Session `json:"Session"`
	View      view.View       `json:"View"`
    Static    static.StaticInfo    `json:"Static"`
}
var config = &configuration{}
func main(){

    fmt.Println("ok")
    jsonconfig.LoadConfig("config" + string(os.PathSeparator)+"config.json", config)
    view.Configure(config.View)
    static.Configure(config.Static)

    server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
