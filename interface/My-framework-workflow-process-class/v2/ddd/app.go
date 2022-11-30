package ddd

import "sort"

type (
	ExamInterface interface {
		GetAnswer(uid, xid int) (aids []int)
		Submit(uid, xid int, data string)
		Update(uid, aid int, data string)
	}
	YikeApp struct {
		exam ExamInterface
	}
	ZuoyebangApp struct {
		exam ExamInterface
	}
	XiaoluApp struct {
		exam ExamInterface
	}
)

func (p *YikeApp) Submit(uid, xid int, data string) bool { //1)一课只能提交一次
	var aids = p.exam.GetAnswer(uid, xid)
	if len(aids) > 0 {
		return false
	}
	p.exam.Submit(uid, xid, data)
	return true
}

func (p *ZuoyebangApp) Submit(uid, xid int, data string) bool { //2)作业帮如果有就update最新的
	var aids = p.exam.GetAnswer(uid, xid)
	if len(aids) == 0 {
		p.exam.Submit(uid, xid, data)
		return true
	}
	sort.Ints(aids)
	p.exam.Update(uid, aids[len(aids)-1], data)
	return true
}

func (p *XiaoluApp) Submit(uid, xid int, data string) bool { //3)小鹿可以提交多次
	p.exam.Submit(uid, xid, data)
	return true
}
