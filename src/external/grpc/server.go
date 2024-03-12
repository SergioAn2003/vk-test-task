package grpc

import (
	"context"
	"fmt"
	"grpc-test/internal/entity/global"
	"grpc-test/internal/entity/user"
	"grpc-test/uimport"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protoUser "grpc-test/proto/user"
)

type Server struct {
	protoUser.UnimplementedUserServiceServer
	log *logrus.Logger
	uimport.UsecaseImports
}

func NewServer(log *logrus.Logger, ui uimport.UsecaseImports) *Server {
	return &Server{
		log:            log,
		UsecaseImports: ui,
	}
}

func (e *Server) CreateUser(ctx context.Context, req *protoUser.CreateRequest) (*protoUser.CreateResponse, error) {
	ts:=e.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		e.log.Errorln("не удалось открыть транзакцию")
	}

	defer func() {
		if ts.TxIsActive() {
			defer ts.Rollback()
		}
	}()

	u := user.User{
		ID:        int(req.GetId()),
		Name:      req.GetName(),
		Age:       int(req.GetAge()),
		IsMarried: req.GetIsMarried(),
	}

	if err := e.Usecase.Info.SaveUser(ts, u); err != nil {
		return &protoUser.CreateResponse{
			StatusCode: 500,
			Success:    false,
		}, global.ErrInternalError
	}

	if err := ts.Commit(); err != nil {
		e.log.Errorln("не удалось закрыть транзакцию")
	}

	return &protoUser.CreateResponse{
		StatusCode: 200,
		Success:    true,
	}, nil
}

func (e *Server) RunGrpcServer() {
	port := os.Getenv("PORT")
	if port == "" {
		e.log.Errorln("не удалось получить номер порта или порт не укаан")
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		e.log.Errorln("не удалось начать прослушивание")
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)

	protoUser.RegisterUserServiceServer(s, e)
	e.log.Infoln("grpc сервер запущен на порту", port)

	if err = s.Serve(lis); err != nil {
		e.log.Errorln("не удалось запустить grpc сервер, ошибка:", err)
		return
	}
}
