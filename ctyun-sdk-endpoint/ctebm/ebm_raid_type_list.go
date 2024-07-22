package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmRaidTypeListApi D:\Project\go-sdk-auto-write\docs\查询物理机本地盘可选择的raid类型
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4576&data=97&isNormal=1
type EbmRaidTypeListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmRaidTypeListApi(client *ctyunsdk.CtyunClient) *EbmRaidTypeListApi {
	return &EbmRaidTypeListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ebm/raid-type-list",
		},
	}
}

func (this *EbmRaidTypeListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmRaidTypeListRequest) (*EbmRaidTypeListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", req.RegionID).
		AddParam("azName", req.AzName).
		AddParam("deviceType", req.DeviceType).
		AddParam("volumeType", req.VolumeType)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmRaidTypeListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EbmRaidTypeListResultsResponse
	for _, res := range realResponse.Results {
		results = append(results, EbmRaidTypeListResultsResponse{
			DeviceType:    res.DeviceType,
			VolumeType:    res.VolumeType,
			Uuid:          res.Uuid,
			NameEn:        res.NameEn,
			NameZh:        res.NameZh,
			VolumeDetail:  res.VolumeDetail,
			DescriptionEn: res.DescriptionEn,
			DescriptionZh: res.DescriptionZh,
		})
	}

	return &EbmRaidTypeListResponse{
		TotalCount: realResponse.TotalCount,
		Results:    results,
	}, nil
}

type EbmRaidTypeListRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	AzName     string `json:"azName,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	VolumeType string `json:"volumeType,omitempty"`
}

type EbmRaidTypeListRequest struct {
	RegionID   string
	AzName     string
	DeviceType string
	VolumeType string
}

type EbmRaidTypeListResultsRealResponse struct {
	DeviceType    string `json:"deviceType,omitempty"`
	VolumeType    string `json:"volumeType,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
	NameEn        string `json:"nameEn,omitempty"`
	NameZh        string `json:"nameZh,omitempty"`
	VolumeDetail  string `json:"volumeDetail,omitempty"`
	DescriptionEn string `json:"descriptionEn,omitempty"`
	DescriptionZh string `json:"descriptionZh,omitempty"`
}

type EbmRaidTypeListRealResponse struct {
	TotalCount int                                  `json:"totalCount,omitempty"`
	Results    []EbmRaidTypeListResultsRealResponse `json:"results,omitempty"`
}

type EbmRaidTypeListResultsResponse struct {
	DeviceType    string
	VolumeType    string
	Uuid          string
	NameEn        string
	NameZh        string
	VolumeDetail  string
	DescriptionEn string
	DescriptionZh string
}

type EbmRaidTypeListResponse struct {
	TotalCount int
	Results    []EbmRaidTypeListResultsResponse
}
