package pack

import (
	"cloud_tinamic/kitex_gen/base/user"
	"cloud_tinamic/rpc/user/model"
)

// Users Convert model.User list to user_gorm.User list
func Users(models []*model.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

// User Convert model.User to user_gorm.User
func User(model *model.User) *user.User {
	if model == nil {
		return nil
	}
	return &user.User{
		UserId:      model.UserId,
		Name:        model.Name,
		Avatar:      string(model.Avatar),
		PhoneNumber: model.PhoneNumber,
	}
}
