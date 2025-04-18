package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Barack Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter:
	c = &Customer{"Barack Obama", &Log{"1 - Yes we can!"}}
	fmt.Println(c.log)
	c.Log().Add("2 - After me, the world will be a better place!")
	fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

/*
1 - Yes we can!
1 - Yes we can!
2 - After me, the world will be a better place!
*/

/*
In the above code, at line 6, we make a struct of type Log with one field of string named msg.
At line 10, we make a struct of type Customer with two fields. The first is a string variable Name,
and the second is a pointer variable log, of type Log. Before going to main, let’s see the other
three functions of the program:

See the header of Add() method at line 27: func(l *Log) Add(s string). This shows that this method can
be called by the pointer variable of type Log. It takes a string as a parameter. Look at line 28.
This method appends a new blank line and then adds a string s to the internal field msg of l.

See the header of String() method at line 31: func(l *Log) String()string. This shows that this method
can be called by the pointer variable of type Log. This method returns the msg of l.

See the header of String() method at line 35: func (c *Customer) Log() *Log . This shows that this
method can be called by the pointer variable of type Customer. This method returns the log of c.

Now, look at main. At line 16, we make a Customer type variable c using the new() function.

In the next few lines, we are giving the fields their values. At line 17, we are giving
Name( a field of Customer c) a value Barack Obama. In the next line, we make a new Log type variable
and assign it to c.log. Every variable of type Log has a msg. So, at line 19, we give the value to
msg of c.log as: c.log.msg = "1 - Yes we can!". It took us four lines to make a proper Customer object.

How about doing it in a single line? Look at line 21. We make a pointer variable of type Customer
using struct-literal and reinitialize c as: &Customer{"Barack Obama", &Log{"1 - Yes we can!"}}.
The part &Log{"1 - Yes we can!"} is used to assign a value to an anonymous Log pointer
variable(msg gets 1 - Yes we can!), and Barack Obama is assigned to the Name string.
The symbol & makes it a pointer variable.

See line 22. We are printing the c.log. This will call the String() function, and we’ll print the
return value. In this case, it’s 1 - Yes we can!. In the next line, we are adding a string to the
msg of the log of c as: c.Log().Add("2 - After me, the world will be a better place!").
The part c.Log() calls the method Log(), and the pointer to log of c is returned, which then
calls the Add() method along with the string. The string gets appended at the end of msg of c.Log()(c.log).
In the next line, we are printing c.Log() to see whether the msg was changed or not. The c.Log()
calls the method Log(), and the pointer to log of c is returned, which will call the String() function.
Then, we’ll print the return value. In this case, it’s:

1 - Yes we can!
2 - After me, the world will be a better place!
*/
