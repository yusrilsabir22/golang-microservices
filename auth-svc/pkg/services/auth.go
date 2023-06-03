package services

import (
	"context"
	"net/http"

	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/db"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/models"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/pb"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/utils"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

// RegisterAuth	godoc
// @Summary		Register user
// @Description	Save user data in DB
// @Param		register body pb.RegisterRequest true "Register User"
// @Tags		Auth
// @Success		200 {object} pb.RegisterResponse
// @Failure     401  {object}  pb.LoginResponse
// @Failure     403  {object}  pb.LoginResponse
// @Router		/auth/register [post]
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if !utils.IsValidMSISDN(req.Msisdn) || req.Name == "" || req.Password == "" || req.Username == "" {
		return &pb.RegisterResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		}, nil
	}

	if result := s.H.DB.Where(&models.User{MSISDN: req.Msisdn}).Or(&models.User{Username: req.Username}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status:  http.StatusBadRequest,
			Message: "MSISDN already exists",
		}, nil
	}

	user.MSISDN = req.Msisdn
	user.Name = req.Name
	user.Username = req.Username
	user.Password = utils.HashPassword(req.Password)

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Message: "Success",
		Status:  http.StatusCreated,
	}, nil
}

// LoginAuth	godoc
// @Summary		Login user
// @Description	Check user data at DB and generate token
// @Param		login body pb.LoginRequest true "Login User"
// @Tags		Auth
// @Success		200 {object} pb.LoginResponse
// @Failure     401  {object}  pb.LoginResponse
// @Router		/auth/login [post]
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if !utils.IsValidMSISDN(req.Msisdn) {
		return &pb.LoginResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		}, nil
	}

	if req.Msisdn == "" || req.Password == "" {
		return &pb.LoginResponse{
			Message: "User not found",
			Status:  http.StatusUnauthorized,
			Token:   "",
		}, nil
	}
	if result := s.H.DB.First(&user, &models.User{MSISDN: req.Msisdn}); result.Error != nil {
		return &pb.LoginResponse{
			Message: "User not found",
			Status:  http.StatusUnauthorized,
			Token:   "",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)
	if !match {
		return &pb.LoginResponse{
			Status:  http.StatusUnauthorized,
			Message: "User not found",
			Token:   "",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status:  http.StatusAccepted,
		Message: "success",
		Token:   token,
	}, nil
}

// ValidateAuth	godoc
// @Summary		Get private claims data of the token
// @Description	Validate the token and return claims data
// @Param		Authorization header string true "Authrization" Authorzation(Bearer random_value)
// @Tags		Auth
// @Success		200 {object} pb.ValidateResponse
// @Failure		401 {object} pb.ValidateResponse
// @Router		/auth/validate [post]
func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    &pb.ClaimResponse{},
		}, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{ID: claims.Id}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status:  http.StatusNotFound,
			Message: "User not found",
			Data:    &pb.ClaimResponse{},
		}, nil
	}

	return &pb.ValidateResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &pb.ClaimResponse{
			UserId: user.ID.String(),
			StandardClaims: &pb.StandardClaims{
				Id:        claims.Id.String(),
				Audience:  claims.Audience,
				Issuer:    claims.Issuer,
				Subject:   claims.Subject,
				ExpiresAt: claims.ExpiresAt,
				IssuedAt:  claims.IssuedAt,
				NotBefore: claims.NotBefore,
			},
		},
	}, nil
}
