FROM golang:1.20-alpine
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -tags netgo -o /app .

FROM debian
RUN apt-get -qqy update
RUN apt-get -qqy install wget
RUN apt-get -qqy install python3
RUN wget -qO /usr/local/bin/yt-dlp https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp
RUN chmod a+rx /usr/local/bin/yt-dlp

COPY --from=0 /app /app
ENTRYPOINT ["/app"]