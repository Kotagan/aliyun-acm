# aliyun-acm

aliyun-acm是对[阿里云应用配置管理](https://help.aliyun.com/product/59604.html)的封装

## Usage

```go
package main

import (
	"fmt"
	aliacm "github.com/xiaojiaoyu100/aliyun-acm/v2"
	"github.com/xiaojiaoyu100/aliyun-acm/v2/config"
	"github.com/xiaojiaoyu100/aliyun-acm/v2/info"
	"github.com/xiaojiaoyu100/aliyun-acm/v2/observer"
)

func handle(coll map[info.Info]*config.Config) {
	for _, o := range coll {
		fmt.Println(string(o.Content))
	}
}

func main() {
	d, err := aliacm.New(
		addr,
		tenant,
		accessKey,
		secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	o1, err := observer.New(
		observer.WithInfo(
			info.Info{Group: "YourGroup", DataID: "YourDataID"},
			),
		observer.WithHandler(handle))
	if err != nil {
		return
	}
	o2, err := observer.New(
		observer.WithInfo(
			info.Info{Group: "YourGroup", DataID: "YourDataID"},
			info.Info{Group: "YourAnotherGroup", DataID: "YourAnotherDataID"},
			),
		observer.WithHandler(handle))
	if err != nil {
		return
	}

	var f = func(h info.Info, err error) {
		fmt.Println(err)
	}
	d.SetHook(f)

	d.Register(o1, o2)

	select{}
}
```