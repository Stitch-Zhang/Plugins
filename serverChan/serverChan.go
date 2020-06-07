package ServerChan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//This package is based on http://sc.ftqq.com/ site
//Simple action

//INFO is about your personal SCKEY
type INFO struct {
	//SCKEY is your key to push information
	SCKEY string
}


type serverRespond struct {
	Errno int `json:"errno"`
	Errmsg string `json:"errmsg"`
}

func (key *INFO)Push(text,desp string) (status bool){
	url :="https://sc.ftqq.com/"+key.SCKEY+".send?"
	resp,err:=http.Get(url+"text="+text+"&desp="+desp)
	if err!=nil{
		fmt.Println("Request to Send msg wrong!",err)
	}
	if resp==nil{
		fmt.Println("Get respond is nil\nMay be send failed!")
	}
	body,_ :=ioutil.ReadAll(resp.Body)
	var respond serverRespond
	err2 :=json.Unmarshal(body, &respond)
	if err2!=nil{
		fmt.Println("U marshal json wrong!",err2)
	}
	if respond.Errno==0{
		//fmt.Println("Sending msg success!")
		status = true
	} else{
		//fmt.Println("Sending msg wrong!")
		status =false
	}
	return

}


