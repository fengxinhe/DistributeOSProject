package jsonconfig

import (
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)
type Parser interface{
    ParseJSON([]byte) error
}
func LoadConfig(configFile string, p Parser){

    var err error
	var absPath string
    var file = io.ReadCloser(os.Stdin)
    if absPath, err = filepath.Abs(configFile); err != nil {
		log.Fatalln(err)
	}

	if file, err = os.Open(absPath); err != nil {
		log.Fatalln(err)
}
    jsonBytes, err := ioutil.ReadAll(file)
    file.Close()
     if err := p.ParseJSON(jsonBytes); err != nil {
         log.Fatalln("Could not parse %q: %v", configFile, err)
     }
}
