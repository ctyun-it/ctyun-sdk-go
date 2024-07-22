package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsListInstancesApi 查询云主机列表
// https://www.ctyun.cn/document/10026730/10106328

type EcsListInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsListInstancesApi(client *ctyunsdk.CtyunClient) *EcsListInstancesApi {
	return &EcsListInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/list-instances",
		},
	}
}

func (this *EcsListInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsListInstancesRequest) (*EcsListInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var labelList []EcsListInstancesLabelListRealRequest
	for _, request := range req.LabelList {
		labelList = append(labelList, EcsListInstancesLabelListRealRequest{
			LabelKey:   request.LabelKey,
			LableValue: request.LabelValue,
		})
	}

	_, err := builder.WriteJson(&EcsListInstancesRealRequest{
		RegionID:        req.RegionID,
		AzName:          req.AzName,
		ProjectID:       req.ProjectID,
		PageNo:          req.PageNo,
		PageSize:        req.PageSize,
		State:           req.State,
		Keyword:         req.Keyword,
		InstanceName:    req.InstanceName,
		InstanceIDList:  req.InstanceIDList,
		SecurityGroupID: req.SecurityGroupID,
		VpcID:           req.VpcID,
		ResourceID:      req.ResourceID,
		LabelList:       labelList,
		Sort:            req.Sort,
		Asc:             req.Asc,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsListInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EcsListInstancesResultsResponse
	for _, result := range realResponse.Results {
		var sgs []EcsListInstancesSecGroupListResponse
		for _, s := range result.SecGroupList {
			sgs = append(sgs, EcsListInstancesSecGroupListResponse{
				SecurityGroupName: s.SecurityGroupName,
				SecurityGroupID:   s.SecurityGroupID,
			})
		}

		results = append(results, EcsListInstancesResultsResponse{
			AzName:         result.AzName,
			ExpiredTime:    result.ExpiredTime,
			CreatedTime:    result.CreatedTime,
			ProjectID:      result.ProjectID,
			InstanceID:     result.InstanceID,
			DisplayName:    result.DisplayName,
			InstanceName:   result.InstanceName,
			OsType:         result.OsType,
			InstanceStatus: result.InstanceStatus,
			OnDemand:       result.OnDemand,
			KeypairName:    result.KeypairName,
			SecGroupList:   sgs,
			Image: EcsListInstancesImageResponse{
				ImageID:   result.Image.ImageID,
				ImageName: result.Image.ImageName,
			},
			Flavor: EcsListInstancesFlavorResponse{
				FlavorID:     result.Flavor.FlavorID,
				FlavorName:   result.Flavor.FlavorName,
				FlavorCPU:    result.Flavor.FlavorCPU,
				FlavorRAM:    result.Flavor.FlavorRAM,
				GpuType:      result.Flavor.GpuType,
				GpuCount:     result.Flavor.GpuCount,
				GpuVendor:    result.Flavor.GpuVendor,
				VideoMemSize: result.Flavor.VideoMemSize,
			},
		})
	}
	return &EcsListInstancesResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EcsListInstancesLabelListRealRequest struct {
	LabelKey   string `json:"labelKey,omitempty"`
	LableValue string `json:"labelValue,omitempty"`
}

type EcsListInstancesRealRequest struct {
	RegionID        string                                 `json:"regionID"`
	AzName          string                                 `json:"azName,omitempty"`
	ProjectID       string                                 `json:"projectID,omitempty"`
	PageNo          *int                                   `json:"pageNo,omitempty"`
	PageSize        *int                                   `json:"pageSize,omitempty"`
	State           string                                 `json:"state,omitempty"`
	Keyword         string                                 `json:"keyword,omitempty"`
	InstanceName    string                                 `json:"instanceName,omitempty"`
	InstanceIDList  string                                 `json:"instanceIDList,omitempty"`
	SecurityGroupID string                                 `json:"securityGroupID,omitempty"`
	VpcID           string                                 `json:"vpcID,omitempty"`
	ResourceID      string                                 `json:"resourceID,omitempty"`
	LabelList       []EcsListInstancesLabelListRealRequest `json:"labelList,omitempty"`
	Sort            string                                 `json:"sort,omitempty"`
	Asc             *bool                                  `json:"asc,omitempty"`
}

type EcsListInstancesLabelListRequest struct {
	LabelKey   string
	LabelValue string
}

type EcsListInstancesRequest struct {
	RegionID        string
	AzName          string
	ProjectID       string
	PageNo          *int
	PageSize        *int
	State           string
	Keyword         string
	InstanceName    string
	InstanceIDList  string
	SecurityGroupID string
	VpcID           string
	ResourceID      string
	LabelList       []EcsListInstancesLabelListRequest
	Sort            string
	Asc             *bool
}

type EcsListInstancesAddressListRealResponse struct {
	Addr    string `json:"addr,omitempty"`
	Version int    `json:"version,omitempty"`
	Type    string `json:"type,omitempty"`
}

type EcsListInstancesAddressesRealResponse struct {
	VpcName     string                                    `json:"vpcName,omitempty"`
	AddressList []EcsListInstancesAddressListRealResponse `json:"addressList,omitempty"`
}

type EcsListInstancesSecGroupListRealResponse struct {
	SecurityGroupName string `json:"securityGroupName,omitempty"`
	SecurityGroupID   string `json:"securityGroupID,omitempty"`
}

type EcsListInstancesNetworkCardListRealResponse struct {
	IPv4Address   string   `json:"IPv4Address,omitempty"`
	IPv6Address   []string `json:"IPv6Address,omitempty"`
	SubnetID      string   `json:"subnetID,omitempty"`
	SubnetCidr    string   `json:"subnetCidr,omitempty"`
	IsMaster      bool     `json:"isMaster,omitempty"`
	Gateway       string   `json:"gateway,omitempty"`
	NetworkCardID string   `json:"networkCardID,omitempty"`
	SecurityGroup []string `json:"securityGroup,omitempty"`
}

type EcsListInstancesVipInfoListRealResponse struct {
	VipID          string `json:"vipID,omitempty"`
	VipAddress     string `json:"vipAddress,omitempty"`
	VipBindNicIP   string `json:"vipBindNicIP,omitempty"`
	VipBindNicIPv6 string `json:"vipBindNicIPv6,omitempty"`
	NicID          string `json:"nicID,omitempty"`
}

type EcsListInstancesAffinityGroupRealResponse struct {
	Policy            string `json:"policy,omitempty"`
	AffinityGroupName string `json:"affinityGroupName,omitempty"`
	AffinityGroupID   string `json:"affinityGroupID,omitempty"`
}

type EcsListInstancesImageRealResponse struct {
	ImageID   string `json:"imageID,omitempty"`
	ImageName string `json:"imageName,omitempty"`
}

type EcsListInstancesFlavorRealResponse struct {
	FlavorID     string `json:"flavorID,omitempty"`
	FlavorName   string `json:"flavorName,omitempty"`
	FlavorCPU    string `json:"flavorCPU,omitempty"`
	FlavorRAM    int    `json:"flavorRAM,omitempty"`
	GpuType      string `json:"gpuType,omitempty"`
	GpuCount     int    `json:"gpuCount,omitempty"`
	GpuVendor    string `json:"gpuVendor,omitempty"`
	VideoMemSize int    `json:"videoMemSize,omitempty"`
}

type EcsListInstancesResultsRealResponse struct {
	ProjectID       string                                        `json:"projectID,omitempty"`
	AzName          string                                        `json:"azName,omitempty"`
	Addresses       []EcsListInstancesAddressesRealResponse       `json:"addresses,omitempty"`
	ResourceID      string                                        `json:"resourceID,omitempty"`
	InstanceID      string                                        `json:"instanceID,omitempty"`
	DisplayName     string                                        `json:"displayName,omitempty"`
	InstanceName    string                                        `json:"instanceName,omitempty"`
	OsType          int                                           `json:"osType,omitempty"`
	InstanceStatus  string                                        `json:"instanceStatus,omitempty"`
	ExpiredTime     string                                        `json:"expiredTime,omitempty"`
	AvailableDay    string                                        `json:"availableDay,omitempty"`
	CreatedTime     string                                        `json:"createdTime,omitempty"`
	UpdatedTime     string                                        `json:"updatedTime,omitempty"`
	ZabbixName      string                                        `json:"zabbixName,omitempty"`
	SecGroupList    []EcsListInstancesSecGroupListRealResponse    `json:"secGroupList,omitempty"`
	PrivateIP       string                                        `json:"privateIP,omitempty"`
	PrivateIPv6     string                                        `json:"privateIPv6,omitempty"`
	NetworkCardList []EcsListInstancesNetworkCardListRealResponse `json:"networkCardList,omitempty"`
	VipInfoList     []EcsListInstancesVipInfoListRealResponse     `json:"vipInfoList,omitempty"`
	VipCount        int                                           `json:"vipCount,omitempty"`
	AffinityGroup   []EcsListInstancesAffinityGroupRealResponse   `json:"affinityGroup,omitempty"`
	Image           EcsListInstancesImageRealResponse             `json:"image,omitempty"`
	Flavor          EcsListInstancesFlavorRealResponse            `json:"flavor,omitempty"`
	OnDemand        bool                                          `json:"onDemand,omitempty"`
	VpcID           string                                        `json:"vpcID,omitempty"`
	FixedIPList     []string                                      `json:"fixedIpList,omitempty"`
	FloatingIP      string                                        `json:"floatingIp,omitempty"`
	SubnetIDList    []string                                      `json:"subnetIDList,omitempty"`
	KeypairName     string                                        `json:"keypairName,omitempty"`
	DelegateName    string                                        `json:"delegateName,omitempty"`
}

type EcsListInstancesRealResponse struct {
	CurrentCount int                                   `json:"currentCount"`
	TotalCount   int                                   `json:"totalCount"`
	TotalPage    int                                   `json:"totalPage"`
	Results      []EcsListInstancesResultsRealResponse `json:"results"`
}

type EcsListInstancesSecGroupListResponse struct {
	SecurityGroupName string
	SecurityGroupID   string
}

type EcsListInstancesNetworkCardListResponse struct {
	IPv4Address   string
	IPv6Address   []string
	SubnetID      string
	SubnetCidr    string
	IsMaster      bool
	Gateway       string
	NetworkCardID string
	SecurityGroup []string
}

type EcsListInstancesVipInfoListResponse struct {
	VipID          string
	VipAddress     string
	VipBindNicIP   string
	VipBindNicIPv6 string
	NicID          string
}

type EcsListInstancesAffinityGroupResponse struct {
	Policy            string
	AffinityGroupName string
	AffinityGroupID   string
}

type EcsListInstancesImageResponse struct {
	ImageID   string
	ImageName string
}

type EcsListInstancesFlavorResponse struct {
	FlavorID     string
	FlavorName   string
	FlavorCPU    string
	FlavorRAM    int
	GpuType      string
	GpuCount     int
	GpuVendor    string
	VideoMemSize int
}

type EcsListInstancesResultsResponse struct {
	ProjectID       string
	AzName          string
	Addresses       []EcsListInstancesAddressesRealResponse
	ResourceID      string
	InstanceID      string
	DisplayName     string
	InstanceName    string
	OsType          int
	InstanceStatus  string
	ExpiredTime     string
	AvailableDay    string
	CreatedTime     string
	UpdatedTime     string
	ZabbixName      string
	SecGroupList    []EcsListInstancesSecGroupListResponse
	PrivateIP       string
	PrivateIPv6     string
	NetworkCardList []EcsListInstancesNetworkCardListResponse
	VipInfoList     []EcsListInstancesVipInfoListResponse
	VipCount        int
	AffinityGroup   []EcsListInstancesAffinityGroupResponse
	Image           EcsListInstancesImageResponse
	Flavor          EcsListInstancesFlavorResponse
	OnDemand        bool
	VpcID           string
	FixedIPList     []string
	FloatingIP      string
	SubnetIDList    []string
	KeypairName     string
	DelegateName    string
}

type EcsListInstancesResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EcsListInstancesResultsResponse
}
