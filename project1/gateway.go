package main

type MemeGetter interface {
	MemeIDs(tags ...string) ([]string, error)
	MemeDownload(id string) ([]byte, error)
}
