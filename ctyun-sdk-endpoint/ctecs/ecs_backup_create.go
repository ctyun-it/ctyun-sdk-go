package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupCreateApi
type EcsBackupCreateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupCreateApi(client *ctyunsdk.CtyunClient) *EcsBackupCreateApi {
	return &EcsBackupCreateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup/create",
		},
	}
}

func (this *EcsBackupCreateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupCreateRequest) (*EcsBackupCreateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupCreateRealRequest{
		RegionID:                  req.RegionID,
		InstanceID:                req.InstanceID,
		InstanceBackupName:        req.InstanceBackupName,
		InstanceBackupDescription: req.InstanceBackupDescription,
		RepositoryID:              req.RepositoryID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupCreateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	var res EcsBackupCreateResultsResponse
	return &EcsBackupCreateResponse{
		Results: EcsBackupCreateResultsResponse{
			InstanceBackupID:          res.InstanceBackupID,
			InstanceBackupName:        res.InstanceBackupName,
			InstanceBackupStatus:      res.InstanceBackupStatus,
			InstanceBackupDescription: res.InstanceBackupDescription,
			InstanceID:                res.InstanceID,
			InstanceName:              res.InstanceName,
			RepositoryID:              res.RepositoryID,
			RepositoryName:            res.RepositoryName,
			DiskTotalSize:             res.DiskTotalSize,
			UsedSize:                  res.UsedSize,
			CreatedTime:               res.CreatedTime,
			ProjectID:                 res.ProjectID,
		},
	}, nil
}

type EcsBackupCreateRealRequest struct {
	RegionID                  string `json:"regionID,omitempty"`
	InstanceID                string `json:"instanceID,omitempty"`
	InstanceBackupName        string `json:"instanceBackupName,omitempty"`
	InstanceBackupDescription string `json:"instanceBackupDescription,omitempty"`
	RepositoryID              string `json:"repositoryID,omitempty"`
}

type EcsBackupCreateRequest struct {
	RegionID                  string
	InstanceID                string
	InstanceBackupName        string
	InstanceBackupDescription string
	RepositoryID              string
}

type EcsBackupCreateResultsRealResponse struct {
	InstanceBackupID          string `json:"instanceBackupID,omitempty"`
	InstanceBackupName        string `json:"instanceBackupName,omitempty"`
	InstanceBackupStatus      string `json:"instanceBackupStatus,omitempty"`
	InstanceBackupDescription string `json:"instanceBackupDescription,omitempty"`
	InstanceID                string `json:"instanceID,omitempty"`
	InstanceName              string `json:"instanceName,omitempty"`
	RepositoryID              string `json:"repositoryID,omitempty"`
	RepositoryName            string `json:"repositoryName,omitempty"`
	DiskTotalSize             int    `json:"diskTotalSize,omitempty"`
	UsedSize                  int    `json:"usedSize,omitempty"`
	CreatedTime               string `json:"createdTime,omitempty"`
	ProjectID                 string `json:"projectID,omitempty"`
}

type EcsBackupCreateRealResponse struct {
	Results EcsBackupCreateResultsRealResponse `json:"results,omitempty"`
}

type EcsBackupCreateResultsResponse struct {
	InstanceBackupID          string
	InstanceBackupName        string
	InstanceBackupStatus      string
	InstanceBackupDescription string
	InstanceID                string
	InstanceName              string
	RepositoryID              string
	RepositoryName            string
	DiskTotalSize             int
	UsedSize                  int
	CreatedTime               string
	ProjectID                 string
}

type EcsBackupCreateResponse struct {
	Results EcsBackupCreateResultsResponse
}
