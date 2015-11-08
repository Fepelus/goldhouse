

FROM codenvy/go

ENV GOPATH /home/user/application

RUN mkdir -p /home/user/application/src/github.com/Fepelus/goldhouse

ENV CODENVY_APP_BIND_DIR /home/user/application/src/github.com/Fepelus/goldhouse

VOLUME ["/home/user/application/src/github.com/Fepelus/goldhouse"]

WORKDIR /home/user/application/src/github.com/Fepelus/goldhouse

# 1. Get dependencies from source code of application
# 2. Start application
# CMD go get -d && go run $executable:-main.go$
CMD go build && ./goldhouse
