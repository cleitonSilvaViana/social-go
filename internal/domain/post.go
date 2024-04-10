package domain

import (
	"net/http"
	"strings"
	"time"

	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

// Criar um esboço de enum...
// Basicamente, haverá um número limitado de reações, tais como:
// Amei
// Odiei
// Engraçado
// Sem graça
// Indiferente

type React struct {
	ProfileUID []byte `json:"profileUid"`
	PostID     int    `json:"postId"`
	react      string `json:"react"`
}

func (r *React) ValidateReact() *fail.ResponseError {
	switch r.react {
		case "loved":
		case "hated":
		case "funny":
		case "boring":
				// definir regra aqui...
		default:
			return &fail.ResponseError{
				StatusCode: http.StatusBadRequest,
				Message: "unknow react",
			}
	}
	return nil
}


type Comment struct {
	ProfileUID []byte    `json:"profileUid"`
	PostID     int       `json:"postId"`
	CreatedAt  time.Time `json:"createdAt"`
	Comment    string    `json:"comment"`
}

type Post struct {
	ID         int       `json:"id"`
	CreatorUID []byte    `json:"creatorUID, required"`
	Title      string    `json:"title, required"`
	Content    string    `json:"content, required"`
	CreatdAt   time.Time `json:"createdAt"`
	React      []React   `json:"react"`
	Comment    []Comment `json:"comment"`
}

func (p *Post) ValidateIfCreatorExistis() (bool, *fail.ResponseError) {
	return false, nil
}

func (p *Post) Trim() {
	p.Title = strings.Trim(p.Title, " ")
	p.Content = strings.Trim(p.Content, " ")
}

func CreatePost(bodyRequest []byte) (int, *fail.ResponseError) {
	return 0, nil
}

func UpdatePost() {}

func DeletePost() {}

func CommentPost() {}

func ReactPost() {}

func SharePost() {}
