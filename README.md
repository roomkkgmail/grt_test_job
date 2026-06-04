# Project Overview
This project is a collection of Go (Golang) modules designed to practice and demonstrate various Go concepts, including concurrency, thread safety, interfaces, logic/algorithms, and HTTP handlers. Each module is intended to be implemented in its own directory with comprehensive testing.

1. Basic Concurrency: Worker Pool 
โจทย์: จงเขียนฟังก์ชันชื่อ RunWorkers โดยรับ parameter 2 ตัวคือ numWorkers (จำนวนคนงาน) และ numJobs (จำนวนงานทั้งหมด)
ให้สร้าง Workers จำนวน numWorkers ตัว ที่ทำงานพร้อมๆ กัน (Concurrent) 
แต่ละ Worker จะคอยรับงานจาก Channel แล้วพิมพ์ข้อความว่า "Worker [{ID}] processing job [{Job ID}]"
ให้จำลองการทำงานด้วย time.Sleep เล็กน้อย(เช่น 1 วินาที) 
Main function ต้องรอให้ทุกงานถูกทำ จนเสร็จสิ้นจริงๆ ก่อนถึงจะจบโปรแกรม

2. Thread-Safety: Safe Counter 
โจทย์: จงสร้าง struct ชื่อ SafeCounter ที่ภายในเก็บค่ำ count (int) และมี Methods 2 ตัวดังนี้: 
Inc(): สำหรับเพิ่มค่า count ทีละ 1
Value(): สำหรับคืนค่า count ปัจจุบันออกมา
เงื่อนไข: โค้ดนี้ต้องรองรับกรณีที่มี Goroutines จำนวนมาก(เช่น 1,000 routines) เรียกใช้Inc() พร้อมกัน โดยที่ค่าสุดท้ายต้องถูกต้องแม่นยำ ห้ามเกิด Race Condition

3. Interfaces: ระบบคำ นวณพ้ืนที่ (Shape)
โจทย์: ให้ประกาศ(Define) Interface ชื่อ Shape ที่มี method ชื่อ Area() float64 สร้าง Struct 2 ตัวชื่อ Rectangle ( มี field Width, Height) และ Circle ( มี field Radius) 
Implement method Area ให้กับ Struct ทั้งสองเพื่อคำนวณพื้นที่
เขียนฟังก์ชันแยกออกมา 1 ตัว ชื่อ PrintArea(s Shape) ที่รับ Shape รูปทรงไหนก็ได้แล้วพิมพ์ขนาดพื้นที่ออกมา

4. Logic & Map: หาคู่ตัวเลข(Two Sum) 
โจทย์: กำหนดให้มี Slice ของตัวเลข nums := []int{2, 7, 11, 15} และค่าเป้าหมาย target := 9  จงเขียนฟังก์ชันที่รับ nums และ target แล้วคืนค่ำ index ของตัวเลข 2 ตัวใน slice ที่บวกกันแล้วได้เท่ากับ target พอดี สมมุติว่ามีคำตอบที่ถูกต้องแน่นอนเพียง 1 คู่
เงื่อนไข: ห้ามใช้ Loop ซ้อน Loop (Double for-loop) ต้องใช้วิธีที่มีประสิทธิภาพ Time Complexity  ดีกว่า O(n^2)

5. HTTP Handler: สร้าง JSON API ง่ายๆ 
โจทย์: จงใช้ package net/http เขียน Web Server ที่รันบน port 8080 และมี endpoint ชื่อ /hello  โดยมีเงื่อนไขดังนี้:
ต้องรับ Request Method เป็น POST เท่ำนั้น
รับ Body เป็น JSON: {"name": "Somchai"}
แกะค่า name ออกมาแล้วตอบกลับ (Response) เป็น JSON: {"message": "Hello Somchai"} กรณีที่ Method ไม่ใช่ POST หรือส่ง JSON มาผิด format ให้ return HTTP Error Code ที่เหมาะสมกลับไป

# Project Manual

This guide provides instructions on how to set up, use, and test the Go practice modules in this project.

## Prerequisites
- **Go:** Version 1.18 or higher is recommended.
- **Git:** To clone the repository.

## Getting Started

1. **Initialize the Go Module:**
   If the project doesn't have a `go.mod` file yet, initialize it:
   ```bash
   go mod init grt_test
   ```

2. **Install Dependencies:**
   This project uses `testify` for assertions. Fetch it using:
   ```bash
   go get github.com/stretchr/testify
   go mod tidy
   ```

## Running the Interactive Menu

The `main.go` file serves as the project's central interactive entry point. It provides a text-based UI menu that allows you to manually trigger and experiment with each module (e.g., specifying worker counts, incrementing counters, or starting the API server) without writing code.

To start the menu:
```bash
go run main.go
```

## Running Tests

All exercises are designed to be verified through tests. **Note: Use three dots (`...`) for recursive testing.**

### Run All Tests
To run tests for every module in the project:
```bash
go test ./...
```

### Run Tests with Race Detection
Since some modules involve concurrency (Worker Pool, Safe Counter), it is highly recommended to run tests with the race detector enabled:
```bash
go test -race ./...
```

### Run Tests for a Specific Module
Navigate to the module directory or specify the path:
```bash
go test ./workerpool/...
```

## Generating Coverage Reports

To generate an HTML report to visualize which parts of your code are covered by tests:

1. **Generate the coverage profile:**
   ```bash
   go test -coverprofile=coverage.out ./...
   ```

2. **Convert the profile to HTML:**
   ```bash
   go tool cover -html=coverage.out -o coverage.html
   ```

3. **Open the report:**
   ```bash
   open coverage.html
   ```

## How to Use/Implement

Each module should be implemented in its own directory.

1. **Create the Directory:** e.g., `mkdir workerpool`
2. **Implement the Logic:** Create a `.go` file (e.g., `workerpool/workerpool.go`).
3. **Write Table-Driven Tests:** Create a test file (e.g., `workerpool/workerpool_test.go`) using the `testify` library.
4. **Verify:** Run `go test ./workerpool` to ensure your implementation meets the requirements.

## Project Structure Reference
- `workerpool/`: Concurrency practice.
- `safecounter/`: Mutex and thread-safety.
- `shapes/`: Interface implementation.
- `twosum/`: Algorithmic optimization.
- `api/`: Web server and JSON handling.
