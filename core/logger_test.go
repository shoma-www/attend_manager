package core

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestOutputLog(t *testing.T) {
	t.Run("All Output", func(t *testing.T) {
		l := NewLogger(Debug)
		buf := &bytes.Buffer{}
		l.SetLogger(log.New(os.Stderr, "", log.Lmsgprefix))
		l.SetOutput(buf)
		l.Debug("Test Debug %s", "hoge")
		l.Info("Test Info %s", "hoge")
		l.Warn("Test Warn %s", "hoge")
		l.Error("Test Error %s", "hoge")

		wants := []string{
			"[DEBUG] Test Debug hoge\n",
			"[INFO ] Test Info hoge\n",
			"[WARN ] Test Warn hoge\n",
			"[ERROR] Test Error hoge\n",
		}

		for _, want := range wants {
			got, err := buf.ReadString('\n')
			if err != nil {
				t.Errorf("Can't be got actual value. %s\n", err)
			}
			if strings.Compare(want, got) != 0 {
				t.Errorf("Did not match value. want=%s, got=%s\n", want, got)
			}
		}
	})

	t.Run("Level extract by Info", func(t *testing.T) {
		l := NewLogger(Info)
		buf := &bytes.Buffer{}
		l.SetLogger(log.New(os.Stderr, "", log.Lmsgprefix))
		l.SetOutput(buf)
		l.Debug("Test Debug %s", "hoge")
		l.Info("Test Info %s", "hoge")
		l.Warn("Test Warn %s", "hoge")
		l.Error("Test Error %s", "hoge")

		wants := []string{
			"[INFO ] Test Info hoge\n",
			"[WARN ] Test Warn hoge\n",
			"[ERROR] Test Error hoge\n",
		}

		for _, want := range wants {
			got, err := buf.ReadString('\n')
			if err != nil {
				t.Errorf("Can't be got actual value. %s\n", err)
			}
			if strings.Compare(want, got) != 0 {
				t.Errorf("Did not match value. want=%s, got=%s\n", want, got)
			}
		}
	})

	t.Run("Level extract by warn", func(t *testing.T) {
		l := NewLogger(Warn)
		buf := &bytes.Buffer{}
		l.SetLogger(log.New(os.Stderr, "", log.Lmsgprefix))
		l.SetOutput(buf)
		l.Debug("Test Debug %s", "hoge")
		l.Info("Test Info %s", "hoge")
		l.Warn("Test Warn %s", "hoge")
		l.Error("Test Error %s", "hoge")

		wants := []string{
			"[WARN ] Test Warn hoge\n",
			"[ERROR] Test Error hoge\n",
		}

		for _, want := range wants {
			got, err := buf.ReadString('\n')
			if err != nil {
				t.Errorf("Can't be got actual value. %s\n", err)
			}
			if strings.Compare(want, got) != 0 {
				t.Errorf("Did not match value. want=%s, got=%s\n", want, got)
			}
		}
	})

	t.Run("Level extract by error", func(t *testing.T) {
		l := NewLogger(Error)
		buf := &bytes.Buffer{}
		l.SetLogger(log.New(os.Stderr, "", log.Lmsgprefix))
		l.SetOutput(buf)
		l.Debug("Test Debug %s", "hoge")
		l.Info("Test Info %s", "hoge")
		l.Warn("Test Warn %s", "hoge")
		l.Error("Test Error %s", "hoge")

		wants := []string{
			"[ERROR] Test Error hoge\n",
		}

		for _, want := range wants {
			got, err := buf.ReadString('\n')
			if err != nil {
				t.Errorf("Can't be got actual value. %s\n", err)
			}
			if strings.Compare(want, got) != 0 {
				t.Errorf("Did not match value. want=%s, got=%s\n", want, got)
			}
		}
	})
}
