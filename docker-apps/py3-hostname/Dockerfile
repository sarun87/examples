FROM alpine:3.7

RUN apk add --no-cache python3 && \
    if [[ ! -e /usr/bin/python ]]; then ln -sf /usr/bin/python3 /usr/bin/python; fi

COPY server.py /usr/src/server/server.py
EXPOSE 80/TCP
CMD [ "python" , "/usr/src/server/server.py"]

