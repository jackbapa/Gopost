# Gopost
httppost-for-golan
===========================
目前提供两个函数
## 完整的post命令发送
```go
func Gopost_full(url string , k_v_map map[string][]string , body_reader io.Reader) (string,error)
```

- [x]url：post目标地址，即base_url
- [x]k_v_map：
