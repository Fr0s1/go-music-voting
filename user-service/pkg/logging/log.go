package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Log            = logrus.New()
	HostName, _    = os.Hostname()
	Env            = os.Getenv("ENV")
	Service        = "voting-grpc-service"
	StandardFields = logrus.Fields{
		"hostname": HostName,
		"env":      Env,
		"service":  Service,
	}
)

func init() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetReportCaller(true)
	Log.WithFields(StandardFields)
}
