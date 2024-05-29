package v1

import (
	"context"

	grpc_v1 "github.com/Miroshinsv/wcharge_mqtt/gen/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *mqttv1server) PushPowerBank(ctx context.Context, cmd_push *grpc_v1.CommandPush) (*grpc_v1.ResponsePush, error) {
	s.logger.Debug("Return_PushPowerbank_method_ok")
	rp, err := s.rb.PushPowerBank(ctx, cmd_push)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	s.logger.Debug("Return_PushPowerbank_mqtt_ok")
	return rp, nil
}

func (s *mqttv1server) ForcePushPowerBank(context.Context, *grpc_v1.CommandPush) (*grpc_v1.ResponsePush, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForcePushPowerBank not implemented")
}

func (s *mqttv1server) QueryInventory(ctx context.Context, cmd_push *grpc_v1.CommandInventory) (*grpc_v1.ResponseInventory, error) {
	//s.logger.Debug("Return_QueryInventory_method_ok")
	//rp, err := s.rb.QueryInventory(ctx, cmd_push)
	//if err != nil {
	//	s.logger.Error(err)
	//	return nil, err
	//}
	//s.logger.Debug("Return_QueryInventory_mqtt_ok")
	return nil, status.Errorf(codes.Unimplemented, "method QueryInventory not implemented")
}

func (s *mqttv1server) QueryServerInformation(context.Context, *grpc_v1.CommandServerInformation) (*grpc_v1.ResponseServerInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryServerInformation not implemented")
}

func (s *mqttv1server) QueryCabinetAPN(context.Context, *grpc_v1.CommandCabinetAPN) (*grpc_v1.ResponseCabinetAPN, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCabinetAPN not implemented")
}

func (s *mqttv1server) QuerySIMCardICCID(context.Context, *grpc_v1.CommandSIMCardICCID) (*grpc_v1.ResponseSIMCardICCID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySIMCardICCID not implemented")
}

func (s *mqttv1server) QueryNetworkInformation(context.Context, *grpc_v1.CommandNetworkInformation) (*grpc_v1.ResponseNetworkInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNetworkInformation not implemented")
}

func (s *mqttv1server) ResetCabinet(context.Context, *grpc_v1.CommandResetCabinet) (*grpc_v1.ResponseResetCabinet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCabinet not implemented")
}

func (s *mqttv1server) Subscribe(context.Context, *grpc_v1.Device) (*grpc_v1.ResponseString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeMqtt not implemented")
}
