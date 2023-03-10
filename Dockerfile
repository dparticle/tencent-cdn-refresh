FROM golang:1.18 as builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o tcr .

FROM alpine
COPY --from=builder /app/tcr /tcr
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]