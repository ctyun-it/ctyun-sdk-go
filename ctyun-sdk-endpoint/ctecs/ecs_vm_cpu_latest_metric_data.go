package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVmCpuLatestMetricDataApi
type EcsVmCpuLatestMetricDataApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVmCpuLatestMetricDataApi(client *ctyunsdk.CtyunClient) *EcsVmCpuLatestMetricDataApi {
	return &EcsVmCpuLatestMetricDataApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vm-cpu-latest-metric-data",
		},
	}
}

func (this *EcsVmCpuLatestMetricDataApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVmCpuLatestMetricDataRequest) (*EcsVmCpuLatestMetricDataResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVmCpuLatestMetricDataRealRequest{
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

	var realResponse EcsVmCpuLatestMetricDataRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var result []EcsVmCpuLatestMetricDataResultResponse
	for _, res := range realResponse.Result {
		result = append(result, EcsVmCpuLatestMetricDataResultResponse{
			FUID:             res.FUID,
			FuserLastUpdated: res.FuserLastUpdated,
			RegionID:         res.RegionID,
			DeviceUUID:       res.DeviceUUID,
			ItemList: EcsVmCpuLatestMetricDataItemListResponse{
				SamplingTime:       res.ItemList.SamplingTime,
				Process_cpu_used:   res.ItemList.Process_cpu_used,
				Cpu_util:           res.ItemList.Cpu_util,
				Cpu_user_time:      res.ItemList.Cpu_user_time,
				Cpu_system_time:    res.ItemList.Cpu_system_time,
				Cpu_interrupt_time: res.ItemList.Cpu_interrupt_time,
				Cpu_iowait_time:    res.ItemList.Cpu_iowait_time,
				Cpu_softirq_time:   res.ItemList.Cpu_softirq_time,
				Cpu_idle_time:      res.ItemList.Cpu_idle_time,
				Other_cpu_util:     res.ItemList.Other_cpu_util,
			},
		})
	}

	return &EcsVmCpuLatestMetricDataResponse{
		Result:       result,
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PageSize:     realResponse.PageSize,
		Page:         realResponse.Page,
	}, nil
}

type EcsVmCpuLatestMetricDataRealRequest struct {
	RegionID     string   `json:"regionID,omitempty"`
	DeviceIDList []string `json:"deviceIDList,omitempty"`
	PageNo       *int     `json:"pageNo,omitempty"`
	Page         *int     `json:"page,omitempty"`
	PageSize     *int     `json:"pageSize,omitempty"`
}

type EcsVmCpuLatestMetricDataRequest struct {
	RegionID     string
	DeviceIDList []string
	PageNo       *int
	Page         *int
	PageSize     *int
}

type EcsVmCpuLatestMetricDataItemListRealResponse struct {
	SamplingTime       int    `json:"samplingTime,omitempty"`
	Process_cpu_used   string `json:"process_cpu_used,omitempty"`
	Cpu_util           string `json:"cpu_util,omitempty"`
	Cpu_user_time      string `json:"cpu_user_time,omitempty"`
	Cpu_system_time    string `json:"cpu_system_time,omitempty"`
	Cpu_interrupt_time string `json:"cpu_interrupt_time,omitempty"`
	Cpu_iowait_time    string `json:"cpu_iowait_time,omitempty"`
	Cpu_softirq_time   string `json:"cpu_softirq_time,omitempty"`
	Cpu_idle_time      string `json:"cpu_idle_time,omitempty"`
	Other_cpu_util     string `json:"other_cpu_util,omitempty"`
}

type EcsVmCpuLatestMetricDataResultRealResponse struct {
	FUID             string                                       `json:"fUID,omitempty"`
	FuserLastUpdated string                                       `json:"fuserLastUpdated,omitempty"`
	RegionID         string                                       `json:"regionID,omitempty"`
	DeviceUUID       string                                       `json:"deviceUUID,omitempty"`
	ItemList         EcsVmCpuLatestMetricDataItemListRealResponse `json:"itemList,omitempty"`
}

type EcsVmCpuLatestMetricDataRealResponse struct {
	Result       []EcsVmCpuLatestMetricDataResultRealResponse `json:"result,omitempty"`
	CurrentCount int                                          `json:"currentCount,omitempty"`
	TotalCount   int                                          `json:"totalCount,omitempty"`
	TotalPage    int                                          `json:"totalPage,omitempty"`
	PageSize     int                                          `json:"pageSize,omitempty"`
	Page         int                                          `json:"page,omitempty"`
}

type EcsVmCpuLatestMetricDataItemListResponse struct {
	SamplingTime       int
	Process_cpu_used   string
	Cpu_util           string
	Cpu_user_time      string
	Cpu_system_time    string
	Cpu_interrupt_time string
	Cpu_iowait_time    string
	Cpu_softirq_time   string
	Cpu_idle_time      string
	Other_cpu_util     string
}

type EcsVmCpuLatestMetricDataResultResponse struct {
	FUID             string
	FuserLastUpdated string
	RegionID         string
	DeviceUUID       string
	ItemList         EcsVmCpuLatestMetricDataItemListResponse
}

type EcsVmCpuLatestMetricDataResponse struct {
	Result       []EcsVmCpuLatestMetricDataResultResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
	PageSize     int
	Page         int
}
