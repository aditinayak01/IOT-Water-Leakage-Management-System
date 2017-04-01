FROM ubuntu

RUN \
    apt-get update -y && \
    apt-get install -y nano influx* grafana*

CMD ["service influxdb start","service grafana start"]