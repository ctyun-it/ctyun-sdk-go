package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmDeleteApi D:\Project\go-sdk-auto-write\docs\物理机删除
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4908&data=97&isNormal=1
type EbmDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmDeleteApi(client *ctyunsdk.CtyunClient) *EbmDeleteApi {
	return &EbmDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/delete",
		},
	}
}

func (this *EbmDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmDeleteRequest) (*EbmDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EbmDeleteRealRequest{
		RegionID:     req.RegionID,
		AzName:       req.AzName,
		InstanceUUID: req.InstanceUUID,
		ClientToken:  req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EbmDeleteResponse{
		RegionID:      realResponse.RegionID,
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
	}, nil
}

type EbmDeleteRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	InstanceUUID string `json:"instanceUUID,omitempty"`
	ClientToken  string `json:"clientToken,omitempty"`
}

type EbmDeleteRequest struct {
	RegionID     string
	AzName       string
	InstanceUUID string
	ClientToken  string
}

type EbmDeleteRealResponse struct {
	RegionID      string `json:"regionID,omitempty"`
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
}

type EbmDeleteResponse struct {
	RegionID      string
	MasterOrderID string
	MasterOrderNO string
}
