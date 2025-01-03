package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVmCpuHistoryMetricDataApi
type EcsVmCpuHistoryMetricDataApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVmCpuHistoryMetricDataApi(client *ctyunsdk.CtyunClient) *EcsVmCpuHistoryMetricDataApi {
	return &EcsVmCpuHistoryMetricDataApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vm-cpu-history-metric-data",
		},
	}
}

func (this *EcsVmCpuHistoryMetricDataApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVmCpuHistoryMetricDataRequest) (*EcsVmCpuHistoryMetricDataResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVmCpuHistoryMetricDataRealRequest{
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

	var realResponse EcsVmCpuHistoryMetricDataRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var result []EcsVmCpuHistoryMetricDataResultResponse
	for _, res := range realResponse.Result {
		var aggregate EcsVmCpuHistoryMetricDataAggregateListResponse
		var item []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
		for _, i := range aggregate.Other_cpu_util {
			item = append(item, EcsVmCpuHistoryMetricDataAggregateListItemValueResponse{
				Value:        i.Value,
				SamplingTime: i.SamplingTime,
			})
		}
		result = append(result, EcsVmCpuHistoryMetricDataResultResponse{
			FUID:             res.FUID,
			FuserLastUpdated: res.FuserLastUpdated,
			RegionID:         res.RegionID,
			DeviceUUID:       res.DeviceUUID,
			ItemAggregateList: EcsVmCpuHistoryMetricDataAggregateListResponse{
				Process_cpu_used:   item,
				Cpu_util:           item,
				Cpu_user_time:      item,
				Cpu_system_time:    item,
				Cpu_interrupt_time: item,
				Cpu_iowait_time:    item,
				Cpu_softirq_time:   item,
				Cpu_idle_time:      item,
				Other_cpu_util:     item,
			},
		})
	}

	return &EcsVmCpuHistoryMetricDataResponse{
		Result:       result,
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PageSize:     realResponse.PageSize,
		Page:         realResponse.Page,
	}, nil
}

type EcsVmCpuHistoryMetricDataRealRequest struct {
	RegionID     *string   `json:"regionID,omitempty"`
	DeviceIDList *[]string `json:"deviceIDList,omitempty"`
	Period       *int      `json:"period,omitempty"`
	StartTime    *string   `json:"startTime,omitempty"`
	EndTime      *string   `json:"endTime,omitempty"`
	PageNo       *int      `json:"pageNo,omitempty"`
	Page         *int      `json:"page,omitempty"`
	PageSize     *int      `json:"pageSize,omitempty"`
}

type EcsVmCpuHistoryMetricDataRequest struct {
	RegionID     *string
	DeviceIDList *[]string
	Period       *int
	StartTime    *string
	EndTime      *string
	PageNo       *int
	Page         *int
	PageSize     *int
}

type EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse struct {
	Value        string `json:"value,omitempty"`
	SamplingTime int    `json:"samplingTime,omitempty"`
}

type EcsVmCpuHistoryMetricDataAggregateListRealResponse struct {
	Process_cpu_used   []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"process_cpu_used,omitempty"`
	Cpu_util           []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"cpu_util,omitempty"`
	Cpu_user_time      []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"cpu_user_time,omitempty"`
	Cpu_system_time    []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"cpu_system_time,omitempty"`
	Cpu_interrupt_time []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"cpu_interrupt_time,omitempty"`
	Cpu_iowait_time    []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"cpu_iowait_time,omitempty"`
	Cpu_softirq_time   []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"cpu_softirq_time,omitempty"`
	Cpu_idle_time      []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"cpu_idle_time,omitempty"`
	Other_cpu_util     []EcsVmCpuHistoryMetricDataAggregateListItemValueRealResponse `json:"other_cpu_util,omitempty"`
}

type EcsVmCpuHistoryMetricDataResultRealResponse struct {
	FUID              string                                             `json:"fUID,omitempty"`
	FuserLastUpdated  string                                             `json:"fuserLastUpdated,omitempty"`
	RegionID          string                                             `json:"regionID,omitempty"`
	DeviceUUID        string                                             `json:"deviceUUID,omitempty"`
	ItemAggregateList EcsVmCpuHistoryMetricDataAggregateListRealResponse `json:"itemAggregateList,omitempty"`
}

type EcsVmCpuHistoryMetricDataRealResponse struct {
	Result       []EcsVmCpuHistoryMetricDataResultRealResponse `json:"result,omitempty"`
	CurrentCount int                                           `json:"currentCount,omitempty"`
	TotalCount   int                                           `json:"totalCount,omitempty"`
	TotalPage    int                                           `json:"totalPage,omitempty"`
	PageSize     int                                           `json:"pageSize,omitempty"`
	Page         int                                           `json:"page,omitempty"`
}

type EcsVmCpuHistoryMetricDataAggregateListItemValueResponse struct {
	Value        string
	SamplingTime int
}

type EcsVmCpuHistoryMetricDataAggregateListResponse struct {
	Process_cpu_used   []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Cpu_util           []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Cpu_user_time      []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Cpu_system_time    []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Cpu_interrupt_time []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Cpu_iowait_time    []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Cpu_softirq_time   []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Cpu_idle_time      []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
	Other_cpu_util     []EcsVmCpuHistoryMetricDataAggregateListItemValueResponse
}

type EcsVmCpuHistoryMetricDataResultResponse struct {
	FUID              string
	FuserLastUpdated  string
	RegionID          string
	DeviceUUID        string
	ItemAggregateList EcsVmCpuHistoryMetricDataAggregateListResponse
}

type EcsVmCpuHistoryMetricDataResponse struct {
	Result       []EcsVmCpuHistoryMetricDataResultResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
	PageSize     int
	Page         int
}
