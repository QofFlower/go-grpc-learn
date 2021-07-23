# Grpc for Golang

## 环境配置以及软件安装

请参考[视频][https://www.bilibili.com/video/BV1GE411A7kp]

P1部分



注意: 中间运行命令`protoc --go_out=../services Prod.proto`出错，根据官方文档，在.proto文件中指定go-package

```protobuf
syntax = "proto3";
package services;
option go_package = "../services";
message ProdRequest {
    int32 prod_id = 1; // 传入的商品id
}
message ProdResponse {
    int32 prod_stock = 1; // 商品库存
}
```



## 快速入门demo

Prod.Proto文件

```protobuf
syntax = "proto3";
package services;
option go_package = "../services"; // 指定go_package，否则执行命令会出错
message ProdRequest {
    int32 prod_id = 1; // 传入的商品id
}
message ProdResponse {
    int32 prod_stock = 1; // 商品库存
}
service ProdService {
    rpc GetProdStock (ProdRequest) returns (ProdResponse);
}
```



运行命令

```shell
protoc --go_out=plugins=grpc:../services Prod.proto
```



继承生成的Prod.pb.go文件里的ProdServiceServer接口

ProdService.go:

```java
package services

import (
	context "context"
)

type ProdService struct {
}

func (t *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}

```

![image-20210721162556461](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210721162556461.png)

在主目录下添加server.go文件

```go
package main

import (
	"go-rpc/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{})

	listenErr, _ := net.Listen("tcp", ":8083")

	rpcServer.Serve(listenErr)
}

```

client工程部分

![image-20210721171135301](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210721171135301.png)

Prod.pb.go是和server工程一样的

main.go

```go
package main

import (
	"context"
	"fmt"
	"go-rpc/grpc-first/client/services"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	prodClient := services.NewProdServiceClient(cc)
	prodRes, err1 := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 10086})
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(prodRes.ProdStock)
}

```

启动server再启动client即可接收到信息



## 自签证书、服务加入证书验证

### 准备工作

[工具下载](https://slproweb.com/products/Win32OpenSSL.html)

在安装目录下的bin里执行cmd

```shell
openssl
```

```shell
genrsa -out ca.key 2048
```

```shell
req -new -x509 -days 3650 -key ca.key -out ca.pem
```



![image-20210721173924982](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210721173924982.png)

> 注意：Common name填写localhost

### 生成服务的证书

```shell
genrsa -out server.key 2048
```

```shell
req -new -key server.key -out server.csr
```

```shell
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
```

生成server.csr文件和server.pem

### 生成客户端证书

```shell
ecparam -genkey -name secp384r1 -out client.key
req -new -key client.key -out client.csr
x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
```



在服务端工程内的cert目录下，将三个文件ca.pem, server.key, server.pem粘贴过去

![image-20210721180011393](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210721180011393.png)

同样的，将ca.pem, client.key, client.pem粘贴到client工程下的cert目录中



GO 1.15 以上版本解决GRPC X509 Common Name field, use SANs or temporarily enable Common Name matching

[解决方法](https://blog.csdn.net/ma_jiang/article/details/111992609)

或者设置系统变量`GODEBUG=x509IngoreNO=0`

### 运行

启动服务端

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"go-rpc/services"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{})

	l, err := net.Listen("tcp", ":10086")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer.Serve(l)
}

```

启动客户端

```go
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-rpc/grpc-first/client/services"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err != nil {
		log.Fatal("Failed to load the credentials key of client", err)
	}
	certPool := x509.NewCertPool()
	ca, err1 := ioutil.ReadFile("cert/ca.pem")
	if err1 != nil {
		log.Fatal("Failed to load the public key", err1)
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("Failed to connect the server", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cc)

	prodClient := services.NewProdServiceClient(cc)
	prodRes, err1 := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 10086})
	if err1 != nil {
		log.Fatal("Encountering error during requestint to the server", err1)
	}
	fmt.Println(prodRes.ProdStock)
}

