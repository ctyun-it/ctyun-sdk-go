package ctvpc

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
	"strconv"
)

// SecurityGroupQueryApi 创建安全组
// https://www.ctyun.cn/document/10026755/10040907
type SecurityGroupQueryApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewSecurityGroupQueryApi(client *ctyunsdk.CtyunClient) *SecurityGroupQueryApi {
	return &SecurityGroupQueryApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/vpc/new-query-security-groups",
		},
	}
}

func (this *SecurityGroupQueryApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *SecurityGroupQueryRequest) (*SecurityGroupQueryResponse, ctyunsdk.CtyunRequestError) {
	builder := this.
		WithCredential(&credential).
		AddParam("regionID", req.RegionId)

	if len(req.VpcID) > 0 {
		builder.AddParam("vpcID", req.VpcID)
	}

	if len(req.QueryContent) > 0 {
		builder.AddParam("queryContent", req.QueryContent)
	}

	if len(req.InstanceID) > 0 {
		builder.AddParam("instanceID", req.InstanceID)
	}

	if req.PageNumber > 0 {
		builder.AddParam("pageNumber", strconv.Itoa(req.PageNumber))
	}

	if req.PageSize > 0 {
		builder.AddParam("pageSize", strconv.Itoa(req.PageSize))
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtvpc, builder)
	if err != nil {
		return nil, err
	}

	result := &SecurityGroupQueryResponse{}
	err = response.ParseByStandardModelWithCheck(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type securityGroupQueryRealRequest struct {
	RegionID     string `json:"regionID"`
	VpcID        string `json:"vpcID,omitempty"`
	QueryContent string `json:"queryContent,omitempty"`
	InstanceID   string `json:"instanceID,omitempty"`
	PageNumber   int    `json:"pageNumber,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

type SecurityGroupQueryRequest struct {
	RegionId     string // 资源池id
	VpcID        string // vpcId
	QueryContent string // 根据名字和 id 查询
	InstanceID   string // 实例 id
	PageNumber   int    // 页码
	PageSize     int    // 分页大小
}

type SecurityGroupRule struct {
	Direction       string `json:"direction"`
	Priority        int    `json:"priority"`
	EtherType       string `json:"etherType"`
	Protocol        string `json:"protocol"`
	Range           string `json:"range"`
	DestCidrIp      string `json:"destCidrIp"`
	Description     string `json:"description"`
	CreateTime      string `json:"createTime"`
	ID              string `json:"id"`
	SecurityGroupID string `json:"securityGroupID"`
	Action          string `json:"action"`
}

type SecurityGroup struct {
	SecurityGroupName     string              `json:"securityGroupName"`
	ID                    string              `json:"id"`
	VmNum                 int                 `json:"vmNum"`
	Origin                string              `json:"origin"`
	VpcName               string              `json:"vpcName"`
	VpcID                 string              `json:"vpcID"`
	CreationTime          string              `json:"creationTime"`
	Description           string              `json:"description"`
	SecurityGroupRuleList []SecurityGroupRule `json:"securityGroupRuleList"`
}

type SecurityGroupQueryResponse struct {
	SecurityGroups []SecurityGroup `json:"securityGroups"`
	CurrentCount   int             `json:"currentCount"`
	TotalCount     int             `json:"totalCount"`
	TotalPage      int             `json:"totalPage"`
}
