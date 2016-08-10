使用 Docker 镜像和仓库
===
## 什么是 Docker 镜像
copy-on-write

## 列出镜像
`docker images`

- 标签 image:tag
- 用户仓库 user/image
- 顶层仓库 image

## 拉取镜像
`docker pull [image]`，默认拉取 `latest` tag

## 查找镜像
`docker search [keywords]`

## 构建镜像
* 使用 `docker commit`
* 使用 `docker build` 和 *Dockerfile* 文件

> 不推荐 commit 命令，另外一般是局域一个已有的基础镜像来构建镜像

### 创建 Docker Hub 帐号
登录 docker hub `docker login`

### 用 Docker 的 commit 命令创建镜像
类似于版本系统里提交变更 `docker commit id user/image`

### 基于 Dockerfile 构建新镜像
[see here](./Dockerfile)

`docker build -t="user/name:tag" .` 或者远程的 repo 也行啊

### Dockerfile 和构建缓存
默认是包含缓存滴，`--no-cache` 禁止缓存

### 查看新镜像
`docker history`

### 从新镜像启动
`docker run -d -p 80 --name=[container-name] [image] [command]` 这里`-p`选项用来控制 Docker 在运行时一个公开那些网络端口给外部，可以随机选择或者指定具体的端口。

通过 `docker port [container-name]` 可以查看 docker 和宿主机的端口映射。

- 绑定特定端口 `-p 8080:80`
- 绑定特定的网络接口 `-p 127.0.0.1:80:80`
- 绑定到随机端口，留空即可

也可以使用 `-P` 用来公开在 Dockerfile 里面 EXPOSE 的端口

## Dockerfile 指令
### `CMD`
用于指定一个容器启动时要运行的命令，这和 `docker run ...` 非常类似，后者能够覆盖前者，并且一个 Dockerfile 中有且仅有一个 CMD 命令。

### `ENTRYPOINT`
类似 `CMD` 但是不能被覆盖，`run` 或者 `CMD` 的指令会追加到 `ENTRYPOINT` 中

### `WORKDIR`
指定的程序会在这个目录下执行，也可以通过 `-w` 在运行时覆盖工作目录。

### `ENV`
设置环境变量，会持久到容器容器运行，能通过 `-e key=val` 传递环境变量

### `USER`
指定 `user:group` 去运行镜像，默认为 root

### `VOLUME`
向容器添加卷，卷可以共享数据或者对数据进行持久化。

- 卷可以在容器见贡献和重用
- 一个容器不一定的和其他容器共享卷
- 对卷的修改是立即生效的
- 对卷的修改不会对更新镜像产生影响
- 卷会一直存在知道没有任何容器再使用它

> 卷功能让我们可以将数据、数据库或者其他内容添加到镜像中而不是将这些内容提交到镜像中，并且允许我们在多个容器间共享这些内容。我们可以利用此功能来测试容器和内部的应用程序代码，管理日志，或者处理容器内部的数据库。

> `docker cp` 能够在容器合宿主机中复制文件

### `ADD`
`ADD from to`

### `COPY`
不会像 `ADD` 那些解压缩

### `LABEL`
`LABEL key=val key2=val2`

### `STOPSOGNAL`
Unix signal

### `ARG`
用于在 `docker build` 命令执行前传递变量，`docker build --build-arg xx=xx balabala...`

> 不要传递隐私的变量

### `ONBUILD`
当一个镜像被用作其他镜像的基础镜像时，`ONBUILD ADD . /app/src`

## 推送的 Docker Hub
TODO

## 删除镜像
`docker rmi [image]`

