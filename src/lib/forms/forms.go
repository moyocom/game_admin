package forms

import (
	"crypto/md5"
	"fmt"
	"io"
	. "lib/Util"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func Reg_email(email string) bool {
	if len(email) > 40 {
		return false
	}
	remail, err := regexp.Compile(`^[\w\-\.]+@[\w\-\.]+(\.\w+)+$`)
	if err != nil {
		panic(err)
	}
	return remail.MatchString(email)
}

func Reg_user(user string) bool {
	ruser, err := regexp.Compile("^[a-zA-Z0-9_\u2E80-\u9FFF]{1,12}$")
	if err != nil {
		panic(err)
	}
	return ruser.MatchString(user)
}

// 加密密码,转成md5
func Tomd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func IsEmpty(w http.ResponseWriter, strs ...string) bool {
	for _, v := range strs {
		if v == "" {
			ExitMsg(w, "不可为空")
			return true
		}
	}
	return false
}

//字符串转换成数字
func Toint(num interface{}) int {
	rnum, err := strconv.Atoi(num.(string))
	CheckErr(err)
	return rnum
}

func Unique(ar []string) []string {
	if ar == nil || len(ar) < 2 {
		return ar
	}
	var rar []string
	for _, v := range ar {
		if !Ishave(rar, v) {
			v = strings.Trim(v, " ")
			if v == "" {
				continue
			}
			rar = append(rar, v)
		}
	}
	return rar
}

func Ishave(ar []string, str string) bool {
	for _, v := range ar {
		if v == str {
			return true
		}
	}
	return false
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func Ip(r *http.Request) string {
	ipinfo := strings.Split(r.RemoteAddr, ":")
	return ipinfo[0]
}
