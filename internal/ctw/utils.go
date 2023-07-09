package ctw

import (
	"regexp"
	"strings"
)

// default white looks like #c5c5c5 on my terminal, C5 = 197
var (
	noColor    string = "\033[0m"
	_bold      string = "\033[1m"
	_dim       string = "\033[2m"
	_italic    string = "\033[3m"
	_underline string = "\033[4m"
	blue       string = "\033[38;2;97;175;239m"
	brown      string = "\033[38;2;192;154;107m"
	gray       string = "\033[38;2;117;117;117m"
	gray_light string = "\033[38;2;197;197;197m"
	green      string = "\033[38;2;055;183;021m"
	green_exe  string = "\033[38;2;76;175;080m"
	pink       string = "\033[38;2;198;48;95m"
	white      string = "\033[38;2;255;255;255m"
	brailEmpty string = "\u2800"
)

var rePerm = regexp.MustCompile(`(\-+)`)
var rePermFirst = regexp.MustCompile(`^([Ld])`)
var reHasIndicator = regexp.MustCompile(`[\/|@=*]$`).MatchString

func DisplayColor(b bool) {
	if b == false {
		noColor = ""
		_bold = ""
		_dim = ""
		_italic = ""
		_underline = ""
		blue = ""
		brailEmpty = ""
		brown = ""
		brown = ""
		gray = ""
		gray = ""
		gray_light = ""
		green = ""
		green_exe = ""
		green_exe = ""
		pink = ""
		pink = ""
		white = ""
	}
}

func DisplayBrailEmpty(b bool) {
	if b == false {
		brailEmpty = " "
	}
}

func StripAnsiCodes(s string) string {
	re := regexp.MustCompile("\\x1b\\[[0-9;]*m")

	return re.ReplaceAllString(s, "")
}

func getOtherColors(str string) string {
	return noColor
}

// color permissions column
// first character can be L or D
func getColoredPerm(str string) string {
	res := rePerm.ReplaceAllString(str, gray+`$1`+noColor)
	res = rePermFirst.ReplaceAllString(res, _bold+`$1`+noColor)
	return res
}

func getSizeColor(size string) string {
	return gray_light + _italic
}

func getSanitizedFilename(filename string) string {
	fn, del, indicator := filename, "", ""

	if split := strings.Split(filename, "\t"); len(split) > 0 {
		fn = split[0]
		if len(split) > 1 {
			indicator = split[1]
		}
	}

	del = map[string]string{
		"@": " ",
		"*": " ",
	}[indicator]

	return fn + del + indicator
}

// hidden files are greyed out, even if they have gitStatus
func getFilenameColor(filename string) string {
	if len(filename) == 0 {
		return ""
	}

	clr, fn, indicator := noColor, filename, ""

	if split := strings.Split(filename, "\t"); len(split) > 0 {
		fn = split[0]
		if len(split) > 1 {
			indicator = split[1]
		}
	}

	clr += map[bool]string{true: gray}[fn[:1] == "."]

	clr += map[string]string{
		"/": _bold,
		"@": _italic + blue,
		"*": _bold + green_exe,
	}[indicator]

	return clr
}

func ColorizeDirHeader(dirname string) string {
	return _bold + _underline + dirname + noColor
}

func getGitColor(gitStatus string) string {
	switch strings.Trim(gitStatus, " ") {
	case "":
		return noColor
	case "A":
		return green
	case "U":
		return green
	case "M":
		return brown
	default:
		return brown
	}
}

func widthsSum(w [][4]int, p int) int {
	s := 0
	for _, v := range w {
		s += v[0] + v[1] + v[2] + v[3] + p
	}
	s -= p
	return s
}
