package wcharge_mqtt

import (
	"github.com/Miroshinsv/wcharge_mqtt/config"
	grpc_v1 "github.com/Miroshinsv/wcharge_mqtt/internal/controller/grpc/v1"
	"github.com/Miroshinsv/wcharge_mqtt/internal/usecase"
	//"github.com/Miroshinsv/wcharge_mqtt/pkg/logger"
)

func Run(cfg *config.Config) {
	//l := logger.New(cfg.Log.Level)

	//l.Debug(fmt.Printf("Test logger !"))
	u := usecase.New()

	grpc_v1.Start(cfg, u) //, nil)
}
