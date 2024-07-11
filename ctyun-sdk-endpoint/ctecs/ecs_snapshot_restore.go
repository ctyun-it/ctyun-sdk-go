package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsRestoreSnapshotApi
type EcsRestoreSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsRestoreSnapshotApi(client *ctyunsdk.CtyunClient) *EcsRestoreSnapshotApi {
	return &EcsRestoreSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-restore",
		},
	}
}

func (this *EcsRestoreSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsRestoreSnapshotRequest) (*EcsRestoreSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsRestoreSnapshotRealRequest{
		RegionID:   req.RegionID,
		SnapshotID: req.SnapshotID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsRestoreSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsRestoreSnapshotResponse{
		SnapshotID: realResponse.SnapshotID,
	}, nil
}

type EcsRestoreSnapshotRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	SnapshotID string `json:"snapshotID,omitempty"`
}

type EcsRestoreSnapshotRequest struct {
	RegionID   string
	SnapshotID string
}

type EcsRestoreSnapshotRealResponse struct {
	SnapshotID string `json:"snapshotID,omitempty"`
}

type EcsRestoreSnapshotResponse struct {
	SnapshotID string
}
