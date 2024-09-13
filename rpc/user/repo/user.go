package repo

import (
	"cloud_tinamic/kitex_gen/base/user"
	"cloud_tinamic/pkg/pg"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type UserRepo interface {
	QueryAccount(account *user.Account) error
	QueryUser(account *user.Account, user *user.User, role *user.UserRole) error
	AddUser(account *user.Account, user *user.User, role *user.UserRole) error
}

type UserRepoImpl struct {
	*pg.PgPool
	dialect goqu.DialectWrapper
}

func NewUserRepoImpl(pool *pg.PgPool) UserRepo {
	return &UserRepoImpl{
		pool,
		goqu.Dialect("postgres"),
	}
}

func (db *UserRepoImpl) QueryUser(account *user.Account, user *user.User, role *user.UserRole) error {
	// 构建查询
	query := db.dialect.From(goqu.I("user_info.user").As("u")).
		Select(
			goqu.I("a.user_id").As("account_user_id"),
			goqu.I("u.user_id"),
			goqu.I("u.state"),
			goqu.I("u.name"),
			goqu.I("u.avatar"),
			goqu.I("u.cell_phone"),
			goqu.I("u.password"),
			goqu.I("ur.user_id").As("role_user_id"),
			goqu.I("ur.role_id"),
		).
		Join(
			goqu.I("user_info.account").As("a"),
			goqu.On(goqu.Ex{
				"u.user_id": goqu.I("a.user_id"),
			}),
		).
		Join(
			goqu.I("user_info.user_role").As("ur"),
			goqu.On(goqu.Ex{
				"ur.user_id": goqu.I("a.user_id"),
			}),
		).
		Where(goqu.Ex{
			"a.deleted":      false,
			"u.state":        true,
			"a.user_account": account.Username,
			"a.category":     account.UserCategory,
		})

	// 生成 SQL 查询语句
	sql, _, err := query.ToSQL()

	if err != nil {
		klog.Errorf("create sql error: %v", err)
		return err
	}

	err = db.QueryRow(context.Background(), sql).Scan(
		&account.UserId,
		&user.UserId, &user.State, &user.Name, &user.Avatar, &user.PhoneNumber, &user.Password,
		&role.UserId, &role.RoleCode)

	if err != nil {
		klog.Errorf("get user error: %v", err)
		return err
	}
	return nil
}

func (db *UserRepoImpl) AddUser(account *user.Account, u *user.User, role *user.UserRole) error {
	// 需要增加三张表 用户表、账号表和用户权限表
	addAccountSQL, _, err := db.dialect.Insert("user_info.account").
		Cols("user_id", "user_account", "category").
		Vals(goqu.Vals{account.UserId, account.Username, account.UserCategory}).
		ToSQL()
	if err != nil {
		klog.Errorf("create sql error: %v", err)
		return err
	}

	addUserSQL, _, err := db.dialect.Insert("user_info.user").
		Cols("user_id", "name", "avatar", "cell_phone", "salt", "password").
		Vals(goqu.Vals{u.UserId, u.Name, u.Avatar, u.PhoneNumber, u.Salt, u.Password}).
		ToSQL()
	if err != nil {
		klog.Errorf("insert user error: %v", err)
		return err
	}

	addUserRoleSQL, _, err := db.dialect.Insert("user_info.user_role").
		Cols("user_id", "role_id").
		Vals(goqu.Vals{role.UserId, role.RoleCode}).
		ToSQL()
	if err != nil {
		klog.Errorf("insert user_role error: %v", err)
		return err
	}

	ctx := context.Background()
	tx, err := db.Begin(ctx)
	if err != nil {
		klog.Errorf("begin transaction error: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			// 回滚事务
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				klog.Errorf("rollback transaction error: %v", rbErr)
			}
			return
		}
		// 提交事务
		if err = tx.Commit(ctx); err != nil {
			klog.Errorf("commit transaction error: %v", err)
		}
	}()

	// 执行插入操作
	_, err = tx.Exec(ctx, addAccountSQL)
	if err != nil {
		klog.Errorf("insert account error: %v", err)
		return err
	}

	_, err = tx.Exec(ctx, addUserSQL)
	if err != nil {
		klog.Errorf("insert user error: %v", err)
		return err
	}

	_, err = tx.Exec(ctx, addUserRoleSQL)
	if err != nil {
		klog.Errorf("insert user role error: %v", err)
		return err
	}
	return nil
}

func (db *UserRepoImpl) QueryAccount(account *user.Account) error {
	toSQL, _, err := db.dialect.From("user_info.account").
		Select("user_id", "category").
		Where(goqu.C("user_account").
			Eq(account.Username)).
		ToSQL()
	if err != nil {
		klog.Errorf("create sql error: %v", err)
		return err
	}
	err = db.QueryRow(context.Background(), toSQL).Scan(&account.UserId, &account.UserCategory)
	if err != nil {
		klog.Errorf("get account error: %v", err)
		return err
	}
	return nil
}
