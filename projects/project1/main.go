package main

func main() {
	giphyProvider := newGiphy("<API KEY>")
	s := newService(giphyProvider)
	initRouter(s)
}
