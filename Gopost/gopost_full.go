package Gopost

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Gopost_full(url string,k_v_map map[string][]string,body_reader io.Reader) (string,error) {

	//构造 base-url
	temp := url + "?"


	//构造查询参数
	for k ,v := range k_v_map {
		for i:=0;i<len(v);i++{
			temp+= ( k + "=" + v[i]+"&" )
		}
	}
	//去除最后的&符号
	temp=strings.TrimRight(temp,"&")


	fmt.Printf(temp+"\n")
	req, err := http.NewRequest("POST", temp, body_reader)

	if err != nil {
		//fmt.Print("生成req失败，正在重试")
		//fmt.Println(err)
		return "", err
	} else {
		//设置body的编码
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		//防止反爬-模拟chorme
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


	} else {
		var bytes, _ = ioutil.ReadAll(respones.Body)
		reponses2str := string(bytes)
		fmt.Print(reponses2str)
		return reponses2str,nil

	}
	return "",err_req

};
