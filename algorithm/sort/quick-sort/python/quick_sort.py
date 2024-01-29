from typing import List

def quick_sort(arr: List[int], low: int, high: int):
    def partition(arr, l: int, h: int) -> int:
        i = l
        pivot = arr[h]
        for j in range(l, h):
            if arr[j] < pivot:
                arr[i], arr[j] = arr[j], arr[i]
                i += 1
        arr[i], arr[h] = arr[h], arr[i]
        return i
    
    if low < high:
        pi = partition(arr, low, high)
        quick_sort(arr, low, pi-1)
        quick_sort(arr, pi+1, high)