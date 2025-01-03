package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyInstanceListApi 查询快照策略绑定云主机列表
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9604&data=87&isNormal=1
type EcsSnapshotPolicyInstanceListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyInstanceListApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyInstanceListApi {
	return &EcsSnapshotPolicyInstanceListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/instance-list",
		},
	}
}

func (this *EcsSnapshotPolicyInstanceListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyInstanceListRequest) (*EcsSnapshotPolicyInstanceListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyInstanceListRealRequest{
		RegionID:         req.RegionID,
		SnapshotPolicyID: req.SnapshotPolicyID,
		PageNo:           req.PageNo,
		PageSize:         req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyInstanceListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var instanceList []EcsSnapshotPolicyInstanceListInstanceListResponse
	for _, res := range realResponse.InstanceList {
		instanceList = append(instanceList, EcsSnapshotPolicyInstanceListInstanceListResponse{
			InstanceID:     res.InstanceID,
			InstanceName:   res.InstanceName,
			DisplayName:    res.DisplayName,
			InstanceStatus: res.InstanceStatus,
			VolumeCount:    res.VolumeCount,
		})
	}

	return &EcsSnapshotPolicyInstanceListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		InstanceList: instanceList,
	}, nil
}

type EcsSnapshotPolicyInstanceListRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
	PageNo           *int    `json:"pageNo,omitempty"`
	PageSize         *int    `json:"pageSize,omitempty"`
}

type EcsSnapshotPolicyInstanceListRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
	PageNo           *int
	PageSize         *int
}

type EcsSnapshotPolicyInstanceListInstanceListRealResponse struct {
	InstanceID     string `json:"instanceID,omitempty"`
	InstanceName   string `json:"instanceName,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	InstanceStatus string `json:"instanceStatus,omitempty"`
	VolumeCount    int    `json:"volumeCount,omitempty"`
}

type EcsSnapshotPolicyInstanceListRealResponse struct {
	CurrentCount int                                                     `json:"currentCount,omitempty"`
	TotalCount   int                                                     `json:"totalCount,omitempty"`
	TotalPage    int                                                     `json:"totalPage,omitempty"`
	InstanceList []EcsSnapshotPolicyInstanceListInstanceListRealResponse `json:"instanceList,omitempty"`
}

type EcsSnapshotPolicyInstanceListInstanceListResponse struct {
	InstanceID     string
	InstanceName   string
	DisplayName    string
	InstanceStatus string
	VolumeCount    int
}

type EcsSnapshotPolicyInstanceListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	InstanceList []EcsSnapshotPolicyInstanceListInstanceListResponse
}
