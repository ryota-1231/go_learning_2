package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	res, err := http.Get("http://example.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	base, err1 := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	if err1 != nil {
		fmt.Println("---------error--------")
		log.Fatalln(err1)
	}
	fmt.Println(base)
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyz"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resq, _ := client.Do(req)
	body1, _ := ioutil.ReadAll(resq.Body)
	fmt.Println(string(body1))

	b := []byte(`{"name":"mike", "age":"20", "height":0, "nicknames":["a","b","c"]}`)
	var p Person
	if err2 := json.Unmarshal(b, &p); err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(p.Name, p.Age, p.Height, p.Nicknames)

	//jsonに変換する
	v, _ := json.Marshal(p)

	fmt.Printf("%T\n", v)
	fmt.Println(string(v))

	const (
		apiKey       = "User1Key"
		apiSecretKey = "User1SecretKey"
	)

	mac := hmac.New(sha256.New, []byte(apiSecretKey))
	message := []byte("data")
	mac.Write(message)
	// expectedMAC := mac.Sum(nil)
	// fmt.Println(expectedMAC)
	sign := hex.EncodeToString(mac.Sum(nil))
	fmt.Println(sign)

}

// json.MarshalをカスタマイズしたいときはMarshalJSONという関数名で定義する
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string `json:"name"`
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

// json.UnmarshalをカスタマイズしたいときはUnmarshalJSONという関数名で定義する
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

type Person struct {
	Name string `json:"name"`
	//unmarshalではintで表示、marshalではstring表示
	Age       int      `json:"age,string"`
	Height    int      `json:"height,omitempty"`
	Nicknames []string `json:"-"`
}
