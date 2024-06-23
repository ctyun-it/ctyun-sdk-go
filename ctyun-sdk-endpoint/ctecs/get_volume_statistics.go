package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// GetVolumeStatisticsApi
type GetVolumeStatisticsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewGetVolumeStatisticsApi(client *ctyunsdk.CtyunClient) *GetVolumeStatisticsApi {
	return &GetVolumeStatisticsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/volume/statistics",
		},
	}
}

func (this *GetVolumeStatisticsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *GetVolumeStatisticsRequest) (*GetVolumeStatisticsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&GetVolumeStatisticsRealRequest{
		RegionID:  req.RegionID,
		ProjectID: req.ProjectID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse GetVolumeStatisticsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &GetVolumeStatisticsResponse{
		VolumeStatistics: GetVolumeStatisticsVolumeStatisticsResponse{
			TotalCount:    realResponse.VolumeStatistics.TotalCount,
			RootDiskCount: realResponse.VolumeStatistics.RootDiskCount,
			DataDiskCount: realResponse.VolumeStatistics.DataDiskCount,
			TotalSize:     realResponse.VolumeStatistics.TotalSize,
			RootDiskSize:  realResponse.VolumeStatistics.RootDiskSize,
			DataDiskSize:  realResponse.VolumeStatistics.DataDiskSize,
		},
	}, nil
}

type GetVolumeStatisticsRealRequest struct {
	RegionID  string `json:"regionID,omitempty"`
	ProjectID string `json:"projectID,omitempty"`
}

type GetVolumeStatisticsRequest struct {
	RegionID  string
	ProjectID string
}

type GetVolumeStatisticsVolumeStatisticsRealResponse struct {
	TotalCount    int `json:"totalCount,omitempty"`
	RootDiskCount int `json:"rootDiskCount,omitempty"`
	DataDiskCount int `json:"dataDiskCount,omitempty"`
	TotalSize     int `json:"totalSize,omitempty"`
	RootDiskSize  int `json:"rootDiskSize,omitempty"`
	DataDiskSize  int `json:"dataDiskSize,omitempty"`
}

type GetVolumeStatisticsRealResponse struct {
	VolumeStatistics GetVolumeStatisticsVolumeStatisticsRealResponse `json:"volumeStatistics,omitempty"`
}

type GetVolumeStatisticsVolumeStatisticsResponse struct {
	TotalCount    int
	RootDiskCount int
	DataDiskCount int
	TotalSize     int
	RootDiskSize  int
	DataDiskSize  int
}

type GetVolumeStatisticsResponse struct {
	VolumeStatistics GetVolumeStatisticsVolumeStatisticsResponse
}
