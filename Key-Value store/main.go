package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var GlobalStore = make(map[string]string)

type Transaction struct {
	store map[string]string
	next  *Transaction
}

type TransactionStack struct {
	top  *Transaction
	size int
}

func (ts *TransactionStack) PushTransaction() {

	temp := Transaction{store: make(map[string]string)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++

}

func (ts *TransactionStack) PopTransaction() {

	if ts.top == nil {
		fmt.Printf("ERROR: Transaction stack is empty")
	} else {
		node := &Transaction{}
		ts.top = ts.top.next
		node.next = nil
		ts.size--
	}

}

func (ts *TransactionStack) Peek() *Transaction {
	return ts.top
}

func (ts *TransactionStack) Commit() {
	ActivationTransaction := ts.Peek()

	if ActivationTransaction != nil {
		for key, value := range ActivationTransaction.store {
			GlobalStore[key] = value
			if ActivationTransaction.next != nil {
				ActivationTransaction.next.store[key] = value
			}
		}
	} else {
		fmt.Printf("INFO: Nothing to commit \n")
	}
}

func Set(key string, value string, T *TransactionStack) {
	ActivationTransaction := T.Peek()

	if ActivationTransaction == nil {
		GlobalStore[key] = value
	} else {
		ActivationTransaction.store[key] = value
	}
}

func Get(key string, T *TransactionStack) {
	ActivationTransaction := T.Peek()

	if ActivationTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
			fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set \n", key)
		}
	} else {
		if val, ok := ActivationTransaction.store[key]; ok {
			fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set \n", key)
		}
	}
}

func DeleteTransaction(key string, T *TransactionStack) {
	ActivationTransaction := T.Peek()
	if ActivationTransaction == nil {
		if _, ok := GlobalStore[key]; ok {
			delete(GlobalStore, key)
		} else {
			fmt.Printf("%s not set", key)
		}
	} else {
		if _, ok := ActivationTransaction.store[key]; ok {
			delete(ActivationTransaction.store, key)
		} else {
			fmt.Printf("%s not set", key)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	items := &TransactionStack{}
	for {
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		operation := strings.Fields(text)

		switch operation[0] {
		case "BEGIN":
			items.PushTransaction()
		case "PEEK":
			items.Peek()
		case "COMMIT":
			items.Commit()
			items.PopTransaction()
		case "SET":
			Set(operation[1], operation[2], items)
		case "GET":
			Get(operation[1], items)
		case "DELETE":
			DeleteTransaction(operation[1], items)
		case "END":
			items.PopTransaction()
		case "STOP":
			os.Exit(0)

		default:
			fmt.Printf("ERROR: Unrecognised Operation %s\n", operation[0])
		}
	}
}
