package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	loggerDev, _ := zap.NewDevelopment()
	defer loggerDev.Sync()

	loggerProd, _ := zap.NewProduction()
	defer loggerProd.Sync()

	conf := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "./log/development.out.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	loggerCostom, _ := conf.Build()

	sugarDev := loggerDev.Sugar()
	sugarProd := loggerProd.Sugar()
	sugarCostom := loggerCostom.Sugar()

	url := "http://example.com"

	loggerDev.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	sugarDev.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	fmt.Println("")

	sugarProd.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	loggerProd.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	fmt.Println("")

	loggerCostom.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	sugarCostom.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	fmt.Println("")

	user := &user{
		Name:      "Zap",
		Email:     "zap@sample.com",
		CreatedAt: time.Now(),
	}
	loggerCostom.Info("object sample", zap.Object("userObj", user))
}

type user struct {
	Name      string
	Email     string
	CreatedAt time.Time
}

func (u user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.Name)
	enc.AddString("email", u.Email)
	enc.AddInt64("created_at", u.CreatedAt.UnixNano())
	return nil
}
