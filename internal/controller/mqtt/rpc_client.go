package mqtt

import (
	"context"
	//"log"
	//"time"

	grpc_v1 "github.com/Miroshinsv/wcharge_mqtt/gen/v1"
	_ "github.com/Miroshinsv/wcharge_mqtt/pkg/rabbitmq_service"
	//amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func (mq *MqttController) PushPowerBank(ctx context.Context, r *grpc_v1.CommandPush) (*grpc_v1.ResponsePush, error) {
	m := &grpc_v1.RequestPush{RlSlot: uint32(4)}
	//m.RlSeq = uint32(1)
	//m.Push.RlSlot = uint32(4)
	//r.Push.RlSlot = uint32(4)
	messageBytes, err := proto.Marshal(m)

	t := mq.Rabbit.PublishMqtt("cabinet/RL3H082111030142/cmd/15", messageBytes)
	//t.Wait()

	if err != nil {
		mq.logger.Info("Failed to declare a queue reply: %w", err)
	}
	tt := mq.Rabbit.SubscribeMqtt("cabinet/RL3H082111030142/reply/15")

	//time.Sleep(10 * time.Second)

	//tt.Wait()
	//tt.WaitTimeout(20 * time.Second)

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	//
	if t == nil {
		mq.logger.Info("Failed to declare a queue reply: %w", err)
	}
	//
	if tt == nil {
		mq.logger.Info("Failed to declare a queue reply: %w", err)
	}

	//msg := &grpc_v1.ResponsePush{}
	//for d := range tt {
	msg := &grpc_v1.ResponsePush{}
	err = proto.Unmarshal(tt, msg)
	//	if err != nil {
	//		continue
	//	}
	//	if uint32(1) == msg.RlSeq {
	//		break
	//	}
	//}

	if err != nil {
		mq.logger.Info("%s: %s", "Failed to encode message", err)
	}

	return msg, nil
	/*
		ch, err := mq.rabbit.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()
		topik := /*r.Device.Cabinet +*/ /*"cabinet/RL3H082111030142" + /*r.Device.DeviceNumber + */ /*"/cmd/" + PUSH_POWER_BANK
	_, err = ch.QueueDeclare(
		topik, // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		mq.logger.Info("Failed to declare a queue cmd: %w", err)
	}

	topikReply := /*r.Device.Cabinet +*/ /*"cabinet/RL3H082111030142" + /*r.Device.DeviceNumber +*/ /*"/reply/" + PUSH_POWER_BANK
	_, err = ch.QueueDeclare(
		topikReply, // name
		false,      // durable
		false,      // delete when unused
		true,       // exclusive
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		mq.logger.Info("Failed to declare a queue reply: %w", err)
	}

	msgs, err := ch.Consume(
		topikReply, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, ok := mq.seqs[topik]
	if !ok {
		mq.seqs[topik] = 1
	} else {
		mq.seqs[topik] += 1
	}
	rlSeq, _ := mq.seqs[topik]
	m := r.Push
	m.RlSeq = uint32(rlSeq)
	m.RlSlot = uint32(4)
	messageBytes, err := proto.Marshal(m)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to encode message", err)
	}

	err = ch.PublishWithContext(
		ctx,
		"",    // exchange
		topik, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/octet-stream",
			//CorrelationId: corrId,
			ReplyTo: topikReply,
			Body:    messageBytes,
		})
	failOnError(err, "Failed to publish a message")

	msg := &grpc_v1.ResponsePush{}
	for d := range msgs {
		msg = &grpc_v1.ResponsePush{}
		err := proto.Unmarshal(d.Body, msg)
		if err != nil {
			continue
		}
		if uint32(rlSeq) == msg.RlSeq {
			break
		}
	}

	return msg, nil*/
}

//func (mq *MqttController) QueryInventory(ctx context.Context, r *grpc_v1.CommandInventory) (*grpc_v1.ResponseInventory, error) {
//	ch, err := mq.rabbit.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//	topik := r.Device.Cabinet + "/" + r.Device.DeviceNumber + "/cmd/" + QUERY_THE_INVENTORY
//	_, err = ch.QueueDeclare(
//		topik, // name
//		false, // durable
//		false, // delete when unused
//		true,  // exclusive
//		false, // noWait
//		nil,   // arguments
//	)
//	if err != nil {
//		mq.logger.Info("Failed to declare a queue cmd: %w", err)
//	}
//
//	topikReply := r.Device.Cabinet + "/" + r.Device.DeviceNumber + "/reply/" + QUERY_THE_INVENTORY
//	_, err = ch.QueueDeclare(
//		topikReply, // name
//		false,      // durable
//		false,      // delete when unused
//		true,       // exclusive
//		false,      // noWait
//		nil,        // arguments
//	)
//	if err != nil {
//		mq.logger.Info("Failed to declare a queue reply: %w", err)
//	}
//
//	msgs, err := ch.Consume(
//		topikReply, // queue
//		"",         // consumer
//		true,       // auto-ack
//		false,      // exclusive
//		false,      // no-local
//		false,      // no-wait
//		nil,        // args
//	)
//	failOnError(err, "Failed to register a consumer")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	_, ok := mq.seqs[topik]
//	if !ok {
//		mq.seqs[topik] = 1
//	} else {
//		mq.seqs[topik] += 1
//	}
//	rlSeq, _ := mq.seqs[topik]
//	m := r.Invent
//	m.RlSeq = uint32(rlSeq)
//	messageBytes, err := proto.Marshal(m)
//	if err != nil {
//		log.Fatalf("%s: %s", "Failed to encode message", err)
//	}
//
//	err = ch.PublishWithContext(ctx,
//		"",    // exchange
//		topik, // routing key
//		false, // mandatory
//		false, // immediate
//		amqp.Publishing{
//			ContentType: "application/octet-stream",
//			//CorrelationId: corrId,
//			ReplyTo: topikReply,
//			Body:    messageBytes,
//		})
//	failOnError(err, "Failed to publish a message")
//
//	msg := &grpc_v1.ResponseInventory{}
//	for d := range msgs {
//		msg = &grpc_v1.ResponseInventory{}
//		err := proto.Unmarshal(d.Body, msg)
//		if err != nil {
//			continue
//		}
//		if uint32(rlSeq) == msg.RlSeq {
//			break
//		}
//	}
//
//	return msg, nil
//}
