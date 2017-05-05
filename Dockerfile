FROM daocloud.io/golang:1.3-onbuild
MAINTAINER  suzhen
ADD ./server   /
CMD /server
