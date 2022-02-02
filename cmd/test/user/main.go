package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/user"
)

var (
	ErrConfigNotFound = errors.New("config not found")
)

func main() {
	service := micro.NewService(
		micro.Name("test-user"),
		micro.Registry(etcd.NewRegistry()),
	)

	service.Init(
		micro.Action(func(cli *cli.Context) (err error) {
			path := cli.String("config")
			if path == "" {
				return ErrConfigNotFound
			}
			log.Printf(" config path:[%s]\n", path)

			root := config.ReadConfigs(path)

			return service.Options().
				Registry.Init(
				registry.Addrs(root.Etcd.Endpoints...),
				registry.TLSConfig(
					util.NewTLSConfig(root.Etcd.Ca, root.Etcd.Cert, root.Etcd.Key),
				),
			)
		}),
	)

	userService := pb.NewUserService(pb.ServiceName, service.Client())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testCreateUser := func() {
		cuRsp, err := userService.CreateUser(ctx, &pb.CreateUserRequest{User: &pb.User{
			Nickname: "jdxj",
			Password: "jdxj",
			Contact: &pb.Contact{
				Mail:     "",
				Telegram: 0,
			},
		}})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("cuRsp: %+v\n", cuRsp)
	}
	testCreateUser()

	testAuthUser := func() {
		auRsp, err := userService.AuthUser(ctx, &pb.AuthUserRequest{
			Nickname: "",
			Password: "iii",
		})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("auRsp: %+v\n", auRsp)
	}
	testAuthUser()

	testGetUser := func() {
		guRsp, err := userService.GetUser(ctx, &pb.GetUserRequest{UserID: 2})
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("guRsp: %s\n", guRsp)
	}
	testGetUser()

	testUpdateUser := func() {
		_, err := userService.UpdateUser(ctx, &pb.UpdateUserRequest{User: &pb.User{
			UserId:   1,
			Nickname: "jdxj-rename",
			Password: "jdxj-repass",
			Contact: &pb.Contact{
				Mail:     "abc",
				Telegram: 8888,
			},
		}})
		if err != nil {
			log.Println(err)
			return
		}
	}
	testUpdateUser()
}
