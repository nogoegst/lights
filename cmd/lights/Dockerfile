# docker run -p 80:80 --rm -d --mount type=bind,source=/sys/,destination=/sys/ lights
FROM scratch
COPY lights /
ENTRYPOINT ["/lights"]
