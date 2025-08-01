FROM golang:1.24 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o mokujin

FROM gcr.io/distroless/static-debian12

WORKDIR /mokujin

# Copy the binary from the builder stage to the final image
COPY --from=builder /build/mokujin /mokujin/mokujin

# Run the binary from the correct path
CMD ["/mokujin/mokujin"]