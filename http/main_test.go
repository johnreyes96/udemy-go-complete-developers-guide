package main

import "testing"

func TestGivenUrlWhenExecuteUrlThenMustReturnOK(t *testing.T) {
	resp := executeUrl("https://www.google.com")

	if resp == nil {
		t.Errorf("Expected response, but got %v", resp)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected 200, but got %v", resp.StatusCode)
	}
	if resp.Status != "200 OK" {
		t.Errorf("Expected 200 OK, but got %v", resp.Status)
	}
	if resp.Body == nil {
		t.Errorf("Expected response body, but got %v", resp.Body)
	}
}

func TestGivenLogWriteWhenWrite6BytesThenMustReturn6Items(t *testing.T) {
	bs := []byte{100, 50, 40, 30, 55, 16}
	lw := logWriter{}

	write, err := lw.Write(bs)

	if err != nil {
		t.Errorf("Expected 6, but got %v", err)
	}
	if write != 6 {
		t.Errorf("Expected 6, but got %v", err)
	}
}
