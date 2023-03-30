package redis

import (
	//  "aimi-landing-back-go/config"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Rdb *redis.Client
}

func (r *Redis) Connect() {
	r.Rdb = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// var (
// 	ErrTokenExpires        = errors.New("token expires")
// 	ErrAuthenticationFaild = errors.New("authentication failed")
// )

// var ctx = context.Background()

// func GenerateToken(id string, exp int) (string, error) {
// 	// check lại thư viên gen xem uuid có bị trùng không?
// 	uuid := uuid.NewString()
// 	expired := time.Duration(exp) * time.Minute
// 	tx := config.DB.Begin()
// 	err := tx.Create(&TokenAes{
// 		Id:      uuid,
// 		Expired: time.UnixMilli(makeTimestamp() + expired.Milliseconds()),
// 		UserId:  id,
// 	}).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return "", nil
// 	}
// 	errRedis := rdb.Set(ctx, uuid, id, expired).Err()

// 	if errRedis != nil {
// 		tx.Rollback()
// 		return "", nil
// 	}
// 	tx.Commit()
// 	return uuid, nil
// }

// func VerifyToken(token string) (*TokenAes, error) {
// 	val, err := rdb.Get(ctx, token).Result()
// 	if err != nil {
// 		return nil, ErrAuthenticationFaild
// 	}
// 	return &TokenAes{
// 		UserId: val,
// 	}, nil
// }

// func makeTimestamp() int64 {
// 	return time.Now().UnixNano() / int64(time.Millisecond)
// }
