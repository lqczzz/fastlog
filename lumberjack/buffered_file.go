package lumberjack

import (
	"bufio"
	"os"
	"sync"
	"time"
)

// BufferedFile is buffer writer than can be reopned
type BufferedFile struct {
	mutex     sync.Mutex
	quitChan  chan bool
	done      bool
	OrigRLog  *os.File
	BufWriter *bufio.Writer
}

// BufferedOptions is otions for BufferedFile
type BufferedOptions struct {
	bufferSize    int
	flushInterval time.Duration
}

// func (opts *bufferedOptions)

var (
	// defaultBufferSize exists so it can be mocked out by tests
	defaultBufferSize = 32 * 1024 * 1024
	// defaultFlushInterval exists so it can be mocked out by tests
	defaultFlushInterval = 1 * time.Second
)

// NewBufferedFile opens a buffered file that is periodically flushed.
func NewBufferedFile(rl *os.File, opts *BufferedOptions) *BufferedFile {
	brl := BufferedFile{
		quitChan: make(chan bool, 1),
		OrigRLog: rl,
	}
	brl.init(opts)
	return &brl
}

func (brl *BufferedFile) init(opts *BufferedOptions) {
	if opts == nil {
		opts = &BufferedOptions{
			bufferSize:    defaultBufferSize,
			flushInterval: defaultFlushInterval,
		}
	}
	if opts.bufferSize == 0 {
		opts.bufferSize = defaultBufferSize
	}
	if opts.flushInterval <= 0 {
		opts.flushInterval = defaultFlushInterval
	}

	brl.BufWriter = bufio.NewWriterSize(brl.OrigRLog, opts.bufferSize)

	go brl.flushDaemon(opts.flushInterval)
}

// flushDaemon periodically flushes the log file buffers
func (brl *BufferedFile) flushDaemon(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-brl.quitChan:
			ticker.Stop()
			return
		case <-ticker.C:
			brl.Flush()
		}
	}
}

// Flush flushes the buffer
func (brl *BufferedFile) Flush() {
	brl.mutex.Lock()

	if brl.done {
		brl.mutex.Unlock()
		return
	}

	brl.BufWriter.Flush()
	brl.OrigRLog.Sync()
	brl.mutex.Unlock()
}

// Write implements io.WriteCloser
func (brl *BufferedFile) Write(p []byte) (int, error) {
	brl.mutex.Lock()
	n, err := brl.BufWriter.Write(p)

	// means flush happenede in the middle of the line
	// and we need to flush the rest of our string at this point
	if brl.BufWriter.Buffered() < len(p) {
		brl.BufWriter.Flush()
	}

	brl.mutex.Unlock()
	return n, err
}

// Close flushes the internal buffer and closes the destination file
func (brl *BufferedFile) Close() error {
	brl.quitChan <- true
	brl.mutex.Lock()

	brl.done = true
	brl.BufWriter.Flush()

	err := brl.OrigRLog.Close()
	if err != nil {
		return err
	}

	brl.mutex.Unlock()
	return nil
}

func (brl *BufferedFile) Stat() (os.FileInfo, error) {
	return brl.OrigRLog.Stat()
}
