// Just like try/catch block in exception in languages like Java, C#, etc.
// are used to catch exception similarly in Go language, recover function is used to handle panic
import(
  "fmt"
  "log"
)

// NOTE :Recover function is always called inside the deferred function
// because deferred func didnot stop at panic

//Recover function only work if you call in the same goroutine in which panic occurs.
//If you call it in a different goroutine, then it will not work

func main()  {
  fmt.Println("Start")
  paniker()
  fmt.Println("End")
}


func paniker()  {
  fmt.Println("about to start panic")
  defer func ()  {
    if err := recovery(); err != nil{ // recovery is like catch func and we say we will handle this panic inthis block
      log.Println("ERROR :",err)// if will not to able handle that panic here we will call painc again with errors
      // panic(err)
    }
  }()
  panic("Something bad happened")
  fmt.Println("done panicking") // you will see here we are done with pnicking but at the time code will run this line will not Print
  // reson behind this panicker func will stop executation and it will execute any deferred func or statement
  //and comeout with recovery and execute the left part of main functions
  // which will print "end"
}
