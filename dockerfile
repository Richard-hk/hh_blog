# 一阶段构建
FROM golang:1.17.5 as builder
# 设置代理
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
# 设置工作目录
WORKDIR /work/hh_blog
# 除.dokcerignore忽略文件外，复制其他文件到工作区
ADD  . .
# 编译go文件
RUN CGO_ENABLED=0  GOOS=linux go build -a -installsuffix cgo -o hh_blog .

# 二阶段构建，生产环境
FROM alpine as prod
WORKDIR /root 
# 从上一阶段builder构建里复制相关文件
COPY --from=builder /work/hh_blog/hh_blog . 
COPY --from=builder /work/hh_blog/web web
COPY --from=builder /work/hh_blog/conf conf
# 执行二进制文件
CMD ["./hh_blog"]
# 暴露端口
EXPOSE 3000