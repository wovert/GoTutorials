package main

/*
#include <stdio.h>
void SayHello() {
	printf("C Program Language");
}
*/
import "C"

func main() {
	C.SayHello()
}
