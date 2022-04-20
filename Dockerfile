# Go version
FROM golang:1.17 AS build-env
RUN mkdir /block_creator

ENV USER=trevorjo
ENV UID=10001

# create a sytstem group dev with no password, no home directory set, and no shell so prevents the user from
# being a login account and reduces the attack vector
RUN adduser \
--disabled-password \
--gecos "" \
--home "/nonexistent" \
--shell "/sbin/nologin" \
--no-create-home \
--uid "${UID}" \
${USER}

WORKDIR /block_creator
COPY . /block_creator
# change ownership of all /block_creator content to created user
RUN chown -R trevorjo /block_creator
RUN go mod download && \
go mod verify && \
CGO_ENABLED=0 go build -o app -mod vendor -trimpath cmd/main.go
USER trevorjo

FROM scratch AS run-env
WORKDIR /build
COPY --from=build-env /block_creator/app /build/
ENTRYPOINT ["/build/app", "-grpcserverhost", "blockreceiver"]