package v1

import (
	"fmt"
	"net"

	"github.com/Miroshinsv/wcharge_mqtt/config"
	grpc_v1 "github.com/Miroshinsv/wcharge_mqtt/gen/v1"
	"github.com/Miroshinsv/wcharge_mqtt/internal/controller/mqtt"
	"github.com/Miroshinsv/wcharge_mqtt/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type mqttv1server struct {
	useCase *usecase.UseCase
	/*logger  logger.Interface*/
	rb *mqtt.MqttController
	grpc_v1.UnimplementedMqttMiddlewareV1Server
}

func NewMqttV1Server(u *usecase.UseCase, /*l logger.Interface, */ cfg *config.Config, server *grpc.Server) *mqttv1server {
	s := &mqttv1server{
		useCase: u,
		/*logger:  l,*/
		rb: mqtt.NewMqttController(cfg.MQTT.URL, /* l, */ server),
	}
	return s
}

func Start(cfg *config.Config, u *usecase.UseCase) { //, l logger.Interface) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		//l.Fatal(err)
	}

	g := grpc.NewServer()
	reflection.Register(g) // разрешить запрос на именование gRPC функций
	s := NewMqttV1Server(u, /*l, */ cfg, g)

	//grpc_v1.MainServerV1Client()

	grpc_v1.RegisterMqttMiddlewareV1Server(g, s)

	//l.Info(fmt.Sprintf("gRPC server listen on %s port ...", cfg.GRPC.Port))

	if err = g.Serve(lis); err != nil {
		//l.Fatal(err)
	}
}
