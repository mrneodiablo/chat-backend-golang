####################################
# STEP 1 | build executable binary #
####################################
FROM golang:1.11 AS builder

RUN mkdir -p /go/src/hope-pet-chat-backend
ADD hope-pet-chat-backend /go/src/hope-pet-chat-backend


ENV CSGOHOME=/go/src/hope-pet-chat-backend

WORKDIR $CSGOHOME

RUN go get -u github.com/kardianos/govendor

RUN govendor sync -v

RUN CGO_ENABLED=0 go build -a -installsuffix cgo  -ldflags '-s -extldflags "-static"' -o chat-backend -v main.go


################################
# STEP 2 | build a small image #
################################
FROM alpine:latest


RUN mkdir -p /data

# Copy file executable from STEP 1.
COPY --from=builder /go/src/hope-pet-chat-backend/chat-backend  /data/chat-backend
RUN chmod +x /data/chat-backend

WORKDIR /data

ADD conf conf
RUN mkdir logs

EXPOSE 2345

CMD ["./chat-backend"]
