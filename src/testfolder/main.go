
package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/api/compute/v1"
	"golang.org/x/net/context"
)
	const projectName = "western-notch-185412"
	const regionName = "us-central1"
const zoneValue = "us-central1-a"
const networkName = "asnesubhadsdwsadav"
const subnetworkName = "networksub1"
const subnetworkNametwo = "networksub2"
type response struct {
	Error *compute.OperationError `json: "error"`
	SelfLink  string `json:"selflink"`
	TargetLink string `json:"targetLink"`
	User string`json:user,omitempty`
	Code int64 `json:"code"`
	Message string `json:"message"`
	Id string `json:"id"`
	isError bool `json: "iserror"`
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
	//first :=&check{
	//	Find: "anubhav",
	//	Process: "dsfb",
	//	Laafa: "dsfbh",
	//}
	//lava := first.func1()
	//fmt.Println(*lava,11)
	//sum,r := add()
	//fmt.Println(sum,r)
	ctx := context.Background()
	finished := make(chan *response)
	go CreatenetSubNetMachine(ctx, finished, networkName,
		true, true, true, projectName)
	response := <-finished
	ab,_ := json.Marshal(response)
	fmt.Println(string(ab))

}