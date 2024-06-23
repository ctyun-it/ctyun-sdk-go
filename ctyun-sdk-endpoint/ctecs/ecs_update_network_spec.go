package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsUpdateNetworkSpecApi 云主机修改带宽
// https://www.ctyun.cn/document/10026730/10106610

type EcsUpdateNetworkSpecApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsUpdateNetworkSpecApi(client *ctyunsdk.CtyunClient) *EcsUpdateNetworkSpecApi {
	return &EcsUpdateNetworkSpecApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/update-network-spec",
		},
	}
}

func (this *EcsUpdateNetworkSpecApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsUpdateNetworkSpecRequest) (*EcsUpdateNetworkSpecResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsUpdateNetworkSpecRealRequest{
		RegionID:    req.RegionID,
		InstanceID:  req.InstanceID,
		Bandwidth:   req.Bandwidth,
		ClientToken: req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsUpdateNetworkSpecRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsUpdateNetworkSpecResponse{
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
		RegionID:      realResponse.RegionID,
	}, nil
}

type EcsUpdateNetworkSpecRealRequest struct {
	RegionID    string `json:"regionID,omitempty"`
	InstanceID  string `json:"instanceID,omitempty"`
	Bandwidth   int    `json:"bandwidth,omitempty"`
	ClientToken string `json:"clientToken,omitempty"`
}

type EcsUpdateNetworkSpecRequest struct {
	RegionID    string
	InstanceID  string
	Bandwidth   int
	ClientToken string
}

type EcsUpdateNetworkSpecRealResponse struct {
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
	RegionID      string `json:"regionID,omitempty"`
}

type EcsUpdateNetworkSpecResponse struct {
	MasterOrderID string
	MasterOrderNO string
	RegionID      string
}
