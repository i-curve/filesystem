# filesystem

基于 golang + gin 的简单文件系统

## Intro

- path-url 固定转换
- 文件属性公有

filesystem 是一套基于规则的文件系统存储系统，根据特定 path 转为特定 url 地址。所有文件为公有属性，所有人可以进行访问。

把文件系统从整个项目中进行抽离，优化项目结构

## Deploy

- 源码部署

克隆项目

```bash
git clone https://github.com/i-curve/filesystem.git
```

进入项目并启动

```bash
cd filesystem && go run main.go
```

- 下载可执行文件

[下载位置](https://github.com/i-curve/filesystem/releases)

- 使用 docker

拉取镜像

```bash
docker pull wjuncurve/filesystem
```

启动镜像

```bash
docker run -d --rm \
 --name filesystem \
 -p 8080:8080 \
 -v /var/www/html/data:/data \
 -e "BASE_URL=http://127.0.0.1" \
 -e "USER=i-curve" \
 -e "AUTH=12345678" \
    wjuncurve/filesystem
```

-v 把路径挂在到前端展示的页面上  
-p 映射想要的端口  
BASE_URL: 网络文件 url  
USER: 用户  
AUTH: 用户的认证

如果未指定 USER 和 AUTH 的话, 为生成一个临时的

- 检测是否部署成功

```bash
# 查看版本信息
curl "http://localhost:8080/version"
```

## Usage

文件上传

```bash
curl 'http://localhost:8080/upload' -X POST \
--form 'file=@"C:\\Users\\curve\\Pictures\\images.jpg"'

```

文件获取

```bash
curl "http://localhost:8080/file?short_url=${url}"
```

文件复制

```bash
curl "http://localhost:8080/copy" -X POST \
    -d "short_url=${old_url}&new_url=${new_url}"
```

文件转移

```bash
curl "http://localhost:8080/move" -X POST \
    -d "short_url=${old_url}&new_url=${new_url}"
```

文件删除

```bash
curl "http://localhost:8080/file?short_url=${short_url}" -X DELETE
```

[详情请参考文档](https://www.apifox.cn/apidoc/shared-e29b73da-4337-4787-8a0f-e31312d8f99e/api-40901537)

## 测试环境

地址: [https://filesystem.ml/test](https://filesystem.ml/test)

user: i-curve
auth: 12345678
