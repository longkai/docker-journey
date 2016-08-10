#!/bin/bash
docker build -t xiaolong/static_web .
docker run -d -P --name=static_web -v $PWD/www:/var/www/html/website:ro xiaolong/static_web nginx
