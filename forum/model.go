package forum

import (
	"github.com/modcloth/sqlutil"
	"time"
)

type Forum struct {
	Id sqlutil.NullInt64 `db:"forum_id" json:"id"`
	Name sqlutil.NullString `db:"forum_name" json:"name"`
	Description sqlutil.NullString `db:"forum_desc" json:"description"`
	Order sqlutil.NullInt64 `db:"forum_order" json:"order"`
	Type sqlutil.NullInt64 `db:"forum_type" json:"type"`
	Topics sqlutil.NullInt64 `db:"forum_topics" json:"topics"`
	ShowTopics sqlutil.NullInt64 `db:"show_topics" json:"showTopics"`
}

type Topic struct {
	Id sqlutil.NullInt64 `db:"topic_id" json:"id"`
	ForumId sqlutil.NullInt64 `db:"forum_id" json:"forumId"`
	Name sqlutil.NullString `db:"topic_name" json:"name"`
	AutorId sqlutil.NullInt64 `db:"first_poster" json:"autorId"`
	AutorName sqlutil.NullString `db:"first_poster_name" json:"autorName"`
	LastPostAutorId sqlutil.NullInt64 `db:"last_poster" json:"lastPostAutorId"`
	LastPostAutorName sqlutil.NullString `db:"last_poster_name" json:"lastPostAutorName"`
	LastPostId sqlutil.NullInt64 `db:"last_post_id" json:"lastPostId"`
	LastPostDate *time.Time `db:"last_post_date" json:"lastPostDate"`
	NbOfPosts sqlutil.NullInt64 `db:"topic_posts" json:"nbOfPosts"`
	NbOfViews sqlutil.NullInt64 `db:"topic_views" json:"obOfViews"`
	IsClosed *int16 `db:"topic_closed" json:"isClosed"`
	IsPinned *int16 `db:"topic_pined" json:"isPinned"`
	IsDeleted *int16 `db:"topic_deleted" json:"isDeleted"`
	VisibleFrom *time.Time `db:"topic_visible_from" json:"visibleFrom"`
	VisibleTo *time.Time `db:"topic_visible_to" json:"visibleTo"`
	ChangeAt *time.Time `db:"change_date" json:"changeAt"`
	ChangerId sqlutil.NullInt64 `db:"change_user_id" json:"changerId"`
	ChangerIP sqlutil.NullString `db:"change_ip" json:"changerIP"`
}