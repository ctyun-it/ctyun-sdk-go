package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupListApi
type EcsBackupListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupListApi(client *ctyunsdk.CtyunClient) *EcsBackupListApi {
	return &EcsBackupListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup/list",
		},
	}
}

func (this *EcsBackupListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupListRequest) (*EcsBackupListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupListRealRequest{
		RegionID:             req.RegionID,
		PageNo:               req.PageNo,
		PageSize:             req.PageSize,
		InstanceID:           req.InstanceID,
		RepositoryID:         req.RepositoryID,
		InstanceBackupID:     req.InstanceBackupID,
		QueryContent:         req.QueryContent,
		InstanceBackupStatus: req.InstanceBackupStatus,
		ProjectID:            req.ProjectID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EcsBackupListResultsResponse
	for _, res := range realResponse.Results {
		results = append(results, EcsBackupListResultsResponse{
			InstanceBackupID:          res.InstanceBackupID,
			InstanceBackupName:        res.InstanceBackupName,
			InstanceBackupStatus:      res.InstanceBackupStatus,
			InstanceBackupDescription: res.InstanceBackupDescription,
			InstanceID:                res.InstanceID,
			InstanceName:              res.InstanceName,
			RepositoryID:              res.RepositoryID,
			RepositoryName:            res.RepositoryName,
			RepositoryExpired:         res.RepositoryExpired,
			RepositoryFreeze:          res.RepositoryFreeze,
			DiskTotalSize:             res.DiskTotalSize,
			UsedSize:                  res.UsedSize,
			DiskCount:                 res.DiskCount,
			RestoreFinishedTime:       res.RestoreFinishedTime,
			CreatedTime:               res.CreatedTime,
			FinishedTime:              res.FinishedTime,
			ProjectID:                 res.ProjectID,
		})
	}

	return &EcsBackupListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EcsBackupListRealRequest struct {
	RegionID             *string `json:"regionID,omitempty"`
	PageNo               *int    `json:"pageNo,omitempty"`
	PageSize             *int    `json:"pageSize,omitempty"`
	InstanceID           *string `json:"instanceID,omitempty"`
	RepositoryID         *string `json:"repositoryID,omitempty"`
	InstanceBackupID     *string `json:"instanceBackupID,omitempty"`
	QueryContent         *string `json:"queryContent,omitempty"`
	InstanceBackupStatus *string `json:"instanceBackupStatus,omitempty"`
	ProjectID            *string `json:"projectID,omitempty"`
}

type EcsBackupListRequest struct {
	RegionID             *string
	PageNo               *int
	PageSize             *int
	InstanceID           *string
	RepositoryID         *string
	InstanceBackupID     *string
	QueryContent         *string
	InstanceBackupStatus *string
	ProjectID            *string
}

type EcsBackupListResultsRealResponse struct {
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

type EcsBackupListRealResponse struct {
	CurrentCount int                                `json:"currentCount,omitempty"`
	TotalCount   int                                `json:"totalCount,omitempty"`
	TotalPage    int                                `json:"totalPage,omitempty"`
	Results      []EcsBackupListResultsRealResponse `json:"results,omitempty"`
}

type EcsBackupListResultsResponse struct {
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

type EcsBackupListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EcsBackupListResultsResponse
}
