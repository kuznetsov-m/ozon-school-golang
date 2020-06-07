result = False

with open('input.txt') as fp: 
    lines = fp.readlines()
    target = int(lines[0])
    values = lines[1].strip().split(' ')
    values.sort()

    first = 0
    last = len(values) - 1
    
    while (first < last):
        s = int(values[first]) + int(values[last])
        if s == target:
            result = True
            break
        else:
            if s < target: first += 1
            else: last -= 1

with open('output.txt', 'w') as f:
    f.write(str(int(result)))