package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// RestoreSnapshotApi
type RestoreSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewRestoreSnapshotApi(client *ctyunsdk.CtyunClient) *RestoreSnapshotApi {
	return &RestoreSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-restore",
		},
	}
}

func (this *RestoreSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *RestoreSnapshotRequest) (*RestoreSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&RestoreSnapshotRealRequest{
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

	var realResponse RestoreSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &RestoreSnapshotResponse{
		SnapshotID: realResponse.SnapshotID,
	}, nil
}

type RestoreSnapshotRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	SnapshotID string `json:"snapshotID,omitempty"`
}

type RestoreSnapshotRequest struct {
	RegionID   string
	SnapshotID string
}

type RestoreSnapshotRealResponse struct {
	SnapshotID string `json:"snapshotID,omitempty"`
}

type RestoreSnapshotResponse struct {
	SnapshotID string
}
