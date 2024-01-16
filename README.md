# filesystem

![GitHub Release](https://img.shields.io/github/v/release/i-curve/filesystem)[![Docker Image CI](https://github.com/i-curve/filesystem/actions/workflows/docker-image.yml/badge.svg)](https://github.com/i-curve/filesystem/actions/workflows/docker-image.yml)![GitHub Tag](https://img.shields.io/github/v/tag/i-curve/filesystem-gosdk?label=filesystem-gosdk)![GitHub Tag](https://img.shields.io/github/v/tag/i-curve/filesystem-pysdk?label=filesystem-pysdk)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [filesystem](#filesystem)
  - [I. web 服务](#i-web-服务)
  - [II. api 服务](#ii-api-服务)
    - [1. user](#1-user)
    - [2. bucket](#2-bucket)
    - [3. filesystem](#3-filesystem)
    - [4. admin](#4-admin)
      - [权限刷新](#权限刷新)
  - [III. sdk](#iii-sdk)
    - [1. go sdk](#1-go-sdk)
  - [IV. deploy](#iv-deploy)
    - [构建程序](#构建程序)
    - [docker 构建](#docker-构建)
    - [docker hub](#docker-hub)
    - [运行 docker](#运行-docker)
    - [5. 检测是否部署成功](#5-检测是否部署成功)

<!-- /code_chunk_output -->

基于 golang + gin 的简单文件系统

## I. web 服务

> 0.0.0.0: 8000

用于 http 资源请求, 只能访问,不能上传

默认挂载目录: /var/www/filesystem

可以直接通过 web 端进行访问

- 下载

下载条件:

1. bucker1 桶下有 test.txt 文件
2. bucket1 属性为公有可访问

[http://localhost:8000/bucket1/test.txt]()

## II. api 服务

0.0.0.0: 8001

api 服务需要进行身份认证, 需要在请求头添加 user,auth 进行身份认证

### 1. user

> user 等级:
> system: 默认会自动生成一个 system 用户
> admin: 管理用户
> user: 不同用户, 只对自己创建的桶拥有权限

- 添加用户

```bash
curl -X POST "${URL}/user" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
--header 'Content-Type: application/json' \
--data-raw '{"name": "i-curve","u_type": 2}'
```

- 删除用户

```bash
curl -X DELETE "${URL}/user" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
--header 'Content-Type: application/json' \
--data-raw '{"name": "i-curve"}'
```

### 2. bucket

- 创建 bucket

```bash
curl -X POST "${URL}/bucket" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
--header 'Content-Type: application/json' \
--data-raw '{"name": "bucket2", "b_type": 3}'
```

- 删除 bucket

```bash
curl -X DELETE "${URL}/bucket" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
--header 'Content-Type: application/json' \
--data-raw '{"name": "bucket2"}'
```

### 3. filesystem

- 文件上传

上传 a.txt 文件到 bucket1 的 test/01 内

```bash
curl -X POST "${URL}/file/upload" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
-F "file=@/root/a.txt" \
-F "bucket=bucket1" \
-F "key=/test/a.txt"
```

- 文件下载

```bash
curl -X GET "${URL}/file/download?bucket=bucket1&key=/test/a.txt" \
--header "user: ${USER}" \
--header "auth: ${AUTH}"
```

- 文件删除

```bash
curl -X DELETE "${URL}/file/delete" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
-d '{"bucket": "bucket", "key": "/test/01"}'
```

- 文件移动

```bash
curl -X POST "${URL}/file/move" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
--header "Content-Type: application/json" \
-d '{"s_bucket": "bucket1",
"s_key": "test/a.txt","d_bucket": "bucket1","d_key":"test/b.txt"}'
```

- 文件复制

```bash
curl -X POST "${URL}/file/copy" \
--header "user: ${USER}" \
--header "auth: ${AUTH}" \
--header "Content-Type: application/json" \
-d '{"s_bucket":"bucket1","s_key":"test/a.txt","d_bucket": "bucket1","d_key": "new_path/dd/b.txt"}'
```

### 4. admin

管理接口

#### 权限刷新

当进行数据库更改后, 并不会直接生效, 需要进行权限刷新

```bash
curl -X POST "${URL}/refresh" \
--header "user: ${USER}" \
--header  "auth: ${AUTH}"
```

- 获取版本信息

```bash
curl "${URL}/version" \
--header  "user: ${USER}" \
--header  "auth: ${AUTH}"
```

## III. sdk

### 1. go sdk

repo: [https://github.com/i-curve/filesystem-gosdk](https://github.com/i-curve/filesystem-gosdk)

## IV. deploy

克隆项目

```bash
git clone https://github.com/i-curve/filesystem.git
```

### 构建程序

生成可执行文件 filesystem

```bash
cd filesystem && make
```

### docker 构建

构建 docker 镜像

```bash
make docker
```

### docker hub

拉取镜像

```bash
docker pull wjuncurve/filesystem:latest
```

### 运行 docker

```bash
docker run --name filesystem -d  \
 -p 8000:8000 -p 8001:8001 \
 -e "MYSQL_HOST=host.docker.internal" \
    filesystem
```

| 变量名         | 说明                                    |
| -------------- | --------------------------------------- |
| LANGUAGE       | 程序运行语言,默认 zh 中文(en 可选)      |
| MODE           | 运行模式 (DEBUG \| RELEASE)             |
| BASE_DIR       | 默认程序存储目录为(/var/www/filesystem) |
| MYSQL_HOST     | 数据库地址                              |
| MYSQL_USER     | 数据库用户                              |
| MYSQL_PASSWORD | 数据库密码                              |
| MYSQL_PORT     | 数据库端口                              |
| DATABASE       | 数据库名 (默认 filesystem)              |

### 5. 检测是否部署成功

查看版本信息

```bash
curl "http://localhost:8001/version"
```

查看生成 system 用户信息

```bash
docker logs filesystem
```

```txt
system user
user: system
auth: d6837f276686cbfe4852c8d8c5104f59
```
