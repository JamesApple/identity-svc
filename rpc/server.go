package rpc

import (
	"context"
	"fmt"
	"log"
	"root"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	root.UnimplementedIdentityServer
}

func (s *Server) CreateAccount(ctx context.Context, ar *root.AccountRequest) (*root.AccountResponse, error) {
	log.Print(ar)
	v := &errdetails.BadRequest_FieldViolation{
		Field:       "username",
		Description: "its bad",
	}

	br := &errdetails.BadRequest{}
	br.FieldViolations = append(br.FieldViolations, v)
	st, err := status.New(codes.InvalidArgument, "invalid").WithDetails(br)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	}

	return nil, st.Err()
}
