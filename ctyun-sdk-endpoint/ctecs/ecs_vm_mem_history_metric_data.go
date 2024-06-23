package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVmMemHistoryMetricDataApi
type EcsVmMemHistoryMetricDataApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVmMemHistoryMetricDataApi(client *ctyunsdk.CtyunClient) *EcsVmMemHistoryMetricDataApi {
	return &EcsVmMemHistoryMetricDataApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vm-mem-history-metric-data",
		},
	}
}

func (this *EcsVmMemHistoryMetricDataApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVmMemHistoryMetricDataRequest) (*EcsVmMemHistoryMetricDataResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVmMemHistoryMetricDataRealRequest{
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

	var realResponse EcsVmMemHistoryMetricDataRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var result []EcsVmMemHistoryMetricDataResultResponse
	for _, res := range realResponse.Result {
		var aggregate EcsVmMemHistoryMetricDataAggregateListResponse
		var item []EcsVmMemHistoryMetricDataAggregateItemValueResponse
		for _, i := range aggregate.Mem_util {
			item = append(item, EcsVmMemHistoryMetricDataAggregateItemValueResponse{
				Value:        i.Value,
				SamplingTime: i.SamplingTime,
			})
		}

		result = append(result, EcsVmMemHistoryMetricDataResultResponse{
			FUID:             res.FUID,
			FuserLastUpdated: res.FuserLastUpdated,
			RegionID:         res.RegionID,
			DeviceUUID:       res.DeviceUUID,
			ItemAggregateList: EcsVmMemHistoryMetricDataAggregateListResponse{
				Mem_util:            item,
				Free_memory:         item,
				Used_memory:         item,
				Buffer_memory:       item,
				Cache_memory:        item,
				Process_memory_used: item,
				Pused_memory:        item,
			},
		})
	}

	return &EcsVmMemHistoryMetricDataResponse{
		Result:       result,
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PageSize:     realResponse.PageSize,
		Page:         realResponse.Page,
	}, nil
}

type EcsVmMemHistoryMetricDataRealRequest struct {
	RegionID     string   `json:"regionID,omitempty"`
	DeviceIDList []string `json:"deviceIDList,omitempty"`
	Period       int      `json:"period,omitempty"`
	StartTime    string   `json:"startTime,omitempty"`
	EndTime      string   `json:"endTime,omitempty"`
	PageNo       int      `json:"pageNo,omitempty"`
	Page         int      `json:"page,omitempty"`
	PageSize     int      `json:"pageSize,omitempty"`
}

type EcsVmMemHistoryMetricDataRequest struct {
	RegionID     string
	DeviceIDList []string
	Period       int
	StartTime    string
	EndTime      string
	PageNo       int
	Page         int
	PageSize     int
}

type EcsVmMemHistoryMetricDataAggregateItemValueRealResponse struct {
	Value        string `json:"value,omitempty"`
	SamplingTime int    `json:"samplingTime,omitempty"`
}

type AggregateListRealResponse struct {
	Mem_util            []EcsVmMemHistoryMetricDataAggregateItemValueRealResponse `json:"mem_util,omitempty"`
	Free_memory         []EcsVmMemHistoryMetricDataAggregateItemValueRealResponse `json:"free_memory,omitempty"`
	Used_memory         []EcsVmMemHistoryMetricDataAggregateItemValueRealResponse `json:"used_memory,omitempty"`
	Buffer_memory       []EcsVmMemHistoryMetricDataAggregateItemValueRealResponse `json:"buffer_memory,omitempty"`
	Cache_memory        []EcsVmMemHistoryMetricDataAggregateItemValueRealResponse `json:"cache_memory,omitempty"`
	Process_memory_used []EcsVmMemHistoryMetricDataAggregateItemValueRealResponse `json:"process_memory_used,omitempty"`
	Pused_memory        []EcsVmMemHistoryMetricDataAggregateItemValueRealResponse `json:"pused_memory,omitempty"`
}

type EcsVmMemHistoryMetricDataResultRealResponse struct {
	FUID              string                    `json:"fUID,omitempty"`
	FuserLastUpdated  string                    `json:"fuserLastUpdated,omitempty"`
	RegionID          string                    `json:"regionID,omitempty"`
	DeviceUUID        string                    `json:"deviceUUID,omitempty"`
	ItemAggregateList AggregateListRealResponse `json:"itemAggregateList,omitempty"`
}

type EcsVmMemHistoryMetricDataRealResponse struct {
	Result       []EcsVmMemHistoryMetricDataResultRealResponse `json:"result,omitempty"`
	CurrentCount int                                           `json:"currentCount,omitempty"`
	TotalCount   int                                           `json:"totalCount,omitempty"`
	TotalPage    int                                           `json:"totalPage,omitempty"`
	PageSize     int                                           `json:"pageSize,omitempty"`
	Page         int                                           `json:"page,omitempty"`
}

type EcsVmMemHistoryMetricDataAggregateItemValueResponse struct {
	Value        string
	SamplingTime int
}

type EcsVmMemHistoryMetricDataAggregateListResponse struct {
	Mem_util            []EcsVmMemHistoryMetricDataAggregateItemValueResponse
	Free_memory         []EcsVmMemHistoryMetricDataAggregateItemValueResponse
	Used_memory         []EcsVmMemHistoryMetricDataAggregateItemValueResponse
	Buffer_memory       []EcsVmMemHistoryMetricDataAggregateItemValueResponse
	Cache_memory        []EcsVmMemHistoryMetricDataAggregateItemValueResponse
	Process_memory_used []EcsVmMemHistoryMetricDataAggregateItemValueResponse
	Pused_memory        []EcsVmMemHistoryMetricDataAggregateItemValueResponse
}

type EcsVmMemHistoryMetricDataResultResponse struct {
	FUID              string
	FuserLastUpdated  string
	RegionID          string
	DeviceUUID        string
	ItemAggregateList EcsVmMemHistoryMetricDataAggregateListResponse
}

type EcsVmMemHistoryMetricDataResponse struct {
	Result       []EcsVmMemHistoryMetricDataResultResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
	PageSize     int
	Page         int
}
