package ast

type logicFunc = func(...bool) bool // and or ...

func StringEqual(s0, s1 string) bool {
	return s0 == s1
}

func or(lst ...bool) bool {
	// zero length ?
	if len(lst) == 1 {
		return lst[0]
	} else {
		return lst[0] || or(lst[1:]...)
	}
}

func and(lst ...bool) bool {
	// zero length ?
	if len(lst) == 1 {
		return lst[0]
	} else {
		return lst[0] && and(lst[1:]...)
	}
}

type StringMatchPart struct {
	op   func(string, string) bool
	arg1 string
}

func (part StringMatchPart) ExprFunc() func(string) bool {
	// like name: 360
	return func(arg0 string) bool {
		return part.op(arg0, part.arg1)
	}
}

type MatchPart struct {
	op   logicFunc
	args []StringMatchPart
}

func (part MatchPart) ExprFunc() func(...string) bool {
	// like name, version
	return func(args ...string) bool {
		bl := make([]bool, len(args))
		for index := range args {
			f := part.args[index].ExprFunc()
			r := f(args[index])
			bl[index] = r
		}
		return part.op(bl...)
	}
}

// this would a rule repr
type MatchPartLst struct {
	op   logicFunc
	args []MatchPart
}

func (lst MatchPartLst) ExprFunc() func(...string) bool {
	// 360shandu, 1.1.2
	return func(inputs ...string) bool {
		for index := range lst.args {
			f := lst.args[index].ExprFunc()
			r := f(inputs...)
			if r {
				return r
			} else {
				continue
			}
		}
		return false
	}
}
