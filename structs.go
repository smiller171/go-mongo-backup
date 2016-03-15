package main

type dumpTarget struct {
	Bucket string `json:"bucket"`
	Path   string `json:"path"`
}

type result struct {
	Result string `json:"result"`
}
