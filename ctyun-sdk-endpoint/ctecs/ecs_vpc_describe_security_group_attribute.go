package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVpcDescribeSecurityGroupAttributeApi 查询用户安全组详情
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5574&data=87&isNormal=1
type EcsVpcDescribeSecurityGroupAttributeApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVpcDescribeSecurityGroupAttributeApi(client *ctyunsdk.CtyunClient) *EcsVpcDescribeSecurityGroupAttributeApi {
	return &EcsVpcDescribeSecurityGroupAttributeApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/vpc/describe-security-group-attribute",
		},
	}
}

func (this *EcsVpcDescribeSecurityGroupAttributeApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVpcDescribeSecurityGroupAttributeRequest) (*EcsVpcDescribeSecurityGroupAttributeResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", *req.RegionID).
		AddParam("securityGroupID", *req.SecurityGroupID).
		AddParam("projectID", *req.ProjectID).
		AddParam("direction", *req.Direction)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVpcDescribeSecurityGroupAttributeRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var securityGroupRuleList []EcsVpcDescribeSecurityGroupAttributeSecurityGroupRuleListResponse
	for _, res := range realResponse.SecurityGroupRuleList {
		securityGroupRuleList = append(securityGroupRuleList, EcsVpcDescribeSecurityGroupAttributeSecurityGroupRuleListResponse{
			Direction:       res.Direction,
			Action:          res.Action,
			Origin:          res.Origin,
			Priority:        res.Priority,
			Ethertype:       res.Ethertype,
			Protocol:        res.Protocol,
			Range:           res.Range,
			DestCidrIp:      res.DestCidrIp,
			Description:     res.Description,
			CreateTime:      res.CreateTime,
			Id:              res.Id,
			SecurityGroupID: res.SecurityGroupID,
		})
	}

	return &EcsVpcDescribeSecurityGroupAttributeResponse{
		SecurityGroupName:     realResponse.SecurityGroupName,
		Id:                    realResponse.Id,
		VmNum:                 realResponse.VmNum,
		Origin:                realResponse.Origin,
		VpcName:               realResponse.VpcName,
		VpcID:                 realResponse.VpcID,
		CreationTime:          realResponse.CreationTime,
		Description:           realResponse.Description,
		SecurityGroupRuleList: securityGroupRuleList,
	}, nil
}

type EcsVpcDescribeSecurityGroupAttributeRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	SecurityGroupID *string `json:"securityGroupID,omitempty"`
	ProjectID       *string `json:"projectID,omitempty"`
	Direction       *string `json:"direction,omitempty"`
}

type EcsVpcDescribeSecurityGroupAttributeRequest struct {
	RegionID        *string
	SecurityGroupID *string
	ProjectID       *string
	Direction       *string
}

type EcsVpcDescribeSecurityGroupAttributeSecurityGroupRuleListRealResponse struct {
	Direction       string `json:"direction,omitempty"`
	Action          string `json:"action,omitempty"`
	Origin          string `json:"origin,omitempty"`
	Priority        int    `json:"priority,omitempty"`
	Ethertype       string `json:"ethertype,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Range           string `json:"range,omitempty"`
	DestCidrIp      string `json:"destCidrIp,omitempty"`
	Description     string `json:"description,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
	Id              string `json:"id,omitempty"`
	SecurityGroupID string `json:"securityGroupID,omitempty"`
}

type EcsVpcDescribeSecurityGroupAttributeRealResponse struct {
	SecurityGroupName     string                                                                  `json:"securityGroupName,omitempty"`
	Id                    string                                                                  `json:"id,omitempty"`
	VmNum                 int                                                                     `json:"vmNum,omitempty"`
	Origin                string                                                                  `json:"origin,omitempty"`
	VpcName               string                                                                  `json:"vpcName,omitempty"`
	VpcID                 string                                                                  `json:"vpcID,omitempty"`
	CreationTime          string                                                                  `json:"creationTime,omitempty"`
	Description           string                                                                  `json:"description,omitempty"`
	SecurityGroupRuleList []EcsVpcDescribeSecurityGroupAttributeSecurityGroupRuleListRealResponse `json:"securityGroupRuleList,omitempty"`
}

type EcsVpcDescribeSecurityGroupAttributeSecurityGroupRuleListResponse struct {
	Direction       string
	Action          string
	Origin          string
	Priority        int
	Ethertype       string
	Protocol        string
	Range           string
	DestCidrIp      string
	Description     string
	CreateTime      string
	Id              string
	SecurityGroupID string
}

type EcsVpcDescribeSecurityGroupAttributeResponse struct {
	SecurityGroupName     string
	Id                    string
	VmNum                 int
	Origin                string
	VpcName               string
	VpcID                 string
	CreationTime          string
	Description           string
	SecurityGroupRuleList []EcsVpcDescribeSecurityGroupAttributeSecurityGroupRuleListResponse
}
