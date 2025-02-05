# Student API - HTTP/S Server in Golang

## Overview
This is a simple **Student Management API** built from scratch using **Golang**. The API allows managing **students, classes, and exams**, enabling operations such as retrieving, adding, and updating data via HTTP methods.
No frameworks will be used for implementing the http server

## Features
- **Student Management**: Add, retrieve, and update student details and marks.
- **Class Management**: Assign students to classes and manage class details.
- **Exam Management**: Create exams and manage student marks.
- **RESTful API**: Built following REST principles.
- **Secure HTTP/S Support**.

## Endpoints
### **Student Endpoints**
#### GET
- `GET /students/` → Retrieve a list of all students.
- `GET /students/{id}` → Retrieve details of a specific student.
- `GET /students/{id}/marks` → Retrieve marks of a student.
- `GET /students/{id}/history` → Retrieve the student's history (e.g., class changes, performance).

#### POST
- `POST /students/` → Create a new student.
- `POST /students/{id}/marks` → Add marks for a student.

#### PUT
- `PUT /students/{id}` → Update student details.

---

### **Class Endpoints**
#### GET
- `GET /classes/` → Retrieve all classes.
- `GET /classes/{id}` → Retrieve details of a specific class.
- `GET /classes/{id}/students` → Retrieve all students in a class.

#### POST
- `POST /classes/` → Create a new class.
- `POST /classes/{id}/students/{student_id}` → Assign a student to a class.

#### PUT
- `PUT /classes/{id}` → Update class details.

---

### **Exam Endpoints**
#### GET
- `GET /exams/` → Retrieve all exams.
- `GET /exams/{id}` → Retrieve details of a specific exam.
- `GET /exams/{id}/marks` → Retrieve all marks for an exam.

#### POST
- `POST /exams/` → Create a new exam.
- `POST /exams/{id}/marks/{student_id}` → Add or update a student's marks for an exam.

#### PUT
- `PUT /exams/{id}` → Update exam details.

## Getting Started
### **Requirements**
- **Go 1.19+**


---

🚀 Happy Coding!

