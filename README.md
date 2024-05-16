# ctyun-sdk-go

天翼云C端sdk的Go语言实现



## 项目结构

- ctyun-sdk-core：sdk核心入口
- ctyun-sdk-endpoint：各个产线端点的sdk实现
  - ctebs：弹性云硬盘相关，`go get github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctebs@latest`
  - ctecs：云主机相关，`go get github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctecs@latest`
  - ctiam：iam相关，`go get github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctiam@latest`
  - ctimage：镜像相关，`go get github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctimage@latest`
  - ctvpc：虚拟私有云相关，`go get github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctvpc@latest`



## 使用方式

- 引入相关的功能包

  ```
  go get github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctiam@latest
  ```

- 注册&使用sdk

  ```
  package main
  
  import (
  	"context"
  	"fmt"
  	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
  	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctiam"
  )
  
  func main() {
  	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
  	credential, _ := ctyunsdk.NewCredential("您的ak", "您的sk")
  	apis := ctiam.NewApis(client)
  	resp, _ := apis.UserGetApi.Do(context.Background(), *credential, &ctiam.UserGetRequest{
  		UserId: "查询的userId",
  	})
  	fmt.Println(resp.UserName)
  }
  ```

  