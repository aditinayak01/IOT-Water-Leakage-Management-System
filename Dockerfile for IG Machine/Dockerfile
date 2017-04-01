FROM ubuntu

RUN \
    apt-get update -y && \
    apt-get install -y nano influx* grafana*

ADD ./configs/influxdb.conf /etc/influxdb/influxdb.conf

CMD ["service influxdb start","service grafana start"]