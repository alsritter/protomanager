package protomanager

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/alsritter/protomanager/util"
	"google.golang.org/grpc"
)

func TestServer(t *testing.T) {
	log.Println("Starting...")
	ctx, cancel := context.WithCancel(context.Background())
	stop := util.RegisterExitHandlers(cancel)
	defer cancel()

	// here is something server.
	server := grpc.NewServer(grpc.UnknownServiceHandler(handleStream))
	listener, _ := net.Listen("tcp", "127.0.0.1:4000")

	util.StartServiceAsync(ctx, log.Default(), cancel, func() error {
		return server.Serve(listener)
	}, func() error {
		server.GracefulStop()
		return nil
	})

	<-stop
	log.Panicln("Goodbye")
}

func startServer() {

}

func handleStream(srv interface{}, stream grpc.ServerStream) error {
	// fullMethodName, ok := grpc.MethodFromServerStream(stream)
	// if !ok {
	// 	return status.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	// }
	// md, _ := metadata.FromIncomingContext(stream.Context())
	// method, ok := s.protoManager.GetMethod(fullMethodName)
	// if !ok {
	// 	return status.Errorf(codes.NotFound, "method not found")
	// }
	// request := dynamic.NewMessage(method.GetInputType())
	// if err := stream.RecvMsg(request); err != nil {
	// 	return status.Errorf(codes.Unknown, "failed to recv request")
	// }
	// data, err := request.MarshalJSONPB(&jsonpb.Marshaler{})
	// if err != nil {
	// 	return status.Errorf(codes.Unknown, "failed to marshal request")
	// }
	// response, err := s.apiManager.MockResponse(context.TODO(), &interact.Request{
	// 	Protocol: interact.ProtocolGRPC,
	// 	Method:   http.MethodPost,
	// 	Host:     getAuthorityFromMetadata(md),
	// 	Path:     fullMethodName,
	// 	Header:   getHeadersFromMetadata(md),
	// 	Body:     interact.NewBytesMessage(data),
	// })
	// if err != nil {
	// 	return err
	// }
	// stream.SetTrailer(metadata.New(response.Trailer))
	// if len(response.Header) > 0 {
	// 	if err := stream.SetHeader(metadata.New(response.Header)); err != nil {
	// 		return status.Errorf(codes.Unavailable, "failed to set header: %s", err)
	// 	}
	// }
	// if response.Code != 0 {
	// 	return status.Errorf(codes.Code(response.Code), "expected code is: %d", response.Code)
	// }
	// if err := stream.SendMsg(response.Body); err != nil {
	// 	return status.Errorf(codes.Internal, "failed to send message: %s", err)
	// }
	return nil
}
