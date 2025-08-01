FROM golang:1.24

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o ./mokujin

FROM gcr.io/distroless/static-debian12

WORKDIR /mokujin
# COPY --from=builder /build/mokujin /mokujin/
CMD ["mokujin"]
