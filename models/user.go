package models

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Jane Doe"},
}

func GetAllUsers() ([]User, error) {
	return users, nil
}

func GetUser(id string) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, ErrNotFound
}

func CreateUser(user User) error {
	users = append(users, user)
	return nil
}

func UpdateUser(id string, user User) error {
	for i, u := range users {
		if u.ID == id {
			users[i] = user
			return nil
		}
	}
	return ErrNotFound
}