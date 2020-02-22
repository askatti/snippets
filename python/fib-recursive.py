import sys



def fib(n):
	if n <=2:
		f =1
	else: 
		f= fib(n-1) + fib(n-2)
	return f

print ( sys.argv[1],"th element of fib series=", fib(int(sys.argv[1])))
