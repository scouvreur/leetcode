from typing import List

valid_board = [
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
]

invalid_board = [
    ["8", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
]

invalid_board_2 = [
    [".", ".", ".", ".", "5", ".", ".", "1", "."],
    [".", "4", ".", "3", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "3", ".", ".", "1"],
    ["8", ".", ".", ".", ".", ".", ".", "2", "."],
    [".", ".", "2", ".", "7", ".", ".", ".", "."],
    [".", "1", "5", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "2", ".", ".", "."],
    [".", "2", ".", "9", ".", ".", ".", ".", "."],
    [".", ".", "4", ".", ".", ".", ".", ".", "."],
]


def remove_empties(flat_array: List[str]) -> List[str]:
    """
    Remove empty cells from rows.
    """
    return list(filter(lambda x: x != ".", flat_array))


def check_rows(board: List[List[str]]) -> bool:
    """
    Check horizontal rows.
    """
    rows_valid = True
    for row in board:
        if len(set(remove_empties(row))) != len(remove_empties(row)):
            rows_valid = False
            break
    return rows_valid


def check_columns(board: List[List[str]]) -> bool:
    """
    Check vertical columns.
    """
    columns: List[List[str]] = []
    for i in range(9):
        column = []
        for row in board:
            column.append(row[i])
        columns.append(column)

    return check_rows(columns)


def check_grids(board: List[List[str]]) -> bool:
    """
    Check 3x3 subgrids.
    """
    grids: List[List[str]] = []
    for i in range(3):
        for j in range(3):
            grid = []
            for row in board[i * 3 : (i + 1) * 3]:
                grid.extend(row[j * 3 : (j + 1) * 3])
            grids.append(grid)

    return check_rows(grids)


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        if not check_rows(board):
            return False
        elif not check_columns(board):
            return False
        elif not check_grids(board):
            return False
        else:
            return True


def test_solution():
    assert Solution().isValidSudoku(valid_board) is True
    assert Solution().isValidSudoku(invalid_board) is False


if __name__ == "__main__":
    test_solution()
