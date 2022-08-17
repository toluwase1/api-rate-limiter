package main

import (
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"time"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func main() {
	server := gin.Default()
	// This makes it so each ip can only make 5 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Hour * 24,
		Limit: 3,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
		//BeforeResponse: nil,
	})
	server.GET("/", mw, func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	server.Run(":8080")
}

//func HelloHandler(w http.ResponseWriter, req *http.Request) {
//	//increment the counter
//
//	_, err := w.Write([]byte("Hello, World!"))
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//}
//
//func main() {
//	// Create a request limiter per handler.
//	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(5, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute}), HelloHandler))
//	err := http.ListenAndServe(":12345", nil)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//}

//var (
//	limit ratelimit.Limiter
//	rps   = flag.Int("rps", 1, "request per second")
//)
//
//func init() {
//	log.SetFlags(0)
//	log.SetPrefix("[GIN] ")
//	log.SetOutput(gin.DefaultWriter)
//}
//
//func leakBucket() gin.HandlerFunc {
//	prev := time.Now()
//	return func(ctx *gin.Context) {
//		now := limit.Take()
//		log.Print(color.CyanString("%v", now.Sub(prev)))
//		prev = now
//	}
//}
//
//func ginRun(rps int) {
//	limit = ratelimit.New(rps)
//
//	app := gin.Default()
//	app.Use(leakBucket())
//
//	app.GET("/rate", func(ctx *gin.Context) {
//		ctx.JSON(200, "rate limiting test")
//	})
//
//	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", rps))
//	app.Run(":8080")
//}
//
//func main() {
//	flag.Parse()
//	ginRun(*rps * 10)
//} //
//func main() {
//	defer fmt.Println("exiting application")
//
//	apiConnection := apiConn.Open()
//
//	var wg sync.WaitGroup
//	wg.Add(20)
//
//	for i := 0; i < 10; i++ {
//		go func() {
//			defer wg.Done()
//
//			v, err := apiConnection.Read(context.Background())
//			if err != nil {
//				fmt.Printf("%v Get Error: %v\n", time.Now().Format("15:04:05"), err)
//				return
//			}
//
//			fmt.Printf("%v %v\n", time.Now().Format("15:04:05"), v)
//		}()
//	}
//
//	for i := 0; i < 10; i++ {
//		go func() {
//			defer wg.Done()
//
//			err := apiConnection.Resolve(context.Background())
//			if err != nil {
//				fmt.Printf("%v Resolve Error: %v\n", time.Now().Format("15:04:05"), err)
//				return
//			}
//
//			fmt.Printf("%v Resolved\n", time.Now().Format("15:04:05"))
//		}()
//	}
//
//	wg.Wait()
//}

//func main() {
//	rl := rate.New(3, time.Second) // 3 times per 24 hours
//	begin := time.Now()
//	for i := 1; i <= 10; i++ {
//		rl.Wait()
//		fmt.Printf("%d started at %s\n", i, time.Now().Sub(begin))
//	}
//	// Output:
//	// 1 started at 12.584us
//	// 2 started at 40.13us
//	// 3 started at 44.92us
//	// 4 started at 1.000125362s
//	// 5 started at 1.000143066s
//	// 6 started at 1.000144707s
//	// 7 started at 2.000224641s
//	// 8 started at 2.000240751s
//	// 9 started at 2.00024244s
//	// 10 started at 3.000314332s
//}

//func main() {
//	limitContainer := []int64{}
//	begin := time.Now()
//	rl1 := rate.New(1, time.Second)   // Once per second
//	rl2 := rate.New(1, time.Second*3) // 2 times per 3 seconds
//	for i := 1; i <= 10; i++ {
//
//		rl1.Wait()
//		rl2.Wait()
//		fmt.Printf("%d started at %s\n", i, time.Now().Sub(begin))
//	}
//	// Output:
//	// 1 started at 11.197us
//	// 2 started at 1.00011941s
//	// 3 started at 3.000105858s
//	// 4 started at 4.000210639s
//	// 5 started at 6.000189578s
//	// 6 started at 7.000289992s
//	// 7 started at 9.000289942s
//	// 8 started at 10.00038286s
//	// 9 started at 12.000386821s
//	// 10 started at 13.000465465s
//}

//var rl = rate.New(3, time.Second) // 3 times per second
//
//func say(message string) {
//	if ok, remaining := rl.Try(); ok {
//		fmt.Printf("You said: %s\n", message)
//	} else {
//		fmt.Printf("Spam filter triggered, please wait %s\n", remaining)
//	}
//}
//
//func main() {
//	for i := 1; i <= 5; i++ {
//		say(fmt.Sprintf("Message %d", i))
//	}
//	time.Sleep(time.Second / 2)
//	say("I waited half a second, is that enough?")
//	time.Sleep(time.Second / 2)
//	say("Okay, I waited a second.")
//	// Output:
//	// You said: Message 1
//	// You said: Message 2
//	// You said: Message 3
//	// Spam filter triggered, please wait 999.980816ms
//	// Spam filter triggered, please wait 999.976704ms
//	// Spam filter triggered, please wait 499.844795ms
//	// You said: Okay, I waited a second.
//}
