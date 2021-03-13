package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var priceStr string
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the first item's name:")
	firstItemName, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	firstItemName = strings.TrimSpace(firstItemName)

	fmt.Println("Enter the first item's price:")
	priceStr, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	firstItemPrice, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter the second item's name:")
	secondItemName, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	secondItemName = strings.TrimSpace(secondItemName)

	fmt.Println("Enter the second item's price:")
	priceStr, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	secondItemPrice, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter the third item's name:")
	thirdItemName, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	thirdItemName = strings.TrimSpace(thirdItemName)

	fmt.Println("Enter the third item's price:")
	priceStr, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	thirdItemPrice, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
	if err != nil {
		panic(err)
	}

	firstItemPriceFormatted := fmt.Sprintf("%.2f$", firstItemPrice)
	secondItemPriceFormatted := fmt.Sprintf("%.2f$", secondItemPrice)
	thirdItemPriceFormatted := fmt.Sprintf("%.2f$", thirdItemPrice)

	fmt.Println("Here's your restaurant's menu:")
	fmt.Printf("| %-14s | %-14s |\n", firstItemName, firstItemPriceFormatted)
	fmt.Printf("| %-14s | %-14s |\n", secondItemName, secondItemPriceFormatted)
	fmt.Printf("| %-14s | %-14s |\n", thirdItemName, thirdItemPriceFormatted)
}
