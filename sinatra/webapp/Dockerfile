from fedora
maintainer xiaolong "im.longkai@gmail.com"
run dnf install ruby redis -y

run gem sources --add https://ruby.taobao.org/ --remove https://rubygems.org/
run gem install --no-rdoc --no-ri sinatra redis

run mkdir -p /opt/webapp

expose 4567

cmd ["/opt/webapp/bin/webapp"]
