package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVmDiskLatestMetricDataApi
type EcsVmDiskLatestMetricDataApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVmDiskLatestMetricDataApi(client *ctyunsdk.CtyunClient) *EcsVmDiskLatestMetricDataApi {
	return &EcsVmDiskLatestMetricDataApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vm-disk-latest-metric-data",
		},
	}
}

func (this *EcsVmDiskLatestMetricDataApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVmDiskLatestMetricDataRequest) (*EcsVmDiskLatestMetricDataResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVmDiskLatestMetricDataRealRequest{
		RegionID:     req.RegionID,
		DeviceIDList: req.DeviceIDList,
		PageNo:       req.PageNo,
		Page:         req.Page,
		PageSize:     req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVmDiskLatestMetricDataRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var result []EcsVmDiskLatestMetricDataResultResponse
	for _, res := range realResponse.Result {
		result = append(result, EcsVmDiskLatestMetricDataResultResponse{
			FUID:             res.FUID,
			FuserLastUpdated: res.FuserLastUpdated,
			RegionID:         res.RegionID,
			DeviceUUID:       res.DeviceUUID,
			ExtendStatus:     res.ExtendStatus,
			ItemList: EcsVmDiskLatestMetricDataItemListResponse{
				SamplingTime:             res.ItemList.SamplingTime,
				Disk_read_bytes_rate:     res.ItemList.Disk_read_bytes_rate,
				Disk_read_requests_rate:  res.ItemList.Disk_read_requests_rate,
				Disk_write_requests_rate: res.ItemList.Disk_write_requests_rate,
				Disk_util:                res.ItemList.Disk_util,
				Disk_write_bytes_rate:    res.ItemList.Disk_write_bytes_rate,
			},
		})
	}

	return &EcsVmDiskLatestMetricDataResponse{
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PageSize:     realResponse.PageSize,
		Page:         realResponse.Page,
		CurrentCount: realResponse.CurrentCount,
		Result:       result,
	}, nil
}

type EcsVmDiskLatestMetricDataRealRequest struct {
	RegionID     *string   `json:"regionID,omitempty"`
	DeviceIDList *[]string `json:"deviceIDList,omitempty"`
	PageNo       *int      `json:"pageNo,omitempty"`
	Page         *int      `json:"page,omitempty"`
	PageSize     *int      `json:"pageSize,omitempty"`
}

type EcsVmDiskLatestMetricDataRequest struct {
	RegionID     *string
	DeviceIDList *[]string
	PageNo       *int
	Page         *int
	PageSize     *int
}

type EcsVmDiskLatestMetricDataItemListRealResponse struct {
	SamplingTime             int    `json:"samplingTime,omitempty"`
	Disk_read_bytes_rate     string `json:"disk_read_bytes_rate,omitempty"`
	Disk_read_requests_rate  string `json:"disk_read_requests_rate,omitempty"`
	Disk_write_requests_rate string `json:"disk_write_requests_rate,omitempty"`
	Disk_util                string `json:"disk_util,omitempty"`
	Disk_write_bytes_rate    string `json:"disk_write_bytes_rate,omitempty"`
}

type EcsVmDiskLatestMetricDataResultRealResponse struct {
	FUID             string                                        `json:"fUID,omitempty"`
	FuserLastUpdated string                                        `json:"fuserLastUpdated,omitempty"`
	RegionID         string                                        `json:"regionID,omitempty"`
	DeviceUUID       string                                        `json:"deviceUUID,omitempty"`
	ExtendStatus     int                                           `json:"extendStatus,omitempty"`
	ItemList         EcsVmDiskLatestMetricDataItemListRealResponse `json:"itemList,omitempty"`
}

type EcsVmDiskLatestMetricDataRealResponse struct {
	Result       []EcsVmDiskLatestMetricDataResultRealResponse `json:"result,omitempty"`
	CurrentCount int                                           `json:"currentCount,omitempty"`
	TotalCount   int                                           `json:"totalCount,omitempty"`
	TotalPage    int                                           `json:"totalPage,omitempty"`
	PageSize     int                                           `json:"pageSize,omitempty"`
	Page         int                                           `json:"page,omitempty"`
}

type EcsVmDiskLatestMetricDataItemListResponse struct {
	SamplingTime             int
	Disk_read_bytes_rate     string
	Disk_read_requests_rate  string
	Disk_write_requests_rate string
	Disk_util                string
	Disk_write_bytes_rate    string
}

type EcsVmDiskLatestMetricDataResultResponse struct {
	FUID             string
	FuserLastUpdated string
	RegionID         string
	DeviceUUID       string
	ExtendStatus     int
	ItemList         EcsVmDiskLatestMetricDataItemListResponse
}

type EcsVmDiskLatestMetricDataResponse struct {
	Result       []EcsVmDiskLatestMetricDataResultResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
	PageSize     int
	Page         int
}
