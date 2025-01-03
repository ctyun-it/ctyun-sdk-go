package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBatchUnsubscribeInstanceApi 批量释放云主机
// https://www.ctyun.cn/document/10026730/10040177
type EcsBatchUnsubscribeInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBatchUnsubscribeInstanceApi(client *ctyunsdk.CtyunClient) *EcsBatchUnsubscribeInstancesApi {
	return &EcsBatchUnsubscribeInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/batch-delete",
		},
	}
}

func (this *EcsBatchUnsubscribeInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBatchUnsubscribeInstanceRequest) (*EcsBatchUnsubscribeInstanceResponse, *EcsBatchUnsubscribeInstanceResponse) {
	builder := this.WithCredential(&credential)
	_, err := builder.WriteJson(&ecsBatchUnsubscribeInstanceRealRequest{
		ClientToken:    req.ClientToken,
		RegionID:       req.RegionID,
		InstanceIDList: req.InstanceIDList,
	})
	if err != nil {
		return nil, nil
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, nil
	}

	var realResponse ecsBatchUnsubscribeInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, nil
	}
	return &EcsBatchUnsubscribeInstanceResponse{
		MasterOrderID: realResponse.MasterOrderID,
		RegionID:      realResponse.RegionID,
		MasterOrderNo: realResponse.MasterOrderNO,
	}, nil
}

type ecsBatchUnsubscribeInstanceRealRequest struct {
	ClientToken    *string `json:"clientToken,omitempty"`
	RegionID       *string `json:"regionID,omitempty"`
	InstanceIDList *string `json:"instanceIDList,omitempty"`
}

type EcsBatchUnsubscribeInstanceRequest struct {
	ClientToken    *string
	RegionID       *string
	InstanceIDList *string
}

type ecsBatchUnsubscribeInstanceRealResponse struct {
	MasterOrderNO string `json:"masterOrderNO"`
	RegionID      string `json:"regionID"`
	MasterOrderID string `json:"masterOrderID"`
}

type EcsBatchUnsubscribeInstanceResponse struct {
	MasterOrderNo string
	RegionID      string
	MasterOrderID string
}
