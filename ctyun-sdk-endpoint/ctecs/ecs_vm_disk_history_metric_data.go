package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVmDiskHistoryMetricDataApi
type EcsVmDiskHistoryMetricDataApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVmDiskHistoryMetricDataApi(client *ctyunsdk.CtyunClient) *EcsVmDiskHistoryMetricDataApi {
	return &EcsVmDiskHistoryMetricDataApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vm-disk-history-metric-data",
		},
	}
}

func (this *EcsVmDiskHistoryMetricDataApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVmDiskHistoryMetricDataRequest) (*EcsVmDiskHistoryMetricDataResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVmDiskHistoryMetricDataRealRequest{
		RegionID:     req.RegionID,
		DeviceIDList: req.DeviceIDList,
		Period:       req.Period,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
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

	var realResponse EcsVmDiskHistoryMetricDataRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var result []EcsVmDiskHistoryMetricDataResultResponse
	for _, res := range realResponse.Result {
		var aggregate EcsVmDiskHistoryMetricDataAggregateListResponse
		var item []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse
		for _, i := range aggregate.Disk_write_requests_rate {
			item = append(item, EcsVmDiskHistoryMetricDataAggregateListItemValueResponse{
				Value:        i.Value,
				SamplingTime: i.SamplingTime,
			})
		}

		result = append(result, EcsVmDiskHistoryMetricDataResultResponse{
			FUID:             res.FUID,
			FuserLastUpdated: res.FuserLastUpdated,
			RegionID:         res.RegionID,
			DeviceUUID:       res.DeviceUUID,
			ItemAggregateList: EcsVmDiskHistoryMetricDataAggregateListResponse{
				Disk_read_bytes_rate:     item,
				Disk_read_requests_rate:  item,
				Disk_write_requests_rate: item,
				Disk_util:                item,
				Disk_write_bytes_rate:    item,
			},
		})
	}

	return &EcsVmDiskHistoryMetricDataResponse{
		Result:       result,
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PageSize:     realResponse.PageSize,
		Page:         realResponse.Page,
	}, nil
}

type EcsVmDiskHistoryMetricDataRealRequest struct {
	RegionID     string   `json:"regionID,omitempty"`
	DeviceIDList []string `json:"deviceIDList,omitempty"`
	Period       *int     `json:"period,omitempty"`
	StartTime    string   `json:"startTime,omitempty"`
	EndTime      string   `json:"endTime,omitempty"`
	PageNo       *int     `json:"pageNo,omitempty"`
	Page         *int     `json:"page,omitempty"`
	PageSize     *int     `json:"pageSize,omitempty"`
}

type EcsVmDiskHistoryMetricDataRequest struct {
	RegionID     string
	DeviceIDList []string
	Period       *int
	StartTime    string
	EndTime      string
	PageNo       *int
	Page         *int
	PageSize     *int
}

type EcsVmDiskHistoryMetricDataAggregateListRealResponse struct {
	Disk_read_bytes_rate     []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse `json:"disk_read_bytes_rate,omitempty"`
	Disk_read_requests_rate  []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse `json:"disk_read_requests_rate,omitempty"`
	Disk_write_requests_rate []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse `json:"disk_write_requests_rate,omitempty"`
	Disk_util                []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse `json:"disk_util,omitempty"`
	Disk_write_bytes_rate    []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse `json:"disk_write_bytes_rate,omitempty"`
}

type EcsVmDiskHistoryMetricDataResultRealResponse struct {
	FUID              string                                              `json:"fUID,omitempty"`
	FuserLastUpdated  string                                              `json:"fuserLastUpdated,omitempty"`
	RegionID          string                                              `json:"regionID,omitempty"`
	DeviceUUID        string                                              `json:"deviceUUID,omitempty"`
	ItemAggregateList EcsVmDiskHistoryMetricDataAggregateListRealResponse `json:"itemAggregateList,omitempty"`
}

type EcsVmDiskHistoryMetricDataRealResponse struct {
	Result       []EcsVmDiskHistoryMetricDataResultRealResponse `json:"result,omitempty"`
	CurrentCount int                                            `json:"currentCount,omitempty"`
	TotalCount   int                                            `json:"totalCount,omitempty"`
	TotalPage    int                                            `json:"totalPage,omitempty"`
	PageSize     int                                            `json:"pageSize,omitempty"`
	Page         int                                            `json:"page,omitempty"`
}

type EcsVmDiskHistoryMetricDataAggregateListItemValueResponse struct {
	Value        string
	SamplingTime int
}

type EcsVmDiskHistoryMetricDataAggregateListResponse struct {
	Disk_read_bytes_rate     []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse
	Disk_read_requests_rate  []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse
	Disk_write_requests_rate []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse
	Disk_util                []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse
	Disk_write_bytes_rate    []EcsVmDiskHistoryMetricDataAggregateListItemValueResponse
}

type EcsVmDiskHistoryMetricDataResultResponse struct {
	FUID              string
	FuserLastUpdated  string
	RegionID          string
	DeviceUUID        string
	ItemAggregateList EcsVmDiskHistoryMetricDataAggregateListResponse
}

type EcsVmDiskHistoryMetricDataResponse struct {
	Result       []EcsVmDiskHistoryMetricDataResultResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
	PageSize     int
	Page         int
}
