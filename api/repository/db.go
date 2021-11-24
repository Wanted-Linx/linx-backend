package repository

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Wanted-Linx/linx-backend/api/config"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/migrate"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func Connect() *ent.Client {
	drv, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=%s",
		config.Config.DB.PostgreSql.User,
		config.Config.DB.PostgreSql.Password,
		config.Config.DB.PostgreSql.Host,
		config.Config.DB.PostgreSql.Port,
		config.Config.DB.PostgreSql.DatabaseName,
		config.Config.DB.PostgreSql.TimeZone,
	))
	if err != nil {
		log.Fatal(err)
	}

	db := drv.DB()
	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	conn, err := db.Conn(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ent.Debug()
	ent.Log(func(i ...interface{}) {
		log.Debug(i...)
	})

	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(context.TODO(), migrate.WithDropIndex(true), migrate.WithForeignKeys(true)); err != nil {
		log.Fatalf("failed creating scheam resources: %v", err)
	}

	log.Info("Connect to database successfully")

	return client
}
