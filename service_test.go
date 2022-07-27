package protomanager

import (
	"os"
	"testing"

	"github.com/alsritter/protomanager/util/logger"
	"github.com/alsritter/protomanager/util/synchronization"
	"golang.org/x/net/context"
)

func TestManager_GetMethod(t *testing.T) {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		clog        = logger.NewDefault("test")
	)

	basePath, _ := os.Getwd()
	pms, err := New(&Config{
		ProtoDir:         basePath,
		ProtoImportPaths: []string{"temporary/alsritter/protobuf-examples"},
		Synchronization: &synchronization.Config{
			Enable:     true,
			StorageDir: "temporary",
			Repository: []*synchronization.Repository{{"git@github.com:alsritter/protobuf-examples.git", "main"}},
		},
	}, clog)
	if err != nil {
		t.Error(err)
	}

	err = pms.Start(ctx, cancel)
	if err != nil {
		t.Error(err)
	}

	d, ext := pms.GetMethod("/hello.Hello/SayHello")
	if ext {
		t.Logf("查询到的服务全地址为：%#v", d.GetFullyQualifiedName())
	} else {
		t.Error("不存在")
	}
}
