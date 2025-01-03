package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupInstanceDetailsApi
type EcsBackupInstanceDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupInstanceDetailsApi(client *ctyunsdk.CtyunClient) *EcsBackupInstanceDetailsApi {
	return &EcsBackupInstanceDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup/instance-details",
		},
	}
}

func (this *EcsBackupInstanceDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupInstanceDetailsRequest) (*EcsBackupInstanceDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupInstanceDetailsRealRequest{
		RegionID:   req.RegionID,
		InstanceID: req.InstanceID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupInstanceDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var volumes []EcsBackupInstanceDetailsVolumesResponse
	for _, res := range realResponse.Volumes {
		volumes = append(volumes, EcsBackupInstanceDetailsVolumesResponse{
			IsBootable: res.IsBootable,
			DiskSize:   res.DiskSize,
			DiskType:   res.DiskType,
			DiskID:     res.DiskID,
			DiskName:   res.DiskName,
		})
	}

	var vipInfoList []EcsBackupInstanceDetailsVipInfoListResponse
	for _, res := range realResponse.VipInfoList {
		vipInfoList = append(vipInfoList, EcsBackupInstanceDetailsVipInfoListResponse{
			VipID:          res.VipID,
			VipAddress:     res.VipAddress,
			VipBindNicIP:   res.VipBindNicIP,
			VipBindNicIPv6: res.VipBindNicIPv6,
			NicID:          res.NicID,
		})
	}

	var networkCardList []EcsBackupInstanceDetailsNetworkCardListResponse
	for _, res := range realResponse.NetworkCardList {
		networkCardList = append(networkCardList, EcsBackupInstanceDetailsNetworkCardListResponse{
			IPv4Address:   res.IPv4Address,
			IPv6Address:   res.IPv6Address,
			SubnetID:      res.SubnetID,
			SubnetCidr:    res.SubnetCidr,
			IsMaster:      res.IsMaster,
			Gateway:       res.Gateway,
			NetworkCardID: res.NetworkCardID,
			SecurityGroup: res.SecurityGroup,
		})
	}

	var secGroupList []EcsBackupInstanceDetailsSecGroupListResponse
	for _, res := range realResponse.SecGroupList {
		secGroupList = append(secGroupList, EcsBackupInstanceDetailsSecGroupListResponse{
			SecurityGroupID:   res.SecurityGroupID,
			SecurityGroupName: res.SecurityGroupName,
		})
	}

	var addressList []EcsBackupInstanceDetailsAddressListResponse
	for _, res := range addressList {
		addressList = append(addressList, EcsBackupInstanceDetailsAddressListResponse{
			Addr:    res.Addr,
			Version: res.Version,
			Type:    res.Type,
		})
	}

	var addresses []EcsBackupInstanceDetailsAddressesResponse
	for _, res := range realResponse.Addresses {
		addresses = append(addresses, EcsBackupInstanceDetailsAddressesResponse{
			VpcName:     res.VpcName,
			AddressList: addressList,
		})
	}

	return &EcsBackupInstanceDetailsResponse{
		ProjectID:       realResponse.ProjectID,
		AzName:          realResponse.AzName,
		AttachedVolume:  realResponse.AttachedVolume,
		Addresses:       addresses,
		ResourceID:      realResponse.ResourceID,
		InstanceID:      realResponse.InstanceID,
		DisplayName:     realResponse.DisplayName,
		InstanceName:    realResponse.InstanceName,
		OsType:          realResponse.OsType,
		InstanceStatus:  realResponse.InstanceStatus,
		ExpiredTime:     realResponse.ExpiredTime,
		AvailableDay:    realResponse.AvailableDay,
		UpdatedTime:     realResponse.UpdatedTime,
		CreatedTime:     realResponse.CreatedTime,
		ZabbixName:      realResponse.ZabbixName,
		SecGroupList:    secGroupList,
		PrivateIP:       realResponse.PrivateIP,
		PrivateIPv6:     realResponse.PrivateIPv6,
		NetworkCardList: networkCardList,
		VipInfoList:     vipInfoList,
		VipCount:        realResponse.VipCount,
		AffinityGroup: EcsBackupInstanceDetailsAffinityGroupResponse{
			Policy:            realResponse.AffinityGroup.Policy,
			AffinityGroupName: realResponse.AffinityGroup.AffinityGroupName,
			AffinityGroupID:   realResponse.AffinityGroup.AffinityGroupID,
		},
		Image: EcsBackupInstanceDetailsImageResponse{
			ImageID:   realResponse.Image.ImageID,
			ImageName: realResponse.Image.ImageName,
		},
		Flavor: EcsBackupInstanceDetailsFlavorResponse{
			FlavorID:     realResponse.Flavor.FlavorID,
			FlavorName:   realResponse.Flavor.FlavorName,
			FlavorCPU:    realResponse.Flavor.FlavorCPU,
			FlavorRAM:    realResponse.Flavor.FlavorRAM,
			GpuType:      realResponse.Flavor.GpuType,
			GpuCount:     realResponse.Flavor.GpuCount,
			GpuVendor:    realResponse.Flavor.GpuVendor,
			VideoMemSize: realResponse.Flavor.VideoMemSize,
		},
		OnDemand:     realResponse.OnDemand,
		VpcName:      realResponse.VpcName,
		VpcID:        realResponse.VpcID,
		FixedIP:      realResponse.FixedIP,
		FloatingIP:   realResponse.FloatingIP,
		SubnetIDList: realResponse.SubnetIDList,
		KeypairName:  realResponse.KeypairName,
		Volumes:      volumes,
	}, nil
}

type EcsBackupInstanceDetailsRealRequest struct {
	RegionID   *string `json:"regionID,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
}

type EcsBackupInstanceDetailsRequest struct {
	RegionID   *string
	InstanceID *string
}

type EcsBackupInstanceDetailsVolumesRealResponse struct {
	IsBootable bool   `json:"isBootable,omitempty"`
	DiskSize   int    `json:"diskSize,omitempty"`
	DiskType   string `json:"diskType,omitempty"`
	DiskID     string `json:"diskID,omitempty"`
	DiskName   string `json:"diskName,omitempty"`
}

type EcsBackupInstanceDetailsFlavorRealResponse struct {
	FlavorID     string `json:"flavorID,omitempty"`
	FlavorName   string `json:"flavorName,omitempty"`
	FlavorCPU    int    `json:"flavorCPU,omitempty"`
	FlavorRAM    int    `json:"flavorRAM,omitempty"`
	GpuType      string `json:"gpuType,omitempty"`
	GpuCount     int    `json:"gpuCount,omitempty"`
	GpuVendor    string `json:"gpuVendor,omitempty"`
	VideoMemSize int    `json:"videoMemSize,omitempty"`
}

type EcsBackupInstanceDetailsImageRealResponse struct {
	ImageID   string `json:"imageID,omitempty"`
	ImageName string `json:"imageName,omitempty"`
}

type EcsBackupInstanceDetailsAffinityGroupRealResponse struct {
	Policy            string `json:"policy,omitempty"`
	AffinityGroupName string `json:"affinityGroupName,omitempty"`
	AffinityGroupID   string `json:"affinityGroupID,omitempty"`
}

type EcsBackupInstanceDetailsVipInfoListRealResponse struct {
	VipID          string `json:"vipID,omitempty"`
	VipAddress     string `json:"vipAddress,omitempty"`
	VipBindNicIP   string `json:"vipBindNicIP,omitempty"`
	VipBindNicIPv6 string `json:"vipBindNicIPv6,omitempty"`
	NicID          string `json:"nicID,omitempty"`
}

type EcsBackupInstanceDetailsNetworkCardListRealResponse struct {
	IPv4Address   string   `json:"IPv4Address,omitempty"`
	IPv6Address   []string `json:"IPv6Address,omitempty"`
	SubnetID      string   `json:"subnetID,omitempty"`
	SubnetCidr    string   `json:"subnetCidr,omitempty"`
	IsMaster      bool     `json:"isMaster,omitempty"`
	Gateway       string   `json:"gateway,omitempty"`
	NetworkCardID string   `json:"networkCardID,omitempty"`
	SecurityGroup []string `json:"securityGroup,omitempty"`
}

type EcsBackupInstanceDetailsSecGroupListRealResponse struct {
	SecurityGroupID   string `json:"securityGroupID,omitempty"`
	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

type EcsBackupInstanceDetailsAddressListRealResponse struct {
	Addr    string `json:"addr,omitempty"`
	Version int    `json:"version,omitempty"`
	Type    string `json:"type,omitempty"`
}

type EcsBackupInstanceDetailsAddressesRealResponse struct {
	VpcName     string                                            `json:"vpcName,omitempty"`
	AddressList []EcsBackupInstanceDetailsAddressListRealResponse `json:"addressList,omitempty"`
}

type EcsBackupInstanceDetailsRealResponse struct {
	ProjectID       string                                                `json:"projectID,omitempty"`
	AzName          string                                                `json:"azName,omitempty"`
	AttachedVolume  []string                                              `json:"attachedVolume,omitempty"`
	Addresses       []EcsBackupInstanceDetailsAddressesRealResponse       `json:"addresses,omitempty"`
	ResourceID      string                                                `json:"resourceID,omitempty"`
	InstanceID      string                                                `json:"instanceID,omitempty"`
	DisplayName     string                                                `json:"displayName,omitempty"`
	InstanceName    string                                                `json:"instanceName,omitempty"`
	OsType          int                                                   `json:"osType,omitempty"`
	InstanceStatus  string                                                `json:"instanceStatus,omitempty"`
	ExpiredTime     string                                                `json:"expiredTime,omitempty"`
	AvailableDay    int                                                   `json:"availableDay,omitempty"`
	UpdatedTime     string                                                `json:"updatedTime,omitempty"`
	CreatedTime     string                                                `json:"createdTime,omitempty"`
	ZabbixName      string                                                `json:"zabbixName,omitempty"`
	SecGroupList    []EcsBackupInstanceDetailsSecGroupListRealResponse    `json:"secGroupList,omitempty"`
	PrivateIP       string                                                `json:"privateIP,omitempty"`
	PrivateIPv6     string                                                `json:"privateIPv6,omitempty"`
	NetworkCardList []EcsBackupInstanceDetailsNetworkCardListRealResponse `json:"networkCardList,omitempty"`
	VipInfoList     []EcsBackupInstanceDetailsVipInfoListRealResponse     `json:"vipInfoList,omitempty"`
	VipCount        int                                                   `json:"vipCount,omitempty"`
	AffinityGroup   EcsBackupInstanceDetailsAffinityGroupRealResponse     `json:"affinityGroup,omitempty"`
	Image           EcsBackupInstanceDetailsImageRealResponse             `json:"image,omitempty"`
	Flavor          EcsBackupInstanceDetailsFlavorRealResponse            `json:"flavor,omitempty"`
	OnDemand        bool                                                  `json:"onDemand,omitempty"`
	VpcName         string                                                `json:"vpcName,omitempty"`
	VpcID           string                                                `json:"vpcID,omitempty"`
	FixedIP         []string                                              `json:"fixedIP,omitempty"`
	FloatingIP      string                                                `json:"floatingIP,omitempty"`
	SubnetIDList    []string                                              `json:"subnetIDList,omitempty"`
	KeypairName     string                                                `json:"keypairName,omitempty"`
	Volumes         []EcsBackupInstanceDetailsVolumesRealResponse         `json:"volumes,omitempty"`
}

type EcsBackupInstanceDetailsVolumesResponse struct {
	IsBootable bool
	DiskSize   int
	DiskType   string
	DiskID     string
	DiskName   string
}

type EcsBackupInstanceDetailsFlavorResponse struct {
	FlavorID     string
	FlavorName   string
	FlavorCPU    int
	FlavorRAM    int
	GpuType      string
	GpuCount     int
	GpuVendor    string
	VideoMemSize int
}

type EcsBackupInstanceDetailsImageResponse struct {
	ImageID   string
	ImageName string
}

type EcsBackupInstanceDetailsAffinityGroupResponse struct {
	Policy            string
	AffinityGroupName string
	AffinityGroupID   string
}

type EcsBackupInstanceDetailsVipInfoListResponse struct {
	VipID          string
	VipAddress     string
	VipBindNicIP   string
	VipBindNicIPv6 string
	NicID          string
}

type EcsBackupInstanceDetailsNetworkCardListResponse struct {
	IPv4Address   string
	IPv6Address   []string
	SubnetID      string
	SubnetCidr    string
	IsMaster      bool
	Gateway       string
	NetworkCardID string
	SecurityGroup []string
}

type EcsBackupInstanceDetailsSecGroupListResponse struct {
	SecurityGroupID   string
	SecurityGroupName string
}

type EcsBackupInstanceDetailsAddressListResponse struct {
	Addr    string
	Version int
	Type    string
}

type EcsBackupInstanceDetailsAddressesResponse struct {
	VpcName     string
	AddressList []EcsBackupInstanceDetailsAddressListResponse
}

type EcsBackupInstanceDetailsResponse struct {
	ProjectID       string
	AzName          string
	AttachedVolume  []string
	Addresses       []EcsBackupInstanceDetailsAddressesResponse
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
	SecGroupList    []EcsBackupInstanceDetailsSecGroupListResponse
	PrivateIP       string
	PrivateIPv6     string
	NetworkCardList []EcsBackupInstanceDetailsNetworkCardListResponse
	VipInfoList     []EcsBackupInstanceDetailsVipInfoListResponse
	VipCount        int
	AffinityGroup   EcsBackupInstanceDetailsAffinityGroupResponse
	Image           EcsBackupInstanceDetailsImageResponse
	Flavor          EcsBackupInstanceDetailsFlavorResponse
	OnDemand        bool
	VpcName         string
	VpcID           string
	FixedIP         []string
	FloatingIP      string
	SubnetIDList    []string
	KeypairName     string
	Volumes         []EcsBackupInstanceDetailsVolumesResponse
}
