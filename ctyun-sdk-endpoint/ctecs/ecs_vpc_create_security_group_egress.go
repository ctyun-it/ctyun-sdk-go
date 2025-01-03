package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVpcCreateSecurityGroupEgressApi 创建安全组出向规则
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5564&data=87&isNormal=1
type EcsVpcCreateSecurityGroupEgressApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVpcCreateSecurityGroupEgressApi(client *ctyunsdk.CtyunClient) *EcsVpcCreateSecurityGroupEgressApi {
	return &EcsVpcCreateSecurityGroupEgressApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/vpc/create-security-group-egress",
		},
	}
}

func (this *EcsVpcCreateSecurityGroupEgressApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVpcCreateSecurityGroupEgressRequest) (*EcsVpcCreateSecurityGroupEgressResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var securityGroupRules []EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRealRequest
	for _, request := range req.SecurityGroupRules {
		securityGroupRules = append(securityGroupRules, EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRealRequest{
			Direction:   request.Direction,
			Action:      request.Action,
			Priority:    request.Priority,
			Protocol:    request.Protocol,
			Ethertype:   request.Ethertype,
			DestCidrIp:  request.DestCidrIp,
			Description: request.Description,
			Range:       request.Range,
		})
	}

	_, err := builder.WriteJson(&EcsVpcCreateSecurityGroupEgressRealRequest{
		RegionID:           req.RegionID,
		SecurityGroupID:    req.SecurityGroupID,
		SecurityGroupRules: securityGroupRules,
		ClientToken:        req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVpcCreateSecurityGroupEgressRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVpcCreateSecurityGroupEgressResponse{}, nil
}

type EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRealRequest struct {
	Direction   *string `json:"direction,omitempty"`
	Action      *string `json:"action,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
	Protocol    *string `json:"protocol,omitempty"`
	Ethertype   *string `json:"ethertype,omitempty"`
	DestCidrIp  *string `json:"destCidrIp,omitempty"`
	Description *string `json:"description,omitempty"`
	Range       *string `json:"range,omitempty"`
}

type EcsVpcCreateSecurityGroupEgressRealRequest struct {
	RegionID           *string                                                        `json:"regionID,omitempty"`
	SecurityGroupID    *string                                                        `json:"securityGroupID,omitempty"`
	SecurityGroupRules []EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRealRequest `json:"securityGroupRules,omitempty"`
	ClientToken        *string                                                        `json:"clientToken,omitempty"`
}

type EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRequest struct {
	Direction   *string
	Action      *string
	Priority    *int
	Protocol    *string
	Ethertype   *string
	DestCidrIp  *string
	Description *string
	Range       *string
}

type EcsVpcCreateSecurityGroupEgressRequest struct {
	RegionID           *string
	SecurityGroupID    *string
	SecurityGroupRules []EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRequest
	ClientToken        *string
}

type EcsVpcCreateSecurityGroupEgressRealResponse struct {
}

type EcsVpcCreateSecurityGroupEgressResponse struct {
}
