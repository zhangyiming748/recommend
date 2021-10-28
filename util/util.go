package util

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Marshal2String(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TimeParse(layout, value string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	_time, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return time.Time{}, err
	}
	//Debugf("_time:%s\r\n",_time)
	return _time, err
}

func String2MD5(data string) string {
	input := []byte(data)
	sum := md5.Sum(input)
	res := ""
	for _, ch := range sum {
		res += fmt.Sprintf("%02x", ch)
	}
	return res
}

func StringInSlice(s string, slist []string) bool {
	for _, v := range slist {
		if s == v {
			return true
		}
	}
	return false
}

func SubString(s string, begin, end int) string {
	sbyte := []byte(s)
	if end > len(s) {
		end = len(s)
	}
	if begin > end {
		begin = end
	}
	if begin <= 0 {
		begin = 0
	}
	return string(sbyte[begin:end])
}

func CookieString(c *http.Cookie) string {
	if c == nil {
		return ""
	}
	var b bytes.Buffer
	if len(c.Path) > 0 {
		b.WriteString(";Path=")
		b.WriteString(c.Path)
	}
	if len(c.Domain) > 0 {
		if c.Domain != "" {
			d := c.Domain
			b.WriteString("; Domain=")
			b.WriteString(d)
		} else {
			Infof("net/http: invalid Cookie.Domain %q; dropping domain attribute", c.Domain)
		}
	}
	if c.MaxAge > 0 {
		b.WriteString("; Max-Age=")
		b2 := b.Bytes()
		b.Reset()
		b.Write(strconv.AppendInt(b2, int64(c.MaxAge), 10))
	} else if c.MaxAge < 0 {
		b.WriteString("; Max-Age=0")
	}
	if c.HttpOnly {
		b.WriteString("; HttpOnly")
	}
	if c.Secure {
		b.WriteString("; Secure")
	}
	return b.String()
}

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// 将ip装成16进制
func Ip2hex(ipstr string) string {
	ds := strings.Split(ipstr, ".")
	ds0, _ := strconv.Atoi(ds[0])
	s := fmt.Sprintf("%x", ds0)
	for i := 1; i < 4; i++ {
		dsi, _ := strconv.Atoi(ds[i])
		s += fmt.Sprintf("%x", dsi)
	}
	return s
}

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
