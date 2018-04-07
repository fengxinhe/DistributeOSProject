package controller

import (
	"fmt"
	"log"
	"net/http"
    //"os"
    //"io"
    "view"
    "model"
    //mux "github.com/julienschmidt/httprouter"
	//"github.com/quasoft/memstore"
	"controller/session"


)


// LoginGET displays the login page
func LoginGet(w http.ResponseWriter, r *http.Request) {
	// Get session
//	sess := session.Instance(r)

	// Display the view
	//sess := session.Instance(r)
	v := view.New(r)
	v.Name = "login"
	if sess.Values["authenticated"]==1{
		v.Data["Username"] = sess.Values["username"]
	}else{
			v.Data["Username"] = "guest"
	}
	// Refill any form fields
	//view.Repopulate([]string{"email"}, r.Form, v.Data)
	v.RenderTemplate(w)
}

// LoginPOST handles the login form submission
func LoginPost(w http.ResponseWriter, r *http.Request) {

	sess:=session.Instance(r)
    username:=r.FormValue("user_name")
	password := r.FormValue("user_psd")

	// Get database result
	idealpsd, err := model.UserByName(username)

	// Determine if user exists
	if err != nil {
		// Display error message
		sess.Save(r,w)
		log.Println(err)

	} else if idealpsd != password {
		sess.Save(r,w)
		fmt.Println("invalid psd!")
	} else {
			// Login successfully
			//session.Values["username"]=username
		//	session.Values["password"]=password
			session.Empty(sess)
			sess.Values["authenticated"]=1
			sess.Values["username"]=username
			sess.Save(r,w)
			http.Redirect(w, r, "/", http.StatusFound)
			return
	}

	sess.Save(r,w)
	// Show the login page again
	LoginGet(w, r)
}

//LogoutGET clears the session and logs the user out
func LogoutGet(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// If user is authenticated
	if sess.Values["authenticated"]==1 {
	//	session.Empty(sess)
		sess.Values["authenticated"]=0
		session.Empty(sess)
		sess.Save(r, w)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
