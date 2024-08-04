package schema

type Character struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Alias       string `json:"alias"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type CharacterPage struct {
	ID       string    `json:"id"`
	CID      string    `json:"cid"`
	Likes    []User    `json:"likes"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID      string `json:"id"`
	UID     string `json:"uid"`
	Message string `json:"message"`
}

type User struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Lastname  string      `json:"lastname"`
	Username  string      `json:"username"`
	Image     string      `json:"image"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Favorites []Character `json:"favorites"`
}
