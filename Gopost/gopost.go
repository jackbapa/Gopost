package Gopost

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Gopost_k_v_urlencode(url string,key []string ,value []string) (string,error) {

	//构建 base-url
	url0 := url + "?"

	//在base-url后面添加查询参数
	var par string
	i := 0
	for k := range key {

		temp:=key[k] + "=" + value[i]
		if i==0{
			par=temp
		}else{
			par+="&"+temp

		}
		i++
	}

	//输出查询参数
	//fmt.Printf(par+"\n")

	//把查询参数写入body
	req, err := http.NewRequest("POST", url0, strings.NewReader(par))

	if err != nil {
		//fmt.Print("生成req失败，正在重试")

		//返回错误
		return "",err
	} else {
		//设置请求头
		//urlencode编码
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		//默认仿chorme浏览器行为，防止反扒
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
		//fmt.Print("请求连接是\n")
		//fmt.Println(req.URL)
		//fmt.Print("\n查询参数是\n")
		//fmt.Println(req.URL.RawQuery)
		//fmt.Print("\n")
	}
	//忽略证书校验
	tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//把“忽略校验”加入客户端
	client:=&http.Client{
		Transport:tr,
	}
	respones, err_req := (client.Do(req))
	if err_req != nil {
		//fmt.Println(err_req)
		//fmt.Print("请求失败")
		//重新请求
		//Gopost_k_v_urlencode(url, key, value)
	} else {
		var bytes, _ = ioutil.ReadAll(respones.Body)
		reponses2str := string(bytes)
		fmt.Print(reponses2str)
		return reponses2str,nil

	}
	return "",err_req

};
