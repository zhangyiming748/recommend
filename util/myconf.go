package util

import (
	"bufio"
	"strings"

	"io"
	"os"
)

type Config struct {
	Redis string
	Elastic string
}

func (c *Config) SetRedis(u string) {
	c.Redis= u
	return
}
func (c *Config) SetES(u string) {
	c.Elastic = u
	return
}
func ReadConfig(f string) *Config {
	var (
		c     Config
		redis string
		es    string
	)
	fi, err := os.Open(f)
	if err != nil {
		//fmt.Printf("Error: %s\n", err)
		panic(err)
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		s := string(a)
		if strings.HasPrefix(s, "redis") {
			redis = s
		}
		if strings.HasPrefix(s, "elastic") {
			es = s
		}
	}
	c.SetRedis(redis)
	c.SetES(es)

	return &c
}
