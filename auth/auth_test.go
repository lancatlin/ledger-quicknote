package auth

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/securecookie"
)

type User struct {
	user   string
	pass   string
	hashed string
}

func TestHtpasswdSuccess(t *testing.T) {
	hashKey := securecookie.GenerateRandomKey(32)
	path := "/tmp/.htpasswd"
	user1 := User{
		user:   "user",
		pass:   "password",
		hashed: "$2a$14$SQSscaF4fVO3e5dp2/.VPuVQDPKqxSagLQnN6OncTRtoQw0ie9ByK",
	}
	err := os.WriteFile(path,
		[]byte(fmt.Sprintf("%s:%s", user1.user, user1.hashed)), 0640)
	if err != nil {
		t.Fatal(err)
	}
	store, err := New(path, hashKey)
	if err != nil {
		t.Fatal(err)
	}
	token, err := store.Login(user1.user, user1.pass)
	if err != nil {
		t.Error(err)
	}

	session, err := store.Verify(token)
	if err != nil {
		t.Error(err)
	}
	if session.User != user1.user {
		t.Fatalf("expected %s, got %s", user1.user, session.User)
	}

	user2 := User{
		user: "foo",
		pass: "bar",
	}
	err = store.Register(user2.user, user2.pass)
	if err != nil {
		t.Error(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	for _, u := range []User{user1, user2} {
		if !strings.Contains(string(data), u.user) {
			t.Errorf("%s not found in htpasswd file: %s", u.user, string(data))
		}
	}
	err = store.Remove(user1.user)
	if err != nil {
		t.Error(err)
	}

	data, err = os.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(data), user1.user) {
		t.Errorf("%s is found in htpasswd file but should be removed: %s", user1.user, string(data))
	}
}
