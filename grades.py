#!/usr/bin/python3

# https://www.hackerrank.com/challenges/grading/problem
def gradingStudents(grades):
    rounded_grades = []
    for grade in grades:
        difference_from_next_multiple = 5 - ((grade + 5) % 5)
        if grade < 38:
            rounded_grades.append(grade)
        elif difference_from_next_multiple < 3:
            rounded_grades.append(grade + difference_from_next_multiple)
        else:
            rounded_grades.append(grade)
    return rounded_grades


print(gradingStudents([20, 36, 39, 41, 43, 74, 75, 76, 98, 100]))
