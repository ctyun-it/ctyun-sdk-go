package ctebs

import (
	"context"
	"net/http"

	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
)

// EbsListApi 云硬盘列表查询
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=35&api=7338&data=48&isNormal=1
type EbsListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbsListApi(client *ctyunsdk.CtyunClient) *EbsListApi {
	return &EbsListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ebs/list-ebs",
		},
	}
}

func (this *EbsListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbsListRequest) (*EbsListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	_, err := builder.WriteJson(&ebsListRealRequest{
		RegionID: req.RegionID,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtebs, builder)
	if err != nil {
		return nil, err
	}

	result := &ebsListRealResponse{}

	err = response.ParseByStandardModelWithCheck(result)
	if err != nil {
		return nil, err
	}

	var disklists []EbsDiskList
	for _, disklist := range result.DiskList {

		var attach []EbsListAttachment
		for _, realAttachment := range disklist.Attachments {
			attach = append(attach, EbsListAttachment{
				InstanceId:   realAttachment.InstanceID,
				AttachmentId: realAttachment.AttachmentID,
				Device:       realAttachment.Device,
			})
		}

		disklists = append(disklists, EbsDiskList{
			DiskName:        disklist.DiskName,
			DiskID:          disklist.DiskID,
			DiskSize:        disklist.DiskSize,
			DiskType:        disklist.DiskType,
			DiskMode:        disklist.DiskMode,
			DiskStatus:      disklist.DiskStatus,
			CreateTime:      disklist.CreateTime,
			UpdateTime:      disklist.UpdateTime,
			ExpireTime:      disklist.ExpireTime,
			IsSystemVolume:  disklist.IsSystemVolume,
			IsPackaged:      disklist.IsPackaged,
			InstanceName:    disklist.InstanceName,
			InstanceID:      disklist.InstanceID,
			InstanceStatus:  disklist.InstanceStatus,
			MultiAttach:     disklist.MultiAttach,
			Attachments:     attach,
			ProjectID:       disklist.ProjectID,
			IsEncrypt:       disklist.IsEncrypt,
			KmsUUID:         disklist.KmsUUID,
			RegionID:        disklist.RegionID,
			AzName:          disklist.AzName,
			DiskFreeze:      disklist.DiskFreeze,
			ProvisionedIops: disklist.ProvisionedIops,
		})
	}

	return &EbsListResponse{
		DiskList:     disklists,
		DiskTotal:    result.DiskTotal,
		CurrentCount: result.CurrentCount,
		TotalCount:   result.TotalCount,
		TotalPage:    result.TotalPage,
	}, nil
}

type ebsListRealResponse struct {
	DiskList     []ebsRealDiskList `json:"diskList"`     // 返回数据集合
	DiskTotal    int64             `json:"diskTotal"`    // 云硬盘总数
	CurrentCount int64             `json:"currentCount"` // 当前页记录数目
	TotalCount   int64             `json:"totalCount"`   // 总记录数
	TotalPage    int64             `json:"totalPage"`    // 总页数
}

type ebsRealDiskList struct {
	DiskName        string                  `json:"diskName"`                  // 磁盘名
	DiskID          string                  `json:"diskID"`                    // 磁盘ID
	DiskSize        int64                   `json:"diskSize"`                  // 磁盘大小（GB）
	DiskType        string                  `json:"diskType"`                  // 磁盘规格类型 SATA/SAS/SSD-genric/SSD/FAST-SSD
	DiskMode        string                  `json:"diskMode"`                  // 磁盘模式。VBD/ISCSI/FCSAN
	DiskStatus      string                  `json:"diskStatus"`                // 云硬盘使用状态 deleting/creating/detaching，具体请参考云硬盘使用状态
	CreateTime      int64                   `json:"createTime"`                // 创建时刻，epoch时戳，精度毫秒
	UpdateTime      int64                   `json:"updateTime"`                // 更新时刻，epoch时戳，精度毫秒
	ExpireTime      int64                   `json:"expireTime"`                // 过期时刻，epoch时戳，精度毫秒
	IsSystemVolume  bool                    `json:"isSystemVolume"`            // 是否系统盘，只有为系统盘时才返回该字段
	IsPackaged      bool                    `json:"isPackaged"`                // 是否是云主机成套资源
	InstanceName    string                  `json:"instanceName"`              // 绑定的云主机名，有挂载时才返回
	InstanceID      string                  `json:"instanceID"`                // 绑定云主机resourceUUID，有挂载时才返回
	InstanceStatus  string                  `json:"instanceStatus"`            // 云主机状态starting/restarting/stopping，具体参考云主机状态，有挂载时才返回
	MultiAttach     bool                    `json:"multiAttach"`               // 是否共享云硬盘
	Attachments     []ebsRealListAttachment `json:"attachments"`               // 挂载信息。如果是共享挂载云硬盘，有多项		参考表attachment
	ProjectID       string                  `json:"projectID,omitempty"`       // 资源所属企业项目id
	IsEncrypt       bool                    `json:"isEncrypt"`                 // 是否加密盘
	KmsUUID         string                  `json:"kmsUUID"`                   // 加密盘密钥UUID，是加密盘时才返回
	RegionID        string                  `json:"regionID"`                  // 资源池ID
	AzName          string                  `json:"azName"`                    // 多可用区下的可用区名字
	DiskFreeze      bool                    `json:"diskFreeze"`                // 是否冻结
	ProvisionedIops string                  `json:"provisionedIops,omitempty"` // XSSD类型盘的预配置iops，未配置返回0，其他类型盘不返回
}

type ebsRealListAttachment struct {
	InstanceID   string `json:"instanceID"`   // 绑定云主机实例UUID
	AttachmentID string `json:"attachmentID"` // 挂载ID
	Device       string `json:"device"`       // 挂载设备名，比如/dev/sda
}

type EbsListRequest struct {
	RegionID string
	PageNo   int64
	PageSize int64
}

type ebsListRealRequest struct {
	RegionID string `json:"regionID"`
	PageNo   int64  `json:"pageNo,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
}

type EbsListResponse struct {
	DiskList     []EbsDiskList // 返回数据集合
	DiskTotal    int64         // 云硬盘总数
	CurrentCount int64         // 当前页记录数目
	TotalCount   int64         // 总记录数
	TotalPage    int64         // 总页数
}

type EbsDiskList struct {
	DiskName        string              // 磁盘名
	DiskID          string              // 磁盘ID
	DiskSize        int64               // 磁盘大小（GB）
	DiskType        string              // 磁盘规格类型 SATA/SAS/SSD-genric/SSD/FAST-SSD
	DiskMode        string              // 磁盘模式。VBD/ISCSI/FCSAN
	DiskStatus      string              // 云硬盘使用状态 deleting/creating/detaching，具体请参考云硬盘使用状态
	CreateTime      int64               // 创建时刻，epoch时戳，精度毫秒
	UpdateTime      int64               // 更新时刻，epoch时戳，精度毫秒
	ExpireTime      int64               // 过期时刻，epoch时戳，精度毫秒
	IsSystemVolume  bool                // 是否系统盘，只有为系统盘时才返回该字段
	IsPackaged      bool                // 是否是云主机成套资源
	InstanceName    string              // 绑定的云主机名，有挂载时才返回
	InstanceID      string              // 绑定云主机resourceUUID，有挂载时才返回
	InstanceStatus  string              // 云主机状态starting/restarting/stopping，具体参考云主机状态，有挂载时才返回
	MultiAttach     bool                // 是否共享云硬盘
	Attachments     []EbsListAttachment // 挂载信息。如果是共享挂载云硬盘，有多项		参考表attachment
	ProjectID       string              // 资源所属企业项目id
	IsEncrypt       bool                // 是否加密盘
	KmsUUID         string              // 加密盘密钥UUID，是加密盘时才返回
	RegionID        string              // 资源池ID
	AzName          string              // 多可用区下的可用区名字
	DiskFreeze      bool                // 是否冻结
	ProvisionedIops string              `json:"ProvisionedIops,omitempty"` // XSSD类型盘的预配置iops，未配置返回0，其他类型盘不返回
}

type EbsListAttachment struct {
	InstanceId   string // 绑定云主机实例UUID
	AttachmentId string // 挂载ID
	Device       string // 挂载设备名，比如/dev/sda
}
