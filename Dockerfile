FROM alpine

RUN mkdir -p /statistics
RUN mkdir -p /statistics/data
WORKDIR /statistics
COPY server_docker /statistics/
COPY config.yaml /statistics/
RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --update --no-cache tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 
RUN echo "Asia/Shanghai" > /etc/timezone
ENV LIBRARY_PATH=/lib:/usr/lib
EXPOSE 13280
CMD ["/statistics/server_docker"]