package ctw

import (
	"bytes"
	"fmt"
)

type LongCTW struct {
	d    [][]string // entire data passed to ctw
	c    []int      // widths of each column
	ic   []string   // color of file icon of a row
	cols int        // zero based no of cols (or total cols -1)
}

func NewLong(cols int) *LongCTW {
	t := new(LongCTW)
	t.cols = cols - 1

	// initialize right size slice
	t.c = make([]int, cols)
	t.d = make([][]string, 0)
	t.ic = make([]string, 0)
	return t
}

func (l *LongCTW) AddRow(args ...string) {
	// add length checking for args
	if len(args) != l.cols+1 {
		return
	}
	for i, v := range args {
		if l.c[i] < len(v) {
			l.c[i] = len(v)
		}
	}
	l.d = append(l.d, args)
}

func (l *LongCTW) IconColor(c string) {
	l.ic = append(l.ic, c)
}

func (l *LongCTW) Flush(buf *bytes.Buffer) {
	var skipCol int = 0
	for i, v := range l.c {
		if v == 0 {
			skipCol |= 1 << i
		}
	}

	// explicitly setting width of git column to 1
	l.c[l.cols] = 1
	// explicitly setting width of icon column to 2
	l.c[l.cols-2] = 1

	for i, r := range l.d {
		f := true
		for j, c := range r {
			if (1<<j)&skipCol > 0 {
				continue
			}
			if f == false {
				fmt.Fprintf(buf, "%s", brailEmpty)
			}

			// icon
			if j == l.cols-2 {
				fmt.Fprintf(buf, "%s%*s%s", l.ic[i], l.c[j], c, noColor)
			} else if j == l.cols && (1<<l.cols)&skipCol == 0 {
				color := getGitColor(r[l.cols])
				fmt.Fprintf(buf, "%s%-*s%s", color, l.c[j], c, noColor)
			} else if j == l.cols-1 {
				color := getFilenameColor(c)
				fmt.Fprintf(buf, "%s%-*s%s", color, l.c[j], getSanitizedFilename(c), noColor)
			} else if j == l.cols-4 {
				// file sizes are right-aligned
				// human-readable formats will have the letter highlighted
				if c[len(c)-1:] > "9" {
					fmt.Fprintf(buf, "%s%*s%s%s%s", getSizeColor(c), l.c[j]-1, c[0:len(c)-1], white+_bold, c[len(c)-1:], noColor)
				} else {
					fmt.Fprintf(buf, "%s%*s%s", getSizeColor(c), l.c[j], c, noColor)
				}
			} else if j == 1 {
				// color permissions
				fmt.Fprintf(buf, "%-*s", l.c[j], getColoredPerm(c))
			} else {
				// the rest of the columns are left-aligned
				fmt.Fprintf(buf, "%s%-*s%s", getOtherColors(c), l.c[j], c, noColor)
			}
			f = false
		}
		fmt.Fprintln(buf)
	}
}
