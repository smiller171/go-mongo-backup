package main

type dumpTarget struct {
	Bucket string `json:"bucket"`
	Path   string `json:"path"`
	Region string `json:"region"`
}

type result struct {
	Result string `json:"result"`
}
