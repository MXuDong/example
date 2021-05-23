package nets

import (
	"github.io/MXuDong/example/pkg/constant"
	"strings"
)

// GetHttpUrl will check the url is start with http, and try to append http to head.
// If input is www.baidu.com, it will return false, http://www.example.com
// If input is http://example.com, it will return true, http://example.com
// If input is h://www.example.com, it will return false, http://www.example.com
// The check flag is '://'
// But if input is wrong, can't parse to http protocol, return false and ""
func GetHttpUrl(url string) (bool, string) {
	urls := strings.Split(url, constant.RequestUrlProtocolFlag)
	lens := len(urls)
	// ://example, SomeProtocol://example, and not the nil input
	if (lens == 1 || lens == 2) && urls[lens-1] != "" {
		// The www.example.com, or http://www.example.com or someProtocol://www.example.com
		// All of these input, return is must one
		isRight := true
		if urls[0] != constant.HttpProtocol {
			isRight = false
		}
		return isRight, strings.Join([]string{constant.HttpProtocol, urls[lens-1]}, constant.RequestUrlProtocolFlag)
	}

	// is more than one, is wrong input
	return false, ""
}
