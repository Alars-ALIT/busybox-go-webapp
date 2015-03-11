FROM progrium/busybox
MAINTAINER Andreas Larsson

ADD ./release/go-app /bin/go-app
RUN chmod 755 /bin/go-app

ENTRYPOINT ["/bin/go-app"]