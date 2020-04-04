package Gopost
import "github.com/axgle/mahonia"


//gbk -> utf8 for strings
func Html_utf8(resp string)(string){
	utf8_str:=mahonia.NewDecoder("gbk").ConvertString(resp)
	return utf8_str
}