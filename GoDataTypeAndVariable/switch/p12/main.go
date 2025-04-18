package main

// switch i {
// 	case 0: //empty case body, nothing is executed when i==0
// 	case 1:
// 		f() // f is not called when i==0!
// 	}

// switch i {
// case 0:
// 	fallthrough
// case 1:
// 	f() // f is now also called when i==0!
// }

/*
In the first case, if i has the value 0, no code will be executed because case 0 has no code body,
and the switch terminates immediately. To obtain the same behavior in C, you need to add a break after case 0.
If, on the other hand, you explicitly want to execute the code from case 1, when i has the value 0, you need to
add the keyword fallthrough at the end of the case 0 branch. This is illustrated in the second case.
With fallthrough, all the remaining case branches are executed until a branch is encountered, which does
not contain a fallthrough.

Fallthrough can also be used in a hierarchy of cases where at each level something has to be done in addition
to the code already executed in the higher cases, and a default action also has to be executed. The (optional)
default branch is executed when no value is found to match var1 with; it resembles the else clause in if-else
statements. It can appear anywhere in the switch (even as the first branch), but it is best written as the last
branch.
*/
