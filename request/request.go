package request

import (
	"crypto/tls"
	"github.com/spf13/viper"
	"net/http"
	"runtime"
	"time"
)

func New() *http.Client {

	tr := &http.Transport{
		MaxIdleConns:      runtime.NumCPU() * 250,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		IdleConnTimeout:   time.Duration(viper.GetInt("request.conntimeout")) * time.Microsecond,
		DisableKeepAlives: true,
	}
	return &http.Client{
		Transport: tr,
		Timeout:   time.Duration(viper.GetInt("request.timeout")) * time.Microsecond,
	}
}
