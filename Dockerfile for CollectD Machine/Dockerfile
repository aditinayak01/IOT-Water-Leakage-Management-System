FROM ubuntu

#LABEL \
 #   author = "Siddheswar Chandrasekhar <innocentdevil.sid007@gmail.com>"
  #  version = "1.0" \
   # description = "Dockerfile to build an Ubuntu machine with CollectD"

RUN \
    apt-get update -y && \
    apt-get install -y collectd*

CMD ["service collectd start"]