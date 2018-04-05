package server

import (
    "fmt"
	"log"
	"net/http"
    "time"
    //"../"
)

type Server struct {
    HostName string
    UseHttp bool
    UseHttps bool
    HttpPort string
    HttpsPort string
}

func Run(httpHandlers http.Handler, httpsHandlers http.Handler, s Server) {
    //shttp.Handle("/", http.FileServer(http.Dir("/goweb/static/")))
    if s.UseHttp && s.UseHttps {
		// go func() {
		// 	//startHTTPS(httpsHandlers, s)
		// }()

		startHTTP(httpHandlers, s)
	} else if s.UseHttp {
		startHTTP(httpHandlers, s)
	} else if s.UseHttps {
		//startHTTPS(httpsHandlers, s)
	} else {
		log.Println("Config file does not specify a listener to start")
    }
}

func startHTTP(handlers http.Handler, s Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP "+httpAddress(s))

	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(s), handlers))
}

// startHTTPs starts the HTTPS listener
// func startHTTPS(handlers http.Handler, s Server) {
// 	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTPS "+httpsAddress(s))
//
// 	// Start the HTTPS listener
// 	log.Fatal(http.ListenAndServeTLS(httpsAddress(s), s.CertFile, s.KeyFile, handler))
// }

// httpAddress returns the HTTP address
func httpAddress(s Server) string {
	return s.HostName + ":" + s.HttpPort
}

// httpsAddress returns the HTTPS address
func httpsAddress(s Server) string {
	return s.HostName + ":" + fmt.Sprintf("%d", s.HttpsPort)
}
