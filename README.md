
# Student Grading Program

This program is designed to read student records from a CSV file, calculate their final scores, assign grades, and determine the topper for each university and overall. The data and operations are implemented in Go, and the program processes a CSV file containing student information and test scores.

## Features
- **Parse CSV File**: Reads a CSV file containing student data including names, universities, and test scores.
- **Calculate Final Score**: For each student, the final score is calculated as the average of their four test scores.
- **Assign Grades**: Grades are assigned based on the final score using the following scale:
  - `F`: Final score < 35
  - `C`: Final score >= 35 and < 50
  - `B`: Final score >= 50 and < 70
  - `A`: Final score >= 70
- **Overall Topper**: Identifies the student with the highest final score across all universities.
- **Topper Per University**: Identifies the topper for each university.

## Setup

To get started with this program, follow these steps:

### 1. Clone the Repository
Fork this repository to your own GitHub account, then clone it locally:

```bash
$ git clone https://github.com/HrithikSawant/go-student-grading.git
```
### 2. Run the program test:
   ```bash
   $ go test .
   ```


## Program Workflow

### Parsing the CSV File
The `grades.csv` file is parsed into a slice of `student` structs. Each struct holds the first name, last name, university, and test scores of a student.

### Calculating Final Scores
The final score for each student is calculated by averaging the four test scores. The final score is used to determine the student's grade.

### Assigning Grades
Grades are assigned based on the final score:
- `F` if the score is less than 35.
- `C` if the score is between 35 and 49.
- `B` if the score is between 50 and 69.
- `A` if the score is 70 or higher.

### Finding Topper for Each University
The program identifies the student with the highest final score for each university and stores them in a map where the key is the university name, and the value is the student object.

### Finding Overall Topper
The program identifies the student with the highest final score across all universities and returns the corresponding student.

## Example CSV Format

The `grades.csv` file should be formatted as follows:

```csv
FirstName,LastName,University,Test1,Test2,Test3,Test4
John,Doe,Harvard,50,60,65,45
Jane,Smith,Harvard,60,70,80,90
Alice,Johnson,MIT,50,55,60,45
Bob,Brown,MIT,40,50,45,35
```
