package main

import (
	"fmt"

	"github.com/pjebs/optimus-go"
)

func main() {
	o := optimus.New(1580030173, 59260789, 1163945558) //Prime Number: 1580030173, Mod Inverse: 59260789, Pure Random Number: 1163945558

	new_id := o.Encode(15) //internal id of 15 being transformed to 1103647397
	fmt.Println(new_id)

	orig_id := o.Decode(new_id) //Returns 15 back
	fmt.Println(orig_id)

}
