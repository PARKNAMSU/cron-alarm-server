FROM golang:1.23.0-alpine3.10 AS builder

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
RUN ./aws/install -i /usr/local/aws-cli -b /usr/local/bin

WORKDIR /cron-alarm-server
COPY . .
RUN mv .aws ~/

EXPOSE 8080

CMD ["go","run","./app/cmd/main.go"]