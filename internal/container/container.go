package container

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"device-manager/internal/config"
	"device-manager/internal/pkg/metrics"
	"device-manager/internal/service/databus"
	"device-manager/internal/service/devices"
	"device-manager/internal/service/storage"
)

type Container struct {
	Database        *mongo.Client
	Logger          *zap.SugaredLogger
	Storage         *storage.Storage
	DatabusProducer *databus.Producer
	DevicesService  *devices.Service
	Metrics         *metrics.StubMetrics
}

func InitContainer(cfg *config.Config) (*Container, error) {
	c := Container{}
	err := c.initDB(cfg)
	if err != nil {
		return nil, err
	}
	err = c.InitLogger()
	if err != nil {
		return nil, err
	}

	c.InitMetrics()

	c.initStorage(cfg)

	err = c.initDatabusProducer(cfg)
	if err != nil {
		return nil, err
	}

	c.InitDevicesService()

	return &c, nil
}

func (c *Container) initDB(cfg *config.Config) error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cfg.MongoUser,
		cfg.MongoPass,
		cfg.MongoHost,
		cfg.MongoPort,
		cfg.MongoDatabase,
	))

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to mongo database: %w", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return fmt.Errorf("failed to ping mongo database: %w", err)
	}

	c.Database = client

	return nil
}

func (c *Container) InitLogger() error {
	logger, err := zap.NewDevelopment() //TODO: выбирать окружение
	if err != nil {
		return err
	}

	c.Logger = logger.Sugar()

	return nil
}

func (c *Container) initDatabusProducer(cfg *config.Config) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.KafkaBroker),
		Topic:    cfg.KafkaDevicesTopic,
		Balancer: &kafka.LeastBytes{},
	}

	c.DatabusProducer = databus.NewProducer(writer, c.Logger)

	return nil
}

func (c *Container) initStorage(cfg *config.Config) {
	c.Storage = storage.New(c.Database, c.Logger, cfg.MongoDatabase)
}

func (c *Container) InitDevicesService() {
	c.DevicesService = devices.NewService(c.Storage, c.DatabusProducer)
}

func (c *Container) InitMetrics() {
	c.Metrics = metrics.NewStubMetrics()
}
