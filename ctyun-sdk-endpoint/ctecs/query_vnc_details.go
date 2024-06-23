package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// QueryVncDetailsApi
type QueryVncDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewQueryVncDetailsApi(client *ctyunsdk.CtyunClient) *QueryVncDetailsApi {
	return &QueryVncDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/vnc/details",
		},
	}
}

func (this *QueryVncDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *QueryVncDetailsRequest) (*QueryVncDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&QueryVncDetailsRealRequest{
		RegionID:   req.RegionID,
		InstanceID: req.InstanceID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse QueryVncDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &QueryVncDetailsResponse{
		Token: realResponse.Token,
	}, nil
}

type QueryVncDetailsRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	InstanceID string `json:"instanceID,omitempty"`
}

type QueryVncDetailsRequest struct {
	RegionID   string
	InstanceID string
}

type QueryVncDetailsRealResponse struct {
	Token string `json:"token,omitempty"`
}

type QueryVncDetailsResponse struct {
	Token string
}
