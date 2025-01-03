package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyDeleteApi 删除云主机快照策略
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9692&data=87&isNormal=1
type EcsSnapshotPolicyDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyDeleteApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyDeleteApi {
	return &EcsSnapshotPolicyDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/delete",
		},
	}
}

func (this *EcsSnapshotPolicyDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyDeleteRequest) (*EcsSnapshotPolicyDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyDeleteRealRequest{
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

	var realResponse EcsSnapshotPolicyDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyDeleteResponse{
		SnapshotPolicyID: realResponse.SnapshotPolicyID,
	}, nil
}

type EcsSnapshotPolicyDeleteRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
}

type EcsSnapshotPolicyDeleteRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
}

type EcsSnapshotPolicyDeleteRealResponse struct {
	SnapshotPolicyID string `json:"snapshotPolicyID,omitempty"`
}

type EcsSnapshotPolicyDeleteResponse struct {
	SnapshotPolicyID string
}
