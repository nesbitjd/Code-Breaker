FROM alpine

COPY ./hangle_server /bin/

ENTRYPOINT [ "/bin/hangle_server" ]