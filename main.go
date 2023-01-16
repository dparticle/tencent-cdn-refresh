package main

import (
	"flag"
	"fmt"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"strings"
)

type sliceValue []string

func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = strings.Split(val, ",")
	return nil
}

func (s *sliceValue) String() string {
	return strings.Join(*s, ",")
}

var (
	paths               []string
	secretId, secretKey string
)

func init() {
	flag.StringVar(&secretId, "id", "", "tencent secret id")
	flag.StringVar(&secretKey, "key", "", "tencent secret key")
	flag.Var(newSliceValue([]string{}, &paths), "paths", "refresh path list")
}

func main() {
	flag.Parse()
	if secretId == "" || secretKey == "" || len(paths) == 0 {
		fmt.Println("At least one parameter is empty, e.g. tcr -id xxx -key xxx -paths url1,url2")
		return
	}

	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := cdn.NewClient(credential, "", cpf)

	request := cdn.NewPurgePathCacheRequest()
	request.Paths = common.StringPtrs(paths)
	request.FlushType = common.StringPtr("flush") // 刷新类型，flush 刷新产生更新的资源，delete 刷新全部资源

	response, err := client.PurgePathCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
