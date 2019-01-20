package logger

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

type fakeLogger struct {
	sb strings.Builder
}

func (fl *fakeLogger) Println(v ...interface{}) {
	fmt.Fprintln(&fl.sb, v...)
}

func (fl *fakeLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(&fl.sb, format, v...)
}

func TestDemoV3(t *testing.T) {
	var fl fakeLogger
	DemoV3(fl.Println)
	want := "error in doTheThing():"
	got := fl.sb.String()
	//fmt.Println("got:", got)
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV4(t *testing.T) {
	var fl fakeLogger
	DemoV4(&fl)
	want := "error in doTheThing():"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV5(t *testing.T) {
	var sb strings.Builder
	testLogger := log.New(&sb, "", 0)
	thing := Thing{
		Logger: testLogger,
	}

	thing.DemoV5()
	want := "error in doTheThing():"
	got := sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV7(t *testing.T) {
	var sb strings.Builder
	testLogger := log.New(&sb, "", 0)

	thing := ThingV2{
		Logger: testLogger,
	}
	thing.DemoV7()
	want := "error in doTheThing():"
	got := sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}

	thing = ThingV2{}
	thing.DemoV7()
	want = "error in doTheThing():"
	got = sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}
