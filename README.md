# shortUrl

## Description
	
	short url Encode/Decode

## Installation

	$ go get github.com/wayne666/shortUrl

## Usage

	package main

	import (
		"fmt"			
		"github.com/wayne666/shortUrl"
	)

	func main() {
		fmt.Println(shortUrl.Encode("hello"))
		fmt.Println(shortUrl.Decode(shortUrl.Encode("hello")))
	}

## COPYRIGHT

	Copyright (c) 2016 WayneZhou. This library is free software; you can redistribute it and/or modify it.
