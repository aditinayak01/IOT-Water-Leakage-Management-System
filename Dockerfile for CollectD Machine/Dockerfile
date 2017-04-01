FROM ubuntu

RUN \
    apt-get update -y && \
    apt-get install -y nano collectd*

ADD ./configs/collectd.conf /etc/collectd/collectd.conf

CMD ["service collectd start"]