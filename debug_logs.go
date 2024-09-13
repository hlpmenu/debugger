package debugger

import "gopkg.hlmpn.dev/pkg/go-logger"

func (d *Debug) Printf(f string, args ...interface{}) {
	if d.IsDebug() {
		logger.Printf(f, args...)
	}
}

func (d *Debug) Print(s string) {
	if d.IsDebug() {
		logger.Print(s)
	}
}

func (d *Debug) PrintRed(s string) {
	if d.IsDebug() {
		logger.LogRed(s)
	}
}

func (d *Debug) LogPurple(s string) {
	if d.IsDebug() {
		logger.LogPurple(s)
	}
}
func (d *Debug) LogOrange(s string) {
	if d.IsDebug() {
		logger.LogOrange(s)
	}
}
func (d *Debug) LogRedf(f string, args ...interface{}) {
	if d.IsDebug() {
		logger.LogRedf(f, args...)
	}
}
func (d *Debug) LogPurplef(f string, args ...interface{}) {
	if d.IsDebug() {
		logger.LogPurplef(f, args...)
	}
}
func (d *Debug) LogOrangef(f string, args ...interface{}) {
	if d.IsDebug() {
		logger.LogOrangef(f, args...)
	}
}

func (d *Debug) LogErrorf(f string, args ...interface{}) {
	if d.IsDebug() {
		logger.LogErrorf(f, args...)
	}
}

func (d *Debug) LogError(s string) {
	if d.IsDebug() {
		logger.LogError(s)
	}
}

func (d *Debug) Fatal(s string) {
	if d.IsDebug() {
		logger.Fatal(s)
	}
}

func (d *Debug) Fatalf(f string, args ...interface{}) {
	if d.IsDebug() {
		logger.Fatalf(f, args...)
	}
}

func (d *Debug) Panic(s string) {
	if d.IsDebug() {
		logger.Panic(s)
	}
}

func (d *Debug) Panicf(f string, args ...interface{}) {
	if d.IsDebug() {
		logger.Panicf(f, args...)
	}
}

func (d *Debug) LogFatalRed(s string, args ...interface{}) {
	if d.IsDebug() {
		d.LogErrorf(s, args...)
	}
}
