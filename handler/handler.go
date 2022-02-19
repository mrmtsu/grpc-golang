package handler

import (
	"grpc-golang/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewAPIHandler(
	helloQueryServer pb.HelloQueryServer,
) *grpc.Server {
	// gRPCサーバーの生成
	server := grpc.NewServer()

	// 自動生成された関数に、サーバと実際に処理を行うメソッドを実装したhandlerを設定
	// protoファイルで定義した`helloQuery`に対応。
	pb.RegisterHelloQueryServer(server, helloQueryServer)

	// サーバーリフレクションを有効
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになる
	reflection.Register(server)

	return server
}
