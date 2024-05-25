package main

import (
	"database/sql"
	"go_eduhub/api"
	db "go_eduhub/db"
	"go_eduhub/model"
	repo "go_eduhub/repository"
	"log"

	_ "embed"

	_ "github.com/lib/pq"
)

func Reset(db *sql.DB, table string) error {
	_, err := db.Exec("TRUNCATE " + table)
	if err != nil {
		return err
	}

	_, err = db.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
	}

	dbConn, err := db.Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	dbConn.AutoMigrate(&model.Student{}, &model.Course{})

	studentRepo := repo.NewStudentRepo(dbConn)
	courseRepo := repo.NewCourseRepo(dbConn)

	mainAPI := api.NewAPI(studentRepo, courseRepo)
	mainAPI.Start()
}
