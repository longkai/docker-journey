Docker 入门
===
本章将学习第一个 Docker 容器，还会介绍如何与 docker 进行交互的基本知识。

## 确保 Docker 已经就绪
`docker info` 该命令会返回所有容器和镜像数量等基本配置信息。

## 运行我们的第一个容器
`docker run` 提供了容器的**创建到启动**的功能。

```sh
macOS > docker run -i -t centos /bin/bash
[root@72c7c02024ff /]#
```

- `-i` 保证了容易中**STDIN**是开启的
- `-t` 为创建的容器分配一个伪 tty 终端
- `centos` 告诉 docker 基于什么镜像来创建容器，这里用 centos
- `/bin/bash` 告诉 docker 在容器中运行什么命令

## 使用第一个容器
以前该怎么用就这么用......

## 命名容器
`docker run --name=[a-zA-Z0-9_\.-]+ balabala...`

> 注意，每个名字的容器都是不一样的，即使镜像&命令一样

## 重新启动已经停止的容器
`docker start [container-name]`

## 附着到容器上
`docker attach [container-name]`

## 创建守护式容器
`docker run --name=[container-name] -d balabala...`

- `-d` 将会把容器放到后台运行

## 容器内部都在干些什么
`socker logs -ft --tail num [container-name]`

## Docker 日志驱动
可以控制 Docker 守护进程和容器所用的日志驱动，比如 `docker run --log-driver="syslog"`

## 查看容器内的进程
`docker top [container-name]`

## Docker 统计信息
`docker stats [container-name]...`

## 在容器内部运行进程
`docker exec -[itd] [container-name] command...`

## 停止守护式容器
`docker stop [container-name]`

## 自动重启容器
`docker run --restart=[always|on-failure:n] balabala...`

> Still not understand this translation...

## 深入容器
`docker inspect [container-name](s) -f {{.Golang-template}}`

## 删除容器
`docker rm [container-name]`

### 删除所有容器
```sh
> docker rm `docker ps -a -q`
```
