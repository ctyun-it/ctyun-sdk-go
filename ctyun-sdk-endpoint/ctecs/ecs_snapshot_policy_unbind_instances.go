package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyUnbindInstancesApi 快照策略解绑云主机
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9597&data=87&isNormal=1
type EcsSnapshotPolicyUnbindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyUnbindInstancesApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyUnbindInstancesApi {
	return &EcsSnapshotPolicyUnbindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/unbind-instances",
		},
	}
}

func (this *EcsSnapshotPolicyUnbindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyUnbindInstancesRequest) (*EcsSnapshotPolicyUnbindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyUnbindInstancesRealRequest{
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

	var realResponse EcsSnapshotPolicyUnbindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyUnbindInstancesResponse{
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type EcsSnapshotPolicyUnbindInstancesRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
	InstanceIDs      *string `json:"instanceIDs,omitempty"`
}

type EcsSnapshotPolicyUnbindInstancesRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
	InstanceIDs      *string
}

type EcsSnapshotPolicyUnbindInstancesRealResponse struct {
	InstanceIDList []string `json:"instanceIDList,omitempty"`
}

type EcsSnapshotPolicyUnbindInstancesResponse struct {
	InstanceIDList []string
}
