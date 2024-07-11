package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsUpdateSnapshotApi
type EcsUpdateSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsUpdateSnapshotApi(client *ctyunsdk.CtyunClient) *EcsUpdateSnapshotApi {
	return &EcsUpdateSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/update",
		},
	}
}

func (this *EcsUpdateSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsUpdateSnapshotRequest) (*EcsUpdateSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsUpdateSnapshotRealRequest{
		RegionID:            req.RegionID,
		SnapshotID:          req.SnapshotID,
		SnapshotName:        req.SnapshotName,
		SnapshotDescription: req.SnapshotDescription,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsUpdateSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsUpdateSnapshotResponse{
		SnapshotID: realResponse.SnapshotID,
	}, nil
}

type EcsUpdateSnapshotRealRequest struct {
	RegionID            string `json:"regionID,omitempty"`
	SnapshotID          string `json:"snapshotID,omitempty"`
	SnapshotName        string `json:"snapshotName,omitempty"`
	SnapshotDescription string `json:"snapshotDescription,omitempty"`
}

type EcsUpdateSnapshotRequest struct {
	RegionID            string
	SnapshotID          string
	SnapshotName        string
	SnapshotDescription string
}

type EcsUpdateSnapshotRealResponse struct {
	SnapshotID string `json:"snapshotID,omitempty"`
}

type EcsUpdateSnapshotResponse struct {
	SnapshotID string
}
