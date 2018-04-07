package route

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/gorilla/context"
    "controller/static"
    "controller"
    "fmt"
)

func LoadHTTPS() http.Handler {
    return middleware(routes())
}

func LoadHTTP() http.Handler {
    return middleware(routes())
}

func routes() *httprouter.Router {
    router := httprouter.New()
    //router.ServeFiles("/static/*filepath", http.Dir("/home/firebug/goweb/"))
    router.GET("/home/firebug/go/src/gobird/static/*filepath",wrapHandler(http.HandlerFunc(static.Static)))
    fmt.Println("cssss sssss")

    router.GET("/", controller.IndexGet)
    // router.GET("/classes", wrapHandler(http.HandlerFunc(controller.ClassGet)))
    //
    // router.GET("/classes/create",wrapHandler(http.HandlerFunc(controller.CreateClassGet)))
    //
    // router.POST("/classes/create",wrapHandler(http.HandlerFunc(controller.CreateClassPost)))

    router.GET("/write",controller.CreateBlogGet)
    router.POST("/write",controller.CreateBlogPost)
    router.GET("/register",wrapHandler(http.HandlerFunc(controller.RegisterGet)))
    router.POST("/register",wrapHandler(http.HandlerFunc(controller.RegisterPost)))
    router.GET("/login",wrapHandler(http.HandlerFunc(controller.LoginGet)))
    router.POST("/login",wrapHandler(http.HandlerFunc(controller.LoginPost)))
    router.GET("/logout",wrapHandler(http.HandlerFunc(controller.LogoutGet)))
    //router.GET("/featured/show/:name",controller.ShowProjectGet)
    router.POST("/likehandler",controller.LikeHandler)


    return router
}

func middleware(h http.Handler) http.Handler {
    //h = logrequest.Handler(h)
    return h
}

func wrapHandler(h http.Handler) httprouter.Handle{
    return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
}
}
