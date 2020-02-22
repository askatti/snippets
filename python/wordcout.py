#!/usr/bin/python
import sys
import json
wcount = {}
def wordcount(file):
    with open(file) as f:
        for line in f:
            wordlist = line.split(" ")
            for i in range(len(wordlist)):
                    if wordlist[i] in wcount:
                        wcount[wordlist[i]] = wcount[wordlist[i]] + 1
                    else:
                        wcount[wordlist[i]] = 1

wordcount(sys.argv[1])
print(json.dumps(wcount, indent=1))
