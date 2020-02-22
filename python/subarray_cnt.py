""" find longest subarry cnt """
a = [1,10,11,12,3,13,2,4,5,6,6,7,8]
print a
fin_list = []
i= 0
sub_cnt = 0
a.sort()
print a
llen = len(a) -1
while i < llen:
 if ((( a[i+1] - a[i] ) == 1)  or (( a[i+1] - a[i]) == 0)):
  sub_cnt = sub_cnt + 1
 else:
  print "update cnt=",sub_cnt
  fin_list.append(sub_cnt+1)
  print "reset cnt"
  sub_cnt = 0
 i = i + 1
if sub_cnt:
 fin_list.append(sub_cnt+1)
print "fin_list=",fin_list
fin_list.sort()
print "longest sub_array=",fin_list[-1]
