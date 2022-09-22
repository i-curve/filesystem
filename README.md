# filesystem

基于 golang + gin 的文件系统

## Intro

- path-url 固定转换
- 文件属性公有

filesystem 是一套基于规则的文件系统存储系统，根据特定 path 转为特定 url 地址。所有文件为公有属性，所有人可以进行访问。

把文件系统从整个项目中进行抽离，优化项目结构

## Usage

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
