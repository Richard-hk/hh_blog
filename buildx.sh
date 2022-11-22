#!/bin/sh
PNAME="hh_blog"
VERSION=`cat version`

# 构建hh_blog项目镜像
docker build -t $PNAME .
# # 根据version版本打镜像tag
docker tag $PNAME:latest $PNAME:$VERSION

function cleanImage() {
    cleanImage=$(docker image ls | grep $1 | awk  '{print $3}' | xargs docker rmi)
    if [[ x$cleanImage == x"" ]]; then
        echo "there is no the tag or repository named $1 to clean..."
    else
        echo $cleanImage " is cleaned"
    fi
}
# 删除空白TAG和空白REPOSITORY
cleanImage "none"
