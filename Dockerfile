FROM scratch

COPY codebreaker /codebreaker

ENTRYPOINT [ "/codebreaker" ]