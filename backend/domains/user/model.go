package user

// User Struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"nonzero, min=4, max=16"`
	Password string `json:"password" validate:"nonzero, min=4, max=16"`
	Alamat   string `json:"alamat"`
	Telp     string `json:"telp"`
	Status   string `json:"status"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type TokenUser struct {
	Token string
	User  User
}
