def Sort(data: list, l: int, r: int):
    if l >= r:
        return
    i, j = l - 1, r + 1
    midValue = data[(l + r) // 2]

    while i < j:
        while True:
            i += 1
            if data[i] >= midValue:
                break
        while True:
            j -= 1
            if data[j] <= midValue:
                break
        if i < j:
            data[i], data[j] = data[j], data[i]
    # print('No.1', data, midValue)
    Sort(data, l, j)
    # print('No.2', data, midValue)
    Sort(data, j + 1, r)


if __name__ == '__main__':
    N = int(input())
    data = list(map(int, input().split()))
    # data = [3, 1, 2, 4, 5]
    # print(data)
    # print(len(data) - 1)
    Sort(data, 0, len(data) - 1)
    for i in data:
        print(i, end=' ')
