package util

import (
	"net"
	"bytes"
	"compress/zlib"
	"io"
)

func GetIntranetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("ip:", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func ZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	r, _ := zlib.NewReader(b)

	if r != nil {
		var out bytes.Buffer
		defer r.Close()
		io.Copy(&out, r)
		return out.Bytes()
	}

	return nil
}