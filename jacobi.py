#!/usr/bin/python3
''' A program to compute the roots of a system of linear equations using the method of Jacobi
'''
from itertools import permutations


class Jacobi:
    def __init__(self, m, n, dp, tol):
        self.m = m
        self.n = n
        self.A = [[0.0 for _ in range(n)] for _ in range(m)]
        self.B = [0.0 for _ in range(m)]
        self.guess = [0.0 for _ in range(n)]
        self.dp = dp
        self.tol = tol

    def initialize_A(self):
        for i in range(self.m):
            for j in range(self.n):
                print(f"Enter entry A{i + 1}{j + 1}:\n")
                self.A[i][j] = float(input())

    def initialize_B(self):
        for j in range(self.m):
            print(f"Enter B{j + 1}")
            self.B[j] = float(input())

    def isDiagonallyDominant(self, matrix):
        for i in range(self.m):
            row_sum = 0
            diag_value = 0
            for j in range(self.n):
                diag_value += abs(matrix[i][i])
                if i != j:
                    row_sum += abs(matrix[i][j])
            if diag_value < row_sum:
                return False
        return True

    def forceDiagonalDominance(self):
        perms = list(permutations(self.A))
        for perm in perms:
            if self.isDiagonallyDominant(perm):
                self.A = [list(row) for row in perm]
                print("Matrix is now diagonally dominant")
                print(self.A)
                return
            print("Matrix couldn't be made diagonally dominant")

    def solveJacobi(self):
        x_new = self.guess[:]
        x_old = self.guess[:]

        iterCount = 0
        while True:
            for i in range(self.m):
                sum_Ax = sum(self.A[i][j] * x_old[j]
                             for j in range(self.n) if i != j)
                x_new[i] = round(
                    ((self.B[i] - sum_Ax) / self.A[i][i]), self.dp)

            # Check for convergence
            err = max(abs(x_new[i] - x_old[i]) for i in range(self.m))
            if err < self.tol:
                break
            x_old = x_new[:]
            iterCount += 1
            # Print the final result after convergence
            print(f"Iterations {iterCount}: {x_new}")

        return x_new, iterCount


# Input
rows = int(input("Number of rows?: "))
cols = int(input("Number of columns?: "))
dp_accuracy = int(input("Decimal point accuracy? "))
tol = float(input("Tolerance? "))


# Driver code
if __name__ == '__main__':
    j = Jacobi(rows, cols, dp_accuracy, tol)
    j.initialize_A()
    print("Is the initial matrix diagonally dominant?",
          j.isDiagonallyDominant(j.A))
    j.forceDiagonalDominance()
    j.initialize_B()
    result, iters = j.solveJacobi()
    print(f"Solved in {iters} iterations. Solution: {result}")
