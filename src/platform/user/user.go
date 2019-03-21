package user

type Adder interface {
	Add(user User)
}

type Getter interface {
	GetAll() []User
}

// User struct
type User struct {
	//Id    int32  `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

// Users list struct
type Users struct {
	Users []User
}

// New Users
func New() *Users {
	return &Users{
		Users: []User{},
	}
}

// Add user to Users
func (u *Users) Add(user User) {
	u.Users = append(u.Users, user)
}

// GetAll user in Users
func (u *Users) GetAll() []User {
	return u.Users
}
