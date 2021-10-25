package main

import (  
    "fmt"
)


type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

type SalaryCalculator interface {  
    CalculateSalary() int
}

type CalculateArea interface{
	CalcArea() int
}
type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId    int
    basicpay int
}

type Rect struct{
	len int
	bre int
}

type Sqr struct{
	len int
}
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}


func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

func (r Rect) CalculateArea() int{
	return r.len*r.bre
}

func (s Sqr) CalculateArea() int{
	return s.len*s.len
}

func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func callArea(a CalculateArea){
	op:=0
	op=op+a.CalcArea()
	fmt.Printf("\n the area= %d",op)

}

func main() {  
    name := MyString("Sam Anderson")
    var v VowelsFinder
    v = name 
    fmt.Printf("Vowels are %c", v.FindVowels())
	pemp1 := Permanent{
        empId:    1,
        basicpay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicpay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicpay: 3000,
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

	r1:=Rect{
		len: 10,
		bre: 12,
	}
	s1:=Sqr{
		len: 10,
	}
	callArea(r1)
	callArea(s1)

}