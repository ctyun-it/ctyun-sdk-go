package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmCreateInstanceApi 创建物理机v4plus
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=6942&data=97&isNormal=1
type EbmCreateInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmCreateInstanceApi(client *ctyunsdk.CtyunClient) *EbmCreateInstanceApi {
	return &EbmCreateInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/create-instance",
		},
	}
}

func (this *EbmCreateInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmCreateInstanceRequest) (*EbmCreateInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	var networkCardList []EbmCreateInstanceNetworkCardListRealRequest
	for _, request := range req.NetworkCardList {
		networkCardList = append(networkCardList, EbmCreateInstanceNetworkCardListRealRequest{
			Title:    request.Title,
			FixedIP:  request.FixedIP,
			Master:   request.Master,
			Ipv6:     request.Ipv6,
			SubnetID: request.SubnetID,
		})
	}

	var diskList []EbmCreateInstanceDiskListRealRequest
	for _, request := range req.DiskList {
		diskList = append(diskList, EbmCreateInstanceDiskListRealRequest{
			DiskType: request.DiskType,
			Title:    request.Title,
			Type:     request.Type,
			Size:     request.Size,
		})
	}

	_, err := builder.WriteJson(&EbmCreateInstanceRealRequest{
		RegionID:             req.RegionID,
		AzName:               req.AzName,
		DeviceType:           req.DeviceType,
		InstanceName:         req.InstanceName,
		Hostname:             req.Hostname,
		ImageUUID:            req.ImageUUID,
		Password:             req.Password,
		VpcID:                req.VpcID,
		ExtIP:                req.ExtIP,
		ProjectID:            req.ProjectID,
		IpType:               req.IpType,
		BandWidth:            req.BandWidth,
		PublicIP:             req.PublicIP,
		SystemVolumeRaidUUID: req.SystemVolumeRaidUUID,
		DataVolumeRaidUUID:   req.DataVolumeRaidUUID,
		SecurityGroupID:      req.SecurityGroupID,
		NetworkCardList:      networkCardList,
		DiskList:             diskList,
		UserData:             req.UserData,
		KeyName:              req.KeyName,
		PayVoucherPrice:      req.PayVoucherPrice,
		AutoRenewStatus:      req.AutoRenewStatus,
		InstanceChargeType:   req.InstanceChargeType,
		CycleCount:           req.CycleCount,
		CycleType:            req.CycleType,
		OrderCount:           req.OrderCount,
		ClientToken:          req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmCreateInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EbmCreateInstanceResponse{
		RegionID:      realResponse.RegionID,
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
	}, nil
}

type EbmCreateInstanceRealRequest struct {
	RegionID             string                                        `json:"regionID,omitempty"`
	AzName               string                                        `json:"azName,omitempty"`
	DeviceType           string                                        `json:"deviceType,omitempty"`
	InstanceName         string                                        `json:"instanceName,omitempty"`
	Hostname             string                                        `json:"hostname,omitempty"`
	ImageUUID            string                                        `json:"imageUUID,omitempty"`
	Password             string                                        `json:"password,omitempty"`
	SystemVolumeRaidUUID string                                        `json:"systemVolumeRaidUUID,omitempty"`
	DataVolumeRaidUUID   string                                        `json:"dataVolumeRaidUUID,omitempty"`
	VpcID                string                                        `json:"vpcID,omitempty"`
	ExtIP                string                                        `json:"extIP,omitempty"`
	ProjectID            string                                        `json:"projectID,omitempty"`
	IpType               string                                        `json:"ipType,omitempty"`
	BandWidth            *int                                          `json:"bandWidth,omitempty"`
	PublicIP             string                                        `json:"publicIP,omitempty"`
	SecurityGroupID      string                                        `json:"securityGroupID,omitempty"`
	DiskList             []EbmCreateInstanceDiskListRealRequest        `json:"diskList,omitempty"`
	NetworkCardList      []EbmCreateInstanceNetworkCardListRealRequest `json:"networkCardList,omitempty"`
	UserData             string                                        `json:"userData,omitempty"`
	KeyName              string                                        `json:"keyName,omitempty"`
	PayVoucherPrice      *float64                                      `json:"payVoucherPrice,omitempty"`
	AutoRenewStatus      *int                                          `json:"autoRenewStatus,omitempty"`
	InstanceChargeType   string                                        `json:"instanceChargeType,omitempty"`
	CycleCount           *int                                          `json:"cycleCount,omitempty"`
	CycleType            string                                        `json:"cycleType,omitempty"`
	OrderCount           *int                                          `json:"orderCount,omitempty"`
	ClientToken          string                                        `json:"clientToken,omitempty"`
}

type EbmCreateInstanceDiskListRealRequest struct {
	DiskType string `json:"diskType,omitempty"`
	Title    string `json:"title,omitempty"`
	Type     string `json:"type,omitempty"`
	Size     *int   `json:"size,omitempty"`
}

type EbmCreateInstanceNetworkCardListRealRequest struct {
	Title    string `json:"title,omitempty"`
	FixedIP  string `json:"fixedIP,omitempty"`
	Master   *bool  `json:"master,omitempty"`
	Ipv6     string `json:"ipv6,omitempty"`
	SubnetID string `json:"subnetID,omitempty"`
}

type EbmCreateInstanceRealResponse struct {
	RegionID      string `json:"regionID"`
	MasterOrderID string `json:"masterOrderID"`
	MasterOrderNO string `json:"masterOrderNO"`
}

type EbmCreateInstanceRequest struct {
	RegionID             string
	AzName               string
	DeviceType           string
	InstanceName         string
	Hostname             string
	ImageUUID            string
	Password             string
	SystemVolumeRaidUUID string
	DataVolumeRaidUUID   string
	VpcID                string
	ExtIP                string
	ProjectID            string
	IpType               string
	BandWidth            *int
	PublicIP             string
	SecurityGroupID      string
	DiskList             []EbmCreateInstanceDiskListRequest
	NetworkCardList      []EbmCreateInstanceNetworkCardListRequest
	UserData             string
	KeyName              string
	PayVoucherPrice      *float64
	AutoRenewStatus      *int
	InstanceChargeType   string
	CycleCount           *int
	CycleType            string
	OrderCount           *int
	ClientToken          string
}

type EbmCreateInstanceDiskListRequest struct {
	DiskType string
	Title    string
	Type     string
	Size     *int
}

type EbmCreateInstanceNetworkCardListRequest struct {
	Title    string
	FixedIP  string
	Master   *bool
	Ipv6     string
	SubnetID string
}

type EbmCreateInstanceResponse struct {
	RegionID      string
	MasterOrderID string
	MasterOrderNO string
}
