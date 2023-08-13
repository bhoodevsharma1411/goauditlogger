package goauditlogger

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberLoggingFormat struct {
	TagPid               string `json:"pid"`
	TagTime              string `json:"time"`
	TagReferer           string `json:"referer"`
	TagProtocol          string `json:"protocol"`
	TagPort              string `json:"port"`
	TagIP                string `json:"ip"`
	TagIPs               string `json:"ips"`
	TagHost              string `json:"host"`
	TagMethod            string `json:"method"`
	TagPath              string `json:"path"`
	TagURL               string `json:"url"`
	TagUA                string `json:"ua"`
	TagLatency           string `json:"latency"`
	TagStatus            string `json:"status"`  // response status
	TagResBody           string `json:"resBody"` // response body
	TagReqHeaders        string `json:"reqHeaders"`
	TagQueryStringParams string `json:"queryParams"` // request query parameters
	TagBody              string `json:"body"`        // request body
	TagBytesSent         string `json:"bytesSent"`
	TagBytesReceived     string `json:"bytesReceived"`
	TagRoute             string `json:"route"`
	TagError             string `json:"error"`
}

func GetAuditLogger() fiber.Handler {
	logger_obj := logger.New(GetDefaultConfig())
	return logger_obj
}

func GetDefaultConfig() logger.Config {

	output := "./audit.log"
	custom_output := os.Getenv("LOGGER_OUTPUT")
	if len(custom_output) > 0 {
		output = custom_output
	}

	custome_format := os.Getenv("LOGGER_FORMAT")

	if len(custome_format) > 0 {
		File, _ := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		config := logger.Config{Format: string(custome_format) + "\n", Output: File}
		return config
	} else {
		default_format := FiberLoggingFormat{
			TagHost:              "${host}",
			TagPath:              "${path}",
			TagMethod:            "${method}",
			TagTime:              "${time}",
			TagStatus:            "${status}",
			TagLatency:           "${latency}",
			TagProtocol:          "${protocol}",
			TagIP:                "${ip}",
			TagURL:               "${url}",
			TagUA:                "${ua}",
			TagRoute:             "${route}",
			TagQueryStringParams: "${queryParams}",
		}

		logStr, _ := json.Marshal(default_format)
		File, _ := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		config := logger.Config{Format: string(logStr) + "\n", Output: File}
		return config
	}
}
