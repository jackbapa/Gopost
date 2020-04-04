# Gopost  
:rainbow::rainbow::rainbow:  
****
httppost-for-golan  
go语言发送post请求 


```diff
+ 内置chorme请求头，模拟浏览器请求，防反爬
- 忽略服务器https验证，防无证书下post失败
! 以 map或切片 传递键值对，可一key配多value（map下）
+ 以body_reader接口读取所有实现reader接口的类
！使用mahonia实现gbk-utf8转码
```

****
## 完整的post命令发送 :high_brightness:
```go
func Gopost_full(url string , k_v_map map[string][]string , body_reader io.Reader) (string , error)
```
-  url：  post目标地址，即base_url
-  k_v_map：  以map类型表述的查询参数,其中key为string类型，value为string切片
-  body_reader:  为实现io_reader接口的数据类型，用于写入request.body，即写入请求体的数据
举例如下，查询键（key）为name，值（value）为wy和wxy，url为1techbook.com，请求体为空字符串：
```go
//声明请求体为空字符串
var req_body = ""
body_reader := strings.NewReader(req_body)
//声明查询参数  k_v_map
var user_pssword = map[string][]string{
    "name":[]string{"wy","wxy"}
}
//声明post地址，即base_url
var base_url = "1techbook.com" 
//发送请求
response , err := Gopost_full(url  , user_psswor , body_reader )
//返回response为字符串，提供了函数Html_utf8实现乱码的utf8转换
```

## 简单的post命令发送 :high_brightness:

```go
func Gopost_k_v_urlencode(url string , key []string , value []string ) (string , error)
``` 
-  url：  post目标地址，即base_url
-  key和value:   post参数：键key，以切片的形式与值value对一一对应



## gbk -> utf8 for strings ,  转换gbk编码到utf8编码，解决乱码，感谢mahonia:high_brightness:
```go
func Html_utf8(resp string)(string)
```
