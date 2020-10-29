# rest-load
rest api 负载测试
## 使用

```bash
rest-load -c http-client.env.json  geoip-m.http 

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
    "url":"http://test-vnet.sunniwell.net:11180/iplist/iplist.txt"
  }
```

引用
- ```${1.data.}```
