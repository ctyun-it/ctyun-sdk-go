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

		var vip_info []EcsListInstancesVipInfoListResponse
		for _, vip_info_list := range result.VipInfoList {
			vip_info = append(vip_info, EcsListInstancesVipInfoListResponse{
				VipID:          vip_info_list.VipID,
				VipAddress:     vip_info_list.VipAddress,
				VipBindNicIP:   vip_info_list.VipBindNicIP,
				VipBindNicIPv6: vip_info_list.VipBindNicIPv6,
				NicID:          vip_info_list.NicID,
			})
		}

		var addresses []EcsListInstancesAddressesResponse
		for _, addr := range result.Addresses {
			var add_list []EcsListInstancesAddressListResponse
			for _, add := range addr.AddressList {
				add_list = append(add_list, EcsListInstancesAddressListResponse{
					Addr:    add.Addr,
					Version: add.Version,
					Type:    add.Type,
				})
			}

			addresses = append(addresses, EcsListInstancesAddressesResponse{
				VpcName:     addr.VpcName,
				AddressList: add_list,
			})
		}

		var nic []EcsListInstancesNetworkCardListResponse
		for _, nic_info := range result.NetworkCardList {
			nic = append(nic, EcsListInstancesNetworkCardListResponse{
				IPv4Address:   nic_info.IPv4Address,
				IPv6Address:   nic_info.IPv6Address,
				SubnetID:      nic_info.SubnetID,
				SubnetCidr:    nic_info.SubnetCidr,
				IsMaster:      nic_info.IsMaster,
				Gateway:       nic_info.Gateway,
				NetworkCardID: nic_info.NetworkCardID,
				SecurityGroup: nic_info.SecurityGroup,
			})
		}

		results = append(results, EcsListInstancesResultsResponse{
			Addresses: addresses,
			AffinityGroup: EcsListInstancesAffinityGroupResponse{
				AffinityGroupID:   result.AffinityGroup.AffinityGroupID,
				AffinityGroupName: result.AffinityGroup.AffinityGroupName,
				Policy:            result.AffinityGroup.Policy,
			},
			AvailableDay: result.AvailableDay,
			AzName:       result.AzName,
			CreatedTime:  result.CreatedTime,
			DelegateName: result.DelegateName,
			DisplayName:  result.DisplayName,
			ExpiredTime:  result.ExpiredTime,
			Flavor: EcsListInstancesFlavorResponse{
				FlavorCPU:    result.Flavor.FlavorCPU,
				FlavorID:     result.Flavor.FlavorID,
				FlavorName:   result.Flavor.FlavorName,
				FlavorRAM:    result.Flavor.FlavorRAM,
				GpuCount:     result.Flavor.GpuCount,
				GpuType:      result.Flavor.GpuType,
				GpuVendor:    result.Flavor.GpuVendor,
				VideoMemSize: result.Flavor.VideoMemSize,
			},
			FixedIPList: result.FixedIPList,
			FloatingIP:  result.FloatingIP,
			Image: EcsListInstancesImageResponse{
				ImageID:   result.Image.ImageID,
				ImageName: result.Image.ImageName,
			},
			InstanceID:      result.InstanceID,
			InstanceName:    result.InstanceName,
			InstanceStatus:  result.InstanceStatus,
			KeypairName:     result.KeypairName,
			NetworkCardList: nic,
			OnDemand:        result.OnDemand,
			OsType:          result.OsType,
			PrivateIP:       result.PrivateIP,
			PrivateIPv6:     result.PrivateIPv6,
			ProjectID:       result.ProjectID,
			ResourceID:      result.ResourceID,
			SecGroupList:    sgs,
			SubnetIDList:    result.SubnetIDList,
			UpdatedTime:     result.UpdatedTime,
			VipCount:        result.VipCount,
			VipInfoList:     vip_info,
			VpcID:           result.VpcID,
			ZabbixName:      result.ZabbixName,
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
	LabelKey   *string `json:"labelKey,omitempty"`
	LableValue *string `json:"labelValue,omitempty"`
}

type EcsListInstancesRealRequest struct {
	RegionID        *string                                `json:"regionID"`
	AzName          *string                                `json:"azName,omitempty"`
	ProjectID       *string                                `json:"projectID,omitempty"`
	PageNo          *int                                   `json:"pageNo,omitempty"`
	PageSize        *int                                   `json:"pageSize,omitempty"`
	State           *string                                `json:"state,omitempty"`
	Keyword         *string                                `json:"keyword,omitempty"`
	InstanceName    *string                                `json:"instanceName,omitempty"`
	InstanceIDList  *string                                `json:"instanceIDList,omitempty"`
	SecurityGroupID *string                                `json:"securityGroupID,omitempty"`
	VpcID           *string                                `json:"vpcID,omitempty"`
	ResourceID      *string                                `json:"resourceID,omitempty"`
	LabelList       []EcsListInstancesLabelListRealRequest `json:"labelList,omitempty"`
	Sort            *string                                `json:"sort,omitempty"`
	Asc             *bool                                  `json:"asc,omitempty"`
}

type EcsListInstancesLabelListRequest struct {
	LabelKey   *string
	LabelValue *string
}

type EcsListInstancesRequest struct {
	RegionID        *string
	AzName          *string
	ProjectID       *string
	PageNo          *int
	PageSize        *int
	State           *string
	Keyword         *string
	InstanceName    *string
	InstanceIDList  *string
	SecurityGroupID *string
	VpcID           *string
	ResourceID      *string
	LabelList       []EcsListInstancesLabelListRequest
	Sort            *string
	Asc             *bool
}

type EcsListInstancesAddressListRealResponse struct {
	Addr    string `json:"addr,omitempty"`
	Type    string `json:"type,omitempty"`
	Version int    `json:"version,omitempty"`
}

type EcsListInstancesAddressesRealResponse struct {
	AddressList []EcsListInstancesAddressListRealResponse `json:"addressList,omitempty"`
	VpcName     string                                    `json:"vpcName,omitempty"`
}

type EcsListInstancesSecGroupListRealResponse struct {
	SecurityGroupID   string `json:"securityGroupID,omitempty"`
	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

type EcsListInstancesNetworkCardListRealResponse struct {
	Gateway       string   `json:"gateway,omitempty"`
	IPv4Address   string   `json:"IPv4Address,omitempty"`
	IPv6Address   []string `json:"IPv6Address,omitempty"`
	IsMaster      bool     `json:"isMaster,omitempty"`
	NetworkCardID string   `json:"networkCardID,omitempty"`
	SecurityGroup []string `json:"securityGroup,omitempty"`
	SubnetID      string   `json:"subnetID,omitempty"`
	SubnetCidr    string   `json:"subnetCidr,omitempty"`
}

type EcsListInstancesVipInfoListRealResponse struct {
	NicID          string `json:"nicID,omitempty"`
	VipAddress     string `json:"vipAddress,omitempty"`
	VipBindNicIP   string `json:"vipBindNicIP,omitempty"`
	VipBindNicIPv6 string `json:"vipBindNicIPv6,omitempty"`
	VipID          string `json:"vipID,omitempty"`
}

type EcsListInstancesAffinityGroupRealResponse struct {
	AffinityGroupID   string `json:"affinityGroupID,omitempty"`
	AffinityGroupName string `json:"affinityGroupName,omitempty"`
	Policy            string `json:"policy,omitempty"`
}

type EcsListInstancesImageRealResponse struct {
	ImageID   string `json:"imageID,omitempty"`
	ImageName string `json:"imageName,omitempty"`
}

type EcsListInstancesFlavorRealResponse struct {
	FlavorCPU    int    `json:"flavorCPU,omitempty"`
	FlavorID     string `json:"flavorID,omitempty"`
	FlavorName   string `json:"flavorName,omitempty"`
	FlavorRAM    int    `json:"flavorRAM,omitempty"`
	GpuCount     int    `json:"gpuCount,omitempty"`
	GpuType      string `json:"gpuType,omitempty"`
	GpuVendor    string `json:"gpuVendor,omitempty"`
	VideoMemSize int    `json:"videoMemSize,omitempty"`
}

type EcsListInstancesResultsRealResponse struct {
	Addresses          []EcsListInstancesAddressesRealResponse       `json:"addresses,omitempty"`
	AffinityGroup      EcsListInstancesAffinityGroupRealResponse     `json:"affinityGroup,omitempty"`
	AvailableDay       int                                           `json:"availableDay,omitempty"`
	AzName             string                                        `json:"azName,omitempty"`
	CreatedTime        string                                        `json:"createdTime,omitempty"`
	DelegateName       string                                        `json:"delegateName,omitempty"`
	DeletionProtection bool                                          `json:"deletionProtection,omitempty"`
	DisplayName        string                                        `json:"displayName,omitempty"`
	ExpiredTime        string                                        `json:"expiredTime,omitempty"`
	Flavor             EcsListInstancesFlavorRealResponse            `json:"flavor,omitempty"`
	FixedIPList        []string                                      `json:"fixedIpList,omitempty"`
	FloatingIP         string                                        `json:"floatingIp,omitempty"`
	Image              EcsListInstancesImageRealResponse             `json:"image,omitempty"`
	InstanceID         string                                        `json:"instanceID,omitempty"`
	InstanceName       string                                        `json:"instanceName,omitempty"`
	InstanceStatus     string                                        `json:"instanceStatus,omitempty"`
	KeypairName        string                                        `json:"keypairName,omitempty"`
	NetworkCardList    []EcsListInstancesNetworkCardListRealResponse `json:"networkCardList,omitempty"`
	OnDemand           bool                                          `json:"onDemand,omitempty"`
	OsType             int                                           `json:"osType,omitempty"`
	PrivateIP          string                                        `json:"privateIP,omitempty"`
	PrivateIPv6        string                                        `json:"privateIPv6,omitempty"`
	ProjectID          string                                        `json:"projectID,omitempty"`
	ReleaseTime        string                                        `json:"releaseTime,omitempty"`
	RemainingDay       int                                           `json:"remainingDay,omitempty"`
	ResourceID         string                                        `json:"resourceID,omitempty"`
	SecGroupList       []EcsListInstancesSecGroupListRealResponse    `json:"secGroupList,omitempty"`
	SubnetIDList       []string                                      `json:"subnetIDList,omitempty"`
	UpdatedTime        string                                        `json:"updatedTime,omitempty"`
	VipCount           int                                           `json:"vipCount,omitempty"`
	VipInfoList        []EcsListInstancesVipInfoListRealResponse     `json:"vipInfoList,omitempty"`
	VpcID              string                                        `json:"vpcID,omitempty"`
	ZabbixName         string                                        `json:"zabbixName,omitempty"`
}

type EcsListInstancesRealResponse struct {
	CurrentCount int                                   `json:"currentCount,omitempty"`
	Results      []EcsListInstancesResultsRealResponse `json:"results,omitempty"`
	TotalCount   int                                   `json:"totalCount,omitempty"`
	TotalPage    int                                   `json:"totalPage,omitempty"`
}

type EcsListInstancesAddressListResponse struct {
	Addr    string
	Type    string
	Version int
}

type EcsListInstancesAddressesResponse struct {
	AddressList []EcsListInstancesAddressListResponse
	VpcName     string
}

type EcsListInstancesSecGroupListResponse struct {
	SecurityGroupID   string
	SecurityGroupName string
}

type EcsListInstancesNetworkCardListResponse struct {
	Gateway       string
	IPv4Address   string
	IPv6Address   []string
	IsMaster      bool
	NetworkCardID string
	SecurityGroup []string
	SubnetCidr    string
	SubnetID      string
}

type EcsListInstancesVipInfoListResponse struct {
	NicID          string
	VipAddress     string
	VipBindNicIP   string
	VipBindNicIPv6 string
	VipID          string
}

type EcsListInstancesAffinityGroupResponse struct {
	AffinityGroupID   string
	AffinityGroupName string
	Policy            string
}

type EcsListInstancesImageResponse struct {
	ImageID   string
	ImageName string
}

type EcsListInstancesFlavorResponse struct {
	FlavorCPU    int
	FlavorID     string
	FlavorName   string
	FlavorRAM    int
	GpuCount     int
	GpuType      string
	GpuVendor    string
	VideoMemSize int
}

type EcsListInstancesResultsResponse struct {
	Addresses          []EcsListInstancesAddressesResponse
	AffinityGroup      EcsListInstancesAffinityGroupResponse
	AvailableDay       int
	AzName             string
	CreatedTime        string
	DelegateName       string
	DeletionProtection bool
	DisplayName        string
	ExpiredTime        string
	Flavor             EcsListInstancesFlavorResponse
	FixedIPList        []string
	FloatingIP         string
	Image              EcsListInstancesImageResponse
	InstanceID         string
	InstanceName       string
	InstanceStatus     string
	KeypairName        string
	NetworkCardList    []EcsListInstancesNetworkCardListResponse
	OnDemand           bool
	OsType             int
	PrivateIP          string
	PrivateIPv6        string
	ProjectID          string
	ReleaseTime        string
	RemainingDay       int
	ResourceID         string
	SecGroupList       []EcsListInstancesSecGroupListResponse
	SubnetIDList       []string
	UpdatedTime        string
	VipCount           int
	VipInfoList        []EcsListInstancesVipInfoListResponse
	VpcID              string
	ZabbixName         string
}

type EcsListInstancesResponse struct {
	CurrentCount int
	Results      []EcsListInstancesResultsResponse
	TotalCount   int
	TotalPage    int
}
