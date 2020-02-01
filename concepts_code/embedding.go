// embedding is as close as you can get to inheritance in Go. But there are some things to watch out for
// we are listing here

// that can be achive by struct embedding

package main

import (
  "fmt"
)

type Animal struct{
  Name string
  Origin string
}

type Bird struct {
  Animal
  speedKMPH int
  canFly bool
}

func main()  {
   b := Bird{} // initiating empty sturct
   // entering emdedded stuct key names
   b.Name = "Emu"
   b.Origin = "Australia"
   // birds fields
   b.speedKMPH = 30
   b.canFly = false
   //Other way is as follwing you can try by commenting above and uncommenting following

   // b := Bird{
   //   Animal : Animal{Name:"emu",Origin:"Australia"}
   //   speedKMPH : 30
   //   canFly : false
   // }

   fmt.Println(b.Name)
   fmt.Println(b) // see the structure that is different but access is same as that
   fmt.Println(b.canFly)

}
