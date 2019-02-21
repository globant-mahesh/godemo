package main
//
//import (
//	"fmt"
//)
//
////func main() {
////	h := hyundai{color: "silver", model: "magna", maxspeed: 120.00}
////	a := audi{color: "white", maxspeed: 170.00}
////	h.accelerate()
////	a.accelerate()
////	a.applybrake()
////	a.switchgear()
////}
//
////Only method declaration is allowed in interface type & no variable declaration allowed ininterface
//type carrier interface {
//	accelerate()
//	applybrake()
//}
//
//type fourwheeler interface {
//	applybrake()
//	switchgear()
//}
//
//type hyundai struct {
//	color    string
//	model    string
//	maxspeed float64
//}
//
//type audi struct {
//	color    string
//	maxspeed float64
//}
//
////Implicit implementation of the interface in struct type i.e. syntax is not required to tell compiler that struct type
////implements interface type
//func (h hyundai) accelerate() {
//	fmt.Println("Accelerating Hyundai Car")
//}
//
//func (a audi) accelerate() {
//	fmt.Println("Accelerating Audi Car")
//}
//
////implicit implements makes it possible to realize multiple inheritance with no ambiguity in method reference
//func (a audi) applybrake() {
//	fmt.Println("Applying brakes")
//}
//
//func (a audi)switchgear() {
//	fmt.Println("Switching to next gear")
//}