```

![image-20210722103242505](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210722103242505.png)

## 双向认证下rpc-gateway的使用

[官网](https://github.com/grpc-ecosystem/grpc-gateway)

由于官网使用的比较新的v2版本，在这里我使用视频内所使用的版本，按照视频的步骤做，如果使用的是官方文档，那么请按照官方文档的步骤进行

```shell
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```



找到下载的grpc-gateway包并将google目录下的api拷贝过来，vscode下不好定位所下载的包，可以进入golang查看

![image-20210722153235553](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210722153235553.png)

![image-20210722153306829](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210722153306829.png)



在Prod.proto文件中定义好路由接口:

```protobuf
syntax = "proto3";
package services;
option go_package = "../services"; // 指定go_package，否则执行命令会出错
import "google/api/annotations.proto";

message ProdRequest {
    int32 prod_id = 1; // 传入的商品id
}
message ProdResponse {
    int32 prod_stock = 1; // 商品库存
}
service ProdService {
    rpc GetProdStock (ProdRequest) returns (ProdResponse) {
        option (google.api.http) = {
            get: "/v1/prod/{prod_id}" // get请求路径映射
        };
    }
}
```

server.go:

```go
package main

import (
	"go-rpc/helper"
	"go-rpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredential()))
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{})

	l, err := net.Listen("tcp", ":10086")
	if err != nil {
		log.Fatal(err)
		return
	}
	rpcServer.Serve(l)
}

```

httpserver.go:

```go
package main

import (
	"context"
	"go-rpc/helper"
	"go-rpc/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}

	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:10086", opt)

	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":9468",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}

```

访问 http://localhost:9468/v1/prod/998

![image-20210722153505435](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210722153505435.png)





## 语法入门

### 结合类型

现在需求更改为返回集合类型的库存类型

修改Prod.proto

```protobuf
syntax = "proto3";
package services;
option go_package = "../services"; // 指定go_package，否则执行命令会出错
import "google/api/annotations.proto";

message ProdRequest {
    int32 prod_id = 1; // 传入的商品id
}
message ProdResponse {
    int32 prod_stock = 1; // 商品库存
}
message QuerySize {
    int32 size = 1; // 页的大小
}
message ProdResponseList { // 返回商品库存集合
    repeated ProdResponse productes = 1; // repeated表示类似于其他语言的东西，术语是啥我就不深究了，无妨
}

service ProdService {
    rpc GetProdStock (ProdRequest) returns (ProdResponse) {
        option (google.api.http) = {
            get: "/v1/prod/{prod_id}" // get请求路径映射
        };
    }
    rpc GetProdStocks(QuerySize) returns (ProdResponseList) {}
}
```

为了方便编译生成文件，我们使用脚本

```bat
@REM cd pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
cd pbfiles && protoc --go_out=../services --go-grpc_out=../services Prod.proto
```

在命令行执行即可

```powershell
.\generate.bat
```



实现生成的接口:

```go
package services

import (
	context "context"
	"fmt"
)

type ProdService struct {
}

func (t *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 114514}, nil
}

func (t *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	return &ProdResponseList{Productes: []*ProdResponse{
		{ProdStock: 114514},
		{ProdStock: 10086},
		{ProdStock: 721},
		{ProdStock: 666},
	}}, nil
}

func (t *ProdService) mustEmbedUnimplementedProdServiceServer() {
	fmt.Println("Fuck you")
}

```



把Prod_grpc.pb.go和Prod.pb.go复制到客户端工程中

客户端运行函数

```go
package main

import (
	"context"
	"fmt"
	"go-rpc/grpc-first/client/helper"
	"go-rpc/grpc-first/client/services"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(helper.GetClientCredentials()))

	if err != nil {
		log.Fatal("Failed to connect the server", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cc)

	prodClient := services.NewProdServiceClient(cc)
	cxt := context.Background()

	res, err := prodClient.GetProdStocks(cxt, &services.QuerySize{Size: 10})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res.Productes)
}

```

完成



### 分区获取

先有A，B，C三个库存，用户需要在请求的时候指定哪一个库，并返回该库的库存信息



枚举类型

```protobuf
enum ProdAreas{
	A = 0;
	B = 1;
	C = 2;
}
```

编辑proto文件

```protobuf
syntax = "proto3";
package services;
option go_package = "../services"; // 指定go_package，否则执行命令会出错
import "google/api/annotations.proto";

enum ProdAreas {
    A = 0;
    B = 1;
    C = 2;
}

