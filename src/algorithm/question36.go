package algorithm

var cnt int = 0
var checked = [21][21][21]int{}

func Question36(n int , a int , b int , c int) {

	if a + b + c == n {
		if a <= b && a <= c && a + b > c && checked[a][b][c] == 0 {
			cnt = cnt + 1
			checked[a][b][c] = 1
		}
		return
	}
	Question36(n, a+1 , b , c)
	Question36(n, a , b+1 , c)
	Question36(n, a , b , c+1)
}
