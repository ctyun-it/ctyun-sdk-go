package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// ListInstanceBackupPolicyBindInstancesApi
type ListInstanceBackupPolicyBindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewListInstanceBackupPolicyBindInstancesApi(client *ctyunsdk.CtyunClient) *ListInstanceBackupPolicyBindInstancesApi {
	return &ListInstanceBackupPolicyBindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup-policy/list-instances",
		},
	}
}

func (this *ListInstanceBackupPolicyBindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *ListInstanceBackupPolicyBindInstancesRequest) (*ListInstanceBackupPolicyBindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&ListInstanceBackupPolicyBindInstancesRealRequest{
		RegionID:     req.RegionID,
		PolicyID:     req.PolicyID,
		InstanceName: req.InstanceName,
		PageNo:       req.PageNo,
		PageSize:     req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse ListInstanceBackupPolicyBindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var instancePolicies []ListInstanceBackupPolicyBindInstancesInstancePoliciesResponse
	for _, res := range realResponse.InstancePolicies {
		instancePolicies = append(instancePolicies, ListInstanceBackupPolicyBindInstancesInstancePoliciesResponse{
			Status:          res.Status,
			AttachedVolumes: res.AttachedVolumes,
			DisplayName:     res.DisplayName,
			InstanceID:      res.InstanceID,
			RegionID:        res.RegionID,
			InstanceName:    res.InstanceName,
			CreateTime:      res.CreateTime,
			UpdateTime:      res.UpdateTime,
		})
	}

	return &ListInstanceBackupPolicyBindInstancesResponse{
		CurrentCount:     realResponse.CurrentCount,
		CurrentPage:      realResponse.CurrentPage,
		TotalCount:       realResponse.TotalCount,
		TotalPage:        realResponse.TotalPage,
		InstancePolicies: instancePolicies,
	}, nil
}

type ListInstanceBackupPolicyBindInstancesRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	PolicyID     string `json:"policyID,omitempty"`
	InstanceName string `json:"instanceName,omitempty"`
	PageNo       int    `json:"pageNo,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

type ListInstanceBackupPolicyBindInstancesRequest struct {
	RegionID     string
	PolicyID     string
	InstanceName string
	PageNo       int
	PageSize     int
}

type ListInstanceBackupPolicyBindInstancesInstancePoliciesRealResponse struct {
	Status          string   `json:"status,omitempty"`
	AttachedVolumes []string `json:"attachedVolumes,omitempty"`
	DisplayName     string   `json:"displayName,omitempty"`
	InstanceID      string   `json:"instanceID,omitempty"`
	RegionID        string   `json:"regionID,omitempty"`
	InstanceName    string   `json:"instanceName,omitempty"`
	CreateTime      string   `json:"createTime,omitempty"`
	UpdateTime      string   `json:"updateTime,omitempty"`
}

type ListInstanceBackupPolicyBindInstancesRealResponse struct {
	CurrentCount     int                                                                 `json:"currentCount,omitempty"`
	CurrentPage      int                                                                 `json:"currentPage,omitempty"`
	TotalCount       int                                                                 `json:"totalCount,omitempty"`
	TotalPage        int                                                                 `json:"totalPage,omitempty"`
	InstancePolicies []ListInstanceBackupPolicyBindInstancesInstancePoliciesRealResponse `json:"instancePolicies,omitempty"`
}

type ListInstanceBackupPolicyBindInstancesInstancePoliciesResponse struct {
	Status          string
	AttachedVolumes []string
	DisplayName     string
	InstanceID      string
	RegionID        string
	InstanceName    string
	CreateTime      string
	UpdateTime      string
}

type ListInstanceBackupPolicyBindInstancesResponse struct {
	CurrentCount     int
	CurrentPage      int
	TotalCount       int
	TotalPage        int
	InstancePolicies []ListInstanceBackupPolicyBindInstancesInstancePoliciesResponse
}
