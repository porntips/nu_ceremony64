package model

type Ceremony struct {
	Studentcode       string `json:"studentcode" binding:"required"`
	Sname             string `json:"sname" binding:"required"`
	Degreecertificate string `json:"degreecertificate" binding:"required"` //
	Facultyname       string `json:"facultyname" binding:"required"` //
	Hornor            int16  `json:"hornor" binding:"required"` //
	Ceremonygroup     int16  `json:"ceremonygroup" binding:"required"`
	Ceremonysequence  int16  `json:"ceremonysequence" binding:"required"`
	Ceremonydate      string `json:"ceremonydate" binding:"required"` //
	Ceremonypack      int16  `json:"ceremonypack" binding:"required"`
	Ceremonypackno    int16  `json:"ceremonypackno" binding:"required"` //
	Ceremonysex       any    `json:"ceremonysex"` //
	Ceremonyprefix    any    `json:"ceremonyprefix"` //
	Ceremony          bool   `json:"ceremony"`
}
type ReturnCeremony struct {
	Ceremony []Ceremony
	Count    int
}

type ReturnGrad struct {
	Pack_count     int        `json:"pack_count"`     //Group_count
	Pack_remain    int        `json:"pack_remain"`    //Group_remain
	Receive_result []Ceremony `json:"receive_result"` //Grad_receive
	Receive_count  int        `json:"receive_count"`  //Grad_count
	Remain_result  []Ceremony `json:"remain_result"`  //False_result
	Ceremonypack   int        `json:"ceremonypack"`   //Ceremonygroup
}

type Running struct {
	Studentcode string `json:"studentcode" binding:"required"`
}