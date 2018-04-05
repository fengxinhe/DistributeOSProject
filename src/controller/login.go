package controller

import (
	"fmt"
	"log"
	"net/http"
    //"os"
    //"io"
    "../view"
    "../model"
	"github.com/gorilla/sessions"
    //mux "github.com/julienschmidt/httprouter"

)

const (
	// Name of the session variable that tracks login attempts
	sessLoginAttempt = "login_attempt"
)

// loginAttempt increments the number of login attempts in sessions variable
func loginAttempt(sess *sessions.Session) {
	// Log the attempt
	if sess.Values[sessLoginAttempt] == nil {
		sess.Values[sessLoginAttempt] = 1
	} else {
		sess.Values[sessLoginAttempt] = sess.Values[sessLoginAttempt].(int) + 1
	}
}

// LoginGET displays the login page
func LoginGet(w http.ResponseWriter, r *http.Request) {
	// Get session
//	sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "login"
	// Refill any form fields
	//view.Repopulate([]string{"email"}, r.Form, v.Data)
	v.RenderTemplate(w)
}

// LoginPOST handles the login form submission
func LoginPost(w http.ResponseWriter, r *http.Request) {
	// // Get session
	// sess := session.Instance(r)
    //
	// // Prevent brute force login attempts by not hitting MySQL and pretending like it was invalid :-)
	// if sess.Values[sessLoginAttempt] != nil && sess.Values[sessLoginAttempt].(int) >= 5 {
	// 	log.Println("Brute force login prevented")
	// 	sess.AddFlash(view.Flash{"Sorry, no brute force :-)", view.FlashNotice})
	// 	sess.Save(r, w)
	// 	LoginGET(w, r)
	// 	return
	// }
    //
	// // Validate with required fields
	// if validate, missingField := view.Validate(r, []string{"email", "password"}); !validate {
	// 	sess.AddFlash(view.Flash{"Field missing: " + missingField, view.FlashError})
	// 	sess.Save(r, w)
	// 	LoginGET(w, r)
	// 	return
	// }

	// Form values
	//email := r.FormValue("email")
    username:=r.FormValue("user_name")
	password := r.FormValue("user_psd")

	// Get database result
	idealpsd, err := model.UserByName(username)

	// Determine if user exists
	if err != nil {
		// Display error message
		log.Println(err)

	} else if idealpsd != password {
		fmt.Println("invalid psd!")
	} else {
			// Login successfully
			http.Redirect(w, r, "/", http.StatusFound)
			return
	}


	// Show the login page again
	LoginGet(w, r)
}

// LogoutGET clears the session and logs the user out
// func LogoutGet(w http.ResponseWriter, r *http.Request) {
// 	// Get session
// 	//sess := session.Instance(r)
//
// 	// If user is authenticated
// 	if sess.Values["id"] != nil {
// 	//	session.Empty(sess)
// 		sess.AddFlash(view.Flash{"Goodbye!", view.FlashNotice})
// 		sess.Save(r, w)
// 	}
//
// 	http.Redirect(w, r, "/", http.StatusFound)
// }
