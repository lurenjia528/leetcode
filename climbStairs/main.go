package main

import (
	"fmt"
)

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2  = 1+0+1
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3  = 1+1+1
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

// 4   =  1+3+1
// 1+1+1+1
// --------
// 1+1+2
// 1+2+1
// 2+1+1
// --------
// 2+2

// 5   =  1+4+3
// 1+1+1+1+1
// ---------
// 1+2+1+1
// 1+1+2+1
// 1+1+1+2
// 2+1+1+1
// ---------
// 1+2+2
// 2+1+2
// 2+2+1

// 6   = 1+5+6+1
// 1+1+1+1+1+1  C60
// -----------
// 2+1+1+1+1   C51
// 1+2+1+1+1
// 1+1+2+1+1
// 1+1+1+2+1
// 1+1+1+1+2
// -----------
// 2+2+1+1  C42
// 2+1+2+1
// 2+1+1+2
// 1+1+2+2
// 1+2+1+2
// 1+2+2+1
// -----------
// 2+2+2 C33

func main() {
	//fmt.Println("++++++++++")
	fmt.Println(climbStairs(35))

	//fmt.Println("C51")
	//C(4, 2)
	//C(34,1)
	//fmt.Println("-----")
	//fmt.Println(CNM(34,1))
}

func climbStairs(n int) int {
	var result int
	if n%2 == 0 {
		fmt.Println("偶数")
		fmt.Printf("最多%d个2\n", n/2)
		for i := 0; i <= n/2; i++ {
			fmt.Printf("C%d %d=", n-i, i)
			result += CNM(n-i, i)
			fmt.Println(CNM(n-i, i))
		}
	} else {
		fmt.Println("奇数")
		fmt.Printf("最多%d个2\n", (n-1)/2)
		for i := 0; i <= (n-1)/2; i++ {
			fmt.Printf("C%d %d=", n-i, i)
			result += CNM(n-i, i)
			fmt.Println(CNM(n-i, i))
		}
	}
	fmt.Println(result)
	return 0
}

func A(n int, m int) uint64 {
	u := factorial(n)
	a := factorial(n - m)
	//fmt.Println(u)
	//fmt.Println(a)
	//fmt.Println(u / a)
	return u / a
}

func C(n int, m int) uint64 {
	a := factorial(n)
	b := factorial(m)
	c := factorial(n - m)
	//fmt.Println(a/(b*c))
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	return a / (b * c)
}

func factorial(n int) uint64 {
	var facVal uint64 = 1
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			facVal *= uint64(i)
		}
	}
	return facVal
}


// C 34 1
// 34! / (1! * (34-1)!)
func CNM(n int, m int) int{
	 k:=1
	 s:=1
	 for i:=n-m+1;i<=n;i++ {
	 	s *= i
	 	for k<=m && s%k==0{
	 		s /= k
	 		k++
		}
	 }
	 return s
}
