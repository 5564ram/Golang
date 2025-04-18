package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}

func main() {
	c := &Customer{"Barack Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me, the world will be a better place!")
	fmt.Println(c)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:\n" + fmt.Sprintln(c.Log)
}

/*
Barack Obama
Log:
{1 - Yes we can!
2 - After me, the world will be a better place!}
*/

/*
In the above code, at line 6, we make a struct of type Log with one field of string named msg.
At line 10, we make a struct of type Customer with two fields. The first is a string variable Name,
and the second is an anonymous variable of type Log. Before going to main, let’s see the other
three functions of the program:

See the header of Add() method at line 21: func(l *Log) Add(s string). This shows that this method
can be called by the pointer variable of type Log. It takes a string as a parameter. Look at line 22.
This method appends a new blank line and then adds a string s to the internal field msg of l.

See the header of String() method at line 25: func(l *Log) String()string. This shows that this method
can be called by the pointer variable of type Log. This method returns the msg of l.

See the header of String() method at line 29: func (c *Customer) String() string . This shows that this
method can be called by the pointer variable of type Customer. This method appends the Name of customer c
and its log (anonymous) and returns this final string.

Now, look at main. At line 16, we make a pointer variable of type Customer using struct-literal and
reinitialize c as: &Customer{"Barack Obama", Log{"1 - Yes we can!"}}. The part Log{"1 - Yes we can!"}
is used to assign value to the anonymous Log variable(msg gets 1 - Yes we can!), and Barack Obama is
assigned to the Name string. The symbol & makes it a pointer variable. In the next line, we are adding
a string to the msg of the log of c as: c.Add("2 - After me, the world will be a better place!").
The log of c will automatically call the Add() method. The string gets appended at the end of msg.
In the next line, we are printing c to see whether the msg was changed or not. The c calls the method
String(), and we’ll print the returned value. In this case, it’s:

Barack Obama

Log: {1 - Yes we can!

2 - After me, the world will be a better place!}
*/
