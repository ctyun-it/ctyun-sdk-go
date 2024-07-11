package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotDeleteApi
type EcsSnapshotDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotDeleteApi(client *ctyunsdk.CtyunClient) *EcsSnapshotDeleteApi {
	return &EcsSnapshotDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-delete",
		},
	}
}

func (this *EcsSnapshotDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotDeleteRequest) (*EcsSnapshotDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotDeleteRealRequest{
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

	var realResponse EcsSnapshotDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotDeleteResponse{
		SnapshotID: realResponse.SnapshotID,
	}, nil
}

type EcsSnapshotDeleteRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	SnapshotID string `json:"snapshotID,omitempty"`
}

type EcsSnapshotDeleteRequest struct {
	RegionID   string
	SnapshotID string
}

type EcsSnapshotDeleteRealResponse struct {
	SnapshotID string `json:"snapshotID,omitempty"`
}

type EcsSnapshotDeleteResponse struct {
	SnapshotID string
}
