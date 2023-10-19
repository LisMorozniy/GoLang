package main

type Boss struct {
    position string
    salary   int
    address  string
}

func (b *Boss) GetPosition() string {
    return b.position
}

func (b *Boss) SetPosition(position string) {
    b.position = position
}

func (b *Boss) GetSalary() int {
    return b.salary
}

func (b *Boss) SetSalary(salary int) {
    b.salary = salary
}

func (b *Boss) GetAddress() string {
    return b.address
}

func (b *Boss) SetAddress(address string) {
    b.address = address
}
