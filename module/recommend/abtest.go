package recommend

import (
	"hash/adler32"
	"strconv"
	"strings"
)

var abtest map[int]string

func getABtest(uuid string) string {
	unum := adler32.Checksum([]byte(uuid))
	umod := int(unum % 10)
	return abtest[umod]
}

func SetABtest(abctrl string) {
	abtest = make(map[int]string)
	abs := strings.Split(abctrl, ";")
	for _, ab := range abs {
		t := strings.Split(ab, "-")[0]
		for i := 0; i < 10; i++ {
			if strings.Contains(ab, strconv.Itoa(i)) {
				abtest[i] = t
			}
		}
	}
}
