# WT-Wallet API Gateway Service

---

**本服务需要同时提供内网和外网IP访问，其他service仅限内网**

## k8s部署参考 (不使用k8s直接跳到单机安装)

* [k8s部署文档](https://go-zero.dev/docs/tutorials/ops/k8s)
* [k8s参考配置代码](https://github.com/zeromicro/zero-examples/tree/main/discovery/k8s)

## 单机安装部署

### golang 1.22.4 Linux 安装命令

目前只写了linux版本，以下脚本在跳板机 Ubuntu 上终端直接执行无问题

```bash
curl -LO https://golang.org/dl/go1.22.4.linux-amd64.tar.gz

sudo rm -rf /usr/local/go

sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz

if ! grep -q 'export PATH=/usr/local/go/bin:$PATH' ~/.bashrc; then
  echo 'export PATH=/usr/local/go/bin:$PATH' >> ~/.bashrc
fi

if ! grep -q 'export GO111MODULE=on' ~/.bashrc; then
  echo 'export GO111MODULE=on' >> ~/.bashrc
fi

rm go1.22.4.linux-amd64.tar.gz

source ~/.bashrc

go version
```

### 编译环境

根目录下执行，初始化只需要执行一次，每次更新代码不需要重复执行

```
make init
```

### 编译二进制运行文件

```bash
make build
```

### 要分发到生产服务器上的文件

两个目录：

* **/bin** bin目录以及bin目录下二进制运行文件
* **/configs** 目录下的全部文件，env.yaml由运维修改，只需初始化复制一次，app.yaml需每次复制

### 运行程序

需要在运行命令上添加 conf 参数来指定configs的目录路径，如下是在项目根目录运行的命令

```bash
./bin/marketing_service -conf ./configs
```

成功运行结果示例如下：
每个程序启动输出不一样，有些只有 GRPC server，有些只有 http server，有些两个可能都有，只要看到有8088或9099的端口启动服务即可

```text
./bin/marketing_service -conf ./configs
2024/07/01 01:38:26 maxprocs: Leaving GOMAXPROCS=12: CPU quota undefined
DEBUG msg=config loaded: app.yaml format: yaml
DEBUG msg=config loaded: env.yaml format: yaml
INFO ts=2024-07-01T01:38:26+07:00 caller=http/server.go:329 service.id=Dragon.local service.name= service.version=cfd4281 trace.id= span.id= msg=[HTTP] server listening on: [::]:8088
INFO ts=2024-07-01T01:38:26+07:00 caller=grpc/server.go:212 service.id=Dragon.local service.name= service.version=cfd4281 trace.id= span.id= msg=[gRPC] server listening on: [::]:9099
```

## Docker

docker配置需改动

```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8088:8088 -p 9099:9099 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```
