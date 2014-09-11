FROM progrium/busybox
MAINTAINER Andreas Larsson

ADD ./stage/tiny-go-webapp /bin/tiny-go-webapp

ENTRYPOINT ["/bin/tiny-go-webapp"]