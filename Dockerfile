FROM progrium/busybox
MAINTAINER Andreas Larsson

ADD ./stage/busybox-go-webapp /bin/busybox-go-webapp
RUN chmod 755 /bin/busybox-go-webapp

ENTRYPOINT ["/bin/busybox-go-webapp"]