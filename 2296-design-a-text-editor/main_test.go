package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func Test_Run1(t *testing.T) {
	ed := Constructor()
	ed.AddText("leetcode")

	assert.Equal(t, ed.DeleteText(4), 4)

	ed.AddText("practice")

	assert.Equal(t, ed.CursorRight(3), "etpractice")
	assert.Equal(t, ed.CursorLeft(8), "leet")

	assert.Equal(t, ed.DeleteText(10), 4)

	assert.Equal(t, ed.CursorLeft(2), "")
	assert.Equal(t, ed.CursorRight(6), "practi")
}

func Test_Submit34(t *testing.T) {
	ed := Constructor()
	ed.AddText("bxyackuncqzcqo")

	assert.Equal(t, ed.CursorLeft(12), "bx")
	assert.Equal(t, ed.DeleteText(3), 2)
	assert.Equal(t, ed.CursorLeft(5), "")

	ed.AddText("osdhyvqxf")

	assert.Equal(t, ed.CursorRight(10), "yackuncqzc")
}

func Test_Submit36(t *testing.T) {
	ed := Constructor()

	assert.Equal(t, ed.DeleteText(9), 0)
	assert.Equal(t, ed.CursorLeft(14), "")

	ed.AddText("mjzxkemqyvfrg")
	fmt.Println(ed.String())

	assert.Equal(t, ed.CursorLeft(5), "mjzxkemq")

	fmt.Println(ed.String())

	ed.AddText("svvyobqebssp")
	fmt.Println(ed.String())

	ed.AddText("xkoznsq")
	fmt.Println(ed.String())

	fmt.Println("-----")
	// ed.CursorLeft(3)
	// fmt.Println(ed.String())
	// ed.CursorRight(3)
	// fmt.Println(ed.String())

	// fmt.Println(ed.CursorRight(3))
	// fmt.Println(ed.CursorRight(1))
	// return

	assert.Equal(t, ed.CursorRight(3), "xkoznsqyvf")
	// fmt.Println(ed.String())
	assert.Equal(t, ed.CursorLeft(15), "emqsvvyobq")
	fmt.Println(ed.String())

	ed.AddText("irgkilfitjuysfgu")

	assert.Equal(t, ed.CursorLeft(18), "xkemqsvvyo")

	ed.AddText("aanfokhwqqjf")

	assert.Equal(t, ed.CursorLeft(18), "mjzxkem")
	assert.Equal(t, ed.DeleteText(7), 7)

	ed.AddText("kjekmkypfqvbsyigry")
	ed.AddText("opkmggjoxng")

	assert.Equal(t, ed.CursorRight(17), "anfokhwqqj")
	assert.Equal(t, ed.CursorLeft(1), "aanfokhwqq")
	assert.Equal(t, ed.CursorLeft(9), "xngqsvvyoa")
	assert.Equal(t, ed.DeleteText(16), 16)
	assert.Equal(t, ed.DeleteText(17), 17)
	assert.Equal(t, ed.DeleteText(2), 2)
	assert.Equal(t, ed.DeleteText(10), 1)
	assert.Equal(t, ed.DeleteText(5), 0)
	assert.Equal(t, ed.CursorLeft(15), "")

	ed.AddText("bskjfptrqpg")

	assert.Equal(t, ed.CursorRight(1), "kjfptrqpga")
	assert.Equal(t, ed.CursorLeft(16), "")

	ed.AddText("dpuzcwbqotoihn")
	ed.AddText("an")
	ed.AddText("pntm")
	ed.AddText("eamnglbbdnencdunao")

	assert.Equal(t, ed.CursorRight(12), "kjfptrqpga")
	assert.Equal(t, ed.CursorRight(6), "qpganfokhw")
	assert.Equal(t, ed.CursorLeft(5), "jfptrqpgan")
	assert.Equal(t, ed.CursorLeft(14), "bdnencduna")

	ed.AddText("utcyzkianwahi")
	assert.Equal(t, ed.CursorRight(4), "anwahiobsk")

	ed.AddText("wcojk")
	ed.AddText("yrfmlfvzev")
	ed.AddText("ncpjvbv")
	ed.AddText("ie")
	ed.AddText("megbbxribprrsojsiofd")
	ed.AddText("witgxkjntmvfvqjxitf")
	ed.AddText("ifvhcabjla")

	assert.Equal(t, ed.DeleteText(10), 10)

	/*
		['TextEditor', [], 'null']
		['deleteText', [9], 0]
		['cursorLeft', [14], '']
		['addText', ['mjzxkemqyvfrg'], 'null']
		['cursorLeft', [5], 'mjzxkemq']
		['addText', ['svvyobqebssp'], 'null']
		['addText', ['xkoznsq'], 'null']
		['cursorRight', [3], 'xkoznsqyvf']
		['cursorLeft', [15], 'emqsvvyobq']
		['addText', ['irgkilfitjuysfgu'], 'null']
		['cursorLeft', [18], 'xkemqsvvyo']
		['addText', ['aanfokhwqqjf'], 'null']
		['cursorLeft', [18], 'mjzxkem']
		['deleteText', [7], 7]
		['addText', ['kjekmkypfqvbsyigry'], 'null']
		['addText', ['opkmggjoxng'], 'null']
		['cursorRight', [17], 'anfokhwqqj']
		['cursorLeft', [1], 'aanfokhwqq']
		['cursorLeft', [9], 'xngqsvvyoa']
		['deleteText', [16], 16]
		['deleteText', [17], 17]
		['deleteText', [2], 2]
		['deleteText', [10], 1]
		['deleteText', [5], 0]
		['cursorLeft', [15], '']
		['addText', ['bskjfptrqpg'], 'null']
		['cursorRight', [1], 'kjfptrqpga']
		['cursorLeft', [16], '']
		['addText', ['dpuzcwbqotoihn'], 'null']
		['addText', ['an'], 'null']
		['addText', ['pntm'], 'null']
		['addText', ['eamnglbbdnencdunao'], 'null']
		['cursorRight', [12], 'kjfptrqpga']
		['cursorRight', [6], 'qpganfokhw']
		['cursorLeft', [5], 'jfptrqpgan']
		['cursorLeft', [14], 'bdnencduna']
		['addText', ['utcyzkianwahi'], 'null']
		['cursorRight', [4], 'anwahiobsk']
		['addText', ['wcojk'], 'null']

		['addText', ['yrfmlfvzev'], 'null']
		['addText', ['ncpjvbv'], 'null']
		['addText', ['ie'], 'null']
		['addText', ['megbbxribprrsojsiofd'], 'null']
		['addText', ['witgxkjntmvfvqjxitf'], 'null']
		['addText', ['ifvhcabjla'], 'null']
		['deleteText', [10], 10]
	*/

	// ["TextEditor","deleteText","cursorLeft","addText","cursorLeft","addText","addText","cursorRight","cursorLeft","addText","cursorLeft","addText","cursorLeft","deleteText","addText","addText","cursorRight","cursorLeft","cursorLeft","deleteText","deleteText","deleteText","deleteText","deleteText","cursorLeft","addText","cursorRight","cursorLeft","addText","addText","addText","addText","cursorRight","cursorRight","cursorLeft","cursorLeft","addText","cursorRight","addText","addText","addText","addText","addText","addText","addText","deleteText"]
	//
	// [[],[9],[14],["mjzxkemqyvfrg"],[5],["svvyobqebssp"],["xkoznsq"],[3],[15],["irgkilfitjuysfgu"],[18],["aanfokhwqqjf"],[18],[7],["kjekmkypfqvbsyigry"],["opkmggjoxng"],[17],[1],[9],[16],[17],[2],[10],[5],[15],["bskjfptrqpg"],[1],[16],["dpuzcwbqotoihn"],["an"],["pntm"],["eamnglbbdnencdunao"],[12],[6],[5],[14],["utcyzkianwahi"],[4],["wcojk"],["yrfmlfvzev"],["ncpjvbv"],["ie"],["megbbxribprrsojsiofd"],["witgxkjntmvfvqjxitf"],["ifvhcabjla"],[10]]
	//
	// ["null",0,"","null","mjzxkemq","null","null","xkoznsqyvf","emqsvvyobq","null","xkemqsvvyo","null","mjzxkem",7,"null","null","anfokhwqqj","aanfokhwqq","xngqsvvyoa",16,17,2,1,0,"","null","kjfptrqpga","","null","null","null","null","kjfptrqpga","qpganfokhw","jfptrqpgan","bdnencduna","null","anwahiobsk","null","null","null","null","null","null","null",10]
}
