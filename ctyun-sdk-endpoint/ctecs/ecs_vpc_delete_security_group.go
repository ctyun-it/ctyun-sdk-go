package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVpcDeleteSecurityGroupApi 删除安全组
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5561&data=87&isNormal=1
type EcsVpcDeleteSecurityGroupApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVpcDeleteSecurityGroupApi(client *ctyunsdk.CtyunClient) *EcsVpcDeleteSecurityGroupApi {
	return &EcsVpcDeleteSecurityGroupApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vpc/delete-security-group",
		},
	}
}

func (this *EcsVpcDeleteSecurityGroupApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVpcDeleteSecurityGroupRequest) (*EcsVpcDeleteSecurityGroupResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVpcDeleteSecurityGroupRealRequest{
		ClientToken:     req.ClientToken,
		RegionID:        req.RegionID,
		ProjectID:       req.ProjectID,
		SecurityGroupID: req.SecurityGroupID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVpcDeleteSecurityGroupRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVpcDeleteSecurityGroupResponse{}, nil
}

type EcsVpcDeleteSecurityGroupRealRequest struct {
	ClientToken     *string `json:"clientToken,omitempty"`
	RegionID        *string `json:"regionID,omitempty"`
	ProjectID       *string `json:"projectID,omitempty"`
	SecurityGroupID *string `json:"securityGroupID,omitempty"`
}

type EcsVpcDeleteSecurityGroupRequest struct {
	ClientToken     *string
	RegionID        *string
	ProjectID       *string
	SecurityGroupID *string
}

type EcsVpcDeleteSecurityGroupRealResponse struct {
}

type EcsVpcDeleteSecurityGroupResponse struct {
}
