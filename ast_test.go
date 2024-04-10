package ast

import (
	"fmt"
	"strings"
	"testing"
)

func TestMatchPartLst_ExprFunc(t *testing.T) {
	spart0 := StringMatchPart{
		op:   strings.Contains,
		arg1: "360",
	}

	spart1 := StringMatchPart{
		op:   StringEqual,
		arg1: "1.1.2",
	}

	part0 := MatchPart{
		and,
		[]StringMatchPart{
			spart0,
			spart1,
		},
	}

	spart3 := StringMatchPart{
		op:   strings.Contains,
		arg1: "QQ",
	}

	spart4 := StringMatchPart{
		op:   strings.Contains,
		arg1: "1.1.1",
	}

	part1 := MatchPart{
		and,
		[]StringMatchPart{
			spart3,
			spart4,
		},
	}

	lst := MatchPartLst{
		args: []MatchPart{
			part0,
			part1,
		},
	}

	f := lst.ExprFunc()
	fmt.Println(f("360shadu", "1.1.2"))
	fmt.Println(f("1QQ1", "1.1.2"))
	fmt.Println(f("Qe1", "1.1.2"))
}
