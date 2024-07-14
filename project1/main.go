package main

import "github.com/brandonyeoxg/project52/project1/config"

func main() {
	config := config.New("config")
	giphyProvider := newGiphy(config.GiphyAPIKEY)
	s := newService(giphyProvider)
	initRouter(s)
}
