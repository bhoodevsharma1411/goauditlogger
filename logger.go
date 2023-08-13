package goauditlogger

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberLoggingFormat struct {
	TagPid               string `json:"pid,omitempty"`
	TagTime              string `json:"time,omitempty"`
	TagReferer           string `json:"referer,omitempty"`
	TagProtocol          string `json:"protocol,omitempty"`
	TagPort              string `json:"port,omitempty"`
	TagIP                string `json:"ip,omitempty"`
	TagIPs               string `json:"ips,omitempty"`
	TagHost              string `json:"host,omitempty"`
	TagMethod            string `json:"method,omitempty"`
	TagPath              string `json:"path,omitempty"`
	TagURL               string `json:"url,omitempty"`
	TagUA                string `json:"ua,omitempty"`
	TagLatency           string `json:"latency,omitempty"`
	TagStatus            string `json:"status,omitempty"`  // response status
	TagResBody           string `json:"resBody,omitempty"` // response body
	TagReqHeaders        string `json:"reqHeaders,omitempty"`
	TagQueryStringParams string `json:"queryParams,omitempty"` // request query parameters
	TagBody              string `json:"body,omitempty"`        // request body
	TagBytesSent         string `json:"bytesSent,omitempty"`
	TagBytesReceived     string `json:"bytesReceived,omitempty"`
	TagRoute             string `json:"route,omitempty"`
	TagError             string `json:"error,omitempty"`
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
