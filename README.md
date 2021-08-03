# 项目介绍

自动签到程序, 纯属娱乐.

# 功能

目前实现的自动签到的网站:

- [Go 语言中文网](https://studygolang.com/)
- [B 站](https://www.bilibili.com/)
- [黑客派](https://hacpai.com/)
- [v2ex](https://v2ex.com/)

# 二进制文件部署

1. 配置 Go 环境
2. 编译

```shell
$ make build.apiserver
```

可执行文件默认输出到 `_output/build/apiserver.out`.

3. 更改配置文件

```yaml
# telegram bot
bot:
  token: ""
  chat_id: 0
logger:
  path: ""
  mode: "" # debug|release
api_server:
  host: ""
  port: ""
  user: "" # http basic auth
  pass: ""
storage:
  path: ""
```

4. 启动

```shell
$ ./apiserver.out -f config.yaml
```

# Kubernetes 部署

k8s 部署配置模板在 `deployments/apiserver` 中.

1. 创建持久卷

```shell
$ kubectl create -f pv.yaml
```

2. 创建持久卷声明

```shell
$ kubectl create -f pvc.yaml
```

3. 创建服务

```shell
$ kubectl create -f svc.yaml
```

4. 创建 ConfigMap

```shell
$ kubectl create configmap apiserver-cm --from-file=config.yaml
```

5. 创建 Deployment

```shell
$ kubectl create -f deployment.yaml
```