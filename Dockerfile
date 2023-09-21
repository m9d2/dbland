FROM alpine:3.18 as backend
ENV GO_VERSION 1.21.1
ENV CGO_ENABLED=1
ENV GOPROXY=https://goproxy.cn
WORKDIR /app/backend
COPY ./ ./
RUN set -eux; \
    apk add --no-cache bash ca-certificates curl gcc libc-dev; \
    OS=$(uname -s | tr '[:upper:]' '[:lower:]') && \
    ARCH=$(uname -m | tr '[:upper:]' '[:lower:]') && \
    if [ "$ARCH" = "x86_64" ]; then \
        TARGET="go$GO_VERSION.linux-amd64.tar.gz"; \
    elif [ "$ARCH" = "aarch64" ]; then \
        TARGET="go$GO_VERSION.linux-arm64.tar.gz"; \
    else \
        exit 1; \
    fi; \
    wget -O go.tgz https://go.dev/dl/$TARGET; \
    tar -C /usr/local -xzf go.tgz; \
    export PATH="/usr/local/go/bin:$PATH"; \
    chmod a+x build.sh; \
    bash build.sh docker

FROM node:14 AS frontend
WORKDIR /app/frontend
COPY web .
RUN npm install
RUN npm run build release

FROM alpine:3.18
WORKDIR /app
COPY --from=frontend /app/frontend/dist ./static
COPY --from=backend /app/backend/dbland .
COPY config.yaml README.md README_CN.md LICENSE ./
EXPOSE 2023
CMD ./dbland start -p 2023