package ctecs

import (
	"context"
	"fmt"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupRepoListApi 查询云主机存储库列表
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=6909&data=87&isNormal=1
type EcsBackupRepoListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupRepoListApi(client *ctyunsdk.CtyunClient) *EcsBackupRepoListApi {
	return &EcsBackupRepoListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup-repo/list",
		},
	}
}

func (this *EcsBackupRepoListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupRepoListRequest) (*EcsBackupRepoListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", *req.RegionID).
		AddParam("projectID", *req.ProjectID).
		AddParam("repositoryName", *req.RepositoryName).
		AddParam("repositoryID", *req.RepositoryID).
		AddParam("status", *req.Status)

	if req.PageNo != nil {
		builder.AddParam("pageNo", fmt.Sprintf("%d", *req.PageNo))
	}

	if req.PageSize != nil {
		builder.AddParam("pageSize", fmt.Sprintf("%d", *req.PageSize))
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupRepoListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EcsBackupRepoListResultsResponse
	for _, res := range realResponse.Results {
		results = append(results, EcsBackupRepoListResultsResponse{
			RegionID:       res.RegionID,
			RepositoryID:   res.RepositoryID,
			ProjectID:      res.ProjectID,
			RepositoryName: res.RepositoryName,
			Status:         res.Status,
			Size:           res.Size,
			FreeSize:       res.FreeSize,
			RemainingSize:  res.RemainingSize,
			UsedSize:       res.UsedSize,
			CreatedAt:      res.CreatedAt,
			ExpiredAt:      res.ExpiredAt,
			Expired:        res.Expired,
			Freeze:         res.Freeze,
			Paas:           res.Paas,
			BackupList:     res.BackupList,
			BackupCount:    res.BackupCount,
		})
	}

	return &EcsBackupRepoListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EcsBackupRepoListRealRequest struct {
	RegionID       *string `json:"regionID,omitempty"`
	ProjectID      *string `json:"projectID,omitempty"`
	RepositoryName *string `json:"repositoryName,omitempty"`
	RepositoryID   *string `json:"repositoryID,omitempty"`
	Status         *string `json:"status,omitempty"`
	PageNo         *int    `json:"pageNo,omitempty"`
	PageSize       *int    `json:"pageSize,omitempty"`
}

type EcsBackupRepoListRequest struct {
	RegionID       *string
	ProjectID      *string
	RepositoryName *string
	RepositoryID   *string
	Status         *string
	PageNo         *int
	PageSize       *int
}

type EcsBackupRepoListResultsRealResponse struct {
	RegionID       string   `json:"regionID,omitempty"`
	RepositoryID   string   `json:"repositoryID,omitempty"`
	ProjectID      string   `json:"projectID,omitempty"`
	RepositoryName string   `json:"repositoryName,omitempty"`
	Status         string   `json:"status,omitempty"`
	Size           int      `json:"size,omitempty"`
	FreeSize       float64  `json:"freeSize,omitempty"`
	RemainingSize  float64  `json:"remainingSize,omitempty"`
	UsedSize       int      `json:"usedSize,omitempty"`
	CreatedAt      string   `json:"createdAt,omitempty"`
	ExpiredAt      string   `json:"expiredAt,omitempty"`
	Expired        bool     `json:"expired,omitempty"`
	Freeze         bool     `json:"freeze,omitempty"`
	Paas           bool     `json:"paas,omitempty"`
	BackupList     []string `json:"backupList,omitempty"`
	BackupCount    int      `json:"backupCount,omitempty"`
}

type EcsBackupRepoListRealResponse struct {
	CurrentCount int                                    `json:"currentCount,omitempty"`
	TotalCount   int                                    `json:"totalCount,omitempty"`
	TotalPage    int                                    `json:"totalPage,omitempty"`
	Results      []EcsBackupRepoListResultsRealResponse `json:"results,omitempty"`
}

type EcsBackupRepoListResultsResponse struct {
	RegionID       string
	RepositoryID   string
	ProjectID      string
	RepositoryName string
	Status         string
	Size           int
	FreeSize       float64
	RemainingSize  float64
	UsedSize       int
	CreatedAt      string
	ExpiredAt      string
	Expired        bool
	Freeze         bool
	Paas           bool
	BackupList     []string
	BackupCount    int
}

type EcsBackupRepoListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EcsBackupRepoListResultsResponse
}
