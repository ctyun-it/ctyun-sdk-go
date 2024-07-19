package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupPolicyListInstancesApi
type EcsBackupPolicyListInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupPolicyListInstancesApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyListInstancesApi {
	return &EcsBackupPolicyListInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup-policy/list-instances",
		},
	}
}

func (this *EcsBackupPolicyListInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyListInstancesRequest) (*EcsBackupPolicyListInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupPolicyListInstancesRealRequest{
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

	var realResponse EcsBackupPolicyListInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var instancePolicies []EcsBackupPolicyListInstancesInstancePoliciesResponse
	for _, res := range realResponse.InstancePolicies {
		instancePolicies = append(instancePolicies, EcsBackupPolicyListInstancesInstancePoliciesResponse{
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

	return &EcsBackupPolicyListInstancesResponse{
		CurrentCount:     realResponse.CurrentCount,
		CurrentPage:      realResponse.CurrentPage,
		TotalCount:       realResponse.TotalCount,
		TotalPage:        realResponse.TotalPage,
		InstancePolicies: instancePolicies,
	}, nil
}

type EcsBackupPolicyListInstancesRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	PolicyID     string `json:"policyID,omitempty"`
	InstanceName string `json:"instanceName,omitempty"`
	PageNo       *int   `json:"pageNo,omitempty"`
	PageSize     *int   `json:"pageSize,omitempty"`
}

type EcsBackupPolicyListInstancesRequest struct {
	RegionID     string
	PolicyID     string
	InstanceName string
	PageNo       *int
	PageSize     *int
}

type EcsBackupPolicyListInstancesInstancePoliciesRealResponse struct {
	Status          string   `json:"status,omitempty"`
	AttachedVolumes []string `json:"attachedVolumes,omitempty"`
	DisplayName     string   `json:"displayName,omitempty"`
	InstanceID      string   `json:"instanceID,omitempty"`
	RegionID        string   `json:"regionID,omitempty"`
	InstanceName    string   `json:"instanceName,omitempty"`
	CreateTime      string   `json:"createTime,omitempty"`
	UpdateTime      string   `json:"updateTime,omitempty"`
}

type EcsBackupPolicyListInstancesRealResponse struct {
	CurrentCount     int                                                        `json:"currentCount,omitempty"`
	CurrentPage      int                                                        `json:"currentPage,omitempty"`
	TotalCount       int                                                        `json:"totalCount,omitempty"`
	TotalPage        int                                                        `json:"totalPage,omitempty"`
	InstancePolicies []EcsBackupPolicyListInstancesInstancePoliciesRealResponse `json:"instancePolicies,omitempty"`
}

type EcsBackupPolicyListInstancesInstancePoliciesResponse struct {
	Status          string
	AttachedVolumes []string
	DisplayName     string
	InstanceID      string
	RegionID        string
	InstanceName    string
	CreateTime      string
	UpdateTime      string
}

type EcsBackupPolicyListInstancesResponse struct {
	CurrentCount     int
	CurrentPage      int
	TotalCount       int
	TotalPage        int
	InstancePolicies []EcsBackupPolicyListInstancesInstancePoliciesResponse
}
