from fedora
maintainer xiaolong "im.longkai@gmail.com"

add *.repo /etc/dnf.repos.d/

run dnf makecache
run dnf install iproute -y
run dnf install redis -y

expose 6379

entrypoint ["/usr/bin/redis-server"]

cmd []
