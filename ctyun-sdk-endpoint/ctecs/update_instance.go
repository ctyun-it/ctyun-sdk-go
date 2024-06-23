package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// UpdateInstanceApi
type UpdateInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewUpdateInstanceApi(client *ctyunsdk.CtyunClient) *UpdateInstanceApi {
	return &UpdateInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/update-instance",
		},
	}
}

func (this *UpdateInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *UpdateInstanceRequest) (*UpdateInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&UpdateInstanceRealRequest{
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

	var realResponse UpdateInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &UpdateInstanceResponse{
		InstanceID:  realResponse.InstanceID,
		DisplayName: realResponse.DisplayName,
	}, nil
}

type UpdateInstanceRealRequest struct {
	RegionID    string `json:"regionID,omitempty"`
	InstanceID  string `json:"instanceID,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type UpdateInstanceRequest struct {
	RegionID    string
	InstanceID  string
	DisplayName string
}

type UpdateInstanceRealResponse struct {
	InstanceID  string `json:"instanceID,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type UpdateInstanceResponse struct {
	InstanceID  string
	DisplayName string
}
