import sys

for line_ in sys.stdin:
    line = line_.rstrip()
    values = line.split(' ')
    print(int(values[0]) + int(values[1]))