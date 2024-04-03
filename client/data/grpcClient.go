package data

import (
	admin "client/transport"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type grpcClient struct {
	endpoint string
	client   admin.AdminServiceClient
}

func NewGrpcClient(endpoint string) (*grpcClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s: %v", endpoint, err)
	}

	client := admin.NewAdminServiceClient(conn)
	return &grpcClient{endpoint: endpoint, client: client}, nil
}

func (client *grpcClient) FetchCitizenPermits() ([]*CitizenPermit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := client.client.FetchAll(ctx, &admin.FetchAllRequest{})
	if err != nil {
		return nil, fmt.Errorf("could not fetch citizen permits: %v", err)
	}
	return convertToCitizenPermits(r.GetCitizenPermits()), nil
}

func convertToCitizenPermits(adminCitizenPermits []*admin.CitizenPermit) []*CitizenPermit {
	citizenPermits := make([]*CitizenPermit, len(adminCitizenPermits))
	for i, cp := range adminCitizenPermits {
		citizenPermits[i] = &CitizenPermit{
			PassportNumber:   cp.PassportNumber,
			Surname:          cp.Surname,
			GivenNames:       cp.GivenNames,
			DateOfBirth:      cp.DateOfBirth,
			PlaceOfBirth:     cp.PlaceOfBirth,
			Gender:           cp.Gender,
			Nationality:      cp.Nationality,
			DateOfIssue:      cp.DateOfIssue,
			ExpiryDate:       cp.DateOfIssue,
			IssuingAuthority: cp.IssuingAuthority,
			PermitDate:       cp.PermitDate.AsTime(),
			PermitLocation:   cp.PermitLocation,
			PermitType:       cp.PermitType,
			PermitState:      cp.PermitState,
		}
	}
	return citizenPermits
}

// Modify the permit state of a citizen permit

// Notify the applicant via email
