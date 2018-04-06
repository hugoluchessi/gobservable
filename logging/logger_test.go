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

func Benchmark10000Chars(b *testing.B) {
	var buffer bytes.Buffer
	l := CreateTestLogger(&buffer)
	exctx := exctx.Create()
	msg := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean tincidunt libero a aliquam luctus. Morbi imperdiet sed mi et porta. Ut venenatis ultricies ipsum, sed rhoncus est dapibus quis. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus imperdiet mattis quam sit amet lobortis. Nulla facilisi. Praesent nec mauris augue. Nam accumsan turpis nisl, et elementum orci hendrerit eget. Etiam ut purus vulputate, luctus erat ac, blandit nisl. Nam consequat ac quam a dapibus. Nunc consectetur felis enim, sed fermentum magna euismod sed. Morbi ut diam tristique, posuere sem quis, faucibus sem. Quisque enim sapien, ornare at condimentum tempor, vehicula eu tellus. Ut porta venenatis nulla eu convallis. Suspendisse aliquam a nulla vitae tempus.Aenean vitae dui condimentum, faucibus lacus id, sagittis tellus. Donec quis tellus sed nisi vestibulum scelerisque eget non ante. Nulla blandit arcu a eros varius, at posuere nibh lacinia. Pellentesque ultricies interdum pellentesque. Morbi et quam blandit, suscipit arcu vitae, gravida nunc. Integer scelerisque leo sit amet placerat vulputate. Integer sit amet interdum odio. Curabitur pulvinar blandit magna sit amet dictum. Vivamus ut ex auctor, dictum purus vitae, condimentum sapien. Vestibulum non pretium lorem. Ut consectetur libero nec leo vestibulum, sed molestie dolor commodo. Fusce dui neque, eleifend ac orci vitae, elementum tristique sapien. In faucibus arcu quam, sit amet malesuada est blandit quis. Nulla pretium leo eget tortor posuere, vel aliquet erat scelerisque. Phasellus in ligula tortor. Integer accumsan dictum arcu, nec imperdiet risus consectetur vel. Nam fringilla eros nec lacus mattis, sed ultrices purus porta. Nam a condimentum nisi, non facilisis metus. In non egestas velit, sed gravida turpis. Aenean lacinia, mauris non hendrerit elementum, mi enim lobortis arcu, et elementum purus quam vitae leo. Aenean lobortis, est ac iaculis elementum, est lacus ultricies felis, vitae elementum orci mi non libero. Phasellus non nunc lobortis sem finibus euismod et eu dui. Morbi rhoncus congue felis, sit amet sagittis elit consectetur eu. Proin condimentum consectetur mollis. Etiam eget metus turpis. Maecenas erat nunc, pulvinar non bibendum non, congue sed ante. In et urna eu risus efficitur tincidunt et vel mauris. Suspendisse cursus risus in mauris euismod, a posuere nibh tristique. Curabitur nisi leo, condimentum vitae congue eget, maximus ut purus. Nam non felis nulla. Etiam et leo eu tortor commodo lobortis. Duis ex libero, tempus at nisl eu, viverra rhoncus libero. Donec at dui lectus. Morbi a tellus posuere, maximus neque ac, suscipit tortor. Proin faucibus condimentum leo in scelerisque. Fusce odio sapien, sollicitudin eu enim ac, dignissim lacinia nulla. Mauris sed sem scelerisque, tempor est at, mattis nulla. Vivamus vitae enim sodales, porta sapien quis, egestas mi. Pellentesque ultricies ex ac finibus hendrerit. Proin sed erat et urna dapibus imperdiet vitae nec sem. Vivamus laoreet lorem eget urna pellentesque varius. Sed semper pharetra neque, at gravida est porta in. Duis mattis eu augue a vulputate. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Duis sed interdum tortor. Aliquam in erat justo. Etiam hendrerit eros sed magna vulputate, non cursus nisi aliquam. Nullam eu vestibulum magna. Maecenas sollicitudin eros sit amet scelerisque scelerisque. Cras sem dolor, tristique vitae rutrum a, laoreet a augue. Maecenas eleifend ligula eget urna lacinia, vel euismod lacus malesuada. Praesent molestie leo ornare augue vulputate vulputate. Nunc eget lobortis est, non pretium magna. Nam finibus arcu eget velit hendrerit, sed facilisis felis cursus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nam ultrices ligula quis sem vulputate, eu convallis arcu suscipit. Morbi vel ultricies leo. Sed vitae eros ex. Maecenas faucibus, nisi semper consequat condimentum, est nisi euismod neque, ornare aliquet lectus nunc id sem. Nunc eu mattis justo, a pulvinar leo. Donec hendrerit neque ut malesuada viverra. Nullam in ex nec odio viverra laoreet. Phasellus nec consectetur libero, cursus sagittis neque. Aenean eu finibus tortor. In nulla urna, volutpat ac hendrerit sit amet, ullamcorper eu nisi. Curabitur nunc massa, rutrum vel pharetra in, rutrum sed diam. Aliquam a aliquam dolor. Nulla hendrerit nibh in vestibulum cursus. Pellentesque facilisis diam id sapien condimentum, a commodo lorem sodales. Pellentesque et dapibus nibh, ut tempor lectus. Donec lacinia quam eget purus tincidunt blandit. Nam elit turpis, scelerisque euismod pharetra eget, malesuada at urna. In suscipit, mi eget placerat scelerisque, neque libero cursus ante, sit amet semper tortor massa vitae justo. Cras viverra risus vitae pellentesque varius. Ut id interdum elit. Nulla finibus luctus nisi lacinia mattis. Donec viverra sodales molestie. Fusce lobortis lacus nunc, a porttitor quam porttitor eget. Curabitur suscipit tellus tellus, sed imperdiet odio sollicitudin sed. Vivamus euismod orci finibus enim dignissim, sed ornare lacus pharetra. Nam maximus nunc et condimentum scelerisque. Aliquam in auctor arcu. Duis condimentum est nec arcu pharetra, sit amet tempor felis aliquam. Ut id ligula erat. Donec non eros mauris. Proin elit magna, molestie non dictum in, elementum a diam. Praesent pulvinar, lorem non tincidunt dictum, nunc enim tempus lectus, ac eleifend sapien sapien a ex. Sed eget arcu vulputate, dignissim urna sit amet, viverra erat. Cras quis nunc a turpis convallis efficitur. Ut vitae efficitur metus. Quisque ullamcorper consectetur orci. Cras mattis at tortor vel congue. Proin ullamcorper dapibus rutrum. Nunc pellentesque porta urna, in mollis ex porta eget. Maecenas scelerisque est tortor, in dignissim nisl rhoncus non. In dignissim ex sit amet lorem porta pellentesque a sit amet metus. Donec efficitur feugiat dui a mollis. Vestibulum sed mauris rhoncus, feugiat arcu eu, posuere velit. In ac faucibus orci. Nunc tempor tincidunt quam nec dignissim. Mauris at lorem pulvinar, auctor mauris ac, laoreet nunc. Donec fermentum diam massa, vitae bibendum eros rutrum id. Etiam sit amet orci nunc. Aliquam ut volutpat quam, a porta arcu. Sed euismod enim consequat sollicitudin facilisis. Ut molestie, est non dictum eleifend, eros justo viverra odio, sed tristique erat sapien et diam. Vivamus in lectus at elit consectetur aliquet. Morbi scelerisque consectetur ipsum, vitae viverra orci tempor ac. Pellentesque mauris sem, posuere vitae volutpat ut, pretium in augue. Aliquam erat volutpat. Nulla blandit elit vel lectus ultricies laoreet tempus non dui. Suspendisse et iaculis leo, eget placerat ex. Interdum et malesuada fames ac ante ipsum primis in faucibus. Quisque sagittis venenatis consectetur. Morbi eget consectetur tellus, non maximus nulla. Maecenas pellentesque nisl vitae nibh aliquet, eget accumsan tortor efficitur. Nullam scelerisque sagittis odio, a tempus sem tristique quis. Morbi a semper odio. Donec volutpat consectetur diam, in condimentum erat imperdiet et. Nam in tincidunt ex. Suspendisse sollicitudin pellentesque viverra. Mauris tincidunt, nisl vitae bibendum tempor, nisi sapien mollis felis, eu efficitur sem diam vehicula nulla. Vivamus eleifend tortor eu libero tempor accumsan. Aliquam id nulla vel erat convallis ornare sit amet in nibh. Curabitur rhoncus consectetur arcu id blandit. Nunc scelerisque vehicula enim. Maecenas ac tellus quis nunc dapibus aliquam. Etiam pulvinar gravida magna sed molestie. Vivamus fermentum in erat nec convallis. Nullam commodo purus vel leo aliquam, eu tincidunt ipsum accumsan. Sed efficitur vulputate purus, et vulputate ante viverra vel. Sed suscipit tincidunt est sit amet cursus. Curabitur efficitur congue mi vel posuere. Praesent maximus venenatis tortor non pulvinar. Praesent eu eleifend arcu, a vestibulum lectus. Pellentesque augue ex, efficitur varius tortor vitae, efficitur mattis nisi. Etiam gravida porta nulla, vel porta tellus. Donec viverra nisl eros, varius bibendum nunc aliquam ut. Nam sollicitudin massa vitae enim ornare viverra. Donec euismod facilisis lacinia. Ut quis dui nec nisl iaculis ornare posuere sed mauris. Nunc accumsan eu magna eu imperdiet. Vestibulum porta nisl fermentum nibh porta semper. Sed nulla urna, blandit quis sodales ac, sagittis sed ligula. Integer dictum sapien imperdiet egestas venenatis. In blandit a lorem eget tempor. Ut egestas, ante quis volutpat placerat, nisi orci accumsan ante, ut feugiat purus lorem tincidunt nisl. Quisque porta purus diam, vel porta ex auctor sed. Curabitur felis ante, auctor sed dui ut, fringilla mollis massa. Nullam eget pellentesque arcu, quis finibus justo. Donec tortor urna, hendrerit eu pellentesque porta, ornare sed arcu. Cras interdum ante eu lorem pharetra, eu lobortis libero maximus. Nam condimentum placerat porttitor. Aenean sit amet est ut eros euismod mollis quis quis erat. Duis lacinia laoreet tellus at mattis. Duis tellus lectus, porttitor vitae tincidunt id, ultricies eget sapien. Aliquam tincidunt commodo venenatis. Sed scelerisque porttitor orci vitae hendrerit. Praesent auctor eu massa nec placerat. Nullam tempus dignissim porta. Quisque nisi nisi, elementum interdum vestibulum sit amet, congue in elit. Sed in nisl nec sapien congue volutpat nec non dui. Maecenas vitae mi vel ligula blandit finibus ut id ex. Donec in tempus dui, et condimentum lorem. Nunc maximus, tortor eu dictum efficitur, magna magna faucibus tellus, a ornare velit sem in eros. In hac habitasse platea dictumst. Cras est odio, dignissim vitae lorem ac, placerat finibus urna. Etiam hendrerit eros vitae felis hendrerit aliquam. Duis et dapibus metus, quis volutpat ante. In at posuere est, ut gravida ante. Cras pharetra mauris at augue iaculis, id consequat arcu gravida. Donec non tincidunt lectus. Cras fringilla, nisi sed blandit euismod, nisl lorem iaculis lorem, sit amet elementum lorem enim et odio. Quisque mauris elit, ultrices eu arcu ac amet."

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		l.Debug(exctx, msg)
	}
}

