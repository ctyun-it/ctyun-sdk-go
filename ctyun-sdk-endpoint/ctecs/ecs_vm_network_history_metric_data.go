package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVmNetworkHistoryMetricDataApi
type EcsVmNetworkHistoryMetricDataApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVmNetworkHistoryMetricDataApi(client *ctyunsdk.CtyunClient) *EcsVmNetworkHistoryMetricDataApi {
	return &EcsVmNetworkHistoryMetricDataApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vm-network-history-metric-data",
		},
	}
}

func (this *EcsVmNetworkHistoryMetricDataApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVmNetworkHistoryMetricDataRequest) (*EcsVmNetworkHistoryMetricDataResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVmNetworkHistoryMetricDataRealRequest{
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

	var realResponse EcsVmNetworkHistoryMetricDataRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var result []EcsVmNetworkHistoryMetricDataResultResponse
	for _, res := range realResponse.Result {
		var aggregate EcsVmNetworkHistoryMetricDataAggregateListResponse
		var item []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
		for _, i := range aggregate.Net_in_bytes_rate {
			item = append(item, EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse{
				Value:        i.Value,
				SamplingTime: i.SamplingTime,
			})
		}
		result = append(result, EcsVmNetworkHistoryMetricDataResultResponse{
			FUID:             res.FUID,
			FuserLastUpdated: res.FuserLastUpdated,
			RegionID:         res.RegionID,
			DeviceUUID:       res.DeviceUUID,
			ItemAggregateList: EcsVmNetworkHistoryMetricDataAggregateListResponse{
				Net_in_bytes_rate:                    item,
				Net_out_bytes_rate:                   item,
				Network_incoming_packets_rate_inband: item,
				Network_outing_packets_rate_inband:   item,
				Network_incoming_errs_rate_inband:    item,
				Network_outing_drop_rate_inband:      item,
				Network_incoming_drop_rate_inband:    item,
				Network_outing_errs_rate_inband:      item,
			},
		})
	}

	return &EcsVmNetworkHistoryMetricDataResponse{
		Result:       result,
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		PageSize:     realResponse.PageSize,
		Page:         realResponse.Page,
	}, nil
}

type EcsVmNetworkHistoryMetricDataRealRequest struct {
	RegionID     *string   `json:"regionID,omitempty"`
	DeviceIDList *[]string `json:"deviceIDList,omitempty"`
	Period       *int      `json:"period,omitempty"`
	StartTime    *string   `json:"startTime,omitempty"`
	EndTime      *string   `json:"endTime,omitempty"`
	PageNo       *int      `json:"pageNo,omitempty"`
	Page         *int      `json:"page,omitempty"`
	PageSize     *int      `json:"pageSize,omitempty"`
}

type EcsVmNetworkHistoryMetricDataRequest struct {
	RegionID     *string
	DeviceIDList *[]string
	Period       *int
	StartTime    *string
	EndTime      *string
	PageNo       *int
	Page         *int
	PageSize     *int
}

type EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse struct {
	Value        string `json:"value,omitempty"`
	SamplingTime int    `json:"samplingTime,omitempty"`
}

type EcsVmNetworkHistoryMetricDataAggregateListRealResponse struct {
	Net_in_bytes_rate                    []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"net_in_bytes_rate,omitempty"`
	Net_out_bytes_rate                   []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"net_out_bytes_rate,omitempty"`
	Network_incoming_packets_rate_inband []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"network_incoming_packets_rate_inband,omitempty"`
	Network_outing_packets_rate_inband   []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"network_outing_packets_rate_inband,omitempty"`
	Network_incoming_errs_rate_inband    []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"network_incoming_errs_rate_inband,omitempty"`
	Network_outing_drop_rate_inband      []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"network_outing_drop_rate_inband,omitempty"`
	Network_outing_errs_rate_inband      []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"network_outing_errs_rate_inband,omitempty"`
	Network_incoming_drop_rate_inband    []EcsVmNetworkHistoryMetricDataAggregateListItemValueRealResponse `json:"network_incoming_drop_rate_inband,omitempty"`
}

type EcsVmNetworkHistoryMetricDataResultRealResponse struct {
	FUID              string                                                 `json:"fUID,omitempty"`
	FuserLastUpdated  string                                                 `json:"fuserLastUpdated,omitempty"`
	RegionID          string                                                 `json:"regionID,omitempty"`
	DeviceUUID        string                                                 `json:"deviceUUID,omitempty"`
	ItemAggregateList EcsVmNetworkHistoryMetricDataAggregateListRealResponse `json:"itemAggregateList,omitempty"`
}

type EcsVmNetworkHistoryMetricDataRealResponse struct {
	Result       []EcsVmNetworkHistoryMetricDataResultRealResponse `json:"result,omitempty"`
	CurrentCount int                                               `json:"currentCount,omitempty"`
	TotalCount   int                                               `json:"totalCount,omitempty"`
	TotalPage    int                                               `json:"totalPage,omitempty"`
	PageSize     int                                               `json:"pageSize,omitempty"`
	Page         int                                               `json:"page,omitempty"`
}

type EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse struct {
	Value        string
	SamplingTime int
}

type EcsVmNetworkHistoryMetricDataAggregateListResponse struct {
	Net_in_bytes_rate                    []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
	Net_out_bytes_rate                   []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
	Network_incoming_packets_rate_inband []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
	Network_outing_packets_rate_inband   []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
	Network_incoming_errs_rate_inband    []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
	Network_outing_drop_rate_inband      []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
	Network_outing_errs_rate_inband      []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
	Network_incoming_drop_rate_inband    []EcsVmNetworkHistoryMetricDataAggregateListItemValueResponse
}

type EcsVmNetworkHistoryMetricDataResultResponse struct {
	FUID              string
	FuserLastUpdated  string
	RegionID          string
	DeviceUUID        string
	ItemAggregateList EcsVmNetworkHistoryMetricDataAggregateListResponse
}

type EcsVmNetworkHistoryMetricDataResponse struct {
	Result       []EcsVmNetworkHistoryMetricDataResultResponse
	CurrentCount int
	TotalCount   int
	TotalPage    int
	PageSize     int
	Page         int
}
