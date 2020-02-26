FROM centos:centos7

MAINTAINER OpenShift Development <dev@lists.openshift.redhat.com>

RUN mkdir -p /home/init
COPY ./src /home/init

ENV RUN_CMD="/home/init/init.sh"

CMD ["bash", "-c", "${RUN_CMD}"]