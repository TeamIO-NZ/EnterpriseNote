package note

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.iosoftworks.com/EnterpriseNote/pkg/web"
)

func Start() {
	// 1. parse config
	// 2. init logger
	err := initLogger(true)
	if err != nil {
		log.Fatalf("Failed to init Zap logger: %v", err)
		return
	}
	zap.S().Info("Logger initialized.")
	// 3. establish connection to postgresql
	// 4. start webserver
	web.Server{}.Start()
	// 5. wait for kill sig
	// 6. shutdown webserver
}

func initLogger(debug bool) (err error) {
	var cfg zap.Config
	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	l, err := cfg.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(l)
	return nil
}
