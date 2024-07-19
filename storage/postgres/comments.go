package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/google/uuid"
)

type CommentRepo struct {
	Db *sql.DB
}

func NewCommentRepo(Db *sql.DB) *CommentRepo {
	return &CommentRepo{Db: Db}
}

func (c *CommentRepo) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	req.Id = uuid.NewString()
	query := `insert into comments(id, memory_id, user_id, content) values($1, $2, $3, $4)`

	_, err := c.Db.ExecContext(ctx, query, req.Id, req.MemoryId, req.UserId, req.Content)
	if err != nil {
		log.Printf("Error while creating comment: %v\n", err)
		return nil, err
	}

	return &pb.AddCommentResponse{}, nil
}

func (c *CommentRepo) GetByMemoryId(ctx context.Context, req *pb.GetByIdMemoryRequest) (*pb.GetByIdMemoryResponse, error) {
	query := `select id, user_id, content, created_at from comments where memory_id = $1 and deleted_at = 0`

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
	fmt.Println(req.Id)
	query := `UPDATE comments SET deleted_at = $1 WHERE id = $2`

	_, err := c.Db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		log.Printf("Error while deleting comment: %v\n")
	}

	return &pb.DeleteCommentResponse{}, nil
}

func (c *CommentRepo) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	query := `update comments set memory_id = $1, content = $2 where id = $3 and deleted_at = 0`

	_, err := c.Db.ExecContext(ctx, query, req.MemoryId, req.Content, req.Id)
	if err != nil {
		log.Printf("Error while updating comment: %v\n")
	}

	return &pb.UpdateCommentResponse{}, nil
}

func (c *CommentRepo) GetById(ctx context.Context, req *pb.GetByCommentIdRequest) (*pb.GetByCommentIdResponse, error) {
	query := `select id, memory_id, user_id, content, created_at from comments where id = $1 and deleted_at = 0`

	row := c.Db.QueryRowContext(ctx, query, req.Id)

	var comment pb.Comment

	err := row.Scan(&comment.Id, &comment.MemoryId, &comment.UserId, &comment.Content, &comment.CreatedAt)
	if err != nil {
		log.Printf("Error while scanning comment: %v\n", err)
		return nil, err
	}

	return &pb.GetByCommentIdResponse{Comment: &comment}, nil
}

func (s *CommentRepo) GetAllCommets(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	query := `SELECT id, memory_id, user_id, content, created_at, updated_at FROM comments`
	conditions := []string{}
	args := []interface{}{}

	if req.GetUserId() != "" {
		conditions = append(conditions, "user_id = $1")
		args = append(args, req.GetUserId())
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY created_at DESC"

	if req.GetLimit() != "" {
		query += fmt.Sprintf(" LIMIT $%d", len(args)+1)
		args = append(args, req.GetLimit())
	}

	if req.GetOfset() != "" {
		query += fmt.Sprintf(" OFFSET $%d", len(args)+1)
		args = append(args, req.GetOfset())
	}

	rows, err := s.Db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*pb.Comment
	for rows.Next() {
		var comment pb.Comment
		if err := rows.Scan(&comment.Id, &comment.MemoryId, &comment.UserId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return &pb.GetCommentsResponse{Comment: comments}, nil
}
