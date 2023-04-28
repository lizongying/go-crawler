package main

import "github.com/lizongying/go-crawler/pkg"

type ExtraNoLimit struct {
	Count int
}

type DataNoLimit struct {
	Count int
}

type ExtraTest struct {
	*pkg.Image
}

type DataTest struct {
	*pkg.Image
}
