// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configrotate

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	testLogFileName = "test.log"
)

func TestRotationEnabledCreate(t *testing.T) {
	maxMegabytes := 1524
	maxDays := 145
	maxBackupds := 155
	localTime := true
	// this test will not try to validate, create or write file.
	filename := testLogFileName
	rotateConfig := Config{
		Enabled:      true,
		MaxMegabytes: maxMegabytes,
		MaxDays:      maxDays,
		MaxBackups:   maxBackupds,
		LocalTime:    localTime,
	}
	writer, err := rotateConfig.NewWriter(filename)
	assert.NoError(t, err)
	rotate, ok := writer.(*lumberjack.Logger)
	assert.True(t, ok)
	assert.Equal(t, rotate.Filename, filename)
	assert.Equal(t, rotate.MaxSize, maxMegabytes)
	assert.Equal(t, rotate.MaxAge, maxDays)
	assert.Equal(t, rotate.LocalTime, localTime)
}

func TestRotateDisabledCreate(t *testing.T) {
	rotateConfig := Config{
		Enabled: false,
	}
	tempDir := t.TempDir()
	filename := path.Join(tempDir, testLogFileName)
	writer, err := rotateConfig.NewWriter(filename)
	assert.NoError(t, err)
	file, ok := writer.(*os.File)
	assert.True(t, ok)
	assert.Equal(t, file.Name(), filename)
	assert.NoError(t, file.Close())
}
