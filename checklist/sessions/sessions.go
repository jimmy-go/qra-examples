// Package sessions contains gorilla sessions cookies.
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
package sessions

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jimmy-go/qra"
)

var (
	store *sessions.CookieStore

	errSessionNotFound = errors.New("sessions: not found")
	errInvalidUserID   = errors.New("sessions: invalid user id")
)

const (
	sessionName = "qra_example1"
)

// Configure sets gorilla sessions key.
func Configure(secret string) error {
	store = sessions.NewCookieStore([]byte(secret))
	return nil
}

// UserToken retrieve user token for session.
func UserToken(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	s, ok := session.Values["session_token"].(string)
	if !ok {
		return "", errSessionNotFound
	}
	return s, nil
}

// SetToken stores the session token in user cookies.
func SetToken(w http.ResponseWriter, r *http.Request, token string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["session_token"] = token
	return session.Save(r, w)
}

// UserID retrieve user token for session.
func UserID(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	s, ok := session.Values["user_id"].(string)
	if !ok {
		return "", errSessionNotFound
	}
	if len(s) < 1 {
		return "", errInvalidUserID
	}
	return s, nil
}

// SetUserID stores the user id in user cookies.
func SetUserID(w http.ResponseWriter, r *http.Request, userID string) error {
	if len(userID) < 1 {
		return errInvalidUserID
	}
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user_id"] = userID
	return session.Save(r, w)
}

// Delete deletes cookie session.
func Delete(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1

	username, _ := session.Values["user_id"].(string)
	token, _ := session.Values["session_token"].(string)
	log.Printf("Delete : username [%s] token [%s]", username, token)

	ctx := Ctx(username, token)
	err = qra.Close(ctx)
	if err != nil {
		return err
	}

	return session.Save(r, w)
}

// Login validates user credentials and generates a session.
func Login(username, password string) (string, error) {
	var token string

	ctx := Ctx(username, "")

	// identity: authentication.

	err := qra.Authenticate(ctx, password, &token)
	if err != nil {
		return "", err
	}
	log.Printf("Login : username [%v] token [%v]", ctx.Me(), token)

	// validation: identity permission for session on web admin.

	buf := bytes.NewBuffer([]byte{})
	err = qra.Search(ctx, buf, "session-on:web")
	if err != nil {
		return "", err
	}
	log.Printf("Login : buf [%v]", buf.String())

	return token, nil
}

// Ctx returns a type that satisfies qra.Identity interface.
func Ctx(username, token string) qra.Identity {
	return User{Username: username}
}

// User struct satisfies qra.Identity interface.
type User struct {
	Username string
	Token    string
}

// Me returns user id.
func (ctx User) Me() string {
	return ctx.Username
}

// Session returns user id.
func (ctx User) Session(dst interface{}) error {
	dst = ctx.Token
	return nil
}