message ProdRequest {
    int32 prod_id = 1; // 传入的商品id
    ProdAreas prod_area = 2;
}
message ProdResponse {
    int32 prod_stock = 1; // 商品库存
}
message QuerySize {
    int32 size = 1; // 页的大小
}
message ProdResponseList { // 返回商品库存集合
    repeated ProdResponse productes = 1; // repeated表示类似于其他语言的东西，术语是啥我就不深究了，无妨
}

service ProdService {
    rpc GetProdStock (ProdRequest) returns (ProdResponse) {
        option (google.api.http) = {
            get: "/v1/prod/{prod_id}" // get请求路径映射
        };
    }
    rpc GetProdStocks(QuerySize) returns (ProdResponseList) {}
}
```

实现接口并编写业务

```go
package services

import (
	context "context"
	"fmt"
)

type ProdService struct {
}

func (t *ProdService) GetProdStock(cxt context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32
	if request.ProdArea == ProdAreas_A {
		stock = 114514
	} else if request.ProdArea == ProdAreas_B {
		stock = 721
	} else if request.ProdArea == ProdAreas_C {
		stock = 666
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (t *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	return &ProdResponseList{Productes: []*ProdResponse{
		{ProdStock: 114514},
		{ProdStock: 10086},
		{ProdStock: 721},
		{ProdStock: 666},
	}}, nil
}

func (t *ProdService) mustEmbedUnimplementedProdServiceServer() {
	fmt.Println("Fuck you")
}

```

客户端请求

```go
package main

import (
	"context"
	"fmt"
	"go-rpc/grpc-first/client/helper"
	. "go-rpc/grpc-first/client/services"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(helper.GetClientCredentials()))

	if err != nil {
		log.Fatal("Failed to connect the server", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cc)

	prodClient := NewProdServiceClient(cc)
	cxt := context.Background()

	res, err := prodClient.GetProdStock(cxt, &ProdRequest{ProdId: 2, ProdArea: 2})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res)
}

```

OJBK



### 导入外部proto

摸了



### 日期类型

引包

```protobuf
import "google/protobuf/timestamp.proto"
```

使用

```protobuf
message OrderMain {
	int32 order_id = 1;
	string order_no = 2;
	int32 user_id = 3;
	float order_money = 4;
	google.protobuf.Timestamp order_time = 5; // 日期类型
}
```



### 综合练习

post请求提交主订单信息，http api实现



三个proto文件:

Prod.proto

```protobuf
syntax = "proto3";
package services;
option go_package = "../services"; // 指定go_package，否则执行命令会出错
import "google/api/annotations.proto";
import "Models.proto";

enum ProdAreas {
    A = 0;
    B = 1;
    C = 2;
}

message ProdRequest {
    int32 prod_id = 1; // 传入的商品id
    ProdAreas prod_area = 2;
}
message ProdResponse {
    int32 prod_stock = 1; // 商品库存
}
message QuerySize {
    int32 size = 1; // 页的大小
}
message ProdResponseList { // 返回商品库存集合
    repeated ProdResponse productes = 1; // repeated表示类似于其他语言的东西，术语是啥我就不深究了，无妨
}

service ProdService {
    rpc GetProdStock (ProdRequest) returns (ProdResponse) {
        option (google.api.http) = {
            get: "/v1/prod/{prod_id}" // get请求路径映射
        };
    }
    rpc GetProdStocks(QuerySize) returns (ProdResponseList) {}
    rpc GetProdInfo(ProdRequest) returns (ProdModel) {}
}
```

Models.proto

```protobuf
syntax="proto3";
package services;
option go_package = "../services"; // 指定go_package，否则执行命令会出错
import "google/protobuf/timestamp.proto";

message OrderRequest {
    OrderMain order_main = 1;
}
message ProdModel {
    int32 prod_id = 1;
    string prod_name = 2;
    float prod_price = 3;
}
message OrderMain {
	int32 order_id = 1;
	string order_no = 2;
	int32 user_id = 3;
	float order_money = 4;
	google.protobuf.Timestamp order_time = 5; // 日期类型
}
```

Order.proto

```protobuf
syntax="proto3";
package services;
option go_package = "../services"; // 指定go_package，否则执行命令会出错
import "Models.proto";
import "google/api/annotations.proto";

message OrderResponse {
    string status = 1;
    string message = 2;
}

service OrderService {
    rpc NewOrder(OrderRequest) returns (OrderResponse) {
        option (google.api.http) = {
            post: "/v1/orders"
            body: "order_main"
        };
    }
}
```



方便编译的bat文件

```bash
@REM cd pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
cd pbfiles && protoc --go_out=../services --go-grpc_out=../services Prod.proto
protoc --go_out=../services --go-grpc_out=../services Models.proto
protoc --go_out=../services --go-grpc_out=../services Order.proto

protoc --grpc-gateway_out=logtostderr=true:../services Prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Order.proto
```



编译之后老样子，自己进行实现

server.go

```go
package main

import (
	"go-rpc/helper"
	"go-rpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredential()))
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{})
	services.RegisterOrderServiceServer(rpcServer, &services.OrderService{})
	l, err := net.Listen("tcp", ":10086")
	if err != nil {
		log.Fatal(err)
		return
	}
	rpcServer.Serve(l)
}

```

httpserver.go

```go
package main

import (
	"context"
	"go-rpc/helper"
	"go-rpc/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}
	endPoint := "localhost:10086"
	cxt := context.Background()

	err := services.RegisterProdServiceHandlerFromEndpoint(cxt, gwmux, endPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	err1 := services.RegisterOrderServiceHandlerFromEndpoint(cxt, gwmux, endPoint, opt)
	if err1 != nil {
		log.Fatal(err1)
	}

	httpServer := &http.Server{
		Addr:    ":9468",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}

```

先后启动server.go, httpserver.go

再像定义的接口发送POST请求

![image-20210722180935480](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210722180935480.png)





### 使用第三方验证工具

[官方文档](https://github.com/envoyproxy/protoc-gen-validate)





### grpc的流模式

在正式应用场景中，往往传输的数据量会非常大，就好比睿站吧，你看看数据量大不大

假如我们要从库里传一批用户id过去然后查询用户的积分

先创建用户模型

```protobuf
message UserInfo {
	int32 user_id = 1;
	int32 user_score = 2;
}
```



#### 服务端流

proto文件里添加服务端流方法

```protobuf
syntax="proto3";
package services;
option go_package="../services";
import "Models.proto";

message UserScoreRequest {
    repeated UserInfo users = 1;
}
message UserScoreResponse {
    repeated UserInfo users = 1;
}
service UserService {
    rpc GetUserScore(UserScoreRequest) returns (UserScoreResponse) {}
    rpc GetUserScoreByServerStream(UserScoreRequest) returns (stream UserScoreResponse) {}
}
```

方法实现

```go
package services

import (
	"context"
	"math/rand"
)

type UserService struct {
}

func (s *UserService) GetUserScore(cxt context.Context, userRequest *UserScoreRequest) (*UserScoreResponse, error) {
	res := make([]*UserInfo, 0)
	for _, user := range userRequest.Users {
		user.UserScore = rand.Int31()
		res = append(res, user)
	}
	return &UserScoreResponse{Users: res}, nil
}

// 服务端流模式返回数据
func (s *UserService) GetUserScoreByServerStream(userRequest *UserScoreRequest, serverStream UserService_GetUserScoreByServerStreamServer) error {
	res := make([]*UserInfo, 0)
	for index, user := range userRequest.Users {
		user.UserScore = rand.Int31n(300)
		res = append(res, user)
		if (index+1)%2 == 0 { // 没隔两条发送一次
			err := serverStream.Send(&UserScoreResponse{Users: res})
			if err != nil {
				return err
			}
			res = (res)[0:0]
		}
	}
	if len(res) > 0 {
		err := serverStream.Send(&UserScoreResponse{Users: res})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *UserService) mustEmbedUnimplementedUserServiceServer() {
}

```

客户端接收

```go
func main() {
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(helper.GetClientCredentials()))

	if err != nil {
		log.Fatal("Failed to connect the server", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cc)

	cxt := context.Background()
	usc := NewUserServiceClient(cc)

	req := &UserScoreRequest{}
	for i := 0; i < 5; i++ {
		req.Users = append(req.Users, &UserInfo{UserId: int32(i + 1)})
	}

	stream, err2 := usc.GetUserScoreByServerStream(cxt, req)
	if err2 != nil {
		log.Fatal(err2)
		return
	}
	// traverse the result
	for {
		usr, err3 := stream.Recv()
		if err3 == io.EOF {
			break
		}
		if err3 != nil {
			log.Fatal(err3)
			break
		}
		fmt.Println(usr.Users)
	}
}

```

结果

![image-20210723093320253](C:\Users\huangwenmeng\AppData\Roaming\Typora\typora-user-images\image-20210723093320253.png)

#### 客户端流

proto文件

```protobuf
syntax="proto3";
package services;
option go_package="../services";
import "Models.proto";

message UserScoreRequest {
    repeated UserInfo users = 1;
}
message UserScoreResponse {
    repeated UserInfo users = 1;
}
service UserService {
    rpc GetUserScore(UserScoreRequest) returns (UserScoreResponse) {}
    rpc GetUserScoreByServerStream(UserScoreRequest) returns (stream UserScoreResponse) {}
    rpc GetUserScoreByClientStream(stream UserScoreRequest) returns (UserScoreResponse) {}
}
```

实现方法

```go
func (s *UserService) GetUserScoreByClientStream(clientStream UserService_GetUserScoreByClientStreamServer) error {
	res := make([]*UserInfo, 0)
	for {
		usr, err := clientStream.Recv()
		if err == io.EOF {
			return clientStream.SendAndClose(&UserScoreResponse{Users: res})
		}
		if err != nil {
			return err
		}
		for _, user := range usr.Users {
			user.UserScore = rand.Int31n(300)
			res = append(res, user)
		}
	}
}
```



客户端使用流分发数据

```go
func main() {
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(helper.GetClientCredentials()))

	if err != nil {
		log.Fatal("Failed to connect the server", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cc)

	cxt := context.Background()
	usc := NewUserServiceClient(cc)

	stream, err3 := usc.GetUserScoreByClientStream(cxt)
	if err3 != nil {
		log.Fatal(err3)
		return
	}

	for j := 0; j < 3; j++ {
		req := &UserScoreRequest{}
		for i := 0; i < 5; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: int32(i + 1)})
		}
		err2 := stream.Send(req)
		if err2 != nil {
			log.Fatal(err2)
			break
		}
	}
	usr, err2 := stream.CloseAndRecv()
	if err2 != nil {
		log.Fatal(err2)
		return
	}
	fmt.Println(usr.Users)
}

