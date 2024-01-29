import pytest
from quick_sort import quick_sort

@pytest.mark.parametrize(("input", "expected"), [
    [[],[]],
    [[1],[1]],
    [[3,1,2], [1,2,3]],
    [[1,2,3,4,5], [1,2,3,4,5]],
    [[5,4,3,2,1], [1,2,3,4,5]],
    [[1,1,1,1], [1,1,1,1]],
    [[4,1,3,2], [1,2,3,4]],
    [[4,1,3], [1,3,4]],
    [[1000,2000,3000,4000], [1000,2000,3000,4000]],
    [[-1, -3, -2, -4, -5], [-5, -4, -3, -2, -1]],
    [[-1, 2, -3, 4, -5], [-5, -3, -1, 2, 4]],
    [[-1, 2, -3, 4, -5], [-5, -3, -1, 2, 4]],
    [[3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5], [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]],
    [[1, 2, 3, 2, 1, 3, 2], [1, 1, 2, 2, 2, 3, 3]],
    [[5, 3, 9, 1, 6, 7, 2, 4, 8], [1, 2, 3, 4, 5, 6, 7, 8, 9]],
    [[-100, 0, 100, -50, 50], [-100, -50, 0, 50, 100]],
])
def test_quick_sort(input, expected):
    quick_sort(input, 0, len(input)-1)
    assert input == expected