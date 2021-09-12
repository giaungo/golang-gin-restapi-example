package main

import (
	"net/http"
	"strconv"
	"strings"
	"sort"

	"github.com/gin-gonic/gin"
)

var postIDMap = make(map[string][]comment)
var idMap = make(map[string]comment)
var nameMap = make(map[string][]comment)
var emailMap = make(map[string][]comment)
var bodyMap = make(map[string][]comment)

func indexComments() {
	commentsRes, _ := client.R().
		SetResult(&[]comment{}).
		Get("https://jsonplaceholder.typicode.com/comments")
	comments := commentsRes.Result().(*[]comment)

	for _,c := range *comments {
		idMap[strconv.Itoa(c.ID)] = c
		
		postIDString := strconv.Itoa(c.PostID)
		postIDMap[postIDString] = append(postIDMap[postIDString], c)
		
		emailMap[c.Email] = append(emailMap[c.Email], c)

		nameMap[c.Name] = append(nameMap[c.Name], c)
		for _,name := range strings.Fields(c.Name) {
			nameMap[name] = append(nameMap[name], c)
		}

		bodyMap[c.Body] = append(bodyMap[c.Body], c)
		for _,body := range strings.Fields(c.Body) {
			bodyMap[body] = append(bodyMap[body], c)
		}
	}
}

func getTopPosts(c *gin.Context) {
	postsRes, _ := client.R().
		SetResult(&[]post{}).
		Get("https://jsonplaceholder.typicode.com/posts")
	posts := postsRes.Result().(*[]post)

	var postsDto []postDto
	for _,p := range *posts {
		postsDto = append(postsDto, postDto{
			PostID: p.ID,
			PostTitle: p.Title,
			PostBody: p.Body,
			NumberOfComments: len(postIDMap[strconv.Itoa(p.ID)]),
		})
	}

	sort.Slice(postsDto, func(i, j int) bool {
		return postsDto[i].NumberOfComments > postsDto[j].NumberOfComments
	})

	c.IndentedJSON(http.StatusOK, postsDto)
}

func search(c *gin.Context) {
	query := c.Query("query")

	result := []comment{}

	if cmt,ok := idMap[query]; ok {
		result = append(result, cmt)
	}
	
	if cmt, ok := postIDMap[query]; ok {
		result = append(result, cmt...)
	}

	if cmt, ok := emailMap[query]; ok {
		result = append(result, cmt...)
	}

	if cmt, ok := nameMap[query]; ok {
		result = append(result, cmt...)
	}

	if cmt, ok := bodyMap[query]; ok {
		result = append(result, cmt...)
	}

	if len(c.Query("pageNum")) > 0 &&
		len(c.Query("pageSize")) > 0 {
		pageNum,_ := strconv.Atoi(c.Query("pageNum"))
		pageSize,_ := strconv.Atoi(c.Query("pageSize"))
		start, end := paginate(pageNum, pageSize, len(result))
		c.IndentedJSON(http.StatusOK, result[start:end])
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func paginate(pageNum int, pageSize int, sliceLength int) (int, int) {
    start := pageNum * pageSize

    if start > sliceLength {
        start = sliceLength
    }

    end := start + pageSize
    if end > sliceLength {
        end = sliceLength
    }

    return start, end
}