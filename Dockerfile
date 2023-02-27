ARG APP_NAME
FROM alpine:edge as builder
LABEL stage=go-builder
ARG VERSION
ARG TYPE
ARG APP_NAME
ENV GITHUB_REF=$VERSION
WORKDIR /app/
COPY ./ ./
RUN apk add --no-cache bash git curl go gcc musl-dev; \
    curl -s -O 'https://raw.githubusercontent.com/libsgh/go-build-action/main/build.sh'; \
    bash build.sh ${TYPE} ${APP_NAME}

FROM alpine:edge
MAINTAINER libsgh
ARG APP_NAME
ENV APP_NAME_ENV=${APP_NAME}
WORKDIR /app
COPY --from=builder /app/bin/${APP_NAME} ./
RUN ls -n
RUN echo $APP_NAME_ENV
CMD ["./$APP_NAME_ENV"]
