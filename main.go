package main

import (
	"cmp"
	"fmt"

	"github.com/Gorillarock/list/list"
)

func main() {
	example_usage()
}

func example_usage() {
	strs := new(list.List[string]).Add("echo")
	strs.Add("delta")
	strs.Add("alpha").Add("charlie")
	strs.Add("bravo").Add("alpha")

	fmt.Printf("String List Example (size: %d):\n", strs.Size())
	print_list(strs, "\t")
	fmt.Printf("Find strings:\n")
	find_print(strs, "bravo")
	find_print(strs, "goob")

	nums := new(list.List[int]).Add(5, 2, 8, 3)
	nums.Add(5) // Won't be added, since it is a duplicate

	fmt.Printf("\nNumber list Example (size: %d):\n", nums.Size())
	print_list(nums, "\t")

	fmt.Printf("\nFind numbers:\n")
	find_print(nums, 5)
	find_print(nums, 7)
	find_print(nums, 8)

	flts := list.New(-5.7, -10, -35.6, 23.2)
	print_list(flts, "\t")

	fmt.Printf("\nPrinting with Values func\n")
	fmt.Println(flts.Values())
	fmt.Println("Removing \"-10\"")
	flts.Remove(-10)
	fmt.Println(flts.Values())

	fmt.Printf("\nFloats list Example (size: %d)\n", flts.Size())
	find_print(flts, 5)
	find_print(flts, flts.Last().Val)
	find_print(flts, flts.Start().Val)

	// Indexed list example
	fmt.Printf("\nIndexed List Example:\n")
	indexedList := list.NewIndexed("red", "blue", "green", "yellow")
	print_list(indexedList, "\t")
	find_print(indexedList, "red")
	fmt.Println("Deleting \"yellow\".")
	indexedList.Remove("yellow")
	find_print(indexedList, "yellow")
	find_print(indexedList, "red")
	find_print(indexedList, "green")
	find_print(indexedList, "blue")
}

func find_print[t cmp.Ordered](l *list.List[t], val t) {
	n := new(list.Node[t])
	if l == nil {
		goto not_found
	}
	fmt.Printf("--find memory address of: %v\n", val)
	n = l.Find(val)
	if n != nil {
		fmt.Printf("  found node: %p with value %v\n", n, n.Val)
		print_node_details(l, n)
		println()
		return
	}

not_found:
	fmt.Printf("  not found\n\n")
}

func print_node_details[t cmp.Ordered](l *list.List[t], n *list.Node[t]) {
	if n != nil && l != nil {
		if l.Start() == n {
			fmt.Printf("\tthis is the first node\n")
		} else {
			fmt.Printf("\tthe node before has a value of: %v\n", n.Prev().Val)
		}

		if l.Last() == n {
			fmt.Printf("\tthis is the last node\n")
		} else {
			fmt.Printf("\tthe node after has a value of: %v\n", n.Next().Val)
		}
	}
}

func print_list[t cmp.Ordered](l *list.List[t], f string) {
	for tw := l.Start(); tw != nil; tw = tw.Next() {
		fmt.Printf(f+"%v\n", tw.Val)
	}
}
