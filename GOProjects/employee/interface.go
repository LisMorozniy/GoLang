package main

type Emlployee interface {
    GetPosition() string
    SetPosition(string)
    GetSalary() int
    SetSalary(int)
    GetAddress() string
    SetAddress(string)
}
