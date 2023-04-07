package types

import "github.com/Amqp-prtcl/snowflakes"

type CkString string

type Header struct {
	ID       snowflakes.ID `json:"id"`
	LastMod  int64         `json:"last-mod"`
	Title    string        `json:"title"`
	Category []string      `json:"categories"`
}

type MCQ struct {
	Header
	Question    CkString      `json:"question"`
	Answers     []CkString    `json:"answers"`
	Correct     []int         `json:"correct"`
	Explanation CkString      `json:"explanation"`
	Author      snowflakes.ID `json:"author"`
	Ref         []string      `json:"ref"`
	LastRef     int64         `json:"ref-mod"`
}

type Question struct {
	Header
	Question    CkString      `json:"question"`
	Explanation CkString      `json:"explanation"`
	Author      snowflakes.ID `json:"author"`
	Ref         []string      `json:"ref"`
	LastRef     int64         `json:"ref-mod"`
}

type Group struct {
	ID        snowflakes.ID `json:"id"`
	Name      string        `json:"name"`
	Mcqs      []Header      `json:"mcqs"`
	Questions []Header      `json:"questions"`
}
