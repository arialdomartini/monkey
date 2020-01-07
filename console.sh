docker run \
       -v $(pwd)/src:/go/src \
       -w /go \
       --rm \
       -ti \
       --name golang \
       golang
