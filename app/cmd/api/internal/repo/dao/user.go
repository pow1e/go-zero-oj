package dao

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/query"
	"gorm.io/gorm"
)

var RecordNotFound = errors.New("查询数据不存在")

type IUser interface {
	FindByUserName(ctx context.Context, username string) (*model.User, error)
}

type UserDao struct {
	*query.Query
}

func NewUserDao(query *query.Query) *UserDao {
	return &UserDao{query}
}

func (dao *UserDao) FindByUserName(ctx context.Context, username string) (*model.User, error) {
	u := dao.User
	user, err := u.WithContext(ctx).
		Where(u.Name.Eq(username)).
		First()
	switch {
	case err == nil:
		return user, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, RecordNotFound
	default:
		return nil, err
	}
}
