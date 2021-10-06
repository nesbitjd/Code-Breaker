FROM scratch

COPY codebreaker /bin/code_breaker

ENTRYPOINT [ "/bin/code_breaker" ]