package main

import (
	"reflect"
	"testing"
)

func Test_getAllArticles(t *testing.T) {
	tests := []struct {
		name string
		want []article
	}{
		// TODO: Add test cases.
		{
			"case1",
			[]article{
				{
					Id:      0,
					Title:   "Tom",
					Content: "Tom is a boy",
				},
				{
					Id:      1,
					Title:   "Tina",
					Content: "Tina is a girl",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllArticles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
