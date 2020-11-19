
package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/compute/v1"
	"time"
)
	const project = "western-notch-185412"
	const region = "us-central1"
const zone = "us-central1-a"
const namea = "sasdxsdada"

type simple struct {
	Error *compute.OperationError `json: "error"`
	SelfLink  string `json:"selflink"`
	TargetLink string `json:"targetLink"`
	User string`json:user,omitempty`
	Code int64 `json:"code"`
	Message string `json:"message"`
}

type check1 struct {
	Process string `json:"process"`
	Laafa string `json:"laafa"`
}
type check struct {
	Find string `json:find,omitempty`
	Process string `json:"process"`
	Laafa string `json:"laafa"`
}

//func func1(data *check) *check1{
//	fmt.Println(*data)
//	return data
//}

func (data *check) func1() *check1{
	h := &check1{}
	res1B, _ := json.Marshal(data)
	json.Unmarshal([]byte(string(res1B)), &h)
	return h
}
//func (data *check) func1() interface{} {
//	fmt.Println(*data)
//	res1B, _ := json.Marshal(data)
//	res := check1{}
//	json.Unmarshal([]byte(string(res1B)), &res)
//	//fmt.Println(res)
//	return res
//}
func add() (x int,y int) {
	x = x +10
	y = y +10
	return  x,y
}

func main() {
	first :=&check{
		Find: "anubhav",
		Process: "dsfb",
		Laafa: "dsfbh",
	}
	lava := first.func1()
	fmt.Println(*lava,11)
	sum,r := add()
	fmt.Println(sum,r)
	finished := make(chan simple)
	go createnet(finished, namea)
	response := <-finished
	ad,_ := json.Marshal(response)
	fmt.Println(string(ad),2)
	time.Sleep(20 * time.Second)
	finished1 := make(chan simple)
	finished2 := make(chan simple)
	go createSubnet(finished1, "subnetwork", "10.0.1.0/24", response.TargetLink)
	go createSubnet(finished2, "subnetwork1", "10.0.2.0/25", response.TargetLink)
	response1 := <-finished1
	response2 := <-finished2
	ad1,_ := json.Marshal(response1)
	fmt.Println(string(ad1),2)
	ad2,_ := json.Marshal(response2)
	fmt.Println(string(ad2),2)

}