FROM ubuntu

RUN \
    apt-get update -y && \
    apt-get install -y collectd*

CMD ["service collectd start"]