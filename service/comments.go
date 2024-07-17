package service



import (
	"context"

	pb "github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/Exam4/4th-month-exam-Memory-Service/storage"
)

type CommentService struct{
	st storage.StorageI
	pb.UnimplementedCommentServiceServer
}

func NewCommentService(st storage.StorageI) *CommentService{
	return &CommentService{st: st}
}

func (s *CommentService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error){
	return s.st.Comment().AddComment(ctx, req)
}

func (s *CommentService) GetByMemoryId(ctx context.Context, req *pb.GetByIdMemoryRequest) (*pb.GetByIdMemoryResponse, error){
	return s.st.Comment().GetByMemoryId(ctx, req)
}

func (s *CommentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error){
	return s.st.Comment().DeleteComment(ctx, req)
}

func (s *CommentService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error){
	return s.st.Comment().UpdateComment(ctx, req)
}

func (s *CommentService) GetById(ctx context.Context, req *pb.GetByCommentIdRequest) (*pb.GetByCommentIdResponse, error){
	return s.st.Comment().GetById(ctx, req)
}

func (s *CommentService) GetAllCommets(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error){
	return s.st.Comment().GetAllCommets(ctx, req)
}