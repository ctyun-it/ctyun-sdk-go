package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
	"strconv"
)

// EbmListInstanceApi 批量查询物理机_v4_plus
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=6941&data=97&isNormal=1
type EbmListInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmListInstanceApi(client *ctyunsdk.CtyunClient) *EbmListInstanceApi {
	return &EbmListInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ebm/list-instance",
		},
	}
}

func (this *EbmListInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmListInstanceRequest) (*EbmListInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", req.RegionID).
		AddParam("azName", req.AzName).
		AddParam("resourceID", req.ResourceID).
		AddParam("ip", req.Ip).
		AddParam("instanceName", req.InstanceName).
		AddParam("vpcID", req.VpcID).
		AddParam("subnetID", req.SubnetID).
		AddParam("deviceType", req.DeviceType).
		AddParam("deviceUUIDList", req.DeviceUUIDList).
		AddParam("queryContent", req.QueryContent).
		AddParam("instanceUUIDList", req.InstanceUUIDList).
		AddParam("instanceUUID", req.InstanceUUID).
		AddParam("status", req.Status).
		AddParam("sort", req.Sort).
		AddParam("vipID", req.VipID).
		AddParam("volumeUUID", req.VolumeUUID)

	if req.Asc != nil {
		builder.AddParam("asc", strconv.FormatBool(*req.Asc))
	}
	if req.PageNo != nil {
		builder.AddParam("pageNo", strconv.Itoa(*req.PageNo))
	}
	if req.PageSize != nil {
		builder.AddParam("pageSize", strconv.Itoa(*req.PageSize))
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmListInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EbmListInstanceResultsResponse
	// 遍历每一个返回结果，构造返回参数
	for _, result := range realResponse.Results {
		var interfaces []EbmListInstanceInterfacesResponse
		// 构造 interfaces中的sgs
		for _, interface_info := range result.Interfaces {
			var sgs []EbmListInstanceSecurityGroupsResponse
			for _, sg := range interface_info.SecurityGroups {
				sgs = append(sgs, EbmListInstanceSecurityGroupsResponse{
					SecurityGroupID:   sg.SecurityGroupID,
					SecurityGroupName: sg.SecurityGroupName,
				})
			}
			interfaces = append(interfaces, EbmListInstanceInterfacesResponse{
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
		var deviceDetail EbmListInstanceDeviceDetailResponse
		var realDeviceDetailResponse EbmListInstanceDeviceDetailRealResponse = result.DeviceDetail
		deviceDetail = EbmListInstanceDeviceDetailResponse{
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
		var flavor EbmListInstanceFlavorResponse
		var realFlavorResponse = result.Flavor
		flavor = EbmListInstanceFlavorResponse{
			NumaNodeAmount:  realFlavorResponse.NumaNodeAmount,
			NicAmount:       realFlavorResponse.NicAmount,
			MemSize:         realFlavorResponse.MemSize,
			Ram:             realFlavorResponse.Ram,
			Vcpus:           realFlavorResponse.Vcpus,
			CpuThreadAmount: realFlavorResponse.CpuThreadAmount,
			DeviceType:      realFlavorResponse.DeviceType,
		}

		// 构造 networkInfo
		var networkInfo []EbmListInstanceNetworkInfoResponse
		for _, network_info := range result.NetworkInfo {
			networkInfo = append(networkInfo, EbmListInstanceNetworkInfoResponse{
				SubnetUUID: network_info.SubnetUUID,
				VpcName:    network_info.VpcName,
				VpcID:      network_info.VpcID,
			})
		}

		// 构造 raidDetail
		var raidDetail EbmListInstanceRaidDetailResponse
		var realRaidDetailResponse = result.RaidDetail
		var systemVolume EbmListInstanceSystemVolumeResponse
		var realSystemVolumeResponse = realRaidDetailResponse.SystemVolume
		systemVolume = EbmListInstanceSystemVolumeResponse{
			Uuid:         realSystemVolumeResponse.Uuid,
			VolumeType:   realSystemVolumeResponse.VolumeType,
			Name:         realSystemVolumeResponse.Name,
			VolumeDetail: realSystemVolumeResponse.VolumeDetail,
		}
		var dataVolume EbmListInstanceDataVolumeResponse
		var realDataVolumeResponse = realRaidDetailResponse.DataVolume
		dataVolume = EbmListInstanceDataVolumeResponse{
			Uuid:         realDataVolumeResponse.Uuid,
			VolumeType:   realDataVolumeResponse.VolumeType,
			Name:         realDataVolumeResponse.Name,
			VolumeDetail: realDataVolumeResponse.VolumeDetail,
		}
		raidDetail = EbmListInstanceRaidDetailResponse{
			SystemVolume: systemVolume,
			DataVolume:   dataVolume,
		}

		// 构建 results中的元素
		results = append(results, EbmListInstanceResultsResponse{
			RegionID:           result.RegionID,
			Region:             result.Region,
			AzName:             result.AzName,
			ResourceID:         result.ResourceID,
			InstanceUUID:       result.InstanceUUID,
			DeviceUUID:         result.DeviceUUID,
			DeviceType:         result.DeviceType,
			DisplayName:        result.DisplayName,
			InstanceName:       result.InstanceName,
			Description:        result.Description,
			ZabbixName:         result.ZabbixName,
			SystemVolumeRaidID: result.SystemVolumeRaidID,
			DataVolumeRaidID:   result.DataVolumeRaidID,
			ImageID:            result.ImageID,
			OsType:             result.OsType,
			OsTypeName:         result.OsTypeName,
			VpcID:              result.VpcID,
			VpcName:            result.VpcName,
			SubnetID:           result.SubnetID,
			PublicIP:           result.PublicIP,
			PrivateIP:          result.PrivateIP,
			PublicIPv6:         result.PublicIPv6,
			PrivateIPv6:        result.PrivateIPv6,
			EbmState:           result.EbmState,
			Flavor:             flavor,
			Interfaces:         interfaces,
			NetworkInfo:        networkInfo,
			RaidDetail:         raidDetail,
			AttachedVolumes:    result.AttachedVolumes,
			DeviceDetail:       deviceDetail,
			Freezing:           result.Freezing,
			Expired:            result.Expired,
			ReleaseDate:        result.ReleaseDate,
			CreateTime:         result.CreateTime,
			UpdatedTime:        result.UpdatedTime,
			ExpiredTime:        result.ExpiredTime,
			OnDemand:           result.OnDemand,
		})
	}
	return &EbmListInstanceResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EbmListInstanceRealRequest struct {
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

type EbmListInstanceRealResponse struct {
	CurrentCount int                                  `json:"currentCount,omitempty"`
	TotalCount   int                                  `json:"totalCount,omitempty"`
	TotalPage    int                                  `json:"totalPage,omitempty"`
	Results      []EbmListInstanceResultsRealResponse `json:"results,omitempty"`
}

type EbmListInstanceResultsRealResponse struct {
	RegionID           string                                   `json:"regionID,omitempty"`
	Region             string                                   `json:"region,omitempty"`
	AzName             string                                   `json:"azName,omitempty"`
	ResourceID         string                                   `json:"resourceID,omitempty"`
	InstanceUUID       string                                   `json:"instanceUUID,omitempty"`
	DeviceUUID         string                                   `json:"deviceUUID,omitempty"`
	DeviceType         string                                   `json:"deviceType,omitempty"`
	DisplayName        string                                   `json:"displayName,omitempty"`
	InstanceName       string                                   `json:"instanceName,omitempty"`
	Description        string                                   `json:"description,omitempty"`
	ZabbixName         string                                   `json:"zabbixName,omitempty"`
	SystemVolumeRaidID string                                   `json:"systemVolumeRaidID,omitempty"`
	DataVolumeRaidID   string                                   `json:"dataVolumeRaidID,omitempty"`
	ImageID            string                                   `json:"imageID,omitempty"`
	OsType             int                                      `json:"osType,omitempty"`
	OsTypeName         string                                   `json:"osTypeName,omitempty"`
	VpcID              string                                   `json:"vpcID,omitempty"`
	VpcName            string                                   `json:"vpcName,omitempty"`
	SubnetID           string                                   `json:"subnetID,omitempty"`
	PublicIP           string                                   `json:"publicIP,omitempty"`
	PrivateIP          string                                   `json:"privateIP,omitempty"`
	PublicIPv6         string                                   `json:"publicIPv6,omitempty"`
	PrivateIPv6        string                                   `json:"privateIPv6,omitempty"`
	EbmState           string                                   `json:"ebmState,omitempty"`
	Flavor             EbmListInstanceFlavorRealResponse        `json:"flavor,omitempty"`
	Interfaces         []EbmListInstanceInterfacesRealResponse  `json:"interfaces,omitempty"`
	NetworkInfo        []EbmListInstanceNetworkInfoRealResponse `json:"networkInfo,omitempty"`
	RaidDetail         EbmListInstanceRaidDetailRealResponse    `json:"raidDetail,omitempty"`
	AttachedVolumes    []string                                 `json:"attachedVolumes,omitempty"`
	DeviceDetail       EbmListInstanceDeviceDetailRealResponse  `json:"deviceDetail,omitempty"`
	Freezing           bool                                     `json:"freezing,omitempty"`
	Expired            bool                                     `json:"expired,omitempty"`
	ReleaseDate        string                                   `json:"releaseDate,omitempty"`
	CreateTime         string                                   `json:"createTime,omitempty"`
	UpdatedTime        string                                   `json:"updatedTime,omitempty"`
	ExpiredTime        string                                   `json:"expiredTime,omitempty"`
	OnDemand           bool                                     `json:"onDemand,omitempty"`
}

type EbmListInstanceFlavorRealResponse struct {
	NumaNodeAmount  int    `json:"numaNodeAmount,omitempty"`
	NicAmount       int    `json:"nicAmount,omitempty"`
	MemSize         int    `json:"memSize,omitempty"`
	Ram             int    `json:"ram,omitempty"`
	Vcpus           int    `json:"vcpus,omitempty"`
	CpuThreadAmount int    `json:"cpuThreadAmount,omitempty"`
	DeviceType      string `json:"deviceType,omitempty"`
}

type EbmListInstanceInterfacesRealResponse struct {
	InterfaceUUID  string                                      `json:"interfaceUUID,omitempty"`
	Master         bool                                        `json:"master,omitempty"`
	VpcUUID        string                                      `json:"vpcUUID,omitempty"`
	SubnetUUID     string                                      `json:"subnetUUID,omitempty"`
	PortUUID       string                                      `json:"portUUID,omitempty"`
	Ipv4           string                                      `json:"ipv4,omitempty"`
	Ipv4Gateway    string                                      `json:"ipv4Gateway,omitempty"`
	Ipv6           string                                      `json:"ipv6,omitempty"`
	Ipv6Gateway    string                                      `json:"ipv6Gateway,omitempty"`
	VipUUIDList    []string                                    `json:"vipUUIDList,omitempty"`
	VipList        []string                                    `json:"vipList,omitempty"`
	SecurityGroups []EbmListInstanceSecurityGroupsRealResponse `json:"securityGroups,omitempty"`
}

type EbmListInstanceSecurityGroupsRealResponse struct {
	SecurityGroupID   string `json:"securityGroupID,omitempty"`
	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

type EbmListInstanceNetworkInfoRealResponse struct {
	SubnetUUID string `json:"subnetUUID,omitempty"`
	VpcName    string `json:"vpcName,omitempty"`
	VpcID      string `json:"vpcID,omitempty"`
}

type EbmListInstanceRaidDetailRealResponse struct {
	SystemVolume EbmListInstanceSystemVolumeRealResponse `json:"systemVolume,omitempty"`
	DataVolume   EbmListInstanceDataVolumeRealResponse   `json:"dataVolume,omitempty"`
}

type EbmListInstanceSystemVolumeRealResponse struct {
	Uuid         string `json:"uuid,omitempty"`
	VolumeType   string `json:"volumeType,omitempty"`
	Name         string `json:"name,omitempty"`
	VolumeDetail string `json:"volumeDetail,omitempty"`
}

type EbmListInstanceDataVolumeRealResponse struct {
	Uuid         string `json:"uuid,omitempty"`
	VolumeType   string `json:"volumeType,omitempty"`
	Name         string `json:"name,omitempty"`
	VolumeDetail string `json:"volumeDetail,omitempty"`
}

type EbmListInstanceDeviceDetailRealResponse struct {
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

type EbmListInstanceRequest struct {
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

type EbmListInstanceResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EbmListInstanceResultsResponse
}

type EbmListInstanceResultsResponse struct {
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
	Flavor             EbmListInstanceFlavorResponse
	Interfaces         []EbmListInstanceInterfacesResponse
	NetworkInfo        []EbmListInstanceNetworkInfoResponse
	RaidDetail         EbmListInstanceRaidDetailResponse
	AttachedVolumes    []string
	DeviceDetail       EbmListInstanceDeviceDetailResponse
	Freezing           bool
	Expired            bool
	ReleaseDate        string
	CreateTime         string
	UpdatedTime        string
	ExpiredTime        string
	OnDemand           bool
}

type EbmListInstanceFlavorResponse struct {
	NumaNodeAmount  int
	NicAmount       int
	MemSize         int
	Ram             int
	Vcpus           int
	CpuThreadAmount int
	DeviceType      string
}

type EbmListInstanceInterfacesResponse struct {
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
	SecurityGroups []EbmListInstanceSecurityGroupsResponse
}

type EbmListInstanceSecurityGroupsResponse struct {
	SecurityGroupID   string
	SecurityGroupName string
}

type EbmListInstanceNetworkInfoResponse struct {
	SubnetUUID string
	VpcName    string
	VpcID      string
}

type EbmListInstanceRaidDetailResponse struct {
	SystemVolume EbmListInstanceSystemVolumeResponse
	DataVolume   EbmListInstanceDataVolumeResponse
}

type EbmListInstanceSystemVolumeResponse struct {
	Uuid         string
	VolumeType   string
	Name         string
	VolumeDetail string
}

type EbmListInstanceDataVolumeResponse struct {
	Uuid         string
	VolumeType   string
	Name         string
	VolumeDetail string
}

type EbmListInstanceDeviceDetailResponse struct {
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
