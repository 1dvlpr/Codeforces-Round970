# 🚀 Codeforces Round 970 - My Solutions

Welcome to my solutions repository for **Codeforces Round 970**! This README provides an overview of the solutions to the problems I tackled during this round. If you're looking to understand how to solve these problems, you're in the right place.

## 📚 Problem A: Sakurako's Exam

**Problem Description:**

Sakurako has a math exam and is given an array consisting of `a` ones and `b` twos. She needs to place either a `+` or a `-` in front of each element so that the sum of the array equals `0`.

**Input:**
- Number of test cases `t` (1 ≤ t ≤ 100).
- For each test case, two integers `a` and `b` (0 ≤ a, b < 10), representing the number of '1's and '2's in the array.

**Output:**
- "Yes" if it's possible to arrange the signs to make the sum equal to `0`.
- "No" otherwise.

**Solution Overview:**

To solve this problem, I analyzed the sum of the ones and twos under different sign arrangements. The goal is to determine if it's possible to balance the array's sum to zero for each test case.

## 🟩 Problem B: Square or Not?

**Problem Description:**

Sakurako creates a binary string `s` by writing down the rows of a beautiful binary matrix. You need to check if this string could have been formed from a square matrix.

**Input:**
- Number of test cases `t` (1 ≤ t ≤ 10⁴).
- For each test case:
  - An integer `n` (2 ≤ n ≤ 2 × 10⁵), which is the length of the string.
  - A string `s` of length `n`.

**Output:**
- "Yes" if the matrix could have been square.
- "No" otherwise.

**Solution Overview:**

This problem requires determining whether the string can represent a square matrix. I implemented a solution to check the validity based on the properties of the string `s` and the given constraints.

## 🔢 Problem C: Longest Good Array

**Problem Description:**

A good array is one where:
1. The array is increasing (i.e., `a[i-1] < a[i]` for all `2 ≤ i ≤ n`).
2. The differences between adjacent elements are increasing (i.e., `a[i] - a[i-1] < a[i+1] - a[i]` for all `2 ≤ i < n`).

Given boundaries `l` and `r`, the task is to find the maximum length of such a good array.

**Input:**
- Number of test cases `t` (1 ≤ t ≤ 10⁴).
- For each test case:
  - Two integers `l` and `r` (1 ≤ l ≤ r ≤ 10⁹).

**Output:**
- The length of the longest good array for the given `l` and `r`.

**Solution Overview:**

This problem involves finding the maximum length of a good array that fits within the given boundaries. I approached the solution by analyzing the constraints and implementing a strategy to generate the array with the maximum possible length.

---

