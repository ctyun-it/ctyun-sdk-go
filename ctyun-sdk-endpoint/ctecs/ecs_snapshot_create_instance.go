package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotCreateInstanceApi 快照创建一台云主机
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=8132&data=87&isNormal=1
type EcsSnapshotCreateInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotCreateInstanceApi(client *ctyunsdk.CtyunClient) *EcsSnapshotCreateInstanceApi {
	return &EcsSnapshotCreateInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/create-instance",
		},
	}
}

func (this *EcsSnapshotCreateInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotCreateInstanceRequest) (*EcsSnapshotCreateInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var labelList []EcsSnapshotCreateInstanceLabelListRealRequest
	for _, request := range req.LabelList {
		labelList = append(labelList, EcsSnapshotCreateInstanceLabelListRealRequest{
			LabelKey:   request.LabelKey,
			LabelValue: request.LabelValue,
		})
	}

	var networkCardList []EcsSnapshotCreateInstanceNetworkCardListRealRequest
	for _, request := range req.NetworkCardList {
		networkCardList = append(networkCardList, EcsSnapshotCreateInstanceNetworkCardListRealRequest{
			NicName:  request.NicName,
			FixedIP:  request.FixedIP,
			IsMaster: request.IsMaster,
			SubnetID: request.SubnetID,
		})
	}

	_, err := builder.WriteJson(&EcsSnapshotCreateInstanceRealRequest{
		ClientToken:     req.ClientToken,
		RegionID:        req.RegionID,
		ProjectID:       req.ProjectID,
		InstanceName:    req.InstanceName,
		DisplayName:     req.DisplayName,
		SnapshotID:      req.SnapshotID,
		VpcID:           req.VpcID,
		OnDemand:        req.OnDemand,
		SecGroupList:    req.SecGroupList,
		NetworkCardList: networkCardList,
		ExtIP:           req.ExtIP,
		IpVersion:       req.IpVersion,
		Bandwidth:       req.Bandwidth,
		Ipv6AddressID:   req.Ipv6AddressID,
		EipID:           req.EipID,
		AffinityGroupID: req.AffinityGroupID,
		KeyPairID:       req.KeyPairID,
		UserPassword:    req.UserPassword,
		CycleCount:      req.CycleCount,
		CycleType:       req.CycleType,
		AutoRenewStatus: req.AutoRenewStatus,
		UserData:        req.UserData,
		LabelList:       labelList,
		MonitorService:  req.MonitorService,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotCreateInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotCreateInstanceResponse{
		RegionId:         realResponse.RegionID,
		MasterOrderId:    realResponse.MasterOrderID,
		MasterResourceId: realResponse.MasterResourceID,
		MasterOrderNo:    realResponse.MasterOrderNO,
	}, nil
}

type EcsSnapshotCreateInstanceLabelListRealRequest struct {
	LabelKey   *string `json:"labelKey,omitempty"`
	LabelValue *string `json:"labelValue,omitempty"`
}

type EcsSnapshotCreateInstanceNetworkCardListRealRequest struct {
	NicName  *string `json:"nicName,omitempty"`
	FixedIP  *string `json:"fixedIP,omitempty"`
	IsMaster *bool   `json:"isMaster,omitempty"`
	SubnetID *string `json:"subnetID,omitempty"`
}

type EcsSnapshotCreateInstanceRealRequest struct {
	ClientToken     *string                                               `json:"clientToken,omitempty"`
	RegionID        *string                                               `json:"regionID,omitempty"`
	ProjectID       *string                                               `json:"projectID,omitempty"`
	InstanceName    *string                                               `json:"instanceName,omitempty"`
	DisplayName     *string                                               `json:"displayName,omitempty"`
	SnapshotID      *string                                               `json:"snapshotID,omitempty"`
	VpcID           *string                                               `json:"vpcID,omitempty"`
	OnDemand        *bool                                                 `json:"onDemand,omitempty"`
	SecGroupList    *[]string                                             `json:"secGroupList,omitempty"`
	NetworkCardList []EcsSnapshotCreateInstanceNetworkCardListRealRequest `json:"networkCardList,omitempty"`
	ExtIP           *string                                               `json:"extIP,omitempty"`
	IpVersion       *string                                               `json:"ipVersion,omitempty"`
	Bandwidth       *int                                                  `json:"bandwidth,omitempty"`
	Ipv6AddressID   *string                                               `json:"ipv6AddressID,omitempty"`
	EipID           *string                                               `json:"eipID,omitempty"`
	AffinityGroupID *string                                               `json:"affinityGroupID,omitempty"`
	KeyPairID       *string                                               `json:"keyPairID,omitempty"`
	UserPassword    *string                                               `json:"userPassword,omitempty"`
	CycleCount      *int                                                  `json:"cycleCount,omitempty"`
	CycleType       *string                                               `json:"cycleType,omitempty"`
	AutoRenewStatus *int                                                  `json:"autoRenewStatus,omitempty"`
	UserData        *string                                               `json:"userData,omitempty"`
	LabelList       []EcsSnapshotCreateInstanceLabelListRealRequest       `json:"labelList,omitempty"`
	MonitorService  *bool                                                 `json:"monitorService,omitempty"`
}

type EcsSnapshotCreateInstanceLabelListRequest struct {
	LabelKey   *string
	LabelValue *string
}

type EcsSnapshotCreateInstanceNetworkCardListRequest struct {
	NicName  *string
	FixedIP  *string
	IsMaster *bool
	SubnetID *string
}

type EcsSnapshotCreateInstanceRequest struct {
	ClientToken     *string
	RegionID        *string
	ProjectID       *string
	InstanceName    *string
	DisplayName     *string
	SnapshotID      *string
	VpcID           *string
	OnDemand        *bool
	SecGroupList    *[]string
	NetworkCardList []EcsSnapshotCreateInstanceNetworkCardListRequest
	ExtIP           *string
	IpVersion       *string
	Bandwidth       *int
	Ipv6AddressID   *string
	EipID           *string
	AffinityGroupID *string
	KeyPairID       *string
	UserPassword    *string
	CycleCount      *int
	CycleType       *string
	AutoRenewStatus *int
	UserData        *string
	LabelList       []EcsSnapshotCreateInstanceLabelListRequest
	MonitorService  *bool
}

type EcsSnapshotCreateInstanceRealResponse struct {
	RegionID         string `json:"regionID"`
	MasterOrderID    string `json:"masterOrderID"`
	MasterResourceID string `json:"masterResourceID"`
	MasterOrderNO    string `json:"masterOrderNO"`
}

type EcsSnapshotCreateInstanceResponse struct {
	RegionId         string
	MasterOrderId    string
	MasterResourceId string
	MasterOrderNo    string
}
