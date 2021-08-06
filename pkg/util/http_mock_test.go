package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMockClient_Delete(t *testing.T) {
	client := MockClient{client: http.Client{}}
	rs, err := client.Get("http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	res, _ := ioutil.ReadAll(rs.Body)
	fmt.Println(string(res))
}
