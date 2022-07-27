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
		ProtoImportPaths: []string{"co3k/protobuf-swagger-example"},
		Synchronization: &synchronization.Config{
			Enable:     true,
			StorageDir: "",
			Repository: []*synchronization.Repository{{"git@github.com:co3k/protobuf-swagger-example.git", "master"}},
		},
	}, clog)
	if err != nil {
		t.Error(err)
	}

	err = pms.Start(ctx, cancel)
	if err != nil {
		t.Error(err)
	}
}
