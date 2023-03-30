package initConnect

import (
	"github.com/bav-demo/micro1/internal/models/authenticate"
	"github.com/bav-demo/micro1/internal/models/society"
	g_ "github.com/bav-demo/micro1/pkg/gorm"
	"github.com/bav-demo/micro1/pkg/redis"
	"gorm.io/gorm"
)

var connection *Connect

type Connect struct {
	Db *gorm.DB

	Redis *redis.Redis
}

func (s *Connect) NewConnect() {
	// init gorm
	g := g_.Gorm{}
	// connect DB
	s.Db = g.NewDb()[0]
	// connect redis
	s.Redis = &redis.Redis{}
	s.Redis.Connect()
	// migrate DB
	s.Db.AutoMigrate(society.Post{}, society.Comment{}, authenticate.TokenAes{})
}

func GetInstance() *Connect {
	if connection == nil {
		connect := Connect{}
		connect.NewConnect()
		connection = &connect
		return &connect
	}
	return connection
}
