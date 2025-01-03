package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
	"strconv"
)

// EcsBackupPolicyListApi
type EcsBackupPolicyListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupPolicyListApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyListApi {
	return &EcsBackupPolicyListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup-policy/list",
		},
	}
}

func (this *EcsBackupPolicyListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyListRequest) (*EcsBackupPolicyListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.
		AddParam("regionID", *req.RegionID).
		AddParam("policyID", *req.PolicyID).
		AddParam("policyName", *req.PolicyName).
		AddParam("projectID", *req.ProjectID)
	if req.PageNo != nil {
		builder.AddParam("pageNo", strconv.Itoa(*req.PageNo))
	}
	if req.PageSize != nil {
		builder.AddParam("pageSize", strconv.Itoa(*req.PageSize))
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupPolicyListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var policyList []EcsBackupPolicyListPolicyListResponse
	for _, res := range realResponse.PolicyList {
		var repositoryList []EcsBackupPolicyListRepositoryListResponse
		for _, repo := range res.RepositoryList {
			repositoryList = append(repositoryList, EcsBackupPolicyListRepositoryListResponse{
				RepositoryID:   repo.RepositoryID,
				RepositoryName: repo.RepositoryName,
			})
		}
		policyList = append(policyList, EcsBackupPolicyListPolicyListResponse{
			RegionID:       res.RegionID,
			Status:         res.Status,
			PolicyID:       res.PolicyID,
			PolicyName:     res.PolicyName,
			CycleType:      res.CycleType,
			CycleDay:       res.CycleDay,
			CycleWeek:      res.CycleWeek,
			Time:           res.Time,
			RetentionType:  res.RetentionType,
			RetentionNum:   res.RetentionNum,
			RetentionDay:   res.RetentionDay,
			ResourceCount:  res.ResourceCount,
			ResourceIDs:    res.ResourceIDs,
			RepositoryList: repositoryList,
			ProjectID:      res.ProjectID,
		})
	}

	return &EcsBackupPolicyListResponse{
		CurrentCount: realResponse.CurrentCount,
		CurrentPage:  realResponse.CurrentPage,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PolicyList:   policyList,
	}, nil
}

type EcsBackupPolicyListRealRequest struct {
	RegionID   *string `json:"regionID,omitempty"`
	PolicyID   *string `json:"policyID,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	ProjectID  *string `json:"projectID,omitempty"`
	PageNo     *int    `json:"pageNo,omitempty"`
	PageSize   *int    `json:"pageSize,omitempty"`
}

type EcsBackupPolicyListRequest struct {
	RegionID   *string
	PolicyID   *string
	PolicyName *string
	ProjectID  *string
	PageNo     *int
	PageSize   *int
}

type EcsBackupPolicyListRepositoryListRealResponse struct {
	RepositoryID   string `json:"repositoryID,omitempty"`
	RepositoryName string `json:"repositoryName,omitempty"`
}

type EcsBackupPolicyListPolicyListRealResponse struct {
	RegionID       string                                          `json:"regionID,omitempty"`
	Status         int                                             `json:"status,omitempty"`
	PolicyID       string                                          `json:"policyID,omitempty"`
	PolicyName     string                                          `json:"policyName,omitempty"`
	CycleType      string                                          `json:"cycleType,omitempty"`
	CycleDay       int                                             `json:"cycleDay,omitempty"`
	CycleWeek      string                                          `json:"cycleWeek,omitempty"`
	Time           string                                          `json:"time,omitempty"`
	RetentionType  string                                          `json:"retentionType,omitempty"`
	RetentionNum   int                                             `json:"retentionNum,omitempty"`
	RetentionDay   int                                             `json:"retentionDay,omitempty"`
	ResourceCount  int                                             `json:"resourceCount,omitempty"`
	ResourceIDs    string                                          `json:"resourceIDs,omitempty"`
	RepositoryList []EcsBackupPolicyListRepositoryListRealResponse `json:"repositoryList,omitempty"`
	ProjectID      string                                          `json:"projectID,omitempty"`
}

type EcsBackupPolicyListRealResponse struct {
	CurrentCount int                                         `json:"currentCount,omitempty"`
	CurrentPage  int                                         `json:"currentPage,omitempty"`
	TotalCount   int                                         `json:"totalCount,omitempty"`
	TotalPage    int                                         `json:"totalPage,omitempty"`
	PolicyList   []EcsBackupPolicyListPolicyListRealResponse `json:"policyList,omitempty"`
}

type EcsBackupPolicyListRepositoryListResponse struct {
	RepositoryID   string
	RepositoryName string
}

type EcsBackupPolicyListPolicyListResponse struct {
	RegionID       string
	Status         int
	PolicyID       string
	PolicyName     string
	CycleType      string
	CycleDay       int
	CycleWeek      string
	Time           string
	RetentionType  string
	RetentionNum   int
	RetentionDay   int
	ResourceCount  int
	ResourceIDs    string
	RepositoryList []EcsBackupPolicyListRepositoryListResponse
	ProjectID      string
}

type EcsBackupPolicyListResponse struct {
	CurrentCount int
	CurrentPage  int
	TotalCount   int
	TotalPage    int
	PolicyList   []EcsBackupPolicyListPolicyListResponse
}
