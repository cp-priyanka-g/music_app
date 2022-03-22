package track

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Track struct {
	Id          int    `json:"track_id"`
	Name        string `json:"name"`
	TrackIndex  int    `json:"track_index"`
	TrackUrl    string `json:"track_url"`
	ImageUrl    string `json:"image_url"`
	IsPublished int    `json:"is_published"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	ArtistId    int    `json:"artist_id"`
}

type TrackRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *TrackRepository {
	return &TrackRepository{Db: db}
}

func (repository *TrackRepository) Create(c *gin.Context) {
	input := Track{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Track (name,track_index,track_url,image_url,is_published,artist_id) VALUES (?,?,?,?,?,?)`, input.Name, input.TrackIndex, input.TrackUrl, input.ImageUrl, input.IsPublished, input.ArtistId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Track Created Successfully"})
}

// Select Track
func (repository *TrackRepository) Read(c *gin.Context) {

	input, err := repository.GetTrack()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, input)

}

func (repository *TrackRepository) GetTrack() (input []Track, err error) {

	err = repository.Db.Select(&input, `SELECT name,track_index,track_url,image_url,is_published,created_at,updated_at,artist_id from Track`)
	if err != nil {
		fmt.Println("error on display")
		return

	}

	return
}

// Get Track by ID

func (repository *TrackRepository) TrackById(c *gin.Context) {

	track := Track{}

	id := c.Param("id")

	err := repository.Db.Get(&track, `SELECT name,track_index,track_url,image_url,is_published,created_at,updated_at,artist_id from Track WHERE track_id=?`, id)

	if err != nil {
		fmt.Println("error query is empty")
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, track)
}

//UPDATE Track
func (repository *TrackRepository) Update(c *gin.Context) {
	input := Track{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`UPDATE Track SET name=?,track_index=?,track_url=?,image_url=? ,is_published=? WHERE track_id=?`, input.Name, input.TrackIndex, input.TrackUrl, input.ImageUrl, input.IsPublished, input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot Update ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Track Updated Successfully"})
}

// DELETE Track
func (repository *TrackRepository) Delete(c *gin.Context) {
	input := Track{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`DELETE From Track  WHERE track_id=?`, input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot DELETE ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Track Removed Successfully"})
}
