FROM scratch

COPY code_breaker /bin/code_breaker

ENTRYPOINT [ "/bin/code_breaker" ]