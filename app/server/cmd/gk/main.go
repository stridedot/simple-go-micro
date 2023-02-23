package main

import (
	pb "app/protobuf/helloworld"
	"context"
	"flag"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill)
		defer signal.Stop(ch)

		select {
		case <-ch:
			log.Printf("received interrupt signal")
		case <-ctx.Done():
			log.Printf("goroutine should cancel")
		}
		return nil
	})

	eg.Go(func() error {
		addr := ":50051"
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		log.Printf("grpc server listen succeeded, addr:%s", listener.Addr())

		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, &server{})
		log.Printf("server listening at %v", listener.Addr())

		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		return nil
	})

	eg.Go(func() error {
		broker := "192.168.2.110:10004" // Kafka 代理地址
		groupID := "test-group"    // 消费者组 ID
		topic := "morning"      // 要消费的主题

		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": broker,
			"group.id":          groupID,
			"auto.offset.reset": "earliest",
		})

		if err != nil {
			return err
		}

		err = consumer.SubscribeTopics([]string{topic}, nil)
		if err != nil {
			return err
		}

		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				log.Printf("Received message: %s\n", string(msg.Value))
			} else {
				log.Printf("Error: %v (%v)\n", err, msg)
			}
		}
	})

	if err := eg.Wait(); err != nil {
		log.Fatalf("get error: %v", err)
	}
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
