package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/settings"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
func main() {
	godotenv.Load()
	client := open("postgresql://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/test")
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//create entries
	u := client.User.Create().SetAge(18).SetName("John").SaveX(context.Background())
	l := client.List.Create().SetUserID(u.ID).SaveX(context.Background())
	s := client.Settings.Create().SetUserID(u.ID).SetListID(l.ID).SetTest("create").SaveX(context.Background())

	//try to update
	client.Settings.Update().Where(settings.ID(s.ID)).SetTest("update").SaveX(context.Background())

	//check new updated settings
	newSettings := client.Settings.GetX(context.Background(), s.ID)
	fmt.Println(newSettings)

}
