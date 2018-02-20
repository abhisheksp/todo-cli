FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/todo-cli /opt/bin/todo-cli

CMD ["/opt/bin/todo-cli"]