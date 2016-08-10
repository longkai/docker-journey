from fedora:latest
maintainer longkai "im.longkai@gmail.com"
run yum install iproute -y
run yum install nginx -y
run mkdir -p /var/www/html/website
add ./nginx/global.conf /etc/nginx/conf.d/
add ./nginx/nginx.conf /etc/nginx/nginx.conf
expose 80
