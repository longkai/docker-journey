from xiaolong/centos

run yum install -y make nodejs npm
run mkdir -p /var/log/nodeapp

add nodeapp /opt/nodeapp

workdir /opt/nodeapp

run npm install

volume ["/var/log/nodeapp"]

expose 3000

entrypoint ["node", "server.js"]
