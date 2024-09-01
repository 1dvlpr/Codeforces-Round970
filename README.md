---

# 🚀 Codeforces Round 970 - My Solutions

Welcome to my solutions repository for **Codeforces Round 970**! Below you'll find my approach to each problem in this round. 

## 📚 Problem A: Sakurako's Exam

**Description:**

Sakurako has an array consisting of `a` ones and `b` twos. She needs to place either a `+` or a `-` in front of each element so that the sum of the array equals `0`.

**Input:**
- The number of test cases `t` (1 ≤ t ≤ 100).
- For each test case, two integers `a` and `b` (0 ≤ a, b < 10) represent the number of '1's and '2's in the array.

**Output:**
- "Yes" if it's possible to arrange the signs to make the sum equal to `0`, otherwise "No".

**Solution Overview:**

To solve this problem, I considered the possible sums of ones and twos under different sign arrangements. The challenge is to check if such an arrangement exists for the given values of `a` and `b`.

## 🟩 Problem B: Square or Not?

**Description:**

Sakurako creates a binary string `s` by writing down the rows of a beautiful binary matrix. You need to check if this string `s` could have been formed from a square matrix.

**Input:**
- The number of test cases `t` (1 ≤ t ≤ 10⁴).
- For each test case, an integer `n` (2 ≤ n ≤ 2 × 10⁵) and a string `s` of length `n`.

**Output:**
- "Yes" if the matrix could have been square, otherwise "No".

**Solution Overview:**

The problem requires checking whether the length of the binary string can form a square matrix, considering the constraints on the binary matrix's structure. I implemented a solution that efficiently checks for this property.

## 🔢 Problem C: Longest Good Array

**Description:**

A good array is an increasing array where the differences between adjacent elements also increase. Given boundaries `l` and `r`, the task is to find the maximum length of such a good array.

**Input:**
- The number of test cases `t` (1 ≤ t ≤ 10⁴).
- For each test case, two integers `l` and `r` (1 ≤ l ≤ r ≤ 10⁹).

**Output:**
- The length of the longest good array.

**Solution Overview:**

To find the maximum length of a good array, I analyzed the properties of such arrays and used the given boundaries to determine the length that meets the conditions.

---

Feel free to explore the code and see how I tackled these problems! If you have any questions or feedback, I'd love to hear from you. 😊

---
