package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmRebuildApi D:\Project\go-sdk-auto-write\docs\物理机重装系统
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4582&data=97&isNormal=1
type EbmRebuildApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmRebuildApi(client *ctyunsdk.CtyunClient) *EbmRebuildApi {
	return &EbmRebuildApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/rebuild",
		},
	}
}

func (this *EbmRebuildApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmRebuildRequest) (*EbmRebuildResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EbmRebuildRealRequest{
		RegionID:             req.RegionID,
		AzName:               req.AzName,
		InstanceUUID:         req.InstanceUUID,
		Hostname:             req.Hostname,
		Password:             req.Password,
		ImageUUID:            req.ImageUUID,
		SystemVolumeRaidUUID: req.SystemVolumeRaidUUID,
		DataVolumeRaidUUID:   req.DataVolumeRaidUUID,
		RedoRaid:             req.RedoRaid,
		UserData:             req.UserData,
		KeyName:              req.KeyName,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmRebuildRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EbmRebuildResponse{}, nil
}

type EbmRebuildRealRequest struct {
	RegionID             string `json:"regionID,omitempty"`
	AzName               string `json:"azName,omitempty"`
	InstanceUUID         string `json:"instanceUUID,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Password             string `json:"password,omitempty"`
	ImageUUID            string `json:"imageUUID,omitempty"`
	SystemVolumeRaidUUID string `json:"systemVolumeRaidUUID,omitempty"`
	DataVolumeRaidUUID   string `json:"dataVolumeRaidUUID,omitempty"`
	RedoRaid             *bool  `json:"redoRaid,omitempty"`
	UserData             string `json:"userData,omitempty"`
	KeyName              string `json:"keyName,omitempty"`
}

type EbmRebuildRequest struct {
	RegionID             string
	AzName               string
	InstanceUUID         string
	Hostname             string
	Password             string
	ImageUUID            string
	SystemVolumeRaidUUID string
	DataVolumeRaidUUID   string
	RedoRaid             *bool
	UserData             string
	KeyName              string
}

type EbmRebuildRealResponse struct {
}

type EbmRebuildResponse struct {
}
