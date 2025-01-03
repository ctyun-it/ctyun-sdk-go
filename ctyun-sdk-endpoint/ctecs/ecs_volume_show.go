package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeShowApi 查询云硬盘
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=12225&data=87&isNormal=1
type EcsVolumeShowApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeShowApi(client *ctyunsdk.CtyunClient) *EcsVolumeShowApi {
	return &EcsVolumeShowApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/volume/show",
		},
	}
}

func (this *EcsVolumeShowApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeShowRequest) (*EcsVolumeShowResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("diskID", *req.DiskID).
		AddParam("regionID", *req.RegionID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeShowRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var attachments []EcsVolumeShowAttachmentsResponse
	for _, res := range realResponse.Attachments {
		attachments = append(attachments, EcsVolumeShowAttachmentsResponse{
			InstanceID:   res.InstanceID,
			AttachmentID: res.AttachmentID,
			Device:       res.Device,
		})
	}

	return &EcsVolumeShowResponse{
		DiskName:       realResponse.DiskName,
		DiskID:         realResponse.DiskID,
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
		InstanceID:     realResponse.InstanceID,
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
	}, nil
}

type EcsVolumeShowRealRequest struct {
	DiskID   *string `json:"diskID,omitempty"`
	RegionID *string `json:"regionID,omitempty"`
}

type EcsVolumeShowRequest struct {
	DiskID   *string
	RegionID *string
}

type EcsVolumeShowAttachmentsRealResponse struct {
	InstanceID   string `json:"instanceID,omitempty"`
	AttachmentID string `json:"attachmentID,omitempty"`
	Device       string `json:"device,omitempty"`
}

type EcsVolumeShowRealResponse struct {
	DiskName       string                                 `json:"diskName,omitempty"`
	DiskID         string                                 `json:"diskID,omitempty"`
	DiskSize       int                                    `json:"diskSize,omitempty"`
	DiskType       string                                 `json:"diskType,omitempty"`
	DiskMode       string                                 `json:"diskMode,omitempty"`
	DiskStatus     string                                 `json:"diskStatus,omitempty"`
	CreateTime     int                                    `json:"createTime,omitempty"`
	UpdateTime     int                                    `json:"updateTime,omitempty"`
	ExpireTime     int                                    `json:"expireTime,omitempty"`
	IsSystemVolume bool                                   `json:"isSystemVolume,omitempty"`
	IsPackaged     bool                                   `json:"isPackaged,omitempty"`
	InstanceName   string                                 `json:"instanceName,omitempty"`
	InstanceID     string                                 `json:"instanceID,omitempty"`
	InstanceStatus string                                 `json:"instanceStatus,omitempty"`
	MultiAttach    bool                                   `json:"multiAttach,omitempty"`
	Attachments    []EcsVolumeShowAttachmentsRealResponse `json:"attachments,omitempty"`
	ProjectID      string                                 `json:"projectID,omitempty"`
	IsEncrypt      bool                                   `json:"isEncrypt,omitempty"`
	KmsUUID        string                                 `json:"kmsUUID,omitempty"`
	OnDemand       bool                                   `json:"onDemand,omitempty"`
	CycleType      string                                 `json:"cycleType,omitempty"`
	CycleCount     int                                    `json:"cycleCount,omitempty"`
	RegionID       string                                 `json:"regionID,omitempty"`
	AzName         string                                 `json:"azName,omitempty"`
}

type EcsVolumeShowAttachmentsResponse struct {
	InstanceID   string
	AttachmentID string
	Device       string
}

type EcsVolumeShowResponse struct {
	DiskName       string
	DiskID         string
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
	InstanceID     string
	InstanceStatus string
	MultiAttach    bool
	Attachments    []EcsVolumeShowAttachmentsResponse
	ProjectID      string
	IsEncrypt      bool
	KmsUUID        string
	OnDemand       bool
	CycleType      string
	CycleCount     int
	RegionID       string
	AzName         string
}
