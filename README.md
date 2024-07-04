# pub-ip

该项目使用 Go 语言，实现了监听外部请求并返回公网 IP 地址，默认监听端口为8001。

过滤了私有保留 IP 地址，支持IPv4。

支持curl访问，如：

```bash
curl ip.rrrzr.top
```

Demo：https://ip.rrrzr.top/
