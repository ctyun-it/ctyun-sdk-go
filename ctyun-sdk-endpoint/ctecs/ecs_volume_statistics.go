package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeStatisticsApi
type EcsVolumeStatisticsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeStatisticsApi(client *ctyunsdk.CtyunClient) *EcsVolumeStatisticsApi {
	return &EcsVolumeStatisticsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/volume/statistics",
		},
	}
}

func (this *EcsVolumeStatisticsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeStatisticsRequest) (*EcsVolumeStatisticsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.AddParam("regionID", *req.RegionID)
	builder.AddParam("projectID", *req.ProjectID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeStatisticsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVolumeStatisticsResponse{
		VolumeStatistics: EcsVolumeStatisticsVolumeStatisticsResponse{
			TotalCount:    realResponse.VolumeStatistics.TotalCount,
			RootDiskCount: realResponse.VolumeStatistics.RootDiskCount,
			DataDiskCount: realResponse.VolumeStatistics.DataDiskCount,
			TotalSize:     realResponse.VolumeStatistics.TotalSize,
			RootDiskSize:  realResponse.VolumeStatistics.RootDiskSize,
			DataDiskSize:  realResponse.VolumeStatistics.DataDiskSize,
		},
	}, nil
}

type EcsVolumeStatisticsRealRequest struct {
	RegionID  *string `json:"regionID,omitempty"`
	ProjectID *string `json:"projectID,omitempty"`
}

type EcsVolumeStatisticsRequest struct {
	RegionID  *string
	ProjectID *string
}

type EcsVolumeStatisticsVolumeStatisticsRealResponse struct {
	TotalCount    int `json:"totalCount,omitempty"`
	RootDiskCount int `json:"rootDiskCount,omitempty"`
	DataDiskCount int `json:"dataDiskCount,omitempty"`
	TotalSize     int `json:"totalSize,omitempty"`
	RootDiskSize  int `json:"rootDiskSize,omitempty"`
	DataDiskSize  int `json:"dataDiskSize,omitempty"`
}

type EcsVolumeStatisticsRealResponse struct {
	VolumeStatistics EcsVolumeStatisticsVolumeStatisticsRealResponse `json:"volumeStatistics,omitempty"`
}

type EcsVolumeStatisticsVolumeStatisticsResponse struct {
	TotalCount    int
	RootDiskCount int
	DataDiskCount int
	TotalSize     int
	RootDiskSize  int
	DataDiskSize  int
}

type EcsVolumeStatisticsResponse struct {
	VolumeStatistics EcsVolumeStatisticsVolumeStatisticsResponse
}
