package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotStatusApi
type EcsSnapshotStatusApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotStatusApi(client *ctyunsdk.CtyunClient) *EcsSnapshotStatusApi {
	return &EcsSnapshotStatusApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/snapshot/status",
		},
	}
}

func (this *EcsSnapshotStatusApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotStatusRequest) (*EcsSnapshotStatusResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotStatusRealRequest{
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

	var realResponse EcsSnapshotStatusRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotStatusResponse{
		SnapshotStatus: realResponse.SnapshotStatus,
	}, nil
}

type EcsSnapshotStatusRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	SnapshotID string `json:"snapshotID,omitempty"`
}

type EcsSnapshotStatusRequest struct {
	RegionID   string
	SnapshotID string
}

type EcsSnapshotStatusRealResponse struct {
	SnapshotStatus string `json:"snapshotStatus,omitempty"`
}

type EcsSnapshotStatusResponse struct {
	SnapshotStatus string
}
