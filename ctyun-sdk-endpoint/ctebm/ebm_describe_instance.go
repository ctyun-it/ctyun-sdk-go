package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmDescribeInstanceApi 查询单台物理机_v4_plus
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=6940&data=97&isNormal=1
type EbmDescribeInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmDescribeInstanceApi(client *ctyunsdk.CtyunClient) *EbmDescribeInstanceApi {
	return &EbmDescribeInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ebm/describe-instance",
		},
	}
}

func (this *EbmDescribeInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmDescribeInstanceRequest) (*EbmDescribeInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", req.RegionID).
		AddParam("azName", req.AzName).
		AddParam("instanceUUID", req.InstanceUUID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmDescribeInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var interfaces []EbmDescribeInstanceInterfacesResponse
	// 构造 interfaces中的sgs
	for _, interface_info := range realResponse.Interfaces {
		var sgs []EbmDescribeInstanceSecurityGroupsResponse
		for _, sg := range interface_info.SecurityGroups {
			sgs = append(sgs, EbmDescribeInstanceSecurityGroupsResponse{
				SecurityGroupID:   sg.SecurityGroupID,
				SecurityGroupName: sg.SecurityGroupName,
			})
		}
		interfaces = append(interfaces, EbmDescribeInstanceInterfacesResponse{
			InterfaceUUID:  interface_info.InterfaceUUID,
			Master:         interface_info.Master,
			VpcUUID:        interface_info.VpcUUID,
			SubnetUUID:     interface_info.SubnetUUID,
			PortUUID:       interface_info.PortUUID,
			Ipv4:           interface_info.Ipv4,
			Ipv4Gateway:    interface_info.Ipv4Gateway,
			Ipv6:           interface_info.Ipv6,
			Ipv6Gateway:    interface_info.Ipv6Gateway,
			VipUUIDList:    interface_info.VipUUIDList,
			VipList:        interface_info.VipList,
			SecurityGroups: sgs,
		})
	}
	//	构造 deviceDetail
	var deviceDetail EbmDescribeInstanceDeviceDetailResponse
	var realDeviceDetailResponse = realResponse.DeviceDetail
	deviceDetail = EbmDescribeInstanceDeviceDetailResponse{
		CpuModel:                realDeviceDetailResponse.CpuModel,
		CpuFrequency:            realDeviceDetailResponse.CpuFrequency,
		MemAmount:               realDeviceDetailResponse.MemAmount,
		MemSize:                 realDeviceDetailResponse.MemSize,
		MemFrequency:            realDeviceDetailResponse.MemFrequency,
		NicAmount:               realDeviceDetailResponse.NicAmount,
		NicRate:                 realDeviceDetailResponse.NicRate,
		SystemVolumeAmount:      realDeviceDetailResponse.SystemVolumeAmount,
		SystemVolumeSize:        realDeviceDetailResponse.SystemVolumeSize,
		SystemVolumeType:        realDeviceDetailResponse.SystemVolumeType,
		SystemVolumeInterface:   realDeviceDetailResponse.SystemVolumeInterface,
		SystemVolumeDescription: realDeviceDetailResponse.SystemVolumeDescription,
		DataVolumeAmount:        realDeviceDetailResponse.DataVolumeAmount,
		DataVolumeSize:          realDeviceDetailResponse.DataVolumeSize,
		DataVolumeInterface:     realDeviceDetailResponse.DataVolumeInterface,
		DataVolumeType:          realDeviceDetailResponse.DataVolumeType,
		DataVolumeDescription:   realDeviceDetailResponse.DataVolumeDescription,
		SmartNicExist:           realDeviceDetailResponse.SmartNicExist,
		NvmeVolumeAmount:        realDeviceDetailResponse.NvmeVolumeAmount,
		NvmeVolumeSize:          realDeviceDetailResponse.NvmeVolumeSize,
		NvmeVolumeType:          realDeviceDetailResponse.NvmeVolumeType,
		NvmeVolumeInterface:     realDeviceDetailResponse.NvmeVolumeInterface,
		GpuAmount:               realDeviceDetailResponse.GpuAmount,
		GpuSize:                 realDeviceDetailResponse.GpuSize,
		GpuManufacturer:         realDeviceDetailResponse.GpuManufacturer,
		GpuModel:                realDeviceDetailResponse.GpuModel,
		ComputeIBAmount:         realDeviceDetailResponse.ComputeIBAmount,
		ComputeIBRate:           realDeviceDetailResponse.ComputeIBRate,
		StorageIBAmount:         realDeviceDetailResponse.StorageIBAmount,
		StorageIBRate:           realDeviceDetailResponse.StorageIBRate,
		ComputeRoCEAmount:       realDeviceDetailResponse.ComputeRoCEAmount,
		ComputeRoCERate:         realDeviceDetailResponse.ComputeRoCERate,
		StorageRoCEAmount:       realDeviceDetailResponse.StorageRoCEAmount,
		StorageRoCERate:         realDeviceDetailResponse.StorageRoCERate,
		SupportCloud:            realDeviceDetailResponse.SupportCloud,
		CloudBoot:               realDeviceDetailResponse.CloudBoot,
	}

	// 构造 flavor
	var flavor EbmDescribeInstanceFlavorResponse
	var realFlavorResponse = realResponse.Flavor
	flavor = EbmDescribeInstanceFlavorResponse{
		NumaNodeAmount:  realFlavorResponse.NumaNodeAmount,
		NicAmount:       realFlavorResponse.NicAmount,
		MemSize:         realFlavorResponse.MemSize,
		Ram:             realFlavorResponse.Ram,
		Vcpus:           realFlavorResponse.Vcpus,
		CpuThreadAmount: realFlavorResponse.CpuThreadAmount,
		DeviceType:      realFlavorResponse.DeviceType,
	}

	// 构造 networkInfo
	var networkInfo []EbmDescribeInstanceNetworkInfoResponse
	for _, network_info := range realResponse.NetworkInfo {
		networkInfo = append(networkInfo, EbmDescribeInstanceNetworkInfoResponse{
			SubnetUUID: network_info.SubnetUUID,
			VpcName:    network_info.VpcName,
			VpcID:      network_info.VpcID,
		})
	}

	// 构造 raidDetail
	var raidDetail EbmDescribeInstanceRaidDetailResponse
	var realRaidDetailResponse = realResponse.RaidDetail
	var systemVolume EbmDescribeInstanceSystemVolumeResponse
	var realSystemVolumeResponse = realRaidDetailResponse.SystemVolume
	systemVolume = EbmDescribeInstanceSystemVolumeResponse{
		Uuid:         realSystemVolumeResponse.Uuid,
		VolumeType:   realSystemVolumeResponse.VolumeType,
		Name:         realSystemVolumeResponse.Name,
		VolumeDetail: realSystemVolumeResponse.VolumeDetail,
	}
	var dataVolume EbmDescribeInstanceDataVolumeResponse
	var realDataVolumeResponse = realRaidDetailResponse.DataVolume
	dataVolume = EbmDescribeInstanceDataVolumeResponse{
		Uuid:         realDataVolumeResponse.Uuid,
		VolumeType:   realDataVolumeResponse.VolumeType,
		Name:         realDataVolumeResponse.Name,
		VolumeDetail: realDataVolumeResponse.VolumeDetail,
	}
	raidDetail = EbmDescribeInstanceRaidDetailResponse{
		SystemVolume: systemVolume,
		DataVolume:   dataVolume,
	}

	// 构建 realResponse
	var result = EbmDescribeInstanceResponse{
		RegionID:           realResponse.RegionID,
		Region:             realResponse.Region,
		AzName:             realResponse.AzName,
		ResourceID:         realResponse.ResourceID,
		InstanceUUID:       realResponse.InstanceUUID,
		DeviceUUID:         realResponse.DeviceUUID,
		DeviceType:         realResponse.DeviceType,
		DisplayName:        realResponse.DisplayName,
		InstanceName:       realResponse.InstanceName,
		Description:        realResponse.Description,
		ZabbixName:         realResponse.ZabbixName,
		SystemVolumeRaidID: realResponse.SystemVolumeRaidID,
		DataVolumeRaidID:   realResponse.DataVolumeRaidID,
		ImageID:            realResponse.ImageID,
		OsType:             realResponse.OsType,
		OsTypeName:         realResponse.OsTypeName,
		VpcID:              realResponse.VpcID,
		VpcName:            realResponse.VpcName,
		SubnetID:           realResponse.SubnetID,
		PublicIP:           realResponse.PublicIP,
		PrivateIP:          realResponse.PrivateIP,
		PublicIPv6:         realResponse.PublicIPv6,
		PrivateIPv6:        realResponse.PrivateIPv6,
		EbmState:           realResponse.EbmState,
		Flavor:             flavor,
		Interfaces:         interfaces,
		NetworkInfo:        networkInfo,
		RaidDetail:         raidDetail,
		AttachedVolumes:    realResponse.AttachedVolumes,
		DeviceDetail:       deviceDetail,
		Freezing:           realResponse.Freezing,
		Expired:            realResponse.Expired,
		ReleaseDate:        realResponse.ReleaseDate,
		CreateTime:         realResponse.CreateTime,
		UpdatedTime:        realResponse.UpdatedTime,
		ExpiredTime:        realResponse.ExpiredTime,
		OnDemand:           realResponse.OnDemand,
	}
	return &result, nil
}

type EbmDescribeInstanceRealRequest struct {
	RegionID         string `json:"regionID,omitempty"`
	AzName           string `json:"azName,omitempty"`
	ResourceID       string `json:"resourceID,omitempty"`
	Ip               string `json:"ip,omitempty"`
	InstanceName     string `json:"instanceName,omitempty"`
	VpcID            string `json:"vpcID,omitempty"`
	SubnetID         string `json:"subnetID,omitempty"`
	DeviceType       string `json:"deviceType,omitempty"`
	DeviceUUIDList   string `json:"deviceUUIDList,omitempty"`
	QueryContent     string `json:"queryContent,omitempty"`
	InstanceUUIDList string `json:"instanceUUIDList,omitempty"`
	InstanceUUID     string `json:"instanceUUID,omitempty"`
	Status           string `json:"status,omitempty"`
	Sort             string `json:"sort,omitempty"`
	Asc              *bool  `json:"asc,omitempty"`
	VipID            string `json:"vipID,omitempty"`
	VolumeUUID       string `json:"volumeUUID,omitempty"`
	PageNo           *int   `json:"pageNo,omitempty"`
	PageSize         *int   `json:"pageSize,omitempty"`
}

type EbmDescribeInstanceRealResponse struct {
	RegionID           string                                       `json:"regionID,omitempty"`
	Region             string                                       `json:"region,omitempty"`
	AzName             string                                       `json:"azName,omitempty"`
	ResourceID         string                                       `json:"resourceID,omitempty"`
	InstanceUUID       string                                       `json:"instanceUUID,omitempty"`
	DeviceUUID         string                                       `json:"deviceUUID,omitempty"`
	DeviceType         string                                       `json:"deviceType,omitempty"`
	DisplayName        string                                       `json:"displayName,omitempty"`
	InstanceName       string                                       `json:"instanceName,omitempty"`
	Description        string                                       `json:"description,omitempty"`
	ZabbixName         string                                       `json:"zabbixName,omitempty"`
	SystemVolumeRaidID string                                       `json:"systemVolumeRaidID,omitempty"`
	DataVolumeRaidID   string                                       `json:"dataVolumeRaidID,omitempty"`
	ImageID            string                                       `json:"imageID,omitempty"`
	OsType             int                                          `json:"osType,omitempty"`
	OsTypeName         string                                       `json:"osTypeName,omitempty"`
	VpcID              string                                       `json:"vpcID,omitempty"`
	VpcName            string                                       `json:"vpcName,omitempty"`
	SubnetID           string                                       `json:"subnetID,omitempty"`
	PublicIP           string                                       `json:"publicIP,omitempty"`
	PrivateIP          string                                       `json:"privateIP,omitempty"`
	PublicIPv6         string                                       `json:"publicIPv6,omitempty"`
	PrivateIPv6        string                                       `json:"privateIPv6,omitempty"`
	EbmState           string                                       `json:"ebmState,omitempty"`
	Flavor             EbmDescribeInstanceFlavorRealResponse        `json:"flavor,omitempty"`
	Interfaces         []EbmDescribeInstanceInterfacesRealResponse  `json:"interfaces,omitempty"`
	NetworkInfo        []EbmDescribeInstanceNetworkInfoRealResponse `json:"networkInfo,omitempty"`
	RaidDetail         EbmDescribeInstanceRaidDetailRealResponse    `json:"raidDetail,omitempty"`
	AttachedVolumes    []string                                     `json:"attachedVolumes,omitempty"`
	DeviceDetail       EbmDescribeInstanceDeviceDetailRealResponse  `json:"deviceDetail,omitempty"`
	Freezing           bool                                         `json:"freezing,omitempty"`
	Expired            bool                                         `json:"expired,omitempty"`
	ReleaseDate        string                                       `json:"releaseDate,omitempty"`
	CreateTime         string                                       `json:"createTime,omitempty"`
	UpdatedTime        string                                       `json:"updatedTime,omitempty"`
	ExpiredTime        string                                       `json:"expiredTime,omitempty"`
	OnDemand           bool                                         `json:"onDemand,omitempty"`
}

type EbmDescribeInstanceFlavorRealResponse struct {
	NumaNodeAmount  int    `json:"numaNodeAmount,omitempty"`
	NicAmount       int    `json:"nicAmount,omitempty"`
	MemSize         int    `json:"memSize,omitempty"`
	Ram             int    `json:"ram,omitempty"`
	Vcpus           int    `json:"vcpus,omitempty"`
	CpuThreadAmount int    `json:"cpuThreadAmount,omitempty"`
	DeviceType      string `json:"deviceType,omitempty"`
}

type EbmDescribeInstanceInterfacesRealResponse struct {
	InterfaceUUID  string                                          `json:"interfaceUUID,omitempty"`
	Master         bool                                            `json:"master,omitempty"`
	VpcUUID        string                                          `json:"vpcUUID,omitempty"`
	SubnetUUID     string                                          `json:"subnetUUID,omitempty"`
	PortUUID       string                                          `json:"portUUID,omitempty"`
	Ipv4           string                                          `json:"ipv4,omitempty"`
	Ipv4Gateway    string                                          `json:"ipv4Gateway,omitempty"`
	Ipv6           string                                          `json:"ipv6,omitempty"`
	Ipv6Gateway    string                                          `json:"ipv6Gateway,omitempty"`
	VipUUIDList    []string                                        `json:"vipUUIDList,omitempty"`
	VipList        []string                                        `json:"vipList,omitempty"`
	SecurityGroups []EbmDescribeInstanceSecurityGroupsRealResponse `json:"securityGroups,omitempty"`
}

type EbmDescribeInstanceSecurityGroupsRealResponse struct {
	SecurityGroupID   string `json:"securityGroupID,omitempty"`
	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

type EbmDescribeInstanceNetworkInfoRealResponse struct {
	SubnetUUID string `json:"subnetUUID,omitempty"`
	VpcName    string `json:"vpcName,omitempty"`
	VpcID      string `json:"vpcID,omitempty"`
}

type EbmDescribeInstanceRaidDetailRealResponse struct {
	SystemVolume EbmDescribeInstanceSystemVolumeRealResponse `json:"systemVolume,omitempty"`
	DataVolume   EbmDescribeInstanceDataVolumeRealResponse   `json:"dataVolume,omitempty"`
}

type EbmDescribeInstanceSystemVolumeRealResponse struct {
	Uuid         string `json:"uuid,omitempty"`
	VolumeType   string `json:"volumeType,omitempty"`
	Name         string `json:"name,omitempty"`
	VolumeDetail string `json:"volumeDetail,omitempty"`
}

type EbmDescribeInstanceDataVolumeRealResponse struct {
	Uuid         string `json:"uuid,omitempty"`
	VolumeType   string `json:"volumeType,omitempty"`
	Name         string `json:"name,omitempty"`
	VolumeDetail string `json:"volumeDetail,omitempty"`
}

type EbmDescribeInstanceDeviceDetailRealResponse struct {
	CpuSockets              int    `json:"cpuSockets,omitempty"`
	NumaNodeAmount          int    `json:"numaNodeAmount,omitempty"`
	CpuAmount               int    `json:"cpuAmount,omitempty"`
	CpuThreadAmount         int    `json:"cpuThreadAmount,omitempty"`
	CpuManufacturer         string `json:"cpuManufacturer,omitempty"`
	CpuModel                string `json:"cpuModel,omitempty"`
	CpuFrequency            string `json:"cpuFrequency,omitempty"`
	MemAmount               int    `json:"memAmount,omitempty"`
	MemSize                 int    `json:"memSize,omitempty"`
	MemFrequency            int    `json:"memFrequency,omitempty"`
	NicAmount               int    `json:"nicAmount,omitempty"`
	NicRate                 int    `json:"nicRate,omitempty"`
	SystemVolumeAmount      int    `json:"systemVolumeAmount,omitempty"`
	SystemVolumeSize        int    `json:"systemVolumeSize,omitempty"`
	SystemVolumeType        string `json:"systemVolumeType,omitempty"`
	SystemVolumeInterface   string `json:"systemVolumeInterface,omitempty"`
	SystemVolumeDescription string `json:"systemVolumeDescription,omitempty"`
	DataVolumeAmount        int    `json:"dataVolumeAmount,omitempty"`
	DataVolumeSize          int    `json:"dataVolumeSize,omitempty"`
	DataVolumeInterface     string `json:"dataVolumeInterface,omitempty"`
	DataVolumeType          string `json:"dataVolumeType,omitempty"`
	DataVolumeDescription   string `json:"dataVolumeDescription,omitempty"`
	SmartNicExist           bool   `json:"smartNicExist,omitempty"`
	NvmeVolumeAmount        int    `json:"nvmeVolumeAmount,omitempty"`
	NvmeVolumeSize          int    `json:"nvmeVolumeSize,omitempty"`
	NvmeVolumeType          string `json:"nvmeVolumeType,omitempty"`
	NvmeVolumeInterface     string `json:"nvmeVolumeInterface,omitempty"`
	GpuAmount               int    `json:"gpuAmount,omitempty"`
	GpuSize                 int    `json:"gpuSize,omitempty"`
	GpuManufacturer         string `json:"gpuManufacturer,omitempty"`
	GpuModel                string `json:"gpuModel,omitempty"`
	ComputeIBAmount         int    `json:"computeIBAmount,omitempty"`
	ComputeIBRate           int    `json:"computeIBRate,omitempty"`
	StorageIBAmount         int    `json:"storageIBAmount,omitempty"`
	StorageIBRate           int    `json:"storageIBRate,omitempty"`
	ComputeRoCEAmount       int    `json:"computeRoCEAmount,omitempty"`
	ComputeRoCERate         int    `json:"computeRoCERate,omitempty"`
	StorageRoCEAmount       int    `json:"storageRoCEAmount,omitempty"`
	StorageRoCERate         int    `json:"storageRoCERate,omitempty"`
	SupportCloud            bool   `json:"supportCloud,omitempty"`
	CloudBoot               bool   `json:"cloudBoot,omitempty"`
}

type EbmDescribeInstanceRequest struct {
	RegionID         string
	AzName           string
	ResourceID       string
	Ip               string
	InstanceName     string
	VpcID            string
	SubnetID         string
	DeviceType       string
	DeviceUUIDList   string
	QueryContent     string
	InstanceUUIDList string
	InstanceUUID     string
	Status           string
	Sort             string
	Asc              *bool
	VipID            string
	VolumeUUID       string
	PageNo           *int
	PageSize         *int
}

type EbmDescribeInstanceResponse struct {
	RegionID           string
	Region             string
	AzName             string
	ResourceID         string
	InstanceUUID       string
	DeviceUUID         string
	DeviceType         string
	DisplayName        string
	InstanceName       string
	Description        string
	ZabbixName         string
	SystemVolumeRaidID string
	DataVolumeRaidID   string
	ImageID            string
	OsType             int
	OsTypeName         string
	VpcID              string
	VpcName            string
	SubnetID           string
	PublicIP           string
	PrivateIP          string
	PublicIPv6         string
	PrivateIPv6        string
	EbmState           string
	Flavor             EbmDescribeInstanceFlavorResponse
	Interfaces         []EbmDescribeInstanceInterfacesResponse
	NetworkInfo        []EbmDescribeInstanceNetworkInfoResponse
	RaidDetail         EbmDescribeInstanceRaidDetailResponse
	AttachedVolumes    []string
	DeviceDetail       EbmDescribeInstanceDeviceDetailResponse
	Freezing           bool
	Expired            bool
	ReleaseDate        string
	CreateTime         string
	UpdatedTime        string
	ExpiredTime        string
	OnDemand           bool
}

type EbmDescribeInstanceFlavorResponse struct {
	NumaNodeAmount  int
	NicAmount       int
	MemSize         int
	Ram             int
	Vcpus           int
	CpuThreadAmount int
	DeviceType      string
}

type EbmDescribeInstanceInterfacesResponse struct {
	InterfaceUUID  string
	Master         bool
	VpcUUID        string
	SubnetUUID     string
	PortUUID       string
	Ipv4           string
	Ipv4Gateway    string
	Ipv6           string
	Ipv6Gateway    string
	VipUUIDList    []string
	VipList        []string
	SecurityGroups []EbmDescribeInstanceSecurityGroupsResponse
}

type EbmDescribeInstanceSecurityGroupsResponse struct {
	SecurityGroupID   string
	SecurityGroupName string
}

type EbmDescribeInstanceNetworkInfoResponse struct {
	SubnetUUID string
	VpcName    string
	VpcID      string
}

type EbmDescribeInstanceRaidDetailResponse struct {
	SystemVolume EbmDescribeInstanceSystemVolumeResponse
	DataVolume   EbmDescribeInstanceDataVolumeResponse
}

type EbmDescribeInstanceSystemVolumeResponse struct {
	Uuid         string
	VolumeType   string
	Name         string
	VolumeDetail string
}

type EbmDescribeInstanceDataVolumeResponse struct {
	Uuid         string
	VolumeType   string
	Name         string
	VolumeDetail string
}

type EbmDescribeInstanceDeviceDetailResponse struct {
	CpuSockets              int
	NumaNodeAmount          int
	CpuAmount               int
	CpuThreadAmount         int
	CpuManufacturer         string
	CpuModel                string
	CpuFrequency            string
	MemAmount               int
	MemSize                 int
	MemFrequency            int
	NicAmount               int
	NicRate                 int
	SystemVolumeAmount      int
	SystemVolumeSize        int
	SystemVolumeType        string
	SystemVolumeInterface   string
	SystemVolumeDescription string
	DataVolumeAmount        int
	DataVolumeSize          int
	DataVolumeInterface     string
	DataVolumeType          string
	DataVolumeDescription   string
	SmartNicExist           bool
	NvmeVolumeAmount        int
	NvmeVolumeSize          int
	NvmeVolumeType          string
	NvmeVolumeInterface     string
	GpuAmount               int
	GpuSize                 int
	GpuManufacturer         string
	GpuModel                string
	ComputeIBAmount         int
	ComputeIBRate           int
	StorageIBAmount         int
	StorageIBRate           int
	ComputeRoCEAmount       int
	ComputeRoCERate         int
	StorageRoCEAmount       int
	StorageRoCERate         int
	SupportCloud            bool
	CloudBoot               bool
}
