package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// ListInstanceBackupPolicyApi
type ListInstanceBackupPolicyApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewListInstanceBackupPolicyApi(client *ctyunsdk.CtyunClient) *ListInstanceBackupPolicyApi {
	return &ListInstanceBackupPolicyApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup-policy/list",
		},
	}
}

func (this *ListInstanceBackupPolicyApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *ListInstanceBackupPolicyRequest) (*ListInstanceBackupPolicyResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&ListInstanceBackupPolicyRealRequest{
		RegionID:   req.RegionID,
		PolicyID:   req.PolicyID,
		PolicyName: req.PolicyName,
		ProjectID:  req.ProjectID,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse ListInstanceBackupPolicyRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var policyList []ListInstanceBackupPolicyPolicyListResponse
	for _, res := range realResponse.PolicyList {
		var repositoryList []ListInstanceBackupPolicyRepositoryListResponse
		for _, repo := range res.RepositoryList {
			repositoryList = append(repositoryList, ListInstanceBackupPolicyRepositoryListResponse{
				RepositoryID:   repo.RepositoryID,
				RepositoryName: repo.RepositoryName,
			})
		}
		policyList = append(policyList, ListInstanceBackupPolicyPolicyListResponse{
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

	return &ListInstanceBackupPolicyResponse{
		CurrentCount: realResponse.CurrentCount,
		CurrentPage:  realResponse.CurrentPage,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PolicyList:   policyList,
	}, nil
}

type ListInstanceBackupPolicyRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	PolicyID   string `json:"policyID,omitempty"`
	PolicyName string `json:"policyName,omitempty"`
	ProjectID  string `json:"projectID,omitempty"`
	PageNo     int    `json:"pageNo,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
}

type ListInstanceBackupPolicyRequest struct {
	RegionID   string
	PolicyID   string
	PolicyName string
	ProjectID  string
	PageNo     int
	PageSize   int
}

type ListInstanceBackupPolicyRepositoryListRealResponse struct {
	RepositoryID   string `json:"repositoryID,omitempty"`
	RepositoryName string `json:"repositoryName,omitempty"`
}

type ListInstanceBackupPolicyPolicyListRealResponse struct {
	RegionID       string                                               `json:"regionID,omitempty"`
	Status         int                                                  `json:"status,omitempty"`
	PolicyID       string                                               `json:"policyID,omitempty"`
	PolicyName     string                                               `json:"policyName,omitempty"`
	CycleType      string                                               `json:"cycleType,omitempty"`
	CycleDay       int                                                  `json:"cycleDay,omitempty"`
	CycleWeek      string                                               `json:"cycleWeek,omitempty"`
	Time           string                                               `json:"time,omitempty"`
	RetentionType  string                                               `json:"retentionType,omitempty"`
	RetentionNum   int                                                  `json:"retentionNum,omitempty"`
	RetentionDay   int                                                  `json:"retentionDay,omitempty"`
	ResourceCount  int                                                  `json:"resourceCount,omitempty"`
	ResourceIDs    string                                               `json:"resourceIDs,omitempty"`
	RepositoryList []ListInstanceBackupPolicyRepositoryListRealResponse `json:"repositoryList,omitempty"`
	ProjectID      string                                               `json:"projectID,omitempty"`
}

type ListInstanceBackupPolicyRealResponse struct {
	CurrentCount int                                              `json:"currentCount,omitempty"`
	CurrentPage  int                                              `json:"currentPage,omitempty"`
	TotalCount   int                                              `json:"totalCount,omitempty"`
	TotalPage    int                                              `json:"totalPage,omitempty"`
	PolicyList   []ListInstanceBackupPolicyPolicyListRealResponse `json:"policyList,omitempty"`
}

type ListInstanceBackupPolicyRepositoryListResponse struct {
	RepositoryID   string
	RepositoryName string
}

type ListInstanceBackupPolicyPolicyListResponse struct {
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
	RepositoryList []ListInstanceBackupPolicyRepositoryListResponse
	ProjectID      string
}

type ListInstanceBackupPolicyResponse struct {
	CurrentCount int
	CurrentPage  int
	TotalCount   int
	TotalPage    int
	PolicyList   []ListInstanceBackupPolicyPolicyListResponse
}
