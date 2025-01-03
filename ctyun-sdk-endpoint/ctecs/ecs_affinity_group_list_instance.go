package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupListInstanceApi 查询主机组内的云主机
// https://www.ctyun.cn/document/10026730/10106071
type EcsAffinityGroupListInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupListInstanceApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupListInstanceApi {
	return &EcsAffinityGroupListInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group/list-instance",
		},
	}
}

func (this *EcsAffinityGroupListInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupListInstanceRequest) (*EcsAffinityGroupListInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupListInstanceRealRequest{
		RegionID:        req.RegionID,
		AffinityGroupID: req.AffinityGroupID,
		PageNo:          req.PageNo,
		PageSize:        req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupListInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EcsAffinityGroupListInstanceResultsResponse
	for _, result := range realResponse.Results {
		var vip_info []EcsAffinityGroupListInstanceVipInfoListResponse
		for _, vip_info_list := range result.VipInfoList {
			vip_info = append(vip_info, EcsAffinityGroupListInstanceVipInfoListResponse{
				VipID:          vip_info_list.VipID,
				VipAddress:     vip_info_list.VipAddress,
				VipBindNicIP:   vip_info_list.VipBindNicIP,
				VipBindNicIPv6: vip_info_list.VipBindNicIPv6,
				NicID:          vip_info_list.NicID,
			})
		}

		var nic []EcsAffinityGroupListInstanceNetworkCardListResponse
		for _, nic_info := range result.NetworkCardList {
			nic = append(nic, EcsAffinityGroupListInstanceNetworkCardListResponse{
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

		var sgs []EcsAffinityGroupListInstanceSecGroupListResponse
		for _, s := range result.SecGroupList {
			sgs = append(sgs, EcsAffinityGroupListInstanceSecGroupListResponse{
				SecurityGroupName: s.SecurityGroupName,
				SecurityGroupID:   s.SecurityGroupID,
			})
		}

		var addresses []EcsAffinityGroupListInstanceAddressesResponse
		for _, addr := range result.Addresses {
			var add_list []EcsAffinityGroupListInstanceAddressListResponse
			for _, add := range addr.AddressList {
				add_list = append(add_list, EcsAffinityGroupListInstanceAddressListResponse{
					Addr:    add.Addr,
					Version: add.Version,
					Type:    add.Type,
				})
			}

			addresses = append(addresses, EcsAffinityGroupListInstanceAddressesResponse{
				VpcName:     addr.VpcName,
				AddressList: add_list,
			})
		}

		results = append(results, EcsAffinityGroupListInstanceResultsResponse{
			Addresses: addresses,
			AffinityGroup: EcsAffinityGroupListInstanceAffinityGroupResponse{
				AffinityGroupID:     result.AffinityGroup.AffinityGroupID,
				AffinityGroupName:   result.AffinityGroup.AffinityGroupName,
				AffinityGroupPolicy: result.AffinityGroup.AffinityGroupPolicy,
			},
			AvailableDay: result.AvailableDay,
			AzName:       result.AzName,
			CreatedTime:  result.CreatedTime,
			DisplayName:  result.DisplayName,
			ExpiredTime:  result.ExpiredTime,
			Flavor: EcsAffinityGroupListInstanceFlavorResponse{
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
			Image: EcsAffinityGroupListInstanceImageResponse{
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

	return &EcsAffinityGroupListInstanceResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EcsAffinityGroupListInstanceRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	AffinityGroupID *string `json:"affinityGroupID,omitempty"`
	PageNo          *int    `json:"pageNo,omitempty"`
	PageSize        *int    `json:"pageSize,omitempty"`
}

type EcsAffinityGroupListInstanceRequest struct {
	RegionID        *string
	AffinityGroupID *string
	PageNo          *int
	PageSize        *int
}

type EcsAffinityGroupListInstanceFlavorRealResponse struct {
	FlavorID     string `json:"flavorID,omitempty"`
	FlavorName   string `json:"flavorName,omitempty"`
	FlavorCPU    int    `json:"flavorCPU,omitempty"`
	FlavorRAM    int    `json:"flavorRAM,omitempty"`
	GpuType      string `json:"gpuType,omitempty"`
	GpuCount     int    `json:"gpuCount,omitempty"`
	GpuVendor    string `json:"gpuVendor,omitempty"`
	VideoMemSize int    `json:"videoMemSize,omitempty"`
}

type EcsAffinityGroupListInstanceImageRealResponse struct {
	ImageID   string `json:"imageID,omitempty"`
	ImageName string `json:"imageName,omitempty"`
}

type EcsAffinityGroupListInstanceAffinityGroupRealResponse struct {
	AffinityGroupPolicy string `json:"affinityGroupPolicy,omitempty"`
	AffinityGroupName   string `json:"affinityGroupName,omitempty"`
	AffinityGroupID     string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupListInstanceVipInfoListRealResponse struct {
	VipID          string `json:"vipID,omitempty"`
	VipAddress     string `json:"vipAddress,omitempty"`
	VipBindNicIP   string `json:"vipBindNicIP,omitempty"`
	VipBindNicIPv6 string `json:"vipBindNicIPv6,omitempty"`
	NicID          string `json:"nicID,omitempty"`
}

type EcsAffinityGroupListInstanceNetworkCardListRealResponse struct {
	IPv4Address   string   `json:"IPv4Address,omitempty"`
	IPv6Address   []string `json:"IPv6Address,omitempty"`
	SubnetID      string   `json:"subnetID,omitempty"`
	SubnetCidr    string   `json:"subnetCidr,omitempty"`
	IsMaster      bool     `json:"isMaster,omitempty"`
	Gateway       string   `json:"gateway,omitempty"`
	NetworkCardID string   `json:"networkCardID,omitempty"`
	SecurityGroup []string `json:"securityGroup,omitempty"`
}

type EcsAffinityGroupListInstanceSecGroupListRealResponse struct {
	SecurityGroupID   string `json:"securityGroupID,omitempty"`
	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

type EcsAffinityGroupListInstanceAddressListRealResponse struct {
	Addr    string `json:"addr,omitempty"`
	Version int    `json:"version,omitempty"`
	Type    string `json:"type,omitempty"`
}

type EcsAffinityGroupListInstanceAddressesRealResponse struct {
	VpcName     string                                                `json:"vpcName,omitempty"`
	AddressList []EcsAffinityGroupListInstanceAddressListRealResponse `json:"addressList,omitempty"`
}

type EcsAffinityGroupListInstanceResultsRealResponse struct {
	ProjectID       string                                                    `json:"projectID,omitempty"`
	AzName          string                                                    `json:"azName,omitempty"`
	AttachedVolume  []string                                                  `json:"attachedVolume,omitempty"`
	Addresses       []EcsAffinityGroupListInstanceAddressesRealResponse       `json:"addresses,omitempty"`
	ResourceID      string                                                    `json:"resourceID,omitempty"`
	InstanceID      string                                                    `json:"instanceID,omitempty"`
	DisplayName     string                                                    `json:"displayName,omitempty"`
	InstanceName    string                                                    `json:"instanceName,omitempty"`
	OsType          int                                                       `json:"osType,omitempty"`
	InstanceStatus  string                                                    `json:"instanceStatus,omitempty"`
	ExpiredTime     string                                                    `json:"expiredTime,omitempty"`
	AvailableDay    int                                                       `json:"availableDay,omitempty"`
	UpdatedTime     string                                                    `json:"updatedTime,omitempty"`
	CreatedTime     string                                                    `json:"createdTime,omitempty"`
	ZabbixName      string                                                    `json:"zabbixName,omitempty"`
	SecGroupList    []EcsAffinityGroupListInstanceSecGroupListRealResponse    `json:"secGroupList,omitempty"`
	PrivateIP       string                                                    `json:"privateIP,omitempty"`
	PrivateIPv6     string                                                    `json:"privateIPv6,omitempty"`
	NetworkCardList []EcsAffinityGroupListInstanceNetworkCardListRealResponse `json:"networkCardList,omitempty"`
	VipInfoList     []EcsAffinityGroupListInstanceVipInfoListRealResponse     `json:"vipInfoList,omitempty"`
	VipCount        int                                                       `json:"vipCount,omitempty"`
	AffinityGroup   EcsAffinityGroupListInstanceAffinityGroupRealResponse     `json:"affinityGroup,omitempty"`
	Image           EcsAffinityGroupListInstanceImageRealResponse             `json:"image,omitempty"`
	Flavor          EcsAffinityGroupListInstanceFlavorRealResponse            `json:"flavor,omitempty"`
	OnDemand        bool                                                      `json:"onDemand,omitempty"`
	VpcName         string                                                    `json:"vpcName,omitempty"`
	VpcID           string                                                    `json:"vpcID,omitempty"`
	FixedIPList     []string                                                  `json:"fixedIPList,omitempty"`
	FloatingIP      string                                                    `json:"floatingIP,omitempty"`
	SubnetIDList    []string                                                  `json:"subnetIDList,omitempty"`
	KeypairName     string                                                    `json:"keypairName,omitempty"`
}

type EcsAffinityGroupListInstanceRealResponse struct {
	CurrentCount int                                               `json:"currentCount,omitempty"`
	TotalCount   int                                               `json:"totalCount,omitempty"`
	TotalPage    int                                               `json:"totalPage,omitempty"`
	Results      []EcsAffinityGroupListInstanceResultsRealResponse `json:"results,omitempty"`
}

type EcsAffinityGroupListInstanceFlavorResponse struct {
	FlavorID     string
	FlavorName   string
	FlavorCPU    int
	FlavorRAM    int
	GpuType      string
	GpuCount     int
	GpuVendor    string
	VideoMemSize int
}

type EcsAffinityGroupListInstanceImageResponse struct {
	ImageID   string
	ImageName string
}

type EcsAffinityGroupListInstanceAffinityGroupResponse struct {
	AffinityGroupPolicy string
	AffinityGroupName   string
	AffinityGroupID     string
}

type EcsAffinityGroupListInstanceVipInfoListResponse struct {
	NicID          string
	VipAddress     string
	VipBindNicIP   string
	VipBindNicIPv6 string
	VipID          string
}

type EcsAffinityGroupListInstanceNetworkCardListResponse struct {
	IPv4Address   string
	IPv6Address   []string
	SubnetID      string
	SubnetCidr    string
	IsMaster      bool
	Gateway       string
	NetworkCardID string
	SecurityGroup []string
}

type EcsAffinityGroupListInstanceSecGroupListResponse struct {
	SecurityGroupID   string
	SecurityGroupName string
}

type EcsAffinityGroupListInstanceAddressListResponse struct {
	Addr    string
	Version int
	Type    string
}

type EcsAffinityGroupListInstanceAddressesResponse struct {
	VpcName     string
	AddressList []EcsAffinityGroupListInstanceAddressListResponse
}

type EcsAffinityGroupListInstanceResultsResponse struct {
	ProjectID       string
	AzName          string
	AttachedVolume  []string
	Addresses       []EcsAffinityGroupListInstanceAddressesResponse
	ResourceID      string
	InstanceID      string
	DisplayName     string
	InstanceName    string
	OsType          int
	InstanceStatus  string
	ExpiredTime     string
	AvailableDay    int
	UpdatedTime     string
	CreatedTime     string
	ZabbixName      string
	SecGroupList    []EcsAffinityGroupListInstanceSecGroupListResponse
	PrivateIP       string
	PrivateIPv6     string
	NetworkCardList []EcsAffinityGroupListInstanceNetworkCardListResponse
	VipInfoList     []EcsAffinityGroupListInstanceVipInfoListResponse
	VipCount        int
	AffinityGroup   EcsAffinityGroupListInstanceAffinityGroupResponse
	Image           EcsAffinityGroupListInstanceImageResponse
	Flavor          EcsAffinityGroupListInstanceFlavorResponse
	OnDemand        bool
	VpcName         string
	VpcID           string
	FixedIPList     []string
	FloatingIP      string
	SubnetIDList    []string
	KeypairName     string
}

type EcsAffinityGroupListInstanceResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EcsAffinityGroupListInstanceResultsResponse
}
