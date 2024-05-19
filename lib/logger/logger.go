package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	pec := zap.NewProductionEncoderConfig()
	pec.EncodeLevel = zapcore.CapitalLevelEncoder
	pec.EncodeCaller = zapcore.FullCallerEncoder
	pec.CallerKey = "cs"
	pec.EncodeDuration = zapcore.SecondsDurationEncoder
	pec.EncodeTime = utcISO8601Encoder

	enc := zapcore.NewJSONEncoder(pec)

	core := zapcore.NewCore(enc, zapcore.Lock(os.Stdout), zap.InfoLevel)

	return zap.New(core, zap.AddCaller())
}

func utcISO8601Encoder(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
	zapcore.ISO8601TimeEncoder(t.UTC(), pae)
}