func Benchmark1000Chars(b *testing.B) {
	var buffer bytes.Buffer
	l := CreateTestLogger(&buffer)
	exctx := exctx.Create()
	msg := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque placerat ullamcorper euismod. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a massa in est suscipit molestie. Morbi vel urna vel orci malesuada suscipit ut vitae ipsum. Maecenas aliquam arcu id elit tincidunt, nec lacinia dolor finibus. Fusce sed velit malesuada, pretium nisl at, aliquet ligula. Vestibulum blandit porta rhoncus. Donec dapibus aliquam dictum. Curabitur cursus, nunc quis maximus vehicula, justo enim porttitor lacus, quis pharetra velit quam ac nulla. Quisque libero nisi, lobortis eu gravida nec, mattis sed erat. Fusce condimentum quam id lobortis pretium. Fusce et facilisis nisl. In in bibendum purus. Praesent lectus nisi, feugiat sed luctus sed, aliquet id diam. Nulla facilisi. Suspendisse sollicitudin turpis vel justo fringilla pharetra. Praesent sollicitudin ante ut lacus interdum, ut mattis neque facilisis. Morbi eleifend massa nec felis consequat tempor. Nunc vitae massa mollis posuere."

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		l.Debug(exctx, msg)
	}
}

func Benchmark100Chars(b *testing.B) {
	var buffer bytes.Buffer
	l := CreateTestLogger(&buffer)
	exctx := exctx.Create()
	msg := "Logger bm. Logger bm. Logger bm. Logger bm. Logger bm. Logger bm. Logger bm. Logger bm. Logger bm. Y"

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		l.Debug(exctx, msg)
	}
}

func Benchmark10Chars(b *testing.B) {
	var buffer bytes.Buffer
	l := CreateTestLogger(&buffer)
	exctx := exctx.Create()
	msg := "Logger bm."

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		l.Debug(exctx, msg)
	}
}
