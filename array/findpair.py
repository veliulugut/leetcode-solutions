def find_pair(arr: list[int], target: int) -> None:

    n = len(arr)

    found = False

    for i in range(n - 1):

        for j in range(i + 1, n):

            if arr[i] + arr[j] == target:

                print(f"Pair found ({arr[i]}, {arr[j]})")

                found = True

    if not found:
        print("Pair not found")


nums = [8, 7, 2, 5, 3, 1]
target = 10

find_pair(nums, target)
