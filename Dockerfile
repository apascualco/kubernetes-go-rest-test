FROM arm64v8/golang:1.14
COPY src /opt/src
WORKDIR /opt/src/
RUN go build -o /bin/app && chmod +x /bin/app && rm -Rf /opt/src
EXPOSE 80
CMD ["/bin/app"]
