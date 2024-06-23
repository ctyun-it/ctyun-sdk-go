package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// QuerySnapshotStatisticsApi
type QuerySnapshotStatisticsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewQuerySnapshotStatisticsApi(client *ctyunsdk.CtyunClient) *QuerySnapshotStatisticsApi {
	return &QuerySnapshotStatisticsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/statistics",
		},
	}
}

func (this *QuerySnapshotStatisticsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *QuerySnapshotStatisticsRequest) (*QuerySnapshotStatisticsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&QuerySnapshotStatisticsRealRequest{
		RegionID:       req.RegionID,
		InstanceIDList: req.InstanceIDList,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse QuerySnapshotStatisticsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &QuerySnapshotStatisticsResponse{
		InstanceID: realResponse.InstanceID,
		Count:      realResponse.Count,
	}, nil
}

type QuerySnapshotStatisticsRealRequest struct {
	RegionID       string `json:"regionID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type QuerySnapshotStatisticsRequest struct {
	RegionID       string
	InstanceIDList string
}

type QuerySnapshotStatisticsRealResponse struct {
	InstanceID string `json:"instanceID,omitempty"`
	Count      int    `json:"count,omitempty"`
}

type QuerySnapshotStatisticsResponse struct {
	InstanceID string
	Count      int
}
