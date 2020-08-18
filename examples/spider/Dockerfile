FROM centos:6.6

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /opt

ADD ./crawler_common.tar.gz /opt/

COPY ./start.sh /opt/
RUN chmod +x /opt/start.sh

CMD ["sh", "/opt/start.sh"]
