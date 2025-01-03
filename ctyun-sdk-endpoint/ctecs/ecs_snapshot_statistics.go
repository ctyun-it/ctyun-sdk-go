package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotStatisticsApi
type EcsSnapshotStatisticsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotStatisticsApi(client *ctyunsdk.CtyunClient) *EcsSnapshotStatisticsApi {
	return &EcsSnapshotStatisticsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/statistics",
		},
	}
}

func (this *EcsSnapshotStatisticsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotStatisticsRequest) (*[]EcsSnapshotStatisticsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotStatisticsRealRequest{
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

	var realResponse []EcsSnapshotStatisticsResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	for _, i := range realResponse {
		realResponse = append(realResponse, EcsSnapshotStatisticsResponse{
			InstanceID: i.InstanceID,
			Count:      i.Count,
		})
	}

	return &realResponse, nil

}

type EcsSnapshotStatisticsRealRequest struct {
	RegionID       *string `json:"regionID,omitempty"`
	InstanceIDList *string `json:"instanceIDList,omitempty"`
}

type EcsSnapshotStatisticsRequest struct {
	RegionID       *string
	InstanceIDList *string
}

type EcsSnapshotStatisticsRealResponse struct {
	InstanceID string `json:"instanceID,omitempty"`
	Count      int    `json:"count,omitempty"`
}

type EcsSnapshotStatisticsResponse struct {
	InstanceID string
	Count      int
}
