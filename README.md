## fastlog

fastlog is a zap based and lumberjack based log utils.

fastlog customizes lumberjack for high performance throught buffering the data and flushing intervally.

### feature

- features all zap has
- features all lumberjack has
- buffered log data for high performance

### usage

    fastlog.Debug("debug msg")

    fastlog.Debug("debug msg", zap.String("name", "fastlog"), zap.Int("age", 1))

    // also, you can custom fastlog before you use it
    fastlog.Init(
        fastlog.Path("custom_log"),
    )

