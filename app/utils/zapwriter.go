package utils

import "github.com/KaitoXCode/fiberApp/app/services"

type ZapWriter struct{}

func (zw *ZapWriter) Write(p []byte) (n int, err error) {
	services.Logger.Info(string(p))
	return len(p), nil
}
