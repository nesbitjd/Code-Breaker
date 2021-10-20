FROM alpine

COPY ./codebreaker /bin/

ENTRYPOINT [ "/bin/codebreaker" ]