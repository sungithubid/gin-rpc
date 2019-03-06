## gin-rpc
----------
基于gin和hprose的跨语言IP位置查询rpc服务端

## Features
- gin
- hprose-golang
- go-ini
- go mod

## config
```configs/main.ini```
```
[app]
# possible values : prod, dev, debug
app_mode = dev

# Runtime Path
runtimePath = runtime/

[server]
# Protocol (http or https or tcp)
protocol = http

# The http port  to use
http_port = 8099

# The tcp port  to use
tcp_port = 1314

```

## run

**server**
```$ go run main.go```
```
[GIN-debug] [WARNING] Now Gin requires Go 1.6 or later and Go 1.7 will be required soon.

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /rpc                      --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8099
``` 

**client**
客户端请求示例：phpunit测试下
```
...
$assert = '{"countryName":"中国","regionName":"江苏","cityName":"常州"}';

$ip = '61.160.251.57';
$client = Client::create('http://127.0.0.1:8099/rpc', false);
$res =  $client->IPLocate($ip);

$this->assertSame($assert, $res);
```

输出结果
```
.                    1 / 1 (100%)

Time: 40 ms, Memory: 4.00MB

OK (1 test, 1 assertion)

```


## to do list
- jwt鉴权
- 优雅的重启服务
- 使用goroutine提高连接数和性能
- docker部署