package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyExecuteApi 立即执行云主机快照策略
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9605&data=87&isNormal=1
type EcsSnapshotPolicyExecuteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyExecuteApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyExecuteApi {
	return &EcsSnapshotPolicyExecuteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/execute",
		},
	}
}

func (this *EcsSnapshotPolicyExecuteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyExecuteRequest) (*EcsSnapshotPolicyExecuteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyExecuteRealRequest{
		RegionID:         req.RegionID,
		SnapshotPolicyID: req.SnapshotPolicyID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyExecuteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyExecuteResponse{
		SnapshotPolicyID: realResponse.SnapshotPolicyID,
	}, nil
}

type EcsSnapshotPolicyExecuteRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
}

type EcsSnapshotPolicyExecuteRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
}

type EcsSnapshotPolicyExecuteRealResponse struct {
	SnapshotPolicyID string `json:"snapshotPolicyID,omitempty"`
}

type EcsSnapshotPolicyExecuteResponse struct {
	SnapshotPolicyID string
}
