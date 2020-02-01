package main

import (
  "fmt"
)
// here we are using iota as a multiplier to convert file size in human readable

const (
  _ = iota // assigned to ignore this value and init iota at 0
  KB = 1 << (10*iota) // 0000000001 << 10 places => 10000000000 => 2^10 > 1024
  MB
  GB
  TB
  PB
  EB
  ZB
)


func main()  {
  filesize := 500000000
  fmt.Printf("%.2f GB",filesize/GB)

  fmt.Printf("%v   %T",MB,MB)
  fmt.Printf("%v   %T",KB,KB)
  fmt.Printf("%v   %T",TB,TB)
  fmt.Printf("%v   %T",PB,PB)

}
