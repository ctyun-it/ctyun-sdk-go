package common

import (
	"context"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctecs"
	"net/http"
)

// RegionGetZonesApi 资源池可用区查询
// https://www.ctyun.cn/document/10026730/10040590
type RegionGetZonesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewRegionGetZonesApi(client *ctyunsdk.CtyunClient) *RegionGetZonesApi {
	return &RegionGetZonesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/region/get-zones",
		},
	}
}

func (this *RegionGetZonesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *RegionGetZonesRequest) (*RegionGetZonesListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.AddParam("regionID", req.RegionID)

	response, err := this.client.RequestToEndpoint(ctx, ctecs.EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse RegionGetZonesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)

	if err != nil {
		return nil, err
	}
	var zoneList []RegionGetZonesResponse
	for _, region := range realResponse.ZoneList {
		zoneList = append(zoneList, RegionGetZonesResponse{
			Name:          region.Name,
			AzDisplayName: region.AzDisplayName,
		})
	}
	return &RegionGetZonesListResponse{
		ZoneList: zoneList,
	}, nil
}

type RegionGetZonesRealResponse struct {
	ZoneList []struct {
		Name          string `json:"name"`          // 可用区名称，其他需要可用区参数的接口需要依赖该名称结果
		AzDisplayName string `json:"azDisplayName"` // 可用区展示名
	} `json:"zoneList"`
}

type RegionGetZonesRequest struct {
	RegionID string // 资源池ID
}

type RegionGetZonesResponse struct {
	Name          string // 可用区名称，其他需要可用区参数的接口需要依赖该名称结果
	AzDisplayName string // 可用区展示名
}

type RegionGetZonesListResponse struct {
	ZoneList []RegionGetZonesResponse // 可用区列表
}
