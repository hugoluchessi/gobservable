package logging

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

func TestNewLoggerAndDebug(t *testing.T) {
	var b bytes.Buffer
	exctx := exctx.Create()
	msg := "New logger!"

	f := NewDefaultFormatter()
	l := NewLogger(&b, f)

	l.Debug(exctx, msg)
	time.Sleep(50 * time.Millisecond)

	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) D - " + msg
	iomsg := b.String()
	match, _ := regexp.MatchString(expectedRegex, iomsg)

	if !match {
		t.Errorf("Test failed, got '%s'", iomsg)
	}
}

func CreateTestLogger(b *bytes.Buffer) *Logger {
	l := NewDefaultLogger(b)
	return l
}

func validateMsgRegex(t *testing.T, expectedRegex string, msg string) {
	match, _ := regexp.MatchString(expectedRegex, msg)

	if !match {
		t.Errorf("Test failed, got '%s'", msg)
	}
}

func TestDefaultLoggerLog(t *testing.T) {
	exctx := exctx.Create()
	var b bytes.Buffer
	l := CreateTestLogger(&b)
	msg := "New logger!"

	l.Log(exctx, msg)

	time.Sleep(50 * time.Millisecond)

	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) L - " + msg
	iomsg := b.String()
	validateMsgRegex(t, expectedRegex, iomsg)
}

func TestDefaultLoggerWarn(t *testing.T) {
	exctx := exctx.Create()
	var b bytes.Buffer
	l := CreateTestLogger(&b)
	msg := "New logger!"

	l.Warn(exctx, msg)

	time.Sleep(50 * time.Millisecond)

	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) W - " + msg
	iomsg := b.String()
	validateMsgRegex(t, expectedRegex, iomsg)
}

func TestDefaultLoggerError(t *testing.T) {
	exctx := exctx.Create()
	var b bytes.Buffer
	l := CreateTestLogger(&b)
	msg := "New logger!"

	l.Error(exctx, msg)

	time.Sleep(50 * time.Millisecond)

	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) E - " + msg
	iomsg := b.String()
	validateMsgRegex(t, expectedRegex, iomsg)
}

func TestDefaultLoggerFatal(t *testing.T) {
	exctx := exctx.Create()
	var b bytes.Buffer
	l := CreateTestLogger(&b)
	msg := "New logger!"

	l.Fatal(exctx, msg)

	time.Sleep(50 * time.Millisecond)

	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) F - " + msg
	iomsg := b.String()
	validateMsgRegex(t, expectedRegex, iomsg)
}

func TestAsyncWrites(t *testing.T) {
	var b bytes.Buffer
	var wg sync.WaitGroup

	l := CreateTestLogger(&b)
	countRoutines := 10000

	wg.Add(countRoutines)

	for i := 0; i < countRoutines; i++ {
		go writeLogs(l, i, &wg)
	}

	wg.Wait()

	time.Sleep(30 * time.Millisecond)

	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) \\w{1} - Logger \\d{1,2}"
	rawmsgs := b.String()
	msgs := strings.Split(rawmsgs, "\r\n")

	if len(msgs)-1 != countRoutines*5 {
		t.Errorf("Test failed, expected '%d' messages, got '%d'", countRoutines*5, len(msgs))
	}

	for _, msg := range msgs {
		if len(msg) == 0 {
			continue
		}
		validateMsgRegex(t, expectedRegex, msg)
	}
}

func writeLogs(l *Logger, i int, wg *sync.WaitGroup) {
	exctx := exctx.Create()
	msg := "Logger " + strconv.Itoa(i)

	l.Debug(exctx, msg)
	l.Log(exctx, msg)
	l.Warn(exctx, msg)
	l.Error(exctx, msg)
	l.Fatal(exctx, msg)
	wg.Done()
}

