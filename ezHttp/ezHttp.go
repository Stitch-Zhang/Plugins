package Ezhttp

//Ezhttp is able to simplify http action based on GO

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

//TargetSite is about to drcraibe a GET action Parms
type TargetSite struct {
	URL          string
	Header 		map[string]string
	UA           string
	RegEnabled   bool
	RegExp       string
	ProxyEnabled bool
	Proxy        string
	Times        int
}

//EzRespond is easy http respond result
type EzRespond struct {
	Body       string
	Selected   []interface{}
	StatusCode int
	ConsistExp bool
}

//Do is Get action
//t is times
func (ts *TargetSite) Do() *EzRespond {
	ezRespond := new(EzRespond)
	if !strings.ContainsAny(ts.URL,"http"){
		log.Fatal("Typed Url didnt contain http or https")
	}
	for i := 0; i < ts.Times+1; i++ {
		//HTTP ACTION
		var proxy func(_ *http.Request) (*url.URL, error)
		//If enabled proxy then proxy config will add to http.Transport
		if ts.ProxyEnabled {
			proxy = func(_ *http.Request) (*url.URL, error) {
				return url.Parse("http://" + ts.Proxy)
			}
		}
		tr := &http.Transport{Proxy: proxy}
		client := &http.Client{Transport: tr}
		target, err := http.NewRequest("GET", ts.URL, nil)
		target.Header.Set("User-Agent", ts.UA)
		for i:=0;i<len(ts.Header);i++{
			for k,v :=range ts.Header{
				target.Header.Set(k,v)
			}
		}
		if err != nil {
			fmt.Println("Gente target site wrong", err)
		}
		resp, err := client.Do(target)
		if err != nil {
			fmt.Println("Get action target site wrong", err)
		}
		//DATA Select
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read resp wrong!", err)
		}
		ezRespond.Body = string(body)
		ezRespond.StatusCode = resp.StatusCode
		fmt.Println("Get status is successful")
		//Regexp function if need

		if ts.RegEnabled {
			findExp := regexp.MustCompile(ts.RegExp)
			result := findExp.FindAllStringSubmatch(string(body), -1)
			if len(result) == 0 {
				log.Println("Cant found any data in consist Expression")
				ezRespond.ConsistExp = false
			} else {
				//Default ignore command (result[0][1])
				for i := 0; i < len(result); i++ {
					ezRespond.Selected = append(ezRespond.Selected, result[i][1])
					ezRespond.ConsistExp = true
				}
			}
		}
		return ezRespond
	}
	defer fmt.Println("Totally Gets:", ts.Times)
	return ezRespond

}
