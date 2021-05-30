package main

import (
	"context"

	"github.com/micro/go-micro/v2/logger"

	postProto "micro-service/proto/post"
	userProto "micro-service/proto/user"

	micro "github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// 使用方式：将下面代码插入到web，api项目中来调用user.QueryUserByName服务。可以参考web/gin.go。
func main() {
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379","127.0.0.1:3379","127.0.0.1:4379"}//地址
		//etcd2.Auth("gangan","123")//密码
	})
	// Create a new service
	service := micro.NewService(micro.Name("user.client"),micro.Registry(reg))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new user client
	user := userProto.NewUserService("go.micro.srv.user", service.Client())
	post := postProto.NewPostService("go.micro.srv.user", service.Client())

	// Call the user
	rsp, err := user.QueryUserByName(context.TODO(), &userProto.Request{UserName: "John"})
	if err != nil {
		logger.Fatal(err)
	}

	// Print response
	logger.Info(rsp.GetUser())

	rsp2, err2 := post.QueryUserPosts(context.TODO(), &postProto.Request{UserID: 1})
	if err2 != nil {
		logger.Fatal(err2)
	}
	// Print response
	logger.Info(rsp2.GetPost().Title)
}
