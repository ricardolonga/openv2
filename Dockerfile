FROM google/golang

WORKDIR /gopath/src/github.com/ricardolonga/openv2
ADD . /gopath/src/github.com/ricardolonga/openv2/

RUN go get github.com/gin-gonic/gin
RUN go get github.com/ricardolonga/openv2

EXPOSE 8080

CMD []

ENTRYPOINT ["/gopath/bin/openv2"]
