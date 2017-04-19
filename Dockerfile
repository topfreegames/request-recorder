FROM alpine
COPY ./bin/recorder-linux-amd64 /app/recorder
EXPOSE 8080
CMD /app/recorder start
