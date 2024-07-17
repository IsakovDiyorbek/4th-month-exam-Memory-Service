package service

import (
	"context"

	pb "github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/Exam4/4th-month-exam-Memory-Service/storage"
)




type MediaService struct{
	st storage.StorageI
	pb.UnimplementedMediaServiceServer
}

func NewMediaService(st storage.StorageI) *MediaService{
	return &MediaService{st: st}
}

func (s *MediaService) AddMedia(ctx context.Context, req *pb.AddMediaRequest) (*pb.AddMediaResponse, error){
	return s.st.Media().AddMedia(ctx, req)
}

func (s *MediaService) GetMediaByMemoryId(ctx context.Context, req *pb.GetMediaByMemoryIdRequest) (*pb.GetMediaByMemoryIdResponse, error){
	return s.st.Media().GetMediaByMemoryId(ctx, req)
}

func (s *MediaService) DeleteMedia(ctx context.Context, req *pb.DeleteMediaRequest) (*pb.DeleteMediaResponse, error){
	return s.st.Media().DeleteMedia(ctx, req)
}

func (s *MediaService) UpdateMedia(ctx context.Context, req *pb.UpdateMediaRequest) (*pb.UpdateMediaResponse, error){
	return s.st.Media().UpdateMedia(ctx, req)
}

func (s *MediaService) GetAllMedia(ctx context.Context, req *pb.GetAllMediaRequest) (*pb.GetAllMediaResponse, error){
	return s.st.Media().GetAllMedia(ctx, req)
}

