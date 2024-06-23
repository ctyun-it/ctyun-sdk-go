package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// DeleteSnapshotApi
type DeleteSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewDeleteSnapshotApi(client *ctyunsdk.CtyunClient) *DeleteSnapshotApi {
	return &DeleteSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-delete",
		},
	}
}

func (this *DeleteSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *DeleteSnapshotRequest) (*DeleteSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&DeleteSnapshotRealRequest{
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

	var realResponse DeleteSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &DeleteSnapshotResponse{
		SnapshotID: realResponse.SnapshotID,
	}, nil
}

type DeleteSnapshotRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	SnapshotID string `json:"snapshotID,omitempty"`
}

type DeleteSnapshotRequest struct {
	RegionID   string
	SnapshotID string
}

type DeleteSnapshotRealResponse struct {
	SnapshotID string `json:"snapshotID,omitempty"`
}

type DeleteSnapshotResponse struct {
	SnapshotID string
}
