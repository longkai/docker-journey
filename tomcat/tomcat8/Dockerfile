from xiaolong/centos

run yum install -y java-1.8.0-openjdk
add http://mirrors.cnnic.cn/apache/tomcat/tomcat-8/v8.5.4/bin/apache-tomcat-8.5.4.tar.gz /opt/
workdir /opt/
run tar xvf apache-tomcat-8.5.4.tar.gz
run mv apache-tomcat-8.5.4 tomcat8

#ENV CATALINA_HOME /usr/share/tomcat8
#ENV CATALINA_BASE /var/lib/tomcat8
#ENV CATALINA_PID /var/run/tomcat8.pid
#ENV CATALINA_SH /usr/share/tomcat8/bin/catalina.sh
#ENV CATALINA_TMPDIR /tmp/tomcat8-tomcat8-tmp

#RUN mkdir -p $CATALINA_TMPDIR

VOLUME [ "/opt/tomcat8/webapps/" ]

EXPOSE 8080

ENTRYPOINT [ "/opt/tomcat8/bin/catalina.sh", "run" ]
