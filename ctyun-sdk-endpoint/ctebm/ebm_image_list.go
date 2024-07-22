package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
	"strconv"
)

// EbmImageListApi D:\Project\go-sdk-auto-write\docs\查询物理机可支持的镜像
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4577&data=97&isNormal=1
type EbmImageListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmImageListApi(client *ctyunsdk.CtyunClient) *EbmImageListApi {
	return &EbmImageListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ebm/image-list",
		},
	}
}

func (this *EbmImageListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmImageListRequest) (*EbmImageListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", req.RegionID).
		AddParam("azName", req.AzName).
		AddParam("deviceType", req.DeviceType).
		AddParam("imageType", req.ImageType).
		AddParam("imageUUID", req.ImageUUID).
		AddParam("osName", req.OsName).
		AddParam("osVersion", req.OsVersion).
		AddParam("osType", req.OsType)
	if req.PageNo != nil {
		builder.AddParam("pageNo", strconv.Itoa(*req.PageNo))
	}
	if req.PageSize != nil {
		builder.AddParam("pageSize", strconv.Itoa(*req.PageSize))
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmImageListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EbmImageListResultsResponse
	for _, res := range realResponse.Results {
		var osRealResponse = res.Os
		results = append(results, EbmImageListResultsResponse{
			NameZh:     res.NameZh,
			Format:     res.Format,
			ImageType:  res.ImageType,
			IsShared:   res.IsShared,
			Version:    res.Version,
			ImageUUID:  res.ImageUUID,
			NameEn:     res.NameEn,
			LayoutType: res.LayoutType,
			Os: EbmImageListOsResponse{
				Uuid:         osRealResponse.Uuid,
				SuperUser:    osRealResponse.SuperUser,
				Platform:     osRealResponse.Platform,
				Version:      osRealResponse.Version,
				Architecture: osRealResponse.Architecture,
				NameEn:       osRealResponse.NameEn,
				Bits:         osRealResponse.Bits,
				OsType:       osRealResponse.OsType,
				NameZh:       osRealResponse.NameZh,
			},
		})
	}

	return &EbmImageListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EbmImageListRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	AzName     string `json:"azName,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	ImageType  string `json:"imageType,omitempty"`
	ImageUUID  string `json:"imageUUID,omitempty"`
	OsName     string `json:"osName,omitempty"`
	OsVersion  string `json:"osVersion,omitempty"`
	OsType     string `json:"osType,omitempty"`
	PageNo     *int   `json:"pageNo,omitempty"`
	PageSize   *int   `json:"pageSize,omitempty"`
}

type EbmImageListRequest struct {
	RegionID   string
	AzName     string
	DeviceType string
	ImageType  string
	ImageUUID  string
	OsName     string
	OsVersion  string
	OsType     string
	PageNo     *int
	PageSize   *int
}

type EbmImageListOsRealResponse struct {
	Uuid         string `json:"uuid,omitempty"`
	SuperUser    string `json:"superUser,omitempty"`
	Platform     string `json:"platform,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	NameEn       string `json:"nameEn,omitempty"`
	Bits         int    `json:"bits,omitempty"`
	OsType       string `json:"osType,omitempty"`
	NameZh       string `json:"nameZh,omitempty"`
}

type EbmImageListResultsRealResponse struct {
	NameZh     string                     `json:"nameZh,omitempty"`
	Format     string                     `json:"format,omitempty"`
	ImageType  string                     `json:"imageType,omitempty"`
	IsShared   bool                       `json:"isShared,omitempty"`
	Version    string                     `json:"version,omitempty"`
	ImageUUID  string                     `json:"imageUUID,omitempty"`
	NameEn     string                     `json:"nameEn,omitempty"`
	LayoutType string                     `json:"layoutType,omitempty"`
	Os         EbmImageListOsRealResponse `json:"os,omitempty"`
}

type EbmImageListRealResponse struct {
	CurrentCount int                               `json:"currentCount,omitempty"`
	TotalCount   int                               `json:"totalCount,omitempty"`
	TotalPage    int                               `json:"totalPage,omitempty"`
	Results      []EbmImageListResultsRealResponse `json:"results,omitempty"`
}

type EbmImageListOsResponse struct {
	Uuid         string
	SuperUser    string
	Platform     string
	Version      string
	Architecture string
	NameEn       string
	Bits         int
	OsType       string
	NameZh       string
}

type EbmImageListResultsResponse struct {
	NameZh     string
	Format     string
	ImageType  string
	IsShared   bool
	Version    string
	ImageUUID  string
	NameEn     string
	LayoutType string
	Os         EbmImageListOsResponse
}

type EbmImageListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EbmImageListResultsResponse
}
