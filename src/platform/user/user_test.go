package user

import (
	"testing"
)

func TestAdd(t *testing.T) {
	users := New()
	users.Add(User{"Cleber", "cleber@gmail.com"})

	if len(users.Users) != 1 {
		t.Errorf("User não foi adicionado")
	}
}

func TestGetAll(t *testing.T) {
	users := New()
	users.Add(User{})
	results := users.GetAll()
	if len(results) != 1 {
		t.Errorf("User não foi adicionado")
	}
}