func BenchmarkBigText(b *testing.B) {
	var buffer bytes.Buffer
	l := CreateTestLogger(&buffer)
	exctx := exctx.Create()
	msg := "Logger benchmark. Trying to put a big text here. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam consequat bibendum auctor. Vivamus ac ipsum tempor, porta elit at, blandit urna. Suspendisse lacus odio, dictum non justo nec, viverra convallis urna. Ut risus enim, egestas in risus ut, elementum volutpat nibh. In maximus placerat enim sed euismod. Proin id dolor nec ligula malesuada bibendum. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla et ornare arcu. Praesent turpis erat, elementum eget iaculis vitae, dignissim sit amet justo. Sed tempus at sem vel sollicitudin. Integer purus felis, facilisis eu purus in, rhoncus pellentesque justo. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras dictum neque nec purus faucibus, sit amet mattis urna ultrices. Fusce ut nibh malesuada, pulvinar massa quis, efficitur leo. Ut varius consectetur nunc vitae finibus. In in aliquet mauris. Sed nec dolor dolor. Nullam turpis nisl, semper eu mauris quis, pretium interdum lectus. Donec placerat lacus est, eu porta orci volutpat sed. Etiam molestie nisl in turpis aliquet posuere. Praesent aliquet sagittis neque, non ornare tortor iaculis non. Nunc et suscipit diam. Aliquam finibus libero felis, quis volutpat mauris bibendum consectetur. Etiam eu odio id libero mollis varius. Aenean aliquam ante metus. Vivamus pellentesque enim justo, vel dapibus sem varius in. Ut ligula mauris, interdum nec bibendum at, volutpat vitae velit. Aliquam at mauris sed ex pretium mollis non imperdiet nisl. Etiam maximus interdum ultricies. Nulla id fringilla mauris. Aliquam erat volutpat. Integer aliquet pellentesque enim, sed vehicula nisl venenatis at. Sed porttitor purus lorem, quis scelerisque lacus rhoncus in. Nam vehicula, tellus at dapibus imperdiet, neque elit consectetur leo, nec maximus ligula sapien non purus. Vestibulum et sollicitudin leo. Ut tincidunt mi ac sodales fringilla. Suspendisse egestas sit amet nunc sit amet vulputate. Fusce lacinia, quam mollis malesuada scelerisque, diam dui eleifend sem, at mollis ipsum magna non magna. Vivamus eget mauris vehicula, fermentum quam in, dignissim risus. Quisque eu fermentum tortor. Maecenas vitae faucibus dui. Aliquam consequat, purus sed iaculis consectetur, nisl tellus iaculis mauris, sed imperdiet lorem diam vitae velit. Donec vitae blandit ante. Phasellus imperdiet efficitur euismod. Etiam fringilla luctus consectetur. Fusce libero justo, vehicula quis porttitor nec, tempor et quam. Morbi neque dolor, commodo quis nisl in, commodo dignissim elit. Curabitur hendrerit enim libero, sed placerat velit sollicitudin at. Integer vestibulum maximus mauris, vitae volutpat sapien sollicitudin sit amet. Mauris ullamcorper, eros eget elementum vulputate, odio lectus tempus augue, ac vehicula felis nisl sit amet eros. Aliquam vitae efficitur dui, at lacinia dui. Etiam facilisis imperdiet nisl, eu finibus lacus bibendum et. Pellentesque cursus eu lorem sit amet pulvinar. Sed congue tortor id lectus congue egestas. Praesent dictum est quis nulla accumsan malesuada. Aliquam laoreet aliquet lorem, sit amet dapibus lectus rutrum non. In dictum pellentesque magna, quis sollicitudin lectus ullamcorper eget. Mauris molestie id odio semper pharetra. Vivamus laoreet nulla ac vehicula accumsan."

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		l.Debug(exctx, msg)
	}
}

func BenchmarkLittleText(b *testing.B) {
	var buffer bytes.Buffer
	l := CreateTestLogger(&buffer)
	exctx := exctx.Create()
	msg := "Logger benchmark."

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		l.Debug(exctx, msg)
	}
}
