package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVpcCreateSecurityGroupApi 创建安全组
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5557&data=87&isNormal=1
type EcsVpcCreateSecurityGroupApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVpcCreateSecurityGroupApi(client *ctyunsdk.CtyunClient) *EcsVpcCreateSecurityGroupApi {
	return &EcsVpcCreateSecurityGroupApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vpc/create-security-group",
		},
	}
}

func (this *EcsVpcCreateSecurityGroupApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVpcCreateSecurityGroupRequest) (*EcsVpcCreateSecurityGroupResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVpcCreateSecurityGroupRealRequest{
		ClientToken: req.ClientToken,
		RegionID:    req.RegionID,
		ProjectID:   req.ProjectID,
		VpcID:       req.VpcID,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVpcCreateSecurityGroupRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVpcCreateSecurityGroupResponse{
		SecurityGroupID: realResponse.SecurityGroupID,
	}, nil
}

type EcsVpcCreateSecurityGroupRealRequest struct {
	ClientToken *string `json:"clientToken,omitempty"`
	RegionID    *string `json:"regionID,omitempty"`
	ProjectID   *string `json:"projectID,omitempty"`
	VpcID       *string `json:"vpcID,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type EcsVpcCreateSecurityGroupRequest struct {
	ClientToken *string
	RegionID    *string
	ProjectID   *string
	VpcID       *string
	Name        *string
	Description *string
}

type EcsVpcCreateSecurityGroupRealResponse struct {
	SecurityGroupID string `json:"securityGroupID,omitempty"`
}

type EcsVpcCreateSecurityGroupResponse struct {
	SecurityGroupID string
}
