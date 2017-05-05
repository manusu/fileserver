FROM daocloud.io/golang:1.3-onbuild
MAINTAINER  suzhen
RUN go build .
ADD ./server   /
CMD /server
