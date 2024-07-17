package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/Exam4/4th-month-exam-Memory-Service/helper"
	"github.com/google/uuid"
)

type MediaRepo struct {
	db *sql.DB
}

func NewMediaRepo(db *sql.DB) *MediaRepo {
	return &MediaRepo{db: db}
}


// Mediya yaratish
func (m *MediaRepo) AddMedia(ctx context.Context, req *pb.AddMediaRequest) (*pb.AddMediaResponse, error) {
	query := `insert into media(id, memory_id, type, url) values($1, $2, $3, $4)`

	req.MediaId = uuid.NewString()
	_, err := m.db.ExecContext(ctx, query, req.MediaId, req.MemoryId, req.Type, req.Url)
	if err != nil {
		log.Printf("Error while creating borrower: %v\n", err)
		return nil, err
	}

	return &pb.AddMediaResponse{}, nil
}


// Memory ni medialarini korish
func (m *MediaRepo) GetMediaByMemoryId(ctx context.Context, req *pb.GetMediaByMemoryIdRequest) (*pb.GetMediaByMemoryIdResponse, error) {
	query := `select id, memory_id, type, url, created_at from media where memory_id = $1`

	rows, err := m.db.QueryContext(ctx, query, req.MemoryId)
	if err != nil {
		log.Printf("Error while getting media: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	var media []*pb.Media

	for rows.Next() {
		var m pb.Media
		err = rows.Scan(&m.Id, &m.MemoryId, &m.Type, &m.Url, &m.CreatedAt)
		if err != nil {
			log.Printf("Error while scanning media: %v\n", err)
			return nil, err
		}
		media = append(media, &m)
	}

	return &pb.GetMediaByMemoryIdResponse{Media: media}, nil
}


// Mediani chop etish
func (m *MediaRepo) DeleteMedia(ctx context.Context, req *pb.DeleteMediaRequest) (*pb.DeleteMediaResponse, error) {
	query := `update set deleted_at = $1 where id = $2`

	_, err := m.db.ExecContext(ctx, query, time.Now().Unix(), req.MediaId)
	if err != nil {
		log.Printf("Error while deleting media: %v\n", err)
		return nil, err
	}

	return &pb.DeleteMediaResponse{}, nil
}


// Malumotlarni o'zgartirish
func (m *MediaRepo) UpdateMedia(ctx context.Context, req *pb.UpdateMediaRequest) (*pb.UpdateMediaResponse, error) {
	query := `update media set memory_id = $1, type = $1, url = $2 where id = $3`

	_, err := m.db.ExecContext(ctx, query, req.MemoryId, req.Type, req.Url, req.Id)
	if err != nil {
		log.Printf("Error while updating media: %v\n", err)
		return nil, err
	}

	return &pb.UpdateMediaResponse{}, nil
}

// Filter va Malumotlarini korish
func (m *MediaRepo) GetAllMedia(ctx context.Context, req *pb.GetAllMediaRequest) (*pb.GetAllMediaResponse, error) {
	query := `select id, memory_id, type, url, created_at from media`

	param := make(map[string]interface{})
	filter := ` where deleted_at = 0`

	if len(req.Type) > 0 {
		param["type"] = req.Type
		filter += ` and type = :type`
	}

	if len(req.Url) > 0 {
		param["url"] = req.Url
		filter += ` and url = :url`
	}

	query += filter

	query, arr := helper.ReplaceQueryParams(query, param)

	rows, err := m.db.QueryContext(ctx, query, arr...)
	if err != nil{
		log.Printf("Error while getting media: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var media []*pb.Media

	for rows.Next() {
		var m pb.Media
		err := rows.Scan(&m.Id, &m.MemoryId, &m.Type, &m.Url, &m.CreatedAt)
		if err != nil{
			log.Printf("Error while scan media: %v\n", err)
			return nil, err
		}
	
		media = append(media, &m)

	}

	return &pb.GetAllMediaResponse{Media: media}, nil
}

