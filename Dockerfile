FROM centos
MAINTAINER longkai "im.longkai@gmail.com"
RUN yum install epel-release -y
RUN yum install iproute -y
RUN yum install nginx -y
#CMD ["nginx", "-g", "daemon off;"]
WORKDIR /opt/test
RUN echo "hello" > test.txt
ENTRYPOINT ["nginx"]
EXPOSE 80
