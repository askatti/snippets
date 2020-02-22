import sys



def fib(n):
	fb = {}	
	for k in range(1, n+1):
		if k <= 2:
			f = 1
		else:
			f = fb[k-1] + fb[k-2]
		fb[k] = f
	return fb[n]


k = fib(int(sys.argv[1]))
print("fib of %d=%d",sys.argv[1],k)
