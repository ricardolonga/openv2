FROM google/golang

WORKDIR /gopath/src/github.com/ricardolonga/openv2
ADD . /gopath/src/github.com/ricardolonga/openv2/

# go get all of the dependencies
RUN go get github.com/gin-gonic/gin

EXPOSE 8080

CMD []

ENTRYPOINT ["/gopath/bin/ricardolonga"]
