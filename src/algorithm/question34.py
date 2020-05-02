INT_MAX = 10000

m = [
    [1,5,3],
    [2,5,7],
    [5,3,5]
]

col_check = [False, False, False]
min_sol = INT_MAX

def Question34(row , score):
    global min_sol

    if row == 3:
        if score < min_sol :
            min_sol = score
        return min_sol

    for i in range(3):
        if col_check[i] == False:
            col_check[i] = True
            Question31(row + 1 , score + m[row][i])
            col_check[i] = False

    return min_sol


if __name__ == "__main__":

    print("min_sol : %d" %(Question34(0,0)))