FROM golang:1.10 as hellobuilder

WORKDIR /

COPY main.go .

RUN CGO_ENABLED=0 go build -o hello -a -ldflags '-s' main.go

FROM scratch

COPY --from=hellobuilder hello .

CMD ["/hello"]
