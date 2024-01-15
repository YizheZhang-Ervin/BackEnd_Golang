package dao

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/database"
	"user/internal/model"
)

type UserDao struct {
	*database.DBConn
}

var cacheUserIdPrefix = "cache:user:id:"

func (d *UserDao) FindById(ctx context.Context, id int64) (user *model.User, err error) {
	user = &model.User{}
	query := fmt.Sprintf("select * from %s where id = ?", user.TableName())
	userIdKey := fmt.Sprintf("%s:%d", cacheUserIdPrefix, id)
	err = d.ConnCache.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	return
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (name,gender) values (?,?)", user.TableName())
	result, err := d.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}
