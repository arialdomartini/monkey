docker run \
       -v $(pwd)/src:/go/src \
       -w /go/src \
       --rm \
       -ti \
       --name golang \
       golang
