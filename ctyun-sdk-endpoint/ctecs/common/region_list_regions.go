package common

import (
	"context"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctecs"
	"net/http"
)

// RegionListRegionsApi 资源池列表查询
// https://www.ctyun.cn/document/10026730/10040588
type RegionListRegionsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewRegionListRegionsApi(client *ctyunsdk.CtyunClient) *RegionListRegionsApi {
	return &RegionListRegionsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/region/list-regions",
		},
	}
}

func (this *RegionListRegionsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *RegionListRequest) (*RegionListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.AddParam("regionName", req.RegionName)

	response, err := this.client.RequestToEndpoint(ctx, ctecs.EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse RegionListRegionsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)

	if err != nil {
		return nil, err
	}
	var regionList []RegionListRegionsResponse
	for _, region := range realResponse.RegionList {
		regionList = append(regionList, RegionListRegionsResponse{
			RegionParent: region.RegionParent,
			RegionId:     region.RegionID,
			RegionType:   region.RegionType,
			RegionName:   region.RegionName,
			IsMultiZones: region.IsMultiZones,
			ZoneList:     region.ZoneList,
		})
	}
	return &RegionListResponse{
		RegionList: regionList,
	}, nil
}

type RegionListRegionsRealResponse struct {
	RegionList []struct {
		RegionID     string   `json:"regionID"`     // 资源池ID
		RegionParent string   `json:"regionParent"` // 资源池所属省份
		RegionName   string   `json:"regionName"`   // 资源池名称
		RegionType   string   `json:"regionType"`   // 资源池类型
		IsMultiZones bool     `json:"isMultiZones"` // 是否多可用区资源池
		ZoneList     []string `json:"zoneList"`     // 可用区列表
	} `json:"regionList"`
}

type RegionListRequest struct {
	RegionName string // 资源池名称
}

type RegionListRegionsResponse struct {
	RegionId     string   // 资源池ID
	RegionParent string   // 资源池所属省份
	RegionName   string   // 资源池名称
	RegionType   string   // 资源池类型
	IsMultiZones bool     // 是否多可用区资源池
	ZoneList     []string // 可用区列表
}

type RegionListResponse struct {
	RegionList []RegionListRegionsResponse // 资源池对象
}
