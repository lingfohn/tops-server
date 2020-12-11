package storage

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lingfohn/lime/ent"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var client *ent.Client

var once sync.Once

func InitClient() {
	var err error
	client, err = ent.Open("mysql",
		viper.GetString("mysql.source"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
func CloseClient() {
	client.Close()
}
func GetDB() *ent.Client {
	return client
}
