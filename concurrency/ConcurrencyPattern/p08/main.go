package main

import "fmt"

func getNews(newsChannel chan string) {
	NewsArray := []string{"Roger Federer wins the Wimbledon", "Space Exploration has reached new heights", "Wandering cat prevents playground accident"}
	for _, news := range NewsArray {
		newsChannel <- news
	}

	newsChannel <- "Done"
	close(newsChannel)

}

func main() {
	myNewsChannel := make(chan string)

	go getNews(myNewsChannel)

	for {
		select {
		case news := <-myNewsChannel:
			fmt.Println(news)
			if news == "Done" {
				return
			}
		default:
			fmt.Println("No news available at the moment.")
		}

	}
}

/*
You might have wondered why the loop did not run infinitely. Well, that’s because I checked for my "Done"
signal and returned from the main routine in lines 23-25 as to prevent an infinite loop from running.

Now the for-loop below enables us to execute the select statement infinitely. This is pretty handy in
situations where we have to keep taking inputs from different channels.

The select statement is used to wait for multiple channel operations. It is similar to the switch statement
in Go. The select statement will block until one of its cases can execute
and then it will execute that case. It is important to note that the select statement will not wait for all
cases to be ready, it will only wait for one case to be ready.
*/
