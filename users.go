package rxclient

type User struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName" `
	LastName  string `json:"LastName" `
	NickName  string `json:"nickName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

func (c *rxClient) UserAll() ([]*User, error) {
	var out []*User
	_, err := c.rx.Send("users/all", nil, 5, &out, "any")
	return out, err

}
