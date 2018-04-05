package controller

import (
	"log"
	"net/http"
    //"encoding/json"
    //"fmt"
    //"os"
    "../view"
    "../model"
)

// RegisterGET displays the register page
func RegisterGet(w http.ResponseWriter, r *http.Request) {
	// Get session
//	sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "register"

	// Refill any form fields
//	view.Repopulate([]string{"first_name", "last_name", "email"}, r.Form, v.Data)
	v.RenderTemplate(w)
}

// RegisterPOST handles the registration form submission
func RegisterPost(w http.ResponseWriter, r *http.Request) {
	// Get session
//	sess := session.Instance(r)


	// Validate with required fields

	// Get form values
	userName := r.FormValue("user_name")
	password := r.FormValue("user_password")

	// Get database result
	_,err := model.UserByName(userName)

	if err != nil { // If success (no user exists with that email)
		model.UserCreate(userName, password)
		// Will only error if there is a problem with the query
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else if err == nil { // Catch all other errors
		log.Println(err)

	} else { // Else the user already exists

	}

	// Display the page
	RegisterGet(w, r)
}
