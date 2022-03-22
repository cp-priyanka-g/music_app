package playlist

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type playlist struct {
	Id          int    `json:"playlist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	IsPublished int    `json:"is_published "`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type playlistRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *playlistRepository {
	return &playlistRepository{Db: db}
}

func (repository *playlistRepository) Create(c *gin.Context) {
	input := playlist{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Playlist(name,description,image_url,is_published) VALUES (?,?,?,?)`, input.Name, input.Description, input.ImageUrl, input.IsPublished)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "playlist Created Successfully"})
}

// Select playlist
func (repository *playlistRepository) Read(c *gin.Context) {

	input, err := repository.Getplaylist()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, input)

}

func (repository *playlistRepository) Getplaylist() (input []playlist, err error) {

	err = repository.Db.Select(&input, `SELECT name,description,image_url from Playlist`)
	if err != nil {
		fmt.Println("error on display")
		return

	}

	return
}

//UPDATE playlist
func (repository *playlistRepository) Update(c *gin.Context) {
	input := playlist{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`UPDATE Playlist SET name=?,description=?,image_url=? ,is_published=? WHERE playlist_id=?`, input.Name, input.Description, input.ImageUrl, input.IsPublished, input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot Update ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "playlist Updated Successfully"})
}

// DELETE playlist
func (repository *playlistRepository) Delete(c *gin.Context) {
	input := playlist{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`DELETE From Playlist  WHERE playlist_id=?`, input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot DELETE ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "playlist Removed Successfully"})
}