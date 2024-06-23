package ctecs

import (
	"context"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsFlavorFamiliesListApi 查询云主机规格族列表

type EcsFlavorFamiliesListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsFlavorFamiliesList(client *ctyunsdk.CtyunClient) *EcsFlavorFamiliesListApi {
	return &EcsFlavorFamiliesListApi{
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/flavor-families/list",
		},
		client: client,
	}
}

func (this *EcsFlavorFamiliesListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsFlavorFamiliesListRequest) (*EcsFlavorFamiliesListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.AddParam("regionID", req.RegionID)
	builder.AddParam("azName", req.AzName)
	resp, requestError := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if requestError != nil {
		return nil, requestError
	}

	var realResponse EcsFlavorFamiliesListRealResponse
	err := resp.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsFlavorFamiliesListResponse{
		FlavorFamilyList: realResponse.FlavorFamilyList,
	}, nil
}

type EcsFlavorFamiliesListRequest struct {
	RegionID string
	AzName   string
}

type EcsFlavorFamiliesListRealResponse struct {
	FlavorFamilyList []string `json:"flavorFamilyList,omitempty"`
}

type EcsFlavorFamiliesListResponse struct {
	FlavorFamilyList []string
}
