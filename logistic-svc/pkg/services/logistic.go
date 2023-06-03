package services

import (
	"context"
	"net/http"

	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/db"
	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/models"
	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedLogisticServiceServer
}

// CreateLogistic	godoc
// @Summary			Create a new logistic data
// @Description		Create a new logistic data and return a message
// @Param			create body pb.CreateLogisticRequest true "Create a new logistic data"
// @Param			Authorization header string true "Authorzation(Bearer random_value)"
// @Tags			Logistic
// @Success			200 {object} pb.CreateLogisticResponse
// @Failure			401 {object} string "Unauthorized"
// @Router			/logistic [post]
func (s *Server) CreateLogistic(ctx context.Context, req *pb.CreateLogisticRequest) (*pb.CreateLogisticResponse, error) {
	var logistic models.Logistic

	if req.Amount == 0 || req.LogisticName == "" || req.DestinationName == "" || req.OriginName == "" || req.Duration == "" {
		return &pb.CreateLogisticResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid paramater",
		}, nil
	}

	logistic.LogisticName = req.LogisticName
	logistic.Amount = int(req.Amount)
	logistic.DestinationName = req.DestinationName
	logistic.OriginName = req.OriginName
	logistic.Duration = req.Duration

	if result := s.H.DB.Create(&logistic); result.Error != nil {
		return &pb.CreateLogisticResponse{
			Status:  http.StatusConflict,
			Message: "Duplicates",
		}, nil
	}

	return &pb.CreateLogisticResponse{
		Status:  http.StatusCreated,
		Message: "success",
	}, nil
}

// FindLogistic		godoc
// @Summary			Find a logistic data
// @Description		Get a logistic based on given parameter
// @Param			origin_name query string true "test example"
// @Param			destionation_name query string true "test example"
// @Param			Authorization header string true "Authorzation(Bearer random_value)"
// @Tags			Logistic
// @Success			200 {object} pb.FindOneResponse
// @Failure			401 {object} string "Unauthorized"
// @Router			/logistic [get]
func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var logistic models.Logistic

	if req.DestinationName == "" || req.OriginName == "" {
		return &pb.FindOneResponse{
			Status:  http.StatusNotFound,
			Message: "no data",
		}, nil
	}

	if result := s.H.DB.Where(&models.Logistic{DestinationName: req.DestinationName, OriginName: req.OriginName}).First(&logistic); result.Error != nil {
		return &pb.FindOneResponse{
			Status:  http.StatusNotFound,
			Message: "no data",
		}, nil
	}

	data := &pb.FindOneData{
		LogisticName:    logistic.LogisticName,
		Amount:          int64(logistic.Amount),
		DestinationName: logistic.DestinationName,
		OriginName:      logistic.OriginName,
		Duration:        logistic.Duration,
	}

	return &pb.FindOneResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	}, nil
}
