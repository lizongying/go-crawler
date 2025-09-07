package db

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/utils"
	"go.uber.org/fx"
	"sync"
)

type StorageFactory struct {
	config sync.Map
	logger pkg.Logger

	clients sync.Map

	bucket string
}

func (s *StorageFactory) GetClient(name string) (storage pkg.Store, err error) {
	if v, ok := s.clients.Load(name); ok {
		return v.(pkg.Store), nil
	}

	c, ok := s.config.Load(name)
	if !ok {
		return nil, fmt.Errorf("storage config %s not found", name)
	}

	conf := c.(pkg.Storage)

	bucket := conf.Bucket
	if bucket == "" {
		bucket = s.bucket
	}

	if utils.InSlice(conf.Type, []string{"s3", "oss", "cos", "minio"}) {
		storage, err = NewS3Client(conf, s.logger)
	} else {
		storage, err = NewFileClient(conf, s.logger)
	}

	if err != nil {
		return nil, fmt.Errorf("storage %s error", name)
	}

	actual, loaded := s.clients.LoadOrStore(name, storage)
	if loaded {
		_ = storage.Close(context.Background())
	}

	return actual.(pkg.Store), nil
}

func (s *StorageFactory) Close(ctx context.Context) error {
	s.clients.Range(func(key, value interface{}) bool {
		_ = value.(pkg.Store).Close(ctx)
		return true
	})
	return nil
}

func NewStorageFactory(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (storageFactory *StorageFactory, err error) {
	storageFactory = &StorageFactory{
		logger: logger,
		bucket: config.BotName,
	}
	for _, i := range config.StorageList {
		storageFactory.config.Store(i.Name, i)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) (err error) {
			return storageFactory.Close(ctx)
		},
	})
	return
}
