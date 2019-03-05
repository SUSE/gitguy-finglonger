FROM opensuse/tumbleweed

LABEL Maintainer="SUSE Containers Team <containers@suse.com>"

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN zypper -n up && \
	zypper -n in \
		git \
		go \
		golang-github-cpuguy83-go-md2man \
		make \
		tar && \
	go get -u golang.org/x/lint/golint && \
	go get -u github.com/vbatts/git-validation && type git-validation

VOLUME ["/go/src/github.com/SUSE/gitguy-finglonger"]
WORKDIR /go/src/github.com/SUSE/gitguy-finglonger
COPY . /go/src/github.com/SUSE/gitguy-finglonger