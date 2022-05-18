package pack

import (
	"easy_note/cmd/userdemo/dal/db"
	"easy_note/cmd/userdemo/kitex_gen/userdemo"
)

func User(u *db.User) *userdemo.User {
	if u == nil {
		return nil
	}
	return &userdemo.User{
		UserId:   int64(u.ID),
		UserName: u.UserName,
		Avatar:   "test",
	}
}

func Users(us []*db.User) []*userdemo.User {
	users := make([]*userdemo.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
