package solutions

import (
	"testing"
)

func TestFindMaxConsecutiveOnesCase1(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(0)
	obj.Contains(1)
	obj.Contains(3)
	obj.Add(2)
	obj.Contains(2)
	obj.Remove(2)
	obj.Contains(0)
	/*
		["MyHashSet","add","add","contains","contains","add","contains","remove","contains"]
		[[],[1],[0],[1],[3],[2],[2],[2],[0]]

		Expected: [null,null,null,true,false,null,true,null,true]
	*/
}
