package mongo

import (
	"context"
	"strconv"
	"time"

	pb "github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemoryRepo struct {
	db *mongo.Collection
}

func NewMemoryRepo(db *mongo.Database) *MemoryRepo {
	return &MemoryRepo{db: db.Collection("memories")}
}

func (s *MemoryRepo) AddMemory(ctx context.Context, req *pb.AddMemoryRequest) (*pb.AddMemoryResponse, error) {
	memory := &pb.Memory{
		Id:          uuid.NewString(),
		UserId:      req.GetUserId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Date:        req.GetDate(),
		Tags:        req.GetTags(),
		Latitude:    req.GetLatitude(),
		Longitude:   req.GetLongitude(),
		PlaceName:   req.GetPlaceName(),
		Privacy:     req.GetPrivacy(),
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	_, err := s.db.InsertOne(ctx, memory)
	if err != nil {
		return nil, err
	}

	return &pb.AddMemoryResponse{}, nil
}

func (s *MemoryRepo) GetMemory(ctx context.Context, req *pb.GetMemoryRequest) (*pb.GetMemoryResponse, error) {
	var memory pb.Memory
	err := s.db.FindOne(ctx, bson.M{"id": req.GetId()}).Decode(&memory)
	if err != nil {
		return nil, err
	}
	return &pb.GetMemoryResponse{Memory: &memory}, nil
}

func (s *MemoryRepo) GetAllMemories(ctx context.Context, req *pb.GetAllMemoriesRequest) (*pb.GetAllMemoriesResponse, error) {
	filter := bson.M{}

	if req.GetTags() != "" {
		filter["tags"] = bson.M{"$in": req.GetTags()}
	}
	if req.GetStartDate() != "" && req.GetEndDate() != "" {
		filter["date"] = bson.M{
			"$gte": req.GetStartDate(),
			"$lte": req.GetEndDate(),
		}
	}

	findOptions := options.Find()
	page, err := strconv.Atoi(req.GetPage())
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(req.GetLimit())
	if err != nil {
		limit = 10
	}

	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	cursor, err := s.db.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var memories []*pb.Memory
	for cursor.Next(ctx) {
		var memory pb.Memory
		err := cursor.Decode(&memory)
		if err != nil {
			return nil, err
		}
		memories = append(memories, &memory)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.GetAllMemoriesResponse{Memories: memories}, nil
}

func (s *MemoryRepo) UpdateMemory(ctx context.Context, req *pb.UpdateMemoryRequest) (*pb.UpdateMemoryResponse, error) {
	update := bson.M{}
	if req.GetTitle() != "" {
		update["title"] = req.GetTitle()
	}
	if req.GetTags() != nil {
		update["tags"] = req.GetTags()
	}
	_, err := s.db.UpdateOne(ctx, bson.M{"id": req.GetId()}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateMemoryResponse{}, nil
}

func (s *MemoryRepo) DeleteMemory(ctx context.Context, req *pb.DeleteMemoryRequest) (*pb.DeleteMemoryResponse, error) {
	_, err := s.db.DeleteOne(ctx, bson.M{"id": req.GetId()})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMemoryResponse{Message: "Memory deleted successfully"}, nil
}

func (s *MemoryRepo) ShareMemory(ctx context.Context, req *pb.ShareMemoryRequest) (*pb.ShareMemoryResponse, error) {
	update := bson.M{"$addToSet": bson.M{"tags": bson.M{"$each": req.GetSharedWith()}}}
	_, err := s.db.UpdateOne(ctx, bson.M{"id": req.GetMemoryId()}, update)
	if err != nil {
		return nil, err
	}
	return &pb.ShareMemoryResponse{}, nil
}
