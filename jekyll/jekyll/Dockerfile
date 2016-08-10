from xiaolong/centos
run yum install ruby ruby-devel make nodejs npm -y
run gem sources
run gem sources --remove https://rubygems.org/
run gem sources -a https://mirrors.ustc.edu.cn/rubygems/
run gem install --no-rdoc --no-ri jekyll -v 2.5.3

volume /data
volume /var/www/html
workdir /data

entrypoint ["jekyll", "build", "--destination=/var/www/html"]
