FROM centos-base:master

# Go Setup
#RUN echo $'export GOPRIVATE=path.to.private' > /opt/go113-builder.env
#RUN echo source /opt/go113-builder.env >> ~/.bashrc

RUN source ~/.bashrc

# RUN ["git clone http://github.com/sinngae/golet"]
ADD golet.tar.gz /build
WORKDIR /build/golet
RUN ["/bin/bash", "-c", "source ~/.bashrc; go mod download"]
RUN ["/bin/bash", "-c", "source ~/.bashrc; go build -o /build/goletx cmd/goletx/main.go"]

WORKDIR /build
ENV PORT=8080
EXPOSE 8080

#ADD pipeline /opt/pipeline
ENTRYPOINT ["/build/goletx"]
