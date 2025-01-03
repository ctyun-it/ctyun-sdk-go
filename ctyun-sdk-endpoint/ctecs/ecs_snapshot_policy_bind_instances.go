package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyBindInstancesApi 快照策略绑定云主机
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9594&data=87&isNormal=1
type EcsSnapshotPolicyBindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyBindInstancesApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyBindInstancesApi {
	return &EcsSnapshotPolicyBindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/bind-instances",
		},
	}
}

func (this *EcsSnapshotPolicyBindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyBindInstancesRequest) (*EcsSnapshotPolicyBindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyBindInstancesRealRequest{
		RegionID:         req.RegionID,
		SnapshotPolicyID: req.SnapshotPolicyID,
		InstanceIDs:      req.InstanceIDs,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyBindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyBindInstancesResponse{
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type EcsSnapshotPolicyBindInstancesRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
	InstanceIDs      *string `json:"instanceIDs,omitempty"`
}

type EcsSnapshotPolicyBindInstancesRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
	InstanceIDs      *string
}

type EcsSnapshotPolicyBindInstancesRealResponse struct {
	InstanceIDList []string `json:"instanceIDList,omitempty"`
}

type EcsSnapshotPolicyBindInstancesResponse struct {
	InstanceIDList []string
}
