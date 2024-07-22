package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBatchCreateInstancesApi 批量创建按量付费或包年包月的云主机
// https://www.ctyun.cn/document/10026730/10106585
type EcsBatchCreateInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBatchCreateInstancesApi(client *ctyunsdk.CtyunClient) *EcsBatchCreateInstancesApi {
	return &EcsBatchCreateInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/batch-create-instances",
		},
	}
}

func (this *EcsBatchCreateInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBatchCreateInstancesRequest) (*EcsBatchCreateInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	var networkCardList []EcsBatchCreateInstancesNetworkCardListRealRequest
	for _, request := range req.NetworkCardList {
		networkCardList = append(networkCardList, EcsBatchCreateInstancesNetworkCardListRealRequest{
			NicName:  request.NicName,
			FixedIP:  request.FixedIP,
			IsMaster: request.IsMaster,
			SubnetID: request.SubnetID,
		})
	}

	var dataDiskList []EcsBatchCreateInstancesDataDiskListRealRequest
	for _, request := range req.DataDiskList {
		dataDiskList = append(dataDiskList, EcsBatchCreateInstancesDataDiskListRealRequest{
			DiskMode: request.DiskMode,
			DiskName: request.DiskName,
			DiskType: request.DiskType,
			DiskSize: request.DiskSize,
		})
	}
	var LabelList []EcsBatchCreateInstancesLabelListRealRequest
	for _, request := range req.LabelList {
		LabelList = append(LabelList, EcsBatchCreateInstancesLabelListRealRequest{
			LabelKey:   request.LabelKey,
			LabelValue: request.LabelValue,
		})
	}
	_, err := builder.WriteJson(&EcsBatchCreateInstancesRealRequest{
		ClientToken:     req.ClientToken,
		RegionID:        req.RegionID,
		AzName:          req.AzName,
		InstanceName:    req.InstanceName,
		DisplayName:     req.DisplayName,
		FlavorID:        req.FlavorID,
		ImageType:       req.ImageType,
		ImageID:         req.ImageID,
		BootDiskType:    req.BootDiskType,
		BootDiskSize:    req.BootDiskSize,
		Bandwidth:       req.Bandwidth,
		OrderCount:      req.OrderCount,
		VpcID:           req.VpcID,
		OnDemand:        req.OnDemand,
		NetworkCardList: networkCardList,
		DataDiskList:    dataDiskList,
		ExtIP:           req.ExtIP,
		ProjectID:       req.ProjectID,
		SecGroupList:    req.SecGroupList,
		KeyPairID:       req.KeyPairID,
		UserPassword:    req.UserPassword,
		CycleCount:      req.CycleCount,
		CycleType:       req.CycleType,
		AutoRenewStatus: req.AutoRenewStatus,
		PayVoucherPrice: req.PayVoucherPrice,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBatchCreateInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsBatchCreateInstancesResponse{
		RegionID:         realResponse.RegionID,
		MasterOrderID:    realResponse.MasterOrderID,
		MasterResourceID: realResponse.MasterResourceID,
		MasterOrderNO:    realResponse.MasterOrderNO,
	}, nil
}

type EcsBatchCreateInstancesNetworkCardListRealRequest struct {
	NicName  string `json:"nicName,omitempty"`
	FixedIP  string `json:"fixedIP,omitempty"`
	IsMaster *bool  `json:"isMaster,omitempty"`
	SubnetID string `json:"subnetID,omitempty"`
}

type EcsBatchCreateInstancesDataDiskListRealRequest struct {
	DiskMode string `json:"diskMode,omitempty"`
	DiskName string `json:"diskName,omitempty"`
	DiskType string `json:"diskType,omitempty"`
	DiskSize *int   `json:"diskSize,omitempty"`
}

type EcsBatchCreateInstancesLabelListRealRequest struct {
	LabelKey   string `json:"labelKey,omitempty"`
	LabelValue string `json:"labelValue,omitempty"`
}

type EcsBatchCreateInstancesRealRequest struct {
	ClientToken     string                                              `json:"clientToken,omitempty"`
	RegionID        string                                              `json:"regionID,omitempty"`
	AzName          string                                              `json:"azName,omitempty"`
	InstanceName    string                                              `json:"instanceName,omitempty"`
	DisplayName     string                                              `json:"displayName,omitempty"`
	FlavorID        string                                              `json:"flavorID,omitempty"`
	ImageType       *int                                                `json:"imageType,omitempty"`
	ImageID         string                                              `json:"imageID,omitempty"`
	BootDiskType    string                                              `json:"bootDiskType,omitempty"`
	BootDiskSize    *int                                                `json:"bootDiskSize,omitempty"`
	VpcID           string                                              `json:"vpcID,omitempty"`
	OnDemand        *bool                                               `json:"onDemand,omitempty"`
	NetworkCardList []EcsBatchCreateInstancesNetworkCardListRealRequest `json:"networkCardList,omitempty"`
	ExtIP           string                                              `json:"extIP,omitempty"`
	ProjectID       string                                              `json:"projectID,omitempty"`
	SecGroupList    []string                                            `json:"secGroupList,omitempty"`
	DataDiskList    []EcsBatchCreateInstancesDataDiskListRealRequest    `json:"dataDiskList,omitempty"`
	IpVersion       string                                              `json:"ipVersion,omitempty"`
	Bandwidth       *int                                                `json:"bandwidth,omitempty"`
	Ipv6AddressID   string                                              `json:"ipv6AddressID,omitempty"`
	EipID           string                                              `json:"eipID,omitempty"`
	AffinityGroupID string                                              `json:"affinityGroupID,omitempty"`
	KeyPairID       string                                              `json:"keyPairID,omitempty"`
	UserPassword    string                                              `json:"userPassword,omitempty"`
	CycleCount      *int                                                `json:"cycleCount,omitempty"`
	CycleType       string                                              `json:"cycleType,omitempty"`
	AutoRenewStatus *int                                                `json:"autoRenewStatus,omitempty"`
	UserData        string                                              `json:"userData,omitempty"`
	PayVoucherPrice *float64                                            `json:"payVoucherPrice,omitempty"`
	LabelList       []EcsBatchCreateInstancesLabelListRealRequest       `json:"labelList,omitempty"`
	GpuDriverKits   string                                              `json:"gpuDriverKits,omitempty"`
	MonitorService  *bool                                               `json:"monitorService,omitempty"`
	OrderCount      *int                                                `json:"orderCount,omitempty"`
}

type EcsBatchCreateInstancesRealResponse struct {
	RegionID         string `json:"regionID"`
	MasterOrderID    string `json:"masterOrderID"`
	MasterResourceID string `json:"masterResourceID"`
	MasterOrderNO    string `json:"masterOrderNO"`
}

type EcsBatchCreateInstancesRequest struct {
	ClientToken     string                                          // 客户端存根
	RegionID        string                                          // 资源池ID
	AzName          string                                          // 可用区名称
	InstanceName    string                                          // 云主机名称，不可以使用已存在的云主机名称
	DisplayName     string                                          // 云主机显示名称
	FlavorID        string                                          // 云主机规格ID
	ImageType       *int                                            // 镜像类型
	ImageID         string                                          // 镜像ID
	BootDiskType    string                                          // 系统盘类型，取值范围： SATA（普通IO）， SAS（高IO）， SSD（超高IO）， SSD-genric（通用型SSD）， FAST-SSD（极速型SSD），您可以查看磁盘类型及性能介绍来了解磁盘类型及其对应性能指标
	BootDiskSize    *int                                            // 系统盘大小单位为GiB，取值范围：[40, 32768]
	VpcID           string                                          // 虚拟私有云ID
	OnDemand        *bool                                           // 购买方式，取值范围： false（按周期）， true（按需）
	NetworkCardList []EcsBatchCreateInstancesNetworkCardListRequest // 网卡信息列表
	ExtIP           string                                          // 是否使用弹性公网IP，取值范围： 0（不使用）， 1（自动分配）， 2（使用已有）
	ProjectID       string                                          // 企业项目ID
	SecGroupList    []string                                        // 安全组ID列表
	DataDiskList    []EcsBatchCreateInstancesDataDiskListRequest    // 数据盘信息列表，注：同一云主机下最多可挂载8块数据盘
	IpVersion       string                                          // 弹性IP版本，取值范围：<br />ipv4（v4地址），<br />ipv6（v6地址），<br />不指定默认为ipv4。注：请先确认该资源池是否支持ipv6（多可用区类资源池暂不支持）
	Bandwidth       *int                                            // 带宽大小，单位为Mbit/s，取值范围：[1, 2000]
	Ipv6AddressID   string                                          // 弹性公网IPv6的ID，注：多可用区类资源池暂不支持；填写该参数时请填写ipVersion为ipv6
	EipID           string                                          // 弹性公网IP的ID
	AffinityGroupID string                                          // 云主机组ID
	KeyPairID       string                                          // 密钥对ID
	UserPassword    string                                          // 用户密码
	CycleCount      *int                                            // 订购时长
	CycleType       string                                          // 订购周期类型
	AutoRenewStatus *int                                            // 是否自动续订，取值范围： 0（不续费）， 1（自动续费）， 注：按月购买，自动续订周期为3个月；按年购买，自动续订周期为1年
	UserData        string                                          // 用户自定义数据，需要以Base64方式编码，Base64编码后的长度限制为1-16384字符
	PayVoucherPrice *float64                                        // 代金券，满足以下规则：两位小数，不足两位自动补0，超过两位小数无效；不可为负数；注：字段为0时表示不使用代金券，默认不使用
	LabelList       []EcsBatchCreateInstancesLabelListRequest       // 标签信息列表，注：单台云主机最多可绑定10个标签；主机创建完成后，云主机变为运行状态，此时标签仍可能未绑定，需等待一段时间（0~10分钟）。
	GpuDriverKits   string                                          // GPU云主机安装驱动的工具包，仅在同时选择NVIDIA显卡、计算加速型、linux公共镜像三个条件下，支持安装驱动
	MonitorService  *bool                                           // 监控参数，支持通过该参数指定云主机在创建后是否开启详细监控，取值范围： false（不开启），true（开启）；注：若指定该参数为true或不指定该参数，云主机内默认开启最新详细监控服务。若指定该参数为false，默认公共镜像不开启最新监控服务；私有镜像使用镜像中保留的监控服务。说明：仅部分资源池支持monitorService参数，
	OrderCount      *int
}

type EcsBatchCreateInstancesNetworkCardListRequest struct {
	NicName  string // 长度2~32，支持拉丁字母、中文、数字、下划线、连字符，中文或英文字母开头，不能以http:或https:开头
	FixedIP  string // 内网IPv4地址，注：不可使用已占用IP
	IsMaster *bool  // 是否主网卡，取值范围： true（表示主网卡）， false（表示扩展网卡） 注：只能含有一个主网卡
	SubnetID string // 子网ID
}

type EcsBatchCreateInstancesDataDiskListRequest struct {
	DiskMode string
	DiskName string
	DiskType string
	DiskSize *int
}

type EcsBatchCreateInstancesLabelListRequest struct {
	LabelKey   string
	LabelValue string
}

type EcsBatchCreateInstancesResponse struct {
	RegionID         string
	MasterOrderID    string
	MasterResourceID string
	MasterOrderNO    string
}
