# rest-load
rest api 负载测试
## 使用

```bash
rest-load curl -c http-client.env.json  test.http

# 选择配置环境
rest-load curl -c http-client.env.json  -e baidu  test.http
# 命令额外参数
rest-load curl -c http-client.env.json   -e baidu \
 --extArgs "-k -v"  test.http
 
rest-load ab -c http-client.env.json   -e baidu \
 --extArgs "-n 1000 -c 100  -T application/x-www-form-urlencoded"  test.http
```

## template 

### timestamp

```
{% timestamp %}
```

### 结果替换

支持利用上一个接口的结果

示例:
前面接口结果集
```json
{ 
  "code":0,
  "message":"0",
  "ttl":1,
  "data":{
    "url":"http://127.0.0.1"
  }
  
```

```
# 接口序列号从1开始.
{% result "1.data" %}
```

## 依赖

- curl 用于http请求
- wrk http压力测试
- ab  http压力测试
