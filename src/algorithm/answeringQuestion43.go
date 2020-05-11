package algorithm
var table [101]int

func Question43(k int){
	for i := 1 ; i <= k ; i ++{
		if i == 1{
			table[i] = 1
		} else {
			table[i] = table[i-1] + i;
		}
	}
}

