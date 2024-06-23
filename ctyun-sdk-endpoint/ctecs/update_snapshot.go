package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// UpdateSnapshotApi
type UpdateSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewUpdateSnapshotApi(client *ctyunsdk.CtyunClient) *UpdateSnapshotApi {
	return &UpdateSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/update",
		},
	}
}

func (this *UpdateSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *UpdateSnapshotRequest) (*UpdateSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&UpdateSnapshotRealRequest{
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

	var realResponse UpdateSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &UpdateSnapshotResponse{
		SnapshotID: realResponse.SnapshotID,
	}, nil
}

type UpdateSnapshotRealRequest struct {
	RegionID            string `json:"regionID,omitempty"`
	SnapshotID          string `json:"snapshotID,omitempty"`
	SnapshotName        string `json:"snapshotName,omitempty"`
	SnapshotDescription string `json:"snapshotDescription,omitempty"`
}

type UpdateSnapshotRequest struct {
	RegionID            string
	SnapshotID          string
	SnapshotName        string
	SnapshotDescription string
}

type UpdateSnapshotRealResponse struct {
	SnapshotID string `json:"snapshotID,omitempty"`
}

type UpdateSnapshotResponse struct {
	SnapshotID string
}
