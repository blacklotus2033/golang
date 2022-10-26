package main

import "fmt"

func main() {
	//for i := 1; i <= 42; i++ {
	//	fmt.Printf("JRecr(%d):%d JPower(%d):%d JK(%d,2):%d J(%d):%d Ring:%s\n", i, JRecr(i), i, JPower(i), i, JK(i, 2), i, J(i), CreateRing(i))
	//}
	for k := 2; k <= 6; k++ {
		for i := 1; i <= 20; i++ {
			fmt.Printf("JK(%d,%d):%d JKRecrBy1(%d,%d):%d JKIterBy1(%d,%d):%d\n", i, k, JK(i, k), i, k, JKRecrBy1(i, k), i, k, JKIterBy1(i, k))
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) String() string {
	var accessed = make(map[*ListNode]bool)
	var ret = "ListNode: "
	for p != nil {
		if accessed[p] {
			ret += fmt.Sprintf("accessed(%d)", p.Val)
			return ret
		}
		accessed[p] = true
		ret += fmt.Sprintf("%d ", p.Val)
		p = p.Next
	}
	return ret
}

func CreateRing(n int) *ListNode {
	if n <= 0 {
		return nil
	}
	var all = make([]ListNode, n)
	for i := range all {
		all[i].Val = i + 1 //1)值为1-started
		all[i].Next = &all[(i+1)%n]
	}
	//fmt.Println(all)
	return &all[0]
}

// J 最简单的，淘汰下一个人的写法。
func J(n int) int {
	var head = CreateRing(n)
	for head.Next != head {
		head.Next, head = head.Next.Next, head.Next
	}
	return head.Val
}

// JPower 用2的幂次法来解题，并和J()做对拍
func JPower(n int) int {
	var f = 1
	for m := n; m > 1; m >>= 1 {
		f <<= 1
	}
	//return 1 + (n^f)*2
	return 1 + (n^f)<<1
}

// JRecr 用递推来解决问题
func JRecr(n int) int {
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return 2*JRecr(n/2) - 1
	}
	return 2*JRecr(n/2) + 1
}

// JK 用链表来模拟k>=2的情况
func JK(n, k int) int {
	var head = CreateRing(n)
	var cnt int
	for head.Next != head {
		cnt++
		if cnt%(k-1) == 0 {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return head.Val
}

// JKRecrBy1 猜一个和JKRecr(n-1)算法的算法
func JKRecrBy1(n, k int) int {
	if n == 1 {
		return 1
	}
	return (JKRecrBy1(n-1, k)+k-1)%n + 1
}

// JKIterBy1 猜一个和JKRecr(n-1)算法的算法
func JKIterBy1(n, k int) int {
	var ret = 1
	for i := 2; i <= n; i++ {
		ret = (ret+k-1)%i + 1
	}
	return ret
}
