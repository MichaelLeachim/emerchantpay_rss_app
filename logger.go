// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-19 21:48 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	logrus "github.com/sirupsen/logrus"
	"testing"
)

type logger interface {
	Warn(args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
}

type mockLogger struct{}

func newMockLogger() logger {
	return mockLogger{}
}

func (m mockLogger) Error(args ...interface{}) {
	return
}
func (m mockLogger) Warn(args ...interface{}) {
	return
}

func (m mockLogger) Info(args ...interface{}) {
	return
}

type testLogger struct {
	t *testing.T
}

func newTestLogger(t *testing.T) logger {
	return testLogger{t: t}
}

func (m testLogger) Error(args ...interface{}) {
	m.t.Log(args...)
	return
}

func (m testLogger) Warn(args ...interface{}) {
	m.t.Log(args...)
}
func (m testLogger) Info(args ...interface{}) {
	m.t.Log(args...)
}

type logrusLogger struct {
}

func newLogrusLogger() logger {
	return logrusLogger{}
}

func (m logrusLogger) Error(args ...interface{}) {
	logrus.Error(args...)

}

func (m logrusLogger) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func (m logrusLogger) Info(args ...interface{}) {
	logrus.Info(args...)

}
