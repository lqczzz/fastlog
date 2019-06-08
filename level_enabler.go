package fastlog

import (
	"go.uber.org/zap/zapcore"
)

// LevelEnabler is default log level enabler
type LevelEnabler struct {
	outLevel *zapcore.Level
	minLevel zapcore.Level
	maxLevel zapcore.Level
	logLevel zapcore.Level
}

func (l *LevelEnabler) Enabled(level zapcore.Level) bool {
	//fmt.Println(level, *l.outLevel <= l.minLevel, l.minLevel <= level && level <= l.maxLevel, *l.outLevel, l.logLevel)
	if *l.outLevel <= l.minLevel && // enable loggger
		l.minLevel <= level && level <= l.maxLevel { // log to right logger
		return true
	}
	return false
}

func NewLevelEnabler(outLevel *zapcore.Level, logger_level zapcore.Level) *LevelEnabler {
	var (
		minLevel zapcore.Level
		maxLevel zapcore.Level
	)

	switch logger_level {
	case zapcore.DebugLevel:
		minLevel = zapcore.DebugLevel
		maxLevel = zapcore.DebugLevel
	case zapcore.InfoLevel:
		minLevel = zapcore.InfoLevel
		maxLevel = zapcore.WarnLevel
	case zapcore.ErrorLevel:
		minLevel = zapcore.ErrorLevel
		maxLevel = zapcore.FatalLevel
	}

	return &LevelEnabler{
		outLevel: outLevel,
		minLevel: minLevel,
		maxLevel: maxLevel,
		logLevel: logger_level,
	}
}
