FROM alpine

ADD go-lang /

EXPOSE ${PORT}

CMD ./go-lang --environment=${ENVIRONMENT} --port=${PORT}
