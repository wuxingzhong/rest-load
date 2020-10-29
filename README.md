# rest-load
rest api 负载测试
## 使用

```bash
rest-load -c http-client.env.json  test.http

# 选择配置环境
rest-load -c http-client.env.json  -e baidu  test.http

```

## 结果替换

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

引用: ```${1.data.}```