```

#### 双向流

客户端批量查询数据

1. 客户端分批把用户列表发送过去(客户端获取列表比较慢)
2. 服务端查询积分也很慢，所以分批发送过去

```protobuf
rpc GetUserScoreByDoubleEndStream(stream UserScoreRequest) returns (stream UserScoreResponse) {}
```

实现接口

```go
func (s *UserService) GetUserScoreByDoubleEndStream(stream UserService_GetUserScoreByDoubleEndStreamServer) error {
	res := make([]*UserInfo, 0)
	for {
		usr, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
			return nil
		}
		for _, user := range usr.Users {
			user.UserScore = rand.Int31n(100)
			res = append(res, user)
		}
		err2 := stream.Send(&UserScoreResponse{Users: res})
		if err2 != nil {
			log.Fatal(err2)
			return err2
		}
	}
}
```

客户交互

```go
func main() {
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(helper.GetClientCredentials()))

	if err != nil {
		log.Fatal("Failed to connect the server", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cc)

	cxt := context.Background()
	usc := NewUserServiceClient(cc)

	stream, err3 := usc.GetUserScoreByDoubleEndStream(cxt)
	if err3 != nil {
		log.Fatal(err3)
		return
	}

	for j := 0; j < 3; j++ {
		req := &UserScoreRequest{}
		for i := 0; i < 5; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: int32(i + 1)})
		}
		err2 := stream.Send(req)
		if err2 != nil {
			log.Fatal(err2)
			break
		}
		res, err4 := stream.Recv()
		if err4 == io.EOF {
			break
		}
		if err4 != nil {
			log.Fatal(err4)
		}
		fmt.Println(res.Users)
	}
}

```

