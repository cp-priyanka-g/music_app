package artist

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Artist struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Image_url string `json:"image_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ArtistRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *ArtistRepository {
	return &ArtistRepository{Db: db}
}

type error interface {
	Error() string
}

func (repository *ArtistRepository) Create(c *gin.Context) {
	input := Artist{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Artist(name,image_url) VALUES (?,?)`, input.Name, input.Image_url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Artist Added Successfully"})
}

// Select ARTIST
func (repository *ArtistRepository) Read(c *gin.Context) {

	input, err := repository.GetArtist()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, input)

}

func (repository *ArtistRepository) GetArtist() (input []Artist, err error) {

	err = repository.Db.Select(&input, `SELECT name,image_url,created_at,updated_at from Artist`)
	if err != nil {
		fmt.Println("error on display")
		return

	}

	return
}

//UPDATE artist
func (repository *ArtistRepository) Update(c *gin.Context) {
	input := Artist{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`UPDATE Artist SET name=? ,image_url=? WHERE id=?`, input.Name, input.Image_url, input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot Update ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Artist Updated Successfully"})
}

// DELETE ARTIST
func (repository *ArtistRepository) Delete(c *gin.Context) {
	input := Artist{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`DELETE From Artist  WHERE id=?`, input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot DELETE ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Artist Removed Successfully"})
}
