package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupRepoDeleteApi 退订云主机备份存储库
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=6907&data=87&isNormal=1
type EcsBackupRepoDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupRepoDeleteApi(client *ctyunsdk.CtyunClient) *EcsBackupRepoDeleteApi {
	return &EcsBackupRepoDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-repo/delete",
		},
	}
}

func (this *EcsBackupRepoDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupRepoDeleteRequest) (*EcsBackupRepoDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupRepoDeleteRealRequest{
		RegionID:     req.RegionID,
		RepositoryID: req.RepositoryID,
		ClientToken:  req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupRepoDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupRepoDeleteResponse{
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
		RegionID:      realResponse.RegionID,
	}, nil
}

type EcsBackupRepoDeleteRealRequest struct {
	RegionID     *string `json:"regionID,omitempty"`
	RepositoryID *string `json:"repositoryID,omitempty"`
	ClientToken  *string `json:"clientToken,omitempty"`
}

type EcsBackupRepoDeleteRequest struct {
	RegionID     *string
	RepositoryID *string
	ClientToken  *string
}

type EcsBackupRepoDeleteRealResponse struct {
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
	RegionID      string `json:"regionID,omitempty"`
}

type EcsBackupRepoDeleteResponse struct {
	MasterOrderID string
	MasterOrderNO string
	RegionID      string
}
