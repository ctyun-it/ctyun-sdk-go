package ctecs

import (
	"context"
	"fmt"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVpcQuerySecurityGroupsApi 查询用户安全组列表
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5559&data=87&isNormal=1
type EcsVpcQuerySecurityGroupsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVpcQuerySecurityGroupsApi(client *ctyunsdk.CtyunClient) *EcsVpcQuerySecurityGroupsApi {
	return &EcsVpcQuerySecurityGroupsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/vpc/query-security-groups",
		},
	}
}

func (this *EcsVpcQuerySecurityGroupsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVpcQuerySecurityGroupsRequest) (*[]EcsVpcQuerySecurityGroupsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", *req.RegionID).
		AddParam("vpcID", *req.VpcID).
		AddParam("queryContent", *req.QueryContent).
		AddParam("projectID", *req.ProjectID).
		AddParam("instanceID", *req.InstanceID).
		AddParam("pageNumber", fmt.Sprintf("%d", *req.PageNumber)).
		AddParam("pageNo", fmt.Sprintf("%d", *req.PageNo)).
		AddParam("pageSize", fmt.Sprintf("%d", *req.PageSize))

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse []EcsVpcQuerySecurityGroupsResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	for _, r := range realResponse {
		var securityGroupRuleList []EcsVpcQuerySecurityGroupsSecurityGroupRuleListResponse
		for _, res := range r.SecurityGroupRuleList {
			securityGroupRuleList = append(securityGroupRuleList, EcsVpcQuerySecurityGroupsSecurityGroupRuleListResponse{
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

		realResponse = append(realResponse, EcsVpcQuerySecurityGroupsResponse{
			SecurityGroupName:     r.SecurityGroupName,
			Id:                    r.Id,
			VmNum:                 r.VmNum,
			Origin:                r.Origin,
			VpcName:               r.VpcName,
			VpcID:                 r.VpcID,
			CreationTime:          r.CreationTime,
			Description:           r.Description,
			SecurityGroupRuleList: securityGroupRuleList,
		})
	}

	return &realResponse, nil
}

type EcsVpcQuerySecurityGroupsRealRequest struct {
	RegionID     *string `json:"regionID,omitempty"`
	VpcID        *string `json:"vpcID,omitempty"`
	QueryContent *string `json:"queryContent,omitempty"`
	ProjectID    *string `json:"projectID,omitempty"`
	InstanceID   *string `json:"instanceID,omitempty"`
	PageNumber   *string `json:"pageNumber,omitempty"`
	PageNo       *int    `json:"pageNo,omitempty"`
	PageSize     *int    `json:"pageSize,omitempty"`
}

type EcsVpcQuerySecurityGroupsRequest struct {
	RegionID     *string
	VpcID        *string
	QueryContent *string
	ProjectID    *string
	InstanceID   *string
	PageNumber   *int
	PageNo       *int
	PageSize     *int
}

type EcsVpcQuerySecurityGroupsSecurityGroupRuleListRealResponse struct {
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

type EcsVpcQuerySecurityGroupsRealResponse struct {
	SecurityGroupName     string                                                       `json:"securityGroupName,omitempty"`
	Id                    string                                                       `json:"id,omitempty"`
	VmNum                 int                                                          `json:"vmNum,omitempty"`
	Origin                string                                                       `json:"origin,omitempty"`
	VpcName               string                                                       `json:"vpcName,omitempty"`
	VpcID                 string                                                       `json:"vpcID,omitempty"`
	CreationTime          string                                                       `json:"creationTime,omitempty"`
	Description           string                                                       `json:"description,omitempty"`
	SecurityGroupRuleList []EcsVpcQuerySecurityGroupsSecurityGroupRuleListRealResponse `json:"securityGroupRuleList,omitempty"`
}

type EcsVpcQuerySecurityGroupsSecurityGroupRuleListResponse struct {
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

type EcsVpcQuerySecurityGroupsResponse struct {
	SecurityGroupName     string
	Id                    string
	VmNum                 int
	Origin                string
	VpcName               string
	VpcID                 string
	CreationTime          string
	Description           string
	SecurityGroupRuleList []EcsVpcQuerySecurityGroupsSecurityGroupRuleListResponse
}
