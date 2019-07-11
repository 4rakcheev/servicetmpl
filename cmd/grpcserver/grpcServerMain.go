package main

import (
	"context"
	"github.com/jfeng45/servicetmpl/adapter/userclient"
	uspb "github.com/jfeng45/servicetmpl/adapter/userclient/generatedclient"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/container/servicecontainer"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)
type UserService struct {
	container container.Container
}

func (uss *UserService) RegisterUser(ctx context.Context, req *uspb.RegisterUserReq) (*uspb.RegisterUserResp, error) {
	logger.Log.Debug("RegisterUser called")

	ruci, err := uss.container.RetrieveRegistration()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	mu, err := userclient.GrpcToUser(req.User)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("mu:", mu)
	resultUser, err :=ruci.RegisterUser(mu)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("resultUser:", resultUser)
	gu, err := userclient.UserToGrpc(resultUser)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	logger.Log.Debug("user registered: ", gu)

	return &uspb.RegisterUserResp{User: gu}, nil

}

func (uss *UserService) ListUser(ctx context.Context, in *uspb.ListUserReq) (*uspb.ListUserResp, error) {
	logger.Log.Debug("ListUser called")

	luci, err := uss.container.RetrieveListUser()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	lu, err :=luci.ListUser()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	gu, err := userclient.UserListToGrpc(lu)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	logger.Log.Debug("user list: ", gu)

	return &uspb.ListUserResp{User: gu}, nil

}
func runServer(c container.Container) error {
	logger.Log.Debug("start runserver")

	srv:=grpc.NewServer()

	cs:= &UserService{c}
	uspb.RegisterUserServiceServer(srv, cs)
	l, err:=net.Listen("tcp", "localhost:5052")

	if err!=nil {
		return errors.Wrap(err, "")
	} else {
		logger.Log.Debug("server listening")
	}
	return srv.Serve(l)
}

func main () {
	filename := "../../configs/appConfigDev.yaml"
	container, err := buildContainer(filename)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		panic(err)
	}
	if err := runServer(container); err != nil {
		logger.Log.Errorf("Failed to run cache server: %+v\n", err)
		panic(err)
	} else {
		logger.Log.Info("server started")
	}
}
func buildContainer (filename string) (container.Container, error){

	factoryMap :=make(map[string]interface{})
	config := configs.AppConfig{}
	container := servicecontainer.ServiceContainer{factoryMap, &config}

	err:= container.InitApp( filename)
	if err!=nil  {
		return nil, errors.Wrap(err, "")
	}
	return &container, nil
}
//func initContainer() container.Container{
//	filename := "../../configs/appConfigProd.yaml"
//	container, err := buildContainer(filename)
//	if err!=nil  {
//		logger.Log.Errorf("%+v\n", err)
//		return
//	}
//	return
//	//testFindById(container)
//}
