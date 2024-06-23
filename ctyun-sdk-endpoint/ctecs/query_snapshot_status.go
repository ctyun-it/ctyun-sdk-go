package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// QuerySnapshotStatusApi
type QuerySnapshotStatusApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewQuerySnapshotStatusApi(client *ctyunsdk.CtyunClient) *QuerySnapshotStatusApi {
	return &QuerySnapshotStatusApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/snapshot/status",
		},
	}
}

func (this *QuerySnapshotStatusApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *QuerySnapshotStatusRequest) (*QuerySnapshotStatusResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&QuerySnapshotStatusRealRequest{
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

	var realResponse QuerySnapshotStatusRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &QuerySnapshotStatusResponse{
		SnapshotStatus: realResponse.SnapshotStatus,
	}, nil
}

type QuerySnapshotStatusRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	SnapshotID string `json:"snapshotID,omitempty"`
}

type QuerySnapshotStatusRequest struct {
	RegionID   string
	SnapshotID string
}

type QuerySnapshotStatusRealResponse struct {
	SnapshotStatus string `json:"snapshotStatus,omitempty"`
}

type QuerySnapshotStatusResponse struct {
	SnapshotStatus string
}
