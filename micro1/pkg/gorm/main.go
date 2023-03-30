package gorm

import (
	"github.com/bav-demo/micro1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

type Gorm struct {
	c *config.Config
}

func (g *Gorm) NewDb() []*gorm.DB {
	g.c = &config.Config{}

	c, _ := g.c.NewConfig()

	fmt.Println(c)

	arrayConnections := make([]*gorm.DB, 0)

	var db *gorm.DB

	var err error

	if len(c.DB.Postgres) > 0 {
		for _, v := range c.DB.Postgres {

			dsn := "host=" + v.Host + " user=" + v.UserName + " password=" + v.Password + " dbname=" + v.Database + " port=" + v.Port + " sslmode=disable"

			// Connect to the database
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

			if err != nil {
				fmt.Printf("err postgres = %v \n", err)
			}

			arrayConnections = append(arrayConnections, db)

			db = &gorm.DB{}
		}
	}
	return arrayConnections
}
