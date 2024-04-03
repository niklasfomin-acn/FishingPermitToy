package api

import (
	"context"
	"log"
	"net/http"
	"rest-backend/storage"
	admin "rest-backend/transport"
	"rest-backend/types"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	listenAddr string
	handlers   map[string]func(http.ResponseWriter, *http.Request)
	store      storage.Storage
	grpcServer *grpc.Server
}

type AdminServer struct {
	admin.UnimplementedAdminServiceServer
	store storage.Storage
}

func NewAdminServer(store storage.Storage) *AdminServer {
	return &AdminServer{store: store}
}

// This is the constructor function for the new server
func NewServer(listenAddr string, handlers map[string]func(w http.ResponseWriter, r *http.Request), store storage.Storage) *Server {
	s := grpc.NewServer()
	as := NewAdminServer(store)
	admin.RegisterAdminServiceServer(s, as)

	return &Server{
		listenAddr: listenAddr,
		handlers:   handlers,
		store:      store,
		grpcServer: s,
	}
}

func (s *Server) Start() error {
	log.Printf("Backend Server listening on Port 3000\n")
	for route, handler := range s.handlers {
		http.HandleFunc(route, handler)
	}
	return http.ListenAndServe(s.listenAddr, nil)
}

func convertToAdminCitizenPermits(citizenPermits []types.CitizenPermit) []*admin.CitizenPermit {
	adminCitizenPermits := make([]*admin.CitizenPermit, len(citizenPermits))
	for i, cp := range citizenPermits {
		adminCitizenPermits[i] = &admin.CitizenPermit{
			PassportNumber:   cp.PassportNumber,
			Surname:          cp.Surname,
			GivenNames:       cp.GivenNames,
			DateOfBirth:      cp.DateOfBirth,
			PlaceOfBirth:     cp.PlaceOfBirth,
			Gender:           cp.Gender,
			Nationality:      cp.Nationality,
			DateOfIssue:      cp.DateOfIssue,
			ExpiryDate:       cp.ExpiryDate,
			IssuingAuthority: cp.IssuingAuthority,
			PermitDate:       timestamppb.New(cp.PermitDate),
			PermitLocation:   cp.PermitLocation,
			PermitType:       cp.PermitType,
			PermitState:      cp.PermitState,
		}
	}
	return adminCitizenPermits
}

func (s *Server) FetchAll(ctx context.Context, req *admin.FetchAllRequest) (*admin.FetchAllResponse, error) {
	// FetchAll from db
	citizenPermits, err := s.store.FetchAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch citizen permits: %v", err)
	}

	adminCitizenPermits := convertToAdminCitizenPermits(citizenPermits)

	res := &admin.FetchAllResponse{
		CitizenPermits: adminCitizenPermits,
	}
	return res, nil
}
