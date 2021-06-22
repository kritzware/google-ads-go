FROM namely/protoc-all:latest

RUN apt-get install wget -y
RUN wget https://dl.google.com/go/go1.16.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz

# ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/usr/local/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
RUN go version

RUN apt-get install git -y
RUN git clone https://github.com/googleapis/googleapis /gapi
RUN rm -rf /opt/include/google/ads/googleads/
RUN mv /gapi/google/ads/googleads/ /opt/include/google/ads/googleads/

RUN protoc --version

RUN go get github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic \ 
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc


RUN mkdir /build
COPY build.sh /

RUN sh /build.sh