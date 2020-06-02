# GOproject
template for golang

1. docker run --rm -it -v $PWD/gopath:/gopath:ro -v $PWD/gen:/gen grpc/go /bin/bash
2. run buildenv.sh in the docker `export GOPATH=$GOPATH:/gopath`
3. in docker please need to git checkout tensorfow v2.1.0 and install [tensorflow C libs](https://www.tensorflow.org/install/lang_c) and go get
4. sudo tar -C /usr/local -xzf libtensorflow-cpu-linux-x86_64-1.15.0.tar.gz
5. sudo ldconfig