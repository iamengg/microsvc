FROM golang:latest as builder
WORKDIR /app 
COPY go.mod go.sum ./  
RUN go mod download
RUN go mod tidy  
COPY . .  
RUN go build main.go 
# EXPOSE 80
# ENTRYPOINT [ "./main" ]

FROM gcr.io/distroless/base-debian11 
COPY --from=builder /app/main . 
EXPOSE 80
CMD [ "/main" ]