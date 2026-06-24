package main

import "fmt"

func hasCycle( graph map[string]string)

func main() {
	fmt.Println("adjacency matrix")

	graph:=map[string][]string{
		"A":{"B","C"},  
		"B":{"A","C"},  
		"C":{"A","D"},   
		"D":{"B"},     
	}     
	for c, v := range graph {
		fmt.Println(c,"->",v)        
	}        
   
	fmt.Println("A is conected to",graph["A"])
}     
