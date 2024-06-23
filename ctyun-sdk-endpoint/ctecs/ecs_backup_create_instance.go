package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupCreateInstanceApi 备份创建一台按量付费或包年包月的云主机
type EcsBackupCreateInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupCreateInstanceApi(client *ctyunsdk.CtyunClient) *EcsBackupCreateInstanceApi {
	return &EcsBackupCreateInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup/create-instance",
		},
	}
}

func (this *EcsBackupCreateInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupCreateInstanceRequest) (*EcsBackupCreateInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	var networkCardList []ecsBackupCreateInstanceNetworkCardListRealRequest
	for _, request := range req.NetworkCardList {
		networkCardList = append(networkCardList, ecsBackupCreateInstanceNetworkCardListRealRequest{
			NicName:  request.NicName,
			FixedIP:  request.FixedIP,
			IsMaster: request.IsMaster,
			SubnetID: request.SubnetID,
		})
	}

	var labelList []ecsBackupCreateInstanceLabelListRealRequest
	for _, request := range req.LabelList {
		labelList = append(labelList, ecsBackupCreateInstanceLabelListRealRequest{
			LabelKey:   request.LabelKey,
			LabelValue: request.LabelValue,
		})
	}

	_, err := builder.WriteJson(&ecsBackupCreateInstanceRealRequest{
		ClientToken:      req.ClientToken,
		RegionID:         req.RegionID,
		AzName:           req.AzName,
		InstanceName:     req.InstanceName,
		DisplayName:      req.DisplayName,
		InstanceBackupID: req.InstanceBackupID,
		FlavorID:         req.FlavorID,
		VpcID:            req.VpcID,
		OnDemand:         req.OnDemand,
		SecGroupList:     req.SecGroupList,
		NetworkCardList:  networkCardList,
		ExtIP:            req.ExtIP,
		IpVersion:        req.IpVersion,
		Bandwidth:        req.Bandwidth,
		EipID:            req.EipID,
		AffinityGroupID:  req.AffinityGroupID,
		KeyPairID:        req.KeyPairID,
		UserPassword:     req.UserPassword,
		CycleCount:       req.CycleCount,
		CycleType:        req.CycleType,
		AutoRenewStatus:  req.AutoRenewStatus,
		UserData:         req.UserData,
		ProjectID:        req.ProjectID,
		PayVoucherPrice:  req.PayVoucherPrice,
		LabelList:        labelList,
		MonitorService:   req.MonitorService,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse ecsBackupCreateInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)

	if err != nil {
		return nil, err
	}

	return &EcsBackupCreateInstanceResponse{
		RegionID:         realResponse.RegionID,
		MasterOrderID:    realResponse.MasterOrderID,
		MasterResourceID: realResponse.MasterResourceID,
		MasterOrderNo:    realResponse.MasterOrderNO,
	}, nil
}

type ecsBackupCreateInstanceNetworkCardListRealRequest struct {
	NicName  string `json:"nicName,omitempty"`
	FixedIP  string `json:"fixedIP,omitempty"`
	IsMaster bool   `json:"isMaster,omitempty"`
	SubnetID string `json:"subnetID,omitempty"`
}

type ecsBackupCreateInstanceLabelListRealRequest struct {
	LabelKey   string `json:"labelKey,omitempty"`
	LabelValue string `json:"labelValue,omitempty"`
}

type ecsBackupCreateInstanceRealRequest struct {
	ClientToken      string                                              `json:"clientToken,omitempty"`
	RegionID         string                                              `json:"regionID,omitempty"`
	AzName           string                                              `json:"azName,omitempty"`
	InstanceName     string                                              `json:"instanceName,omitempty"`
	DisplayName      string                                              `json:"displayName,omitempty"`
	InstanceBackupID string                                              `json:"instanceBackupID,omitempty"`
	FlavorID         string                                              `json:"flavorID,omitempty"`
	VpcID            string                                              `json:"vpcID,omitempty"`
	OnDemand         bool                                                `json:"onDemand"`
	SecGroupList     []string                                            `json:"secGroupList,omitempty"`
	NetworkCardList  []ecsBackupCreateInstanceNetworkCardListRealRequest `json:"networkCardList,omitempty"`
	ExtIP            string                                              `json:"extIP,omitempty"`
	IpVersion        string                                              `json:"ipVersion,omitempty"`
	Bandwidth        int                                                 `json:"bandwidth"`
	Ipv6AddressID    string                                              `json:"ipv6AddressID,omitempty"`
	EipID            string                                              `json:"eipID,omitempty"`
	AffinityGroupID  string                                              `json:"affinityGroupID,omitempty"`
	KeyPairID        string                                              `json:"keyPairID,omitempty"`
	UserPassword     string                                              `json:"userPassword,omitempty"`
	CycleCount       int                                                 `json:"cycleCount,omitempty"`
	CycleType        string                                              `json:"cycleType,omitempty"`
	AutoRenewStatus  int                                                 `json:"autoRenewStatus,omitempty"`
	UserData         string                                              `json:"userData,omitempty"`
	ProjectID        string                                              `json:"projectID,omitempty"`
	PayVoucherPrice  float64                                             `json:"payVoucherPrice,omitempty"`
	LabelList        []ecsBackupCreateInstanceLabelListRealRequest       `json:"labelList,omitempty"`
	MonitorService   bool                                                `json:"monitorService,omitempty"`
}

type ecsBackupCreateInstanceRealResponse struct {
	RegionID         string `json:"regionID"`
	MasterOrderID    string `json:"masterOrderID"`
	MasterResourceID string `json:"masterResourceID"`
	MasterOrderNO    string `json:"masterOrderNO"`
}

type EcsBackupCreateInstanceRequest struct {
	ClientToken      string                                          // 客户端存根
	RegionID         string                                          // 资源池ID
	AzName           string                                          // 可用区名称
	InstanceName     string                                          // 云主机名称，不可以使用已存在的云主机名称
	DisplayName      string                                          // 云主机显示名称
	InstanceBackupID string                                          // 云主机备份ID
	FlavorID         string                                          // 云主机规格ID
	VpcID            string                                          // 虚拟私有云ID
	OnDemand         bool                                            // 购买方式，取值范围： false（按周期）， true（按需）
	SecGroupList     []string                                        // 安全组ID列表
	NetworkCardList  []EcsBackupCreateInstanceNetworkCardListRequest // 网卡信息列表
	ExtIP            string                                          // 是否使用弹性公网IP，取值范围： 0（不使用）， 1（自动分配）， 2（使用已有）
	IpVersion        string                                          // 弹性IP版本，取值范围：<br />ipv4（v4地址），<br />ipv6（v6地址），<br />不指定默认为ipv4。注：请先确认该资源池是否支持ipv6（多可用区类资源池暂不支持）
	Bandwidth        int                                             // 带宽大小，单位为Mbit/s，取值范围：[1, 2000]
	Ipv6AddressID    string                                          // 弹性公网IPv6的ID，注：多可用区类资源池暂不支持；填写该参数时请填写ipVersion为ipv6
	EipID            string                                          // 弹性公网IP的ID
	AffinityGroupID  string                                          // 云主机组ID
	KeyPairID        string                                          // 密钥对ID
	UserPassword     string                                          // 用户密码
	CycleCount       int                                             // 订购时长
	CycleType        string                                          // 订购周期类型
	AutoRenewStatus  int                                             // 是否自动续订，取值范围： 0（不续费）， 1（自动续费）， 注：按月购买，自动续订周期为3个月；按年购买，自动续订周期为1年
	UserData         string                                          // 用户自定义数据，需要以Base64方式编码，Base64编码后的长度限制为1-16384字符
	ProjectID        string                                          // 企业项目ID
	PayVoucherPrice  float64                                         // 代金券，满足以下规则：两位小数，不足两位自动补0，超过两位小数无效；不可为负数；注：字段为0时表示不使用代金券，默认不使用
	LabelList        []EcsBackupCreateInstanceLabelListRequest       // 标签信息列表，注：单台云主机最多可绑定10个标签；主机创建完成后，云主机变为运行状态，此时标签仍可能未绑定，需等待一段时间（0~10分钟）。
	MonitorService   bool                                            // 监控参数，支持通过该参数指定云主机在创建后是否开启详细监控，取值范围： false（不开启），true（开启）；注：若指定该参数为true或不指定该参数，云主机内默认开启最新详细监控服务。若指定该参数为false，默认公共镜像不开启最新监控服务；私有镜像使用镜像中保留的监控服务。说明：仅部分资源池支持monitorService参数，
}

type EcsBackupCreateInstanceNetworkCardListRequest struct {
	NicName  string // 长度2~32，支持拉丁字母、中文、数字、下划线、连字符，中文或英文字母开头，不能以http:或https:开头
	FixedIP  string // 内网IPv4地址，注：不可使用已占用IP
	IsMaster bool   // 是否主网卡，取值范围： true（表示主网卡）， false（表示扩展网卡） 注：只能含有一个主网卡
	SubnetID string // 子网ID
}

type EcsBackupCreateInstanceDataDiskListRequest struct {
	DiskMode string
	DiskName string
	DiskType string
	DiskSize int
}

type EcsBackupCreateInstanceLabelListRequest struct {
	LabelKey   string
	LabelValue string
}

type EcsBackupCreateInstanceResponse struct {
	RegionID         string
	MasterOrderID    string
	MasterResourceID string
	MasterOrderNo    string
}
