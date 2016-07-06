FROM golang:1.5
EXPOSE 9090
WORKDIR /app
COPY . /app
RUN chmod a+x .shipped/build .shipped/run .shipped/test
RUN [".shipped/build"]
CMD .shipped/run
