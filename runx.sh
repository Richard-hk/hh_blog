#!bin/sh
CNAME="hh_blog"
PORT=3000
# 停止容器
function stopContainer(){
    stop=$(docker stop $1)
    if [[ x$stop == x"" ]]; then
        echo "stopContainer err"
    else
        echo "stopContainer: " $stop
    fi
}
# 删除容器
function rmContainer(){
    rm=$(docker rm $1)
    if [[ x$rm == x"" ]];  then
        echo "rmContainer err"
    else 
        echo "rmContainer: " $rm
    fi
}
# 创建容器
function runContainer(){
    run=$(docker run $1)
    if [[ x$run == x"" ]]; then
        echo "runContainer err"
    else
        echo "runContainer: " $run
    fi
}

stopContainer $CNAME 
rmContainer $CNAME
runContainer "-itd --name $CNAME -p $PORT:$PORT $CNAME:latest"