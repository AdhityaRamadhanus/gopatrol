FROM golang:1.7

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/AdhityaRamadhanus/checkupd

WORKDIR /go/src/github.com/AdhityaRamadhanus/checkupd

VOLUME ["/checkup", "/checkup/logs"]

COPY checkup.json /checkup/checkup.json

RUN make

# Run the outyet command by default when the container starts.
ENTRYPOINT ["./build/linux/checkupd", "--config", "/checkup/checkup.json"] 

# Expose ports.
EXPOSE 9009