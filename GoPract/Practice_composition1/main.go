package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (p blogPost) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName()) //no need ti use p.author.fullName cuz go allows us to use directly as if its part of outer struct=>blog
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}
func main() {
	author1 := author{
		"Agatha",
		"Christie",
		"The Secret of chimneys",
	}
	blogPost1 := blogPost{
		"Mystries week",
		"Best recommendation by us",
		author1,
	}
	blogPost2 := blogPost{
		"Life as to go on",
		"Said so",
		author1,
	}
	blogPost3 := blogPost{
		"Life itself is a mystery",
		"Authors says so",
		author1,
	}
	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()

}
