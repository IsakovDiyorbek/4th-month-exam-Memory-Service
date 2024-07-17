package service

import (
	"context"

	"github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/Exam4/4th-month-exam-Memory-Service/storage"
)

type MemoryService struct {
	st storage.StorageMongo
	genproto.UnimplementedMemoryServiceServer
}

func NewMemoryService(st storage.StorageMongo) *MemoryService {
	return &MemoryService{st: st}
}

func (m *MemoryService) AddMemory(ctx context.Context, req *genproto.AddMemoryRequest) (*genproto.AddMemoryResponse, error) {
	return m.st.Memory().AddMemory(ctx, req)
}


func (m *MemoryService) GetMemory(ctx context.Context, req *genproto.GetMemoryRequest) (*genproto.GetMemoryResponse, error) {
	return m.st.Memory().GetMemory(ctx, req)
}

func (m *MemoryService) GetAllMemories(ctx context.Context, req *genproto.GetAllMemoriesRequest) (*genproto.GetAllMemoriesResponse, error) {
	return m.st.Memory().GetAllMemories(ctx, req)
}

func (m *MemoryService) UpdateMemory(ctx context.Context, req *genproto.UpdateMemoryRequest) (*genproto.UpdateMemoryResponse, error) {
	return m.st.Memory().UpdateMemory(ctx, req)
}

func (m *MemoryService) DeleteMemory(ctx context.Context, req *genproto.DeleteMemoryRequest) (*genproto.DeleteMemoryResponse, error) {
	return m.st.Memory().DeleteMemory(ctx, req)
}

func (m *MemoryService) ShareMemory(ctx context.Context, req *genproto.ShareMemoryRequest) (*genproto.ShareMemoryResponse, error) {
	return m.st.Memory().ShareMemory(ctx, req)
}