package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
	"strconv"
)

// EbmDeviceStockListApi D:\Project\go-sdk-auto-write\docs\查询物理机库存
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4575&data=97&isNormal=1
type EbmDeviceStockListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmDeviceStockListApi(client *ctyunsdk.CtyunClient) *EbmDeviceStockListApi {
	return &EbmDeviceStockListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ebm/device-stock-list",
		},
	}
}

func (this *EbmDeviceStockListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmDeviceStockListRequest) (*EbmDeviceStockListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", req.RegionID).
		AddParam("azName", req.AzName).
		AddParam("deviceType", req.DeviceType)

	if req.Count != nil {
		builder.AddParam("count", strconv.Itoa(*req.Count))
	}
	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmDeviceStockListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EbmDeviceStockListResultsResponse
	for _, res := range realResponse.Results {
		results = append(results, EbmDeviceStockListResultsResponse{
			Available: res.Available,
			Success:   res.Success,
		})
	}

	return &EbmDeviceStockListResponse{
		TotalCount: realResponse.TotalCount,
		Results:    results,
	}, nil
}

type EbmDeviceStockListRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	AzName     string `json:"azName,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	Count      *int   `json:"count,omitempty"`
}

type EbmDeviceStockListRequest struct {
	RegionID   string
	AzName     string
	DeviceType string
	Count      *int
}

type EbmDeviceStockListResultsRealResponse struct {
	Available int  `json:"available,omitempty"`
	Success   bool `json:"success,omitempty"`
}

type EbmDeviceStockListRealResponse struct {
	TotalCount int                                     `json:"totalCount,omitempty"`
	Results    []EbmDeviceStockListResultsRealResponse `json:"results,omitempty"`
}

type EbmDeviceStockListResultsResponse struct {
	Available int
	Success   bool
}

type EbmDeviceStockListResponse struct {
	TotalCount int
	Results    []EbmDeviceStockListResultsResponse
}
