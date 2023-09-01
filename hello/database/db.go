package database

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func BuscaPorDoc(doc int) (string, error) {
	//doc := 12310400000182

	url := "http://localhost:3004/clientes?doc=" + strconv.Itoa(doc)
	method := "GET"

	//fmt.Println(url)

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		//return

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return

	}
	fmt.Println(string(body))

	return string(body), nil
}
