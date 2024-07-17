package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/Exam4/4th-month-exam-Memory-Service/helper"
)

type CommentRepo struct {
	Db *sql.DB
}

func NewCommentRepo(Db *sql.DB) *CommentRepo {
	return &CommentRepo{Db: Db}
}

func (c *CommentRepo) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	query := `insert into comments(id, memory_id, user_id, content) values($1, $2, $3, $4)`

	_, err := c.Db.ExecContext(ctx, query, req.CommentId, req.MemoryId, req.UserId, req.Content)
	if err != nil {
		log.Printf("Error while creating comment: %v\n", err)
		return nil, err
	}

	return &pb.AddCommentResponse{}, nil
}

func (c *CommentRepo) GetByMemoryId(ctx context.Context, req *pb.GetByIdMemoryRequest) (*pb.GetByIdMemoryResponse, error) {
	query := `select id, user_id, content, created_at from comments where memory_id = $1`

	rows, err := c.Db.QueryContext(ctx, query, req.MemoryId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comment []*pb.ByMemoryId
	for rows.Next() {
		var c pb.ByMemoryId
		err := rows.Scan(&c.Id, &c.UserId, &c.Content, &c.CreatedAt)
		if err != nil {
			log.Printf("Error while scanning comment: %v\n", err)
			return nil, err
		}
		comment = append(comment, &c)
	}

	return &pb.GetByIdMemoryResponse{Comments: comment}, nil
}

func (c *CommentRepo) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	query := `update comments set deleted_at = $1 where id = $2`

	_, err := c.Db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		log.Printf("Error while deleting comment: %v\n")
	}

	return &pb.DeleteCommentResponse{}, nil
}

func (c *CommentRepo) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	query := `update comments set memory_id = $1, content = $3 where id = $4 and deleted_at = 0`

	_, err := c.Db.ExecContext(ctx, query, req.MemoryId, req.Content, req.CommentId)
	if err != nil {
		log.Printf("Error while updating comment: %v\n")
	}

	return &pb.UpdateCommentResponse{}, nil
}

func (c *CommentRepo) GetById(ctx context.Context, req *pb.GetByCommentIdRequest) (*pb.GetByCommentIdResponse, error) {
	query := `select id, memory_id, user_id, content, created_at from comments where id = $1 and deleted_at = 0`

	row := c.Db.QueryRowContext(ctx, query, req.CommentId)

	var comment pb.Comment

	err := row.Scan(&comment.CommentId, &comment.MemoryId, &comment.UserId, &comment.Content, &comment.CreatedAt)
	if err != nil {
		log.Printf("Error while scanning comment: %v\n")
		return nil, err
	}

	return &pb.GetByCommentIdResponse{Comment: &comment}, nil
}

func (c *CommentRepo) GetAllCommets(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	query := `select id, memory_id, user_id, content, created_at, updated_at from comments`

	param := make(map[string]interface{})
	filter := ` where deleted_at = 0`

	if len(req.UserId) > 0 {
		param["user_id"] = req.UserId
		filter += ` and user_id = :user_id`
	}


	query += filter

	query, arr := helper.ReplaceQueryParams(query, param)

	rows, err := c.Db.QueryContext(ctx, query, arr...)
	if err != nil{
		log.Printf("Error while getting comment: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var comment []*pb.Comment

	for rows.Next() {
		var c pb.Comment
		err := rows.Scan(&c.CommentId, &c.MemoryId, &c.UserId, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil{
			log.Printf("Error while scan comment: %v\n", err)
			return nil, err
		}
		comment = append(comment, &c)

	}

	return &pb.GetCommentsResponse{Comment: comment}, nil
}


