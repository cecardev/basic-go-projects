package main

import "fmt"

type pc struct {
	ram   int
	cpu   int
	brand string
}

func (mypc pc) show() {
	fmt.Println("Ram: ", mypc.ram, "GB")
	fmt.Println("CPU: ", mypc.cpu, "GHz")
	fmt.Println("Brand: ", mypc.brand)
}

func (mypc *pc) duplicateRam() {
	mypc.ram = mypc.ram * 2
}

func main() {
	a := 50
	b := &a
	fmt.Println(a, b)
	*b = *b + 1
	fmt.Println(a, b)

	myPc := pc{ram: 4, cpu: 8, brand: "Asus"}
	myPc.show()
	myPc.duplicateRam()
	myPc.show()
	myPc.duplicateRam()
	myPc.show()

}
