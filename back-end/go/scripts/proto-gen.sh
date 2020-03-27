#: '
protoc -I $GOPATH/src \
    -I $GOPATH/src/github.com/googleapis \
    --go_out=$GOPATH/src \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/entities/user.proto \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/entities/product.proto \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/entities/promotion.proto \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/entities/discounted-date.proto
#'

#: '
protoc -I $GOPATH/src \
    -I $GOPATH/src/github.com/googleapis \
    --go_out=plugins=grpc:$GOPATH/src \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/user.proto \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/product.proto \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/promotion.proto \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/discounted-date.proto
#'

#: '
protoc -I $GOPATH/src \
    -I $GOPATH/src/github.com/googleapis \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true,grpc_api_configuration=../internal/yaml/services/user.yaml:$GOPATH/src \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/user.proto \
    --grpc-gateway_out=logtostderr=true,grpc_api_configuration=../internal/yaml/services/product.yaml:$GOPATH/src \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/product.proto \
    --grpc-gateway_out=logtostderr=true,grpc_api_configuration=../internal/yaml/services/promotion.yaml:$GOPATH/src \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/promotion.proto \
    --grpc-gateway_out=logtostderr=true,grpc_api_configuration=../internal/yaml/services/discounted-date.yaml:$GOPATH/src \
    $GOPATH/src/github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/discounted-date.proto
#'