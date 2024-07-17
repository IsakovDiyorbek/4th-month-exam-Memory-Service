package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Exam4/4th-month-exam-Memory-Service/storage"
	"github.com/Exam4/4th-month-exam-Memory-Service/config"

	_ "github.com/lib/pq"

)

type Storage struct {
	Db        *sql.DB
	Commets storage.CommentI
	Medias storage.MediaI
	
}

func DbConnection() (storage.StorageI, error) {
	cfg := config.Load()
	con := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresPort)
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal("Error while db connection", err)
		return nil, nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error while db ping connection", err)
		return nil, nil

	}
	return &Storage{Db: db}, nil

}




func (s *Storage) Comment() storage.CommentI {
	if s.Commets == nil {
		s.Commets = &CommentRepo{s.Db}
	}
	return s.Commets
}

func (s *Storage) Media() storage.MediaI {
	if s.Medias == nil {
		s.Medias = &MediaRepo{s.Db}
	}
	return s.Medias
}
