from __future__ import division
import datetime
import sys


converted_d1 = datetime.datetime.fromtimestamp(float(str(sys.argv[1])))
converted_d2 = datetime.datetime.fromtimestamp(float(str(sys.argv[2])))

print((converted_d2 - converted_d1))
print((converted_d2 - converted_d1).total_seconds() / 60)
