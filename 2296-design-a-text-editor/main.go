package main

import (
	"fmt"
)

type TextEditor struct {
	text *Text
}

func (t *TextEditor) getLeft(k int) string {
	text := ""
	if t.text.Char != rune(0) {
		text = string(t.text.Char)
	}

	tmp := t.text
	for i := 0; i < 10-1; i++ {
		if tmp.PrevPtr == nil {
			break
		}

		if tmp.PrevPtr.Char != rune(0) {
			text = string(tmp.PrevPtr.Char) + text
		}

		tmp = tmp.PrevPtr
	}

	return text
}

// func (t *TextEditor) getRight(k int) string {
// 	text := string(t.text.Char)
// 	tmp := t.text
// 	for i := 0; i < k; i++ {
// 		if tmp.NextPtr == nil {
// 			break
// 		}
// 		text = text + string(tmp.NextPtr.Char)
// 		tmp = tmp.NextPtr
// 	}
// 	return text
// }

func (t *TextEditor) String() string {
	prev := ""
	tmp := t.text

	for tmp.PrevPtr != nil {
		if tmp.PrevPtr.Char != rune(0) {
			prev = string(tmp.PrevPtr.Char) + prev
		}
		tmp = tmp.PrevPtr
	}
	tmp = t.text

	next := ""
	for tmp.NextPtr != nil {
		if tmp.NextPtr.Char != rune(0) {
			next = next + string(tmp.NextPtr.Char)
		}
		tmp = tmp.NextPtr
	}

	result := prev
	if t.text.Char != rune(0) {
		result = result + string(t.text.Char)
	}
	result = result + "|" + next

	return result
	// return prev + string(t.text.Char) + "|" + next
}

type Text struct {
	PrevPtr *Text
	NextPtr *Text
	Char    rune
}

func Constructor() TextEditor {
	return TextEditor{
		// text: nil,
		&Text{
			PrevPtr: nil,
			NextPtr: nil,
			// Char:    rune(0),
		},
	}
}

func (t *TextEditor) AddText(text string) {
	for _, c := range text {
		newText := &Text{
			Char:    c,
			PrevPtr: t.text,
			NextPtr: t.text.NextPtr,
		}

		if t.text.NextPtr != nil {
			t.text.NextPtr.PrevPtr = newText
		}

		t.text.NextPtr = newText
		t.text = newText
	}
}

func (t *TextEditor) DeleteText(k int) int {
	counter := 0
	for i := 0; i < k; i++ {
		// if t.text == nil {
		// 	break
		// }

		// if t.text.Char == rune(0) {
		// 	break
		// }

		if t.text.PrevPtr != nil {
			t.text.PrevPtr.NextPtr = t.text.NextPtr
		}

		if t.text.NextPtr != nil {
			t.text.NextPtr.PrevPtr = t.text.PrevPtr
		}

		if t.text.PrevPtr == nil {
			break
			// t.text.Char = rune(0)
			// counter++
		}

		// if t.text.PrevPtr != nil {
		t.text = t.text.PrevPtr
		// }

		counter++
	}

	return counter
}

func (t *TextEditor) CursorLeft(k int) string {
	if t.text == nil {
		return ""
	}

	for i := 0; i < k; i++ {
		if t.text.PrevPtr == nil {
			break
		}
		t.text = t.text.PrevPtr
	}
	return t.getLeft(k)
}

func (t *TextEditor) CursorRight(k int) string {
	if t.text == nil {
		return ""
	}

	for i := 0; i < k; i++ {
		// fmt.Println(
		// t.text.PrevPtr.PrevPtr.PrevPtr.Char,
		// string(t.text.PrevPtr.PrevPtr.PrevPtr.Char),
		// t.text.PrevPtr.PrevPtr.Char,
		// string(t.text.PrevPtr.PrevPtr.Char),
		// t.text.PrevPtr.Char,
		// string(t.text.PrevPtr.Char),
		// " | ",
		// t.text.Char,
		// string(t.text.PrevPtr.Char),
		// string(t.text.PrevPtr.NextPtr.Char),
		// "("+string(t.text.Char)+")",
		// string(t.text.NextPtr.PrevPtr.PrevPtr.Char),
		// string(t.text.NextPtr.PrevPtr.Char),
		// "("+string(t.text.NextPtr.Char)+")",
		// string(t.text.NextPtr.NextPtr.PrevPtr.Char),
		// string(t.text.NextPtr.NextPtr.Char),
		// " | ",
		// " | ",
		// string(t.text.NextPtr.Char),
		// string(t.text.NextPtr.NextPtr.Char),
		// string(t.text.NextPtr.NextPtr.NextPtr.Char),
		// )
		if t.text.NextPtr == nil {
			fmt.Println("break")
			break
		}
		t.text = t.text.NextPtr
	}

	return t.getLeft(k)
}

func main() {
	// testCase0()
	// testCase1()
	// testCase2()
	testCase111()
}

func testCase0() {
	ed := Constructor()

	ed.AddText("Hello")
	fmt.Println(ed.String())

	fmt.Println(ed.DeleteText(2))
	fmt.Println(ed.String())

	ed.CursorLeft(1)
	fmt.Println(ed.String())

	fmt.Println(ed.DeleteText(10))
	fmt.Println(ed.String())

	fmt.Println(ed.CursorLeft(3))
	fmt.Println(ed.String())

	ed.CursorLeft(4)
	fmt.Println(ed.String())

	ed.CursorLeft(10)
	fmt.Println(ed.String())

	ed.CursorRight(1)
	fmt.Println(ed.String())

	ed.CursorRight(2)
	fmt.Println(ed.String())

	ed.CursorRight(10)
	fmt.Println(ed.String())
}

func testCase111() {
	ed := Constructor()

	ed.AddText("abcd")
	ed.AddText("abcd")
	fmt.Println(ed.String())

	ed.CursorLeft(4)
	fmt.Println(ed.String())

	ed.CursorRight(1)
	fmt.Println(ed.String())

	ed.AddText("efgh")
	fmt.Println(ed.String())
	tmpPrev := ed.text
	for {
		tmpPrev = tmpPrev.PrevPtr
		if tmpPrev == nil || tmpPrev.PrevPtr == nil {
			break
		}
		fmt.Println(
			string(tmpPrev.PrevPtr.Char),
			string(tmpPrev.Char),
			string(tmpPrev.NextPtr.Char),
		)
	}
}

func testCase1() {
	ed := Constructor()

	ed.AddText("dsqdmtbhxxltl")
	fmt.Println(ed.String())

	fmt.Println(ed.DeleteText(13))
	fmt.Println(ed.String())

	ed.AddText("iykwsfld")
	fmt.Println(ed.String())

	fmt.Println(ed.DeleteText(17))
	fmt.Println(ed.String())
}
