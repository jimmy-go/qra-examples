// Package session contains Login and Logout methods
// integrated with qra.Manager
//
// MIT License
//
// Copyright (c) 2016 Angel Del Castillo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package session

import (
	"log"
	"net/http"

	"github.com/jimmy-go/qra"
	"github.com/jimmy-go/qra-examples/checklist/sessions"
	"gopkg.in/jimmy-go/srest.v0"
)

// Index endpoint /login GET
func Index(w http.ResponseWriter, r *http.Request) {
	v := map[string]interface{}{}

	err := srest.Render(w, "login.html", v)
	if err != nil {
		log.Printf("Index : Render : err [%s]", err)
	}
}

// Login endpoint /login POST
func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Login : parse form : err [%s]", err)
	}

	u := r.Form.Get("username")
	p := r.Form.Get("password")
	log.Printf("Login : username [%s] password [%s]", u, p)

	// qra.Authenticate calls qra.DefaultManager.Authentication.Authenticate method
	token, err := sessions.Login(u, p)
	if err != nil {
		log.Printf("Login : err [%s]", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	E("set token", sessions.SetToken(w, r, token))
	E("set user id", sessions.SetUserID(w, r, u))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout endpoint /logout GET
func Logout(w http.ResponseWriter, r *http.Request) {
	// take session from somewhere
	sessionID := r.Header.Get("Authorization")

	userCtx := sessions.Ctx(sessionID)

	err := qra.Close(userCtx)
	if err != nil {
		log.Printf("Logout : err [%s]", err)
	}

	E("delete cookie session", sessions.Delete(w, r))

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// E func logs errors if exists.
func E(name string, err error) {
	if err != nil {
		log.Printf("%s : err [%s]", name, err)
	}
}
