FROM alpine
RUN apk add git apache-ant openjdk8 bash \
    && git clone -b 5.0-mysql-support-opt-2.1 https://github.com/pingcap/benchmarksql.git
WORKDIR ./benchmarksql
RUN ant
COPY entrypoint.sh .
RUN chmod +x ./entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]
