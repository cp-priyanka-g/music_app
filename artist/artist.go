package artist

import (
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
}

type ArtistRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *ArtistRepository {
	return &ArtistRepository{Db: db}
}

func (repository *ArtistRepository) CreateArtist(c *gin.Context) {
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

//UPDATE artist
func (repository *ArtistRepository) UpdateArtist(c *gin.Context) {
	input := Artist{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`UPDATE Artist SET name=? ,image_url=? WHERE id=?`, input.Name, input.Image_url, input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Artist Updated Successfully"})
}
