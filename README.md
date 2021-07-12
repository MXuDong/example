# The example project for mock



The project for mock the product application in test environment.

For mock the application run in the product, like `http/bin`.

## Status:
[![Build Status](http://drone.daocloud.cn/api/badges/xudong.meng/example/status.svg)](http://drone.daocloud.cn/xudong.meng/example)

Already in build-on, for CI branch is develop.

# Feature

* Support the Docker environment
* [Support the Kubernetes environment](./doc/Kubernetes-Feature.md)
* Support the Istio with Kubernetes environment
* Support run directly on vm or pc.
* Support protocol mock
    * support tcp protocol

# Usage

## Build and Run

### Source build
    pre-environment: golang support(v1.15+)
1. Clone project to your environment.
2. Use `go` to build binary-product.
3. Run it.

### Deployment on kubernetes
See the [deployment.yaml](./deployment.yaml).

### Deployment with docker environment
#### Local build
To build image:
```shell
$ docker build -t example:v0.0.0 .
```

To run image:
```shell
$ docker run -itd --name=example -p 3000:3000 example:v0.0.0
```

Support tcp port:
```shell
# Tcp server default listen the 8000, change it on config/conf.yaml
$ docker run -itd --name=example -p 3000:3000 -p 8000:8000 example:v0.0.0
```

## Mock Protocol invoke

### TCP Protocol
The example provide `tcp` protocol mock server.

Set config.yaml `server_config.enable_tcp_server = true` to enable tcp server.

    Only the server_config.enable_tcp_server is true, the tcp will listen target port to accept the connect.

Make a post to create a tcp connection:

**POST** 

`http://{Server Ipaddress and Port}}/protocol/tcp`

**BODY**
```json
{
    "connect_time": 10,
    "remote_ip_address": "127.0.0.1",
    "remote_ip_port": "8000",
    "send_byte_count": 10,
    "protocol": "tcp"
}
```

- `connect_time`: how long connect keep
- `remote_ip_address`: target tcp server's ip address
- `remote_ip_port`: port of target tcp server listen
- `send_byte_count`: the tcp client will send body size to target tcp server
- `protocol`: support tcp protocols: tcp, tcp4, tcp6, unix, unixpacket

**RESPONSE-200**
```json
{
    "invoke_time": "2021-06-01T16:57:58.859541+08:00",
    "end_time": "2021-06-01T16:58:09.860653+08:00",
    "value_byte": "AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQ==",
    "value_size": 100,
    "request_protocol": "tcp",
    "invoking_setting": {
        "connect_time": 10,
        "remote_ip_address": "127.0.0.1",
        "remote_ip_port": "8000",
        "send_byte_count": 10,
        "protocol": "tcp"
    }
}
```