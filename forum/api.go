package forum

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"log"
	"net/http"
	"strconv"
)

func getForumHandler(r render.Render, dbMap *gorp.DbMap) {
	log.Println("getForumHandler")
}

func getForumsHandler(r render.Render, dbMap *gorp.DbMap) {

	var forums []Forum

	_, err := dbMap.Select(&forums, "SELECT * FROM forum")

	if err != nil {
		r.Error(http.StatusNotFound)
	}

	r.JSON(http.StatusOK, &forums)
}

func getTopicsHandler(r render.Render, dbMap *gorp.DbMap, params martini.Params) string {
	forumId, _ := strconv.Atoi(params["forumId"])
	var topics []Topic

	_, err := dbMap.Select(&topics, "SELECT * FROM forum_topic WHERE forum_id = $1", forumId)

	if err != nil {
		r.Error(http.StatusNotFound)
		return ""
	}

	r.JSON(http.StatusOK, &topics)
	return ""
}

func getTopicHandler(r render.Render, dbMap *gorp.DbMap, params martini.Params) string {
	topicId, _ := strconv.Atoi(params["id"])
	var topic Topic

	err := dbMap.SelectOne(&topic, "SELECT * FROM forum_topic WHERE topic_id = $1", topicId)

	if err != nil {
		r.Error(http.StatusNotFound)
		return ""
	}

	r.JSON(http.StatusOK, &topic)
	return ""
}

func getPostsHandler(r render.Render, dbMap *gorp.DbMap, params martini.Params) string {
	topicId, _ := strconv.Atoi(params["topicId"])
	var posts []Post

	_, err := dbMap.Select(&posts, "SELECT * FROM forum_post WHERE topic_id = $1", topicId)

	if err != nil {
		panic(err)
		r.Error(http.StatusNotFound)
		return ""
	}

	r.JSON(http.StatusOK, &posts)
	return ""
}
