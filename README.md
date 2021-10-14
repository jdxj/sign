![logo](./docs/images/logo.png)

# 项目介绍

自动签到程序, 纯属娱乐.

# 功能

目前实现的自动签到的网站:

- [Go语言中文网](https://studygolang.com/)
- [B站](https://www.bilibili.com/)
- [黑客派](https://hacpai.com/)
- [v2ex](https://v2ex.com/)

# 部署

## 依赖环境

- MySQL (启用 Binlog)
  - 存储任务
  - 所需表在 `./deployments/sql`
- RabbitMQ
  - 推送任务
- Etcd
  - 注册中心

## 配置文件

模板在 `./configs/configs.yaml.default`

## 编译

```shell
$ make all
```

可执行文件默认输出到 `_output/build/`.

## 启动

```shell
$ ./xxx.out -f config.yaml
```

# Kubernetes 部署

k8s 部署配置模板在 `./deployments` 中.

1. 创建 ConfigMap

```shell
$ kubectl create configmap apiserver-cm --from-file=config.yaml
```

2. 创建 Deployment

```shell
$ kubectl create -f deployment.yaml
```

# 创建任务

## 使用 signctl

1. 构建 signctl

signctl 生成在 `./_ooutput/tools/`.

```shell
$ make ctl
```

2. 创建用户

创建用户后会返回 `token`, 之后任何操作使用 `-T token` 方式.

```shell
$ ./signctl.out create user -H server_address --nickname xxx
```

3. 创建 secret

创建 secret 后会返回 secretID.

```shell
$ ./signctl.out create secret -H server_address -T token --domain 101 --key xxx
```

4. 创建 task

`--secret-id` 用于指定要使用的 secret.

```shell
$ ./signctl.out create task -H server_address -T token --kind 102 --secret-id xxx --spec "0 8 * * *"
```

# 各组件介绍

## apiserver

类似网关, signctl 与其交互来对各资源进行操作.

## crontab

管理任务对象, 创建任务等.

## executor

任务的执行由其负责, 其中定义了各种任务的执行逻辑.

## notice

类似消息推送, 目前使用 telegram bot 做消息接收.

## secret

管理 secret, 一个 secret 可以对应多个 task.

## trigger

触发器, 时间到时将任务发到 RabbitMQ, 再由 executor 执行.

## user

用户管理.

## signctl
 
一个简单的命令行工具, 用于创建任务.
