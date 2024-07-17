package main

import (
	"log"
	"net"

	"github.com/Exam4/4th-month-exam-Memory-Service/config"
	"github.com/Exam4/4th-month-exam-Memory-Service/genproto"
	"github.com/Exam4/4th-month-exam-Memory-Service/service"
	"github.com/Exam4/4th-month-exam-Memory-Service/storage/mongo"
	"github.com/Exam4/4th-month-exam-Memory-Service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	db, err := postgres.DbConnection()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	mongo, err := mongo.SetupMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal("error while listening: %v", err)
	}
	
	s := grpc.NewServer()
	genproto.RegisterMemoryServiceServer(s, service.NewMemoryService(mongo))
	genproto.RegisterMediaServiceServer(s, service.NewMediaService(db))
	genproto.RegisterCommentServiceServer(s, service.NewCommentService(db))

	
	log.Printf("Server started on port: %v", cfg.HTTPPort)
	if err := s.Serve(lis); err != nil {
		log.Fatal("error while serving: %v", err)
	}
}
