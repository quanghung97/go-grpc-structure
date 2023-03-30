package PkgMiddleWare

import (
	"context"
	"errors"
	"time"

	"github.com/bav-demo/micro1/internal/models/authenticate"
	initConnect "github.com/bav-demo/micro1/server/init"
	"github.com/google/uuid"
)

var (
	ErrTokenExpires        = errors.New("token expires")
	ErrAuthenticationFaild = errors.New("authentication failed")
)

var ctx = context.Background()

var connection = initConnect.GetInstance()

func GenerateToken(id string, exp int) (string, error) {

	uuid := uuid.NewString()
	expired := time.Duration(exp) * time.Minute
	tx := connection.Db.Begin()
	err := tx.Create(&authenticate.TokenAes{
		Id:      uuid,
		Expired: time.UnixMilli(makeTimestamp() + expired.Milliseconds()),
		UserId:  id,
	}).Error
	if err != nil {
		tx.Rollback()
		return "", errors.New("Gen token error.")
	}
	errRedis := connection.Redis.Rdb.Set(ctx, uuid, id, expired).Err()

	if errRedis != nil {
		tx.Rollback()
		return "", errors.New("Gen token error.")
	}
	tx.Commit()
	return uuid, nil
}

func VerifyToken(token string) (*authenticate.TokenAes, error) {
	val, err := connection.Redis.Rdb.Get(ctx, token).Result()
	if err != nil {
		return nil, ErrAuthenticationFaild
	}
	return &authenticate.TokenAes{
		UserId: val,
	}, nil
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
