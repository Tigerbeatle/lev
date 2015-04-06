export MGO_HOSTS=ds035428.mongolab.com:35428
export MGO_DATABASE=tigerbeatle
export MGO_USERNAME=guest
export MGO_PASSWORD=welcome
export BUOY_DATABASE=tigerbeatle

cd $GOPATH/src/github.com/tigerbeatle/le/test/endpointTests
go test -v