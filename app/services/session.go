package services

import (
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	Name  = "_alloydev"
	Store = sessions.NewCookieStore(
		[]byte(envy.Get("HASH_KEY", string(securecookie.GenerateRandomKey(64)))),
		[]byte(envy.Get("BLOCK_KEY", string(securecookie.GenerateRandomKey(32)))),
	)
)

type key int

const (
	SessKey key = 0
)

func Session(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, Name)
	return session
}

func EmptySession(sess *sessions.Session) {
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}

func init() {
	Store.Options.HttpOnly = true
	Store.MaxAge(5 * 86400)
}
