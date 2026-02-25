# exam-08-golang

---

## INDEX

- [ภาพรวมโปรเจค](#ภาพรวมโปรเจค)
- [Features](#features)
- [ขั้นตอนการติดตั้ง](#installation)
- [Structure](#โครงสร้างโปรเจค)
---

## ภาพรวมโปรเจค

**exam No 8 Backend** 
มีฟีเจอร์ แสดงข้อมูล , เพิ่ม และ ลบ ข้อสอบ

---

## Features
-  **Create** - เพิ่มข้อมูลข้อสอบ
-  **Get** - ดึงรายการข้อมูลข้อสอบ
-  **Delete** - ลบข้อมูล

---

### Core
- **languagr** - GO
- **DB** - Maria DB

---

## การเริ่มต้นใช้งาน

### Prepare

```
go lang >= 1.25
docker-compose.yml

```

**enviroment** 
compose.yml

### Installation

``` bash
# 1. Install 
go mod tidy

# 2. docker
make docker-up

# 3. Start development server
make run
```

## โครงสร้างโปรเจค
```
exam-08-golang/
├── app/
│    
├──cmd/
│   └──server
│        ├── main.go
│        └── server.go
├──config/
│   ├── app.go
|   ├── config.go
|   ├── constant.go
|   ├── database.go
|   ├── datastore.go
|   └── server.go
├──model/
│   ├── choice.go
|   ├── exam.go
|   └── model.go
|
├── module/
|       ├── frontweb/
|       |   ├── dto
|       |   ├── handler
|       |   ├── mapper/
|       |   ├── ports
|       |   ├── repositories/
|       |   └── service/
|       └── frontweb.go
├── util/
└── Makefile

```

