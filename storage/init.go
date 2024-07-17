package storage

import (
	"context"

	pb "github.com/Exam4/4th-month-exam-Memory-Service/genproto"
)

type StorageI interface{
	Comment() CommentI
	Media() MediaI
}

type StorageMongo interface{
	Memory() MemoryI
}


type CommentI interface {
	AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error)
	GetByMemoryId(ctx context.Context, req *pb.GetByIdMemoryRequest) (*pb.GetByIdMemoryResponse, error)
	UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error)
	DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error)
	GetById(ctx context.Context, req *pb.GetByCommentIdRequest) (*pb.GetByCommentIdResponse, error)
	GetAllCommets(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error)
}

type MediaI interface {
	AddMedia(ctx context.Context, req *pb.AddMediaRequest) (*pb.AddMediaResponse, error)
	GetMediaByMemoryId(ctx context.Context, req *pb.GetMediaByMemoryIdRequest) (*pb.GetMediaByMemoryIdResponse, error)
	DeleteMedia(ctx context.Context, req *pb.DeleteMediaRequest) (*pb.DeleteMediaResponse, error)
	UpdateMedia(ctx context.Context, req *pb.UpdateMediaRequest) (*pb.UpdateMediaResponse, error)
	GetAllMedia(ctx context.Context, req *pb.GetAllMediaRequest) (*pb.GetAllMediaResponse, error)
}

type MemoryI interface {
	AddMemory(ctx context.Context, req *pb.AddMemoryRequest) (*pb.AddMemoryResponse, error)
	GetMemory(ctx context.Context, req *pb.GetMemoryRequest) (*pb.GetMemoryResponse, error)
	GetAllMemories(ctx context.Context, req *pb.GetAllMemoriesRequest) (*pb.GetAllMemoriesResponse, error)
	UpdateMemory(ctx context.Context, req *pb.UpdateMemoryRequest) (*pb.UpdateMemoryResponse, error)
	DeleteMemory(ctx context.Context, req *pb.DeleteMemoryRequest) (*pb.DeleteMemoryResponse, error)
	ShareMemory(ctx context.Context, req *pb.ShareMemoryRequest) (*pb.ShareMemoryResponse, error)
}
