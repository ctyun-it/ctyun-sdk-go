package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupRepoUpgradeApi 扩容云主机备份存储库
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=6912&data=87&isNormal=1
type EcsBackupRepoUpgradeApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupRepoUpgradeApi(client *ctyunsdk.CtyunClient) *EcsBackupRepoUpgradeApi {
	return &EcsBackupRepoUpgradeApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-repo/upgrade",
		},
	}
}

func (this *EcsBackupRepoUpgradeApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupRepoUpgradeRequest) (*EcsBackupRepoUpgradeResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupRepoUpgradeRealRequest{
		RegionID:        req.RegionID,
		RepositoryID:    req.RepositoryID,
		ClientToken:     req.ClientToken,
		Size:            req.Size,
		PayVoucherPrice: req.PayVoucherPrice,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupRepoUpgradeRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupRepoUpgradeResponse{
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
		RegionID:      realResponse.RegionID,
	}, nil
}

type EcsBackupRepoUpgradeRealRequest struct {
	RegionID        *string  `json:"regionID,omitempty"`
	RepositoryID    *string  `json:"repositoryID,omitempty"`
	ClientToken     *string  `json:"clientToken,omitempty"`
	Size            *int     `json:"size,omitempty"`
	PayVoucherPrice *float64 `json:"payVoucherPrice,omitempty"`
}

type EcsBackupRepoUpgradeRequest struct {
	RegionID        *string
	RepositoryID    *string
	ClientToken     *string
	Size            *int
	PayVoucherPrice *float64
}

type EcsBackupRepoUpgradeRealResponse struct {
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
	RegionID      string `json:"regionID,omitempty"`
}

type EcsBackupRepoUpgradeResponse struct {
	MasterOrderID string
	MasterOrderNO string
	RegionID      string
}
