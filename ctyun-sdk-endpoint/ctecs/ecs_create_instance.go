package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsCreateInstanceApi 创建一台按量付费或包年包月的云主机
// https://www.ctyun.cn/document/10026730/10106580
type EcsCreateInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsCreateInstanceApi(client *ctyunsdk.CtyunClient) *EcsCreateInstanceApi {
	return &EcsCreateInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/create-instance",
		},
	}
}

func (this *EcsCreateInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsCreateInstanceRequest) (*EcsCreateInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	var networkCardList []ecsCreateInstanceNetworkCardListRealRequest
	for _, request := range req.NetworkCardList {
		networkCardList = append(networkCardList, ecsCreateInstanceNetworkCardListRealRequest{
			NicName:  request.NicName,
			FixedIP:  request.FixedIp,
			IsMaster: request.IsMaster,
			SubnetID: request.SubnetId,
		})
	}

	var dataDiskList []ecsCreateInstanceDataDiskListRealRequest
	for _, request := range req.DataDiskList {
		dataDiskList = append(dataDiskList, ecsCreateInstanceDataDiskListRealRequest{
			DiskMode: request.DiskMode,
			DiskName: request.DiskName,
			DiskType: request.DiskType,
			DiskSize: request.DiskSize,
		})
	}
	var LabelList []ecsCreateInstanceLabelListRealRequest
	for _, request := range req.LabelList {
		LabelList = append(LabelList, ecsCreateInstanceLabelListRealRequest{
			LabelKey:   request.LabelKey,
			LabelValue: request.LabelValue,
		})
	}
	_, err := builder.WriteJson(&ecsCreateInstanceRealRequest{
		ClientToken:     req.ClientToken,
		RegionID:        req.RegionId,
		AzName:          req.AzName,
		InstanceName:    req.InstanceName,
		DisplayName:     req.DisplayName,
		FlavorID:        req.FlavorId,
		ImageType:       req.ImageType,
		ImageID:         req.ImageId,
		BootDiskType:    req.BootDiskType,
		BootDiskSize:    req.BootDiskSize,
		VpcID:           req.VpcId,
		OnDemand:        req.OnDemand,
		NetworkCardList: networkCardList,
		ExtIP:           req.ExtIp,
		ProjectID:       req.ProjectID,
		SecGroupList:    req.SecGroupList,
		KeyPairID:       req.KeyPairID,
		UserPassword:    req.UserPassword,
		CycleCount:      req.CycleCount,
		CycleType:       req.CycleType,
		AutoRenewStatus: req.AutoRenewStatus,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse ecsCreateInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsCreateInstanceResponse{
		RegionId:         realResponse.RegionID,
		MasterOrderId:    realResponse.MasterOrderID,
		MasterResourceId: realResponse.MasterResourceID,
		MasterOrderNo:    realResponse.MasterOrderNO,
	}, nil
}

type ecsCreateInstanceNetworkCardListRealRequest struct {
	NicName  string `json:"nicName"`
	FixedIP  string `json:"fixedIP"`
	IsMaster bool   `json:"isMaster"`
	SubnetID string `json:"subnetID"`
}

type ecsCreateInstanceDataDiskListRealRequest struct {
	DiskMode string `json:"diskMode"`
	DiskName string `json:"diskName"`
	DiskType string `json:"diskType"`
	DiskSize int    `json:"diskSize"`
}

type ecsCreateInstanceLabelListRealRequest struct {
	LabelKey   string `json:"labelKey"`
	LabelValue string `json:"labelValue"`
}

type ecsCreateInstanceRealRequest struct {
	ClientToken     string                                        `json:"clientToken"`
	RegionID        string                                        `json:"regionID"`
	AzName          string                                        `json:"azName"`
	InstanceName    string                                        `json:"instanceName"`
	DisplayName     string                                        `json:"displayName"`
	FlavorID        string                                        `json:"flavorID"`
	ImageType       int                                           `json:"imageType"`
	ImageID         string                                        `json:"imageID"`
	BootDiskType    string                                        `json:"bootDiskType"`
	BootDiskSize    int                                           `json:"bootDiskSize"`
	VpcID           string                                        `json:"vpcID"`
	OnDemand        bool                                          `json:"onDemand"`
	NetworkCardList []ecsCreateInstanceNetworkCardListRealRequest `json:"networkCardList"`
	ExtIP           string                                        `json:"extIP"`
	ProjectID       string                                        `json:"projectID"`
	SecGroupList    []string                                      `json:"secGroupList"`
	DataDiskList    []ecsCreateInstanceDataDiskListRealRequest    `json:"dataDiskList"`
	IpVersion       string                                        `json:"ipVersion"`
	Bandwidth       int                                           `json:"bandwidth"`
	Ipv6AddressID   string                                        `json:"ipv6AddressID"`
	EipID           string                                        `json:"eipID"`
	AffinityGroupID string                                        `json:"affinityGroupID"`
	KeyPairID       string                                        `json:"keyPairID"`
	UserPassword    string                                        `json:"userPassword"`
	CycleCount      int                                           `json:"cycleCount"`
	CycleType       string                                        `json:"cycleType"`
	AutoRenewStatus int                                           `json:"autoRenewStatus"`
	UserData        string                                        `json:"userData"`
	PayVoucherPrice float64                                       `json:"payVoucherPrice"`
	LabelList       []ecsCreateInstanceLabelListRealRequest       `json:"labelList"`
	GpuDriverKits   string                                        `json:"gpuDriverKits"`
	MonitorService  bool                                          `json:"monitorService"`
}

type ecsCreateInstanceRealResponse struct {
	RegionID         string `json:"regionID"`
	MasterOrderID    string `json:"masterOrderID"`
	MasterResourceID string `json:"masterResourceID"`
	MasterOrderNO    string `json:"masterOrderNO"`
}

type EcsCreateInstanceRequest struct {
	ClientToken     string                                    // 客户端存根
	RegionId        string                                    // 资源池ID
	AzName          string                                    // 可用区名称
	InstanceName    string                                    // 云主机名称，不可以使用已存在的云主机名称
	DisplayName     string                                    // 云主机显示名称
	FlavorId        string                                    // 云主机规格ID
	ImageType       int                                       // 镜像类型
	ImageId         string                                    // 镜像ID
	BootDiskType    string                                    // 系统盘类型，取值范围： SATA（普通IO）， SAS（高IO）， SSD（超高IO）， SSD-genric（通用型SSD）， FAST-SSD（极速型SSD），您可以查看磁盘类型及性能介绍来了解磁盘类型及其对应性能指标
	BootDiskSize    int                                       // 系统盘大小单位为GiB，取值范围：[40, 32768]
	VpcId           string                                    // 虚拟私有云ID
	OnDemand        bool                                      // 购买方式，取值范围： false（按周期）， true（按需）
	NetworkCardList []EcsCreateInstanceNetworkCardListRequest // 网卡信息列表
	ExtIp           string                                    // 是否使用弹性公网IP，取值范围： 0（不使用）， 1（自动分配）， 2（使用已有）
	ProjectID       string                                    // 企业项目ID
	SecGroupList    []string                                  // 安全组ID列表
	DataDiskList    []EcsCreateInstanceDataDiskListRequest    // 数据盘信息列表，注：同一云主机下最多可挂载8块数据盘
	IpVersion       string                                    // 弹性IP版本，取值范围：<br />ipv4（v4地址），<br />ipv6（v6地址），<br />不指定默认为ipv4。注：请先确认该资源池是否支持ipv6（多可用区类资源池暂不支持）
	Bandwidth       int                                       // 带宽大小，单位为Mbit/s，取值范围：[1, 2000]
	Ipv6AddressID   string                                    // 弹性公网IPv6的ID，注：多可用区类资源池暂不支持；填写该参数时请填写ipVersion为ipv6
	EipID           string                                    // 弹性公网IP的ID
	AffinityGroupID string                                    // 云主机组ID
	KeyPairID       string                                    // 密钥对ID
	UserPassword    string                                    // 用户密码
	CycleCount      int                                       // 订购时长
	CycleType       string                                    // 订购周期类型
	AutoRenewStatus int                                       // 是否自动续订，取值范围： 0（不续费）， 1（自动续费）， 注：按月购买，自动续订周期为3个月；按年购买，自动续订周期为1年
	UserData        string                                    // 用户自定义数据，需要以Base64方式编码，Base64编码后的长度限制为1-16384字符
	PayVoucherPrice float64                                   // 代金券，满足以下规则：两位小数，不足两位自动补0，超过两位小数无效；不可为负数；注：字段为0时表示不使用代金券，默认不使用
	LabelList       []EcsCreateInstanceLabelListRequest       // 标签信息列表，注：单台云主机最多可绑定10个标签；主机创建完成后，云主机变为运行状态，此时标签仍可能未绑定，需等待一段时间（0~10分钟）。
	GpuDriverKits   string                                    // GPU云主机安装驱动的工具包，仅在同时选择NVIDIA显卡、计算加速型、linux公共镜像三个条件下，支持安装驱动
	MonitorService  bool                                      // 监控参数，支持通过该参数指定云主机在创建后是否开启详细监控，取值范围： false（不开启），true（开启）；注：若指定该参数为true或不指定该参数，云主机内默认开启最新详细监控服务。若指定该参数为false，默认公共镜像不开启最新监控服务；私有镜像使用镜像中保留的监控服务。说明：仅部分资源池支持monitorService参数，
}

type EcsCreateInstanceNetworkCardListRequest struct {
	NicName  string // 长度2~32，支持拉丁字母、中文、数字、下划线、连字符，中文或英文字母开头，不能以http:或https:开头
	FixedIp  string // 内网IPv4地址，注：不可使用已占用IP
	IsMaster bool   // 是否主网卡，取值范围： true（表示主网卡）， false（表示扩展网卡） 注：只能含有一个主网卡
	SubnetId string // 子网ID
}

type EcsCreateInstanceDataDiskListRequest struct {
	DiskMode string
	DiskName string
	DiskType string
	DiskSize int
}

type EcsCreateInstanceLabelListRequest struct {
	LabelKey   string
	LabelValue string
}

type EcsCreateInstanceResponse struct {
	RegionId         string
	MasterOrderId    string
	MasterResourceId string
	MasterOrderNo    string
}
