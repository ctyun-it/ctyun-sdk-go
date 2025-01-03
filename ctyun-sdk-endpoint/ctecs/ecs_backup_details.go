package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupDetailsApi
type EcsBackupDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupDetailsApi(client *ctyunsdk.CtyunClient) *EcsBackupDetailsApi {
	return &EcsBackupDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup/details",
		},
	}
}

func (this *EcsBackupDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupDetailsRequest) (*EcsBackupDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupDetailsRealRequest{
		RegionID:         req.RegionID,
		InstanceBackupID: req.InstanceBackupID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupDetailsResponse{
		InstanceBackupID:          realResponse.InstanceBackupID,
		InstanceBackupName:        realResponse.InstanceBackupName,
		InstanceBackupStatus:      realResponse.InstanceBackupStatus,
		InstanceBackupDescription: realResponse.InstanceBackupDescription,
		InstanceID:                realResponse.InstanceID,
		InstanceName:              realResponse.InstanceName,
		RepositoryID:              realResponse.RepositoryID,
		RepositoryName:            realResponse.RepositoryName,
		RepositoryExpired:         realResponse.RepositoryExpired,
		RepositoryFreeze:          realResponse.RepositoryFreeze,
		DiskTotalSize:             realResponse.DiskTotalSize,
		UsedSize:                  realResponse.UsedSize,
		DiskCount:                 realResponse.DiskCount,
		RestoreFinishedTime:       realResponse.RestoreFinishedTime,
		CreatedTime:               realResponse.CreatedTime,
		FinishedTime:              realResponse.FinishedTime,
		ProjectID:                 realResponse.ProjectID,
	}, nil
}

type EcsBackupDetailsRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	InstanceBackupID *string `json:"instanceBackupID,omitempty"`
}

type EcsBackupDetailsRequest struct {
	RegionID         *string
	InstanceBackupID *string
}

type EcsBackupDetailsRealResponse struct {
	InstanceBackupID          string `json:"instanceBackupID,omitempty"`
	InstanceBackupName        string `json:"instanceBackupName,omitempty"`
	InstanceBackupStatus      string `json:"instanceBackupStatus,omitempty"`
	InstanceBackupDescription string `json:"instanceBackupDescription,omitempty"`
	InstanceID                string `json:"instanceID,omitempty"`
	InstanceName              string `json:"instanceName,omitempty"`
	RepositoryID              string `json:"repositoryID,omitempty"`
	RepositoryName            string `json:"repositoryName,omitempty"`
	RepositoryExpired         bool   `json:"repositoryExpired,omitempty"`
	RepositoryFreeze          bool   `json:"repositoryFreeze,omitempty"`
	DiskTotalSize             int    `json:"diskTotalSize,omitempty"`
	UsedSize                  int    `json:"usedSize,omitempty"`
	DiskCount                 int    `json:"diskCount,omitempty"`
	RestoreFinishedTime       string `json:"restoreFinishedTime,omitempty"`
	CreatedTime               string `json:"createdTime,omitempty"`
	FinishedTime              string `json:"finishedTime,omitempty"`
	ProjectID                 string `json:"projectID,omitempty"`
}

type EcsBackupDetailsResponse struct {
	InstanceBackupID          string
	InstanceBackupName        string
	InstanceBackupStatus      string
	InstanceBackupDescription string
	InstanceID                string
	InstanceName              string
	RepositoryID              string
	RepositoryName            string
	RepositoryExpired         bool
	RepositoryFreeze          bool
	DiskTotalSize             int
	UsedSize                  int
	DiskCount                 int
	RestoreFinishedTime       string
	CreatedTime               string
	FinishedTime              string
	ProjectID                 string
}
