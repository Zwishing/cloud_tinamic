package repo

import (
	"cloud_tinamic/kitex_gen/base/user"
	"cloud_tinamic/rpc/user/model"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type UserRepo interface {
	QueryUserByAccount(account string, category user.UserCategory) (*model.User, error)
	QueryUserById(userId string) (*model.User, error)
	AddUser(account *user.Account, user *user.User) (bool, error)
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepoImpl(DB *gorm.DB) UserRepo {
	return &UserRepoImpl{DB: DB}
}

func (u *UserRepoImpl) QueryUserByAccount(account string, category user.UserCategory) (*model.User, error) {
	var usr model.User
	err := u.DB.Table("user_info.user").
		Joins("JOIN user_info.account ON user_info.user.user_id = user_info.account.user_id").
		Where("user_info.account.user_account = ? AND user_info.account.category = ?", account, category).
		First(&usr).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		klog.Errorf("failed to query user by account: %s, category: %d, error: %v", account, category, err)
		return nil, err
	}
	return &usr, nil
}

func (u *UserRepoImpl) QueryUserById(userId string) (*model.User, error) {
	var usr model.User

	err := u.DB.Table("user_info.user").
		Where("user_id = ?", userId).
		First(&usr).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil, nil if user not found
		}
		klog.Errorf("failed to query user by ID %s: %v", userId, err)
		return nil, err
	}

	return &usr, nil
}

func (u *UserRepoImpl) AddUser(account *user.Account, user *user.User) (bool, error) {
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		// Create user
		usr := model.User{
			UserId:      user.UserId,
			Name:        user.Name,
			Avatar:      []byte(user.Avatar),
			PhoneNumber: user.PhoneNumber,
			Salt:        user.Salt,
			Password:    user.Password,
		}
		if err := tx.Table("user_info.user").Create(&usr).Error; err != nil {
			klog.Errorf("failed to create user: %v", err)
			return err
		}

		// Create account
		a := model.Account{
			UserId:      account.UserId,
			UserAccount: account.Username,
			Category:    int64(account.UserCategory),
		}
		if err := tx.Table("user_info.account").Create(&a).Error; err != nil {
			klog.Errorf("failed to create account: %v", err)
			return err
		}

		return nil
	})

	if err != nil {
		klog.Errorf("transaction failed: %v", err)
		return false, err
	}
	return true, nil
}
