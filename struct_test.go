package json_compare

import "testing"

type user struct {
	name string
}

type group struct {
	*user
}

func (u *user) getName() string {
	return u.name
}

func (u *user) Init() {
	u.name = "yuuki"
}

func TestStruct(t *testing.T) {
	g := group{}
	print(g.user.getName())
}
