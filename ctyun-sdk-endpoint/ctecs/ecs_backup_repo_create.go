package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupRepoCreateApi 创建云主机备份存储库
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=6910&data=87&isNormal=1
type EcsBackupRepoCreateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupRepoCreateApi(client *ctyunsdk.CtyunClient) *EcsBackupRepoCreateApi {
	return &EcsBackupRepoCreateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-repo/create",
		},
	}
}

func (this *EcsBackupRepoCreateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupRepoCreateRequest) (*EcsBackupRepoCreateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupRepoCreateRealRequest{
		RegionID:        req.RegionID,
		ProjectID:       req.ProjectID,
		RepositoryName:  req.RepositoryName,
		CycleCount:      req.CycleCount,
		CycleType:       req.CycleType,
		ClientToken:     req.ClientToken,
		Size:            req.Size,
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

	var realResponse EcsBackupRepoCreateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupRepoCreateResponse{
		MasterOrderID:    realResponse.MasterOrderID,
		MasterOrderNO:    realResponse.MasterOrderNO,
		RegionID:         realResponse.RegionID,
		MasterResourceID: realResponse.MasterResourceID,
	}, nil
}

type EcsBackupRepoCreateRealRequest struct {
	RegionID        *string  `json:"regionID,omitempty"`
	ProjectID       *string  `json:"projectID,omitempty"`
	RepositoryName  *string  `json:"repositoryName,omitempty"`
	CycleCount      *int     `json:"cycleCount,omitempty"`
	CycleType       *string  `json:"cycleType,omitempty"`
	ClientToken     *string  `json:"clientToken,omitempty"`
	Size            *int     `json:"size,omitempty"`
	AutoRenewStatus *int     `json:"autoRenewStatus,omitempty"`
	PayVoucherPrice *float64 `json:"payVoucherPrice,omitempty"`
}

type EcsBackupRepoCreateRequest struct {
	RegionID        *string
	ProjectID       *string
	RepositoryName  *string
	CycleCount      *int
	CycleType       *string
	ClientToken     *string
	Size            *int
	AutoRenewStatus *int
	PayVoucherPrice *float64
}

type EcsBackupRepoCreateRealResponse struct {
	MasterOrderID    string `json:"masterOrderID,omitempty"`
	MasterOrderNO    string `json:"masterOrderNO,omitempty"`
	RegionID         string `json:"regionID,omitempty"`
	MasterResourceID string `json:"masterResourceID,omitempty"`
}

type EcsBackupRepoCreateResponse struct {
	MasterOrderID    string
	MasterOrderNO    string
	RegionID         string
	MasterResourceID string
}
