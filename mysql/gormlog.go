package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

var _ logger.Interface = (*GormLogger)(nil)

type GormLogger struct {
	logger        logx.Logger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// func GetGormLogger() *GormLogger {
// 	return &GormLogger{
// 		// logger:,
// 	}
// }
var (
	LogxGorm = GormLogger{
		LogLevel: logger.Warn,
		// logger:   logx.Logger,
		SlowThreshold: 200 * time.Millisecond,
	}
)

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newloger := *l
	newloger.LogLevel = level
	return &newloger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		logx.WithContext(ctx).Info(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	}
}
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		logx.WithContext(ctx).Error(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	}

}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	// logx.WithContext(ctx).Error(s, i)
	if l.LogLevel >= logger.Warn {
		logx.WithContext(ctx).Info(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	}

}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				logx.WithContext(ctx).Error(utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				logx.WithContext(ctx).Error(utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				logx.WithContext(ctx).Info(utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				logx.WithContext(ctx).Info(utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				logx.WithContext(ctx).Info(utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				logx.WithContext(ctx).Info(utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			}
		}
	}

}

// func WarpGormLogger(logger logx.Logger) GormLogger {
// 	return GormLogger{logger}
// }

// func (logger GormLogger) LogMode(logger.LogLevel) logger.Interface {
// 	return logger
// }

// func (logger GormLogger) Info(ctx context.Context, format string, args ...interface{}) {
// 	logger.logger.Info(format, args)
// }

// func (logger GormLogger) Warn(ctx context.Context, format string, args ...interface{}) {
// 	logger.logger.Info(format, args)
// }

// func (logger GormLogger) Error(ctx context.Context, format string, args ...interface{}) {
// 	logger.logger.Info(format, args)

// // func (logger GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
// // 	elapsed := time.Since(begin)
// // 	sql, row := fc()
// // 	logger.logger.Trace().Dur("elapsed", elapsed).Str("sql", sql).Int64("row", row).Err(err).Msg("trace sql")
// // }

// // func (logger GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
// // 	return
// // }

// func WarpGormLogger(logger logx.Logger) GormLogger {
// 	return GormLogger{logger}
// }
