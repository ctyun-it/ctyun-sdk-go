package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeShowOldApi 查询已经创建好的云硬盘
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=5539&data=87&isNormal=1
type EcsVolumeShowOldApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeShowOldApi(client *ctyunsdk.CtyunClient) *EcsVolumeShowOldApi {
	return &EcsVolumeShowOldApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/volume_show",
		},
	}
}

func (this *EcsVolumeShowOldApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeShowOldRequest) (*EcsVolumeShowOldResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("resourceID", *req.ResourceID).
		AddParam("regionID", *req.RegionID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeShowOldRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var attachments []EcsVolumeShowOldAttachmentsResponse
	for _, res := range realResponse.Attachments {
		attachments = append(attachments, EcsVolumeShowOldAttachmentsResponse{
			InstanceUUID: res.InstanceUUID,
			AttachmentID: res.AttachmentID,
			Device:       res.Device,
		})
	}

	return &EcsVolumeShowOldResponse{
		Name:           realResponse.Name,
		ResourceID:     realResponse.ResourceID,
		DiskSize:       realResponse.DiskSize,
		DiskType:       realResponse.DiskType,
		DiskMode:       realResponse.DiskMode,
		DiskStatus:     realResponse.DiskStatus,
		CreateTime:     realResponse.CreateTime,
		UpdateTime:     realResponse.UpdateTime,
		ExpireTime:     realResponse.ExpireTime,
		IsSystemVolume: realResponse.IsSystemVolume,
		IsPackaged:     realResponse.IsPackaged,
		InstanceName:   realResponse.InstanceName,
		InstanceUUID:   realResponse.InstanceUUID,
		InstanceStatus: realResponse.InstanceStatus,
		MultiAttach:    realResponse.MultiAttach,
		Attachments:    attachments,
		ProjectID:      realResponse.ProjectID,
		IsEncrypt:      realResponse.IsEncrypt,
		KmsUUID:        realResponse.KmsUUID,
		OnDemand:       realResponse.OnDemand,
		CycleType:      realResponse.CycleType,
		CycleCount:     realResponse.CycleCount,
		RegionID:       realResponse.RegionID,
		AzName:         realResponse.AzName,
		DiskFreeze:     realResponse.DiskFreeze,
		ResourceUUID:   realResponse.ResourceUUID,
	}, nil
}

type EcsVolumeShowOldRealRequest struct {
	ResourceID *string `json:"resourceID,omitempty"`
	RegionID   *string `json:"regionID,omitempty"`
}

type EcsVolumeShowOldRequest struct {
	ResourceID *string
	RegionID   *string
}

type EcsVolumeShowOldAttachmentsRealResponse struct {
	InstanceUUID string `json:"instanceUUID,omitempty"`
	AttachmentID string `json:"attachmentID,omitempty"`
	Device       string `json:"device,omitempty"`
}

type EcsVolumeShowOldRealResponse struct {
	Name           string                                    `json:"name,omitempty"`
	ResourceID     string                                    `json:"resourceID,omitempty"`
	DiskSize       int                                       `json:"diskSize,omitempty"`
	DiskType       string                                    `json:"diskType,omitempty"`
	DiskMode       string                                    `json:"diskMode,omitempty"`
	DiskStatus     string                                    `json:"diskStatus,omitempty"`
	CreateTime     int                                       `json:"createTime,omitempty"`
	UpdateTime     int                                       `json:"updateTime,omitempty"`
	ExpireTime     int                                       `json:"expireTime,omitempty"`
	IsSystemVolume bool                                      `json:"isSystemVolume,omitempty"`
	IsPackaged     bool                                      `json:"isPackaged,omitempty"`
	InstanceName   string                                    `json:"instanceName,omitempty"`
	InstanceUUID   string                                    `json:"instanceUUID,omitempty"`
	InstanceStatus string                                    `json:"instanceStatus,omitempty"`
	MultiAttach    bool                                      `json:"multiAttach,omitempty"`
	Attachments    []EcsVolumeShowOldAttachmentsRealResponse `json:"attachments,omitempty"`
	ProjectID      string                                    `json:"projectID,omitempty"`
	IsEncrypt      bool                                      `json:"isEncrypt,omitempty"`
	KmsUUID        string                                    `json:"kmsUUID,omitempty"`
	OnDemand       bool                                      `json:"onDemand,omitempty"`
	CycleType      string                                    `json:"cycleType,omitempty"`
	CycleCount     int                                       `json:"cycleCount,omitempty"`
	RegionID       string                                    `json:"regionID,omitempty"`
	AzName         string                                    `json:"azName,omitempty"`
	DiskFreeze     bool                                      `json:"diskFreeze,omitempty"`
	ResourceUUID   string                                    `json:"resourceUUID,omitempty"`
}

type EcsVolumeShowOldAttachmentsResponse struct {
	InstanceUUID string
	AttachmentID string
	Device       string
}

type EcsVolumeShowOldResponse struct {
	Name           string
	ResourceID     string
	DiskSize       int
	DiskType       string
	DiskMode       string
	DiskStatus     string
	CreateTime     int
	UpdateTime     int
	ExpireTime     int
	IsSystemVolume bool
	IsPackaged     bool
	InstanceName   string
	InstanceUUID   string
	InstanceStatus string
	MultiAttach    bool
	Attachments    []EcsVolumeShowOldAttachmentsResponse
	ProjectID      string
	IsEncrypt      bool
	KmsUUID        string
	OnDemand       bool
	CycleType      string
	CycleCount     int
	RegionID       string
	AzName         string
	DiskFreeze     bool
	ResourceUUID   string
}
