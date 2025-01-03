package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeCreateApi 创建一块按量付费或包年包月云硬盘
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=11960&data=87&isNormal=1
type EcsVolumeCreateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeCreateApi(client *ctyunsdk.CtyunClient) *EcsVolumeCreateApi {
	return &EcsVolumeCreateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/volume/create",
		},
	}
}

func (this *EcsVolumeCreateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeCreateRequest) (*EcsVolumeCreateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVolumeCreateRealRequest{
		RegionID:    req.RegionID,
		DiskMode:    req.DiskMode,
		DiskType:    req.DiskType,
		DiskName:    req.DiskName,
		DiskSize:    req.DiskSize,
		ClientToken: req.ClientToken,
		AzName:      req.AzName,
		MultiAttach: req.MultiAttach,
		OnDemand:    req.OnDemand,
		CycleType:   req.CycleType,
		CycleCount:  req.CycleCount,
		IsEncrypt:   req.IsEncrypt,
		KmsUUID:     req.KmsUUID,
		ProjectID:   req.ProjectID,
		ImageID:     req.ImageID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeCreateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVolumeCreateResponse{
		MasterOrderID:        realResponse.MasterOrderID,
		MasterOrderNO:        realResponse.MasterOrderNO,
		MasterResourceID:     realResponse.MasterResourceID,
		MasterResourceStatus: realResponse.MasterResourceStatus,
		RegionID:             realResponse.RegionID,
		Resources: EcsVolumeCreateResourcesResponse{
			DiskID:       realResponse.Resources.DiskID,
			OrderID:      realResponse.Resources.OrderID,
			StartTime:    realResponse.Resources.StartTime,
			CreateTime:   realResponse.Resources.CreateTime,
			UpdateTime:   realResponse.Resources.UpdateTime,
			Status:       realResponse.Resources.Status,
			IsMaster:     realResponse.Resources.IsMaster,
			ItemValue:    realResponse.Resources.ItemValue,
			ResourceType: realResponse.Resources.ResourceType,
			DiskName:     realResponse.Resources.DiskName,
		},
	}, nil
}

type EcsVolumeCreateRealRequest struct {
	RegionID    *string `json:"regionID,omitempty"`
	DiskMode    *string `json:"diskMode,omitempty"`
	DiskType    *string `json:"diskType,omitempty"`
	DiskName    *string `json:"diskName,omitempty"`
	DiskSize    *int    `json:"diskSize,omitempty"`
	ClientToken *string `json:"clientToken,omitempty"`
	AzName      *string `json:"azName,omitempty"`
	MultiAttach *bool   `json:"multiAttach,omitempty"`
	OnDemand    *bool   `json:"onDemand,omitempty"`
	CycleType   *string `json:"cycleType,omitempty"`
	CycleCount  *int    `json:"cycleCount,omitempty"`
	IsEncrypt   *bool   `json:"isEncrypt,omitempty"`
	KmsUUID     *string `json:"kmsUUID,omitempty"`
	ProjectID   *string `json:"projectID,omitempty"`
	ImageID     *string `json:"imageID,omitempty"`
}

type EcsVolumeCreateRequest struct {
	RegionID    *string
	DiskMode    *string
	DiskType    *string
	DiskName    *string
	DiskSize    *int
	ClientToken *string
	AzName      *string
	MultiAttach *bool
	OnDemand    *bool
	CycleType   *string
	CycleCount  *int
	IsEncrypt   *bool
	KmsUUID     *string
	ProjectID   *string
	ImageID     *string
}

type EcsVolumeCreateResourcesRealResponse struct {
	DiskID       string `json:"diskID,omitempty"`
	OrderID      string `json:"orderID,omitempty"`
	StartTime    int    `json:"startTime,omitempty"`
	CreateTime   int    `json:"createTime,omitempty"`
	UpdateTime   int    `json:"updateTime,omitempty"`
	Status       int    `json:"status,omitempty"`
	IsMaster     bool   `json:"isMaster,omitempty"`
	ItemValue    int    `json:"itemValue,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	DiskName     string `json:"diskName,omitempty"`
}

type EcsVolumeCreateRealResponse struct {
	MasterOrderID        string                               `json:"masterOrderID,omitempty"`
	MasterOrderNO        string                               `json:"masterOrderNO,omitempty"`
	MasterResourceID     string                               `json:"masterResourceID,omitempty"`
	MasterResourceStatus string                               `json:"masterResourceStatus,omitempty"`
	RegionID             string                               `json:"regionID,omitempty"`
	Resources            EcsVolumeCreateResourcesRealResponse `json:"resources,omitempty"`
}

type EcsVolumeCreateResourcesResponse struct {
	DiskID       string
	OrderID      string
	StartTime    int
	CreateTime   int
	UpdateTime   int
	Status       int
	IsMaster     bool
	ItemValue    int
	ResourceType string
	DiskName     string
}

type EcsVolumeCreateResponse struct {
	MasterOrderID        string
	MasterOrderNO        string
	MasterResourceID     string
	MasterResourceStatus string
	RegionID             string
	Resources            EcsVolumeCreateResourcesResponse
}
