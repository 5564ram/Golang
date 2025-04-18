package main

import "fmt"

func updatePosition(name string) <-chan string {
	positionChannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionChannel <- fmt.Sprintf("%s %d", name, i)
		}
	}()

	return positionChannel
}

func main() {
	positionChannel1 := updatePosition("Legolas :")
	positionChannel2 := updatePosition("Gandalf :")

	for i := 0; i < 5; i++ {
		fmt.Println(<-positionChannel1)
		fmt.Println(<-positionChannel2)
	}

	fmt.Println("Done with getting updates on positions.")
}

/*
In the code above, we are getting updates on the position of Legolas and Gandalf using the updatePosition function.
We again launch the goroutine from inside the function, which is the key thing to notice in both the examples we have
seen so far. Thus by returning a channel, the function enables us to communicate with the service it provides which,
in this case, is giving position updates.

Moreover, we can have more than one instance of the function, which can also be seen in the example above, as we get
position updates for both Legolas and Gandalf.

However, we still have a slight problem, which is when the following statements are blocking each other:
fmt.Println(<-positionChannel1)
fmt.Println(<-positionChannel2)
*/
