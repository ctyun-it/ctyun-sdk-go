package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsUpdateInstanceApi
type EcsUpdateInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsUpdateInstanceApi(client *ctyunsdk.CtyunClient) *EcsUpdateInstanceApi {
	return &EcsUpdateInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/update-instance",
		},
	}
}

func (this *EcsUpdateInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsUpdateInstanceRequest) (*EcsUpdateInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsUpdateInstanceRealRequest{
		RegionID:    req.RegionID,
		InstanceID:  req.InstanceID,
		DisplayName: req.DisplayName,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsUpdateInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsUpdateInstanceResponse{
		InstanceID:  realResponse.InstanceID,
		DisplayName: realResponse.DisplayName,
	}, nil
}

type EcsUpdateInstanceRealRequest struct {
	RegionID    *string `json:"regionID,omitempty"`
	InstanceID  *string `json:"instanceID,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

type EcsUpdateInstanceRequest struct {
	RegionID    *string
	InstanceID  *string
	DisplayName *string
}

type EcsUpdateInstanceRealResponse struct {
	InstanceID  string `json:"instanceID,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type EcsUpdateInstanceResponse struct {
	InstanceID  string
	DisplayName string
}
