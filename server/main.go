package main

import "server/delivery"

func main(){
	delivery.NewServer().Run()
}