package favourite

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Favourite struct {
	FavTrackId    int `json:"fav_track_id"`
	TrackId       int `json:"track_id"`
	UserID        int `json:"user_id"`
	FavTrackIndex int `json:"fav_track_index"`
}

type FavRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *FavRepository {
	return &FavRepository{Db: db}
}

func (repository *FavRepository) Create(c *gin.Context) {
	input := Favourite{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Favourite_tracks (track_id,user_id,fav_track_index) VALUES (?,?,?)`, input.TrackId, input.UserID, input.FavTrackIndex)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot insert ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Favourite Track created Successfully"})
}

// Select Track
func (repository *FavRepository) Read(c *gin.Context) {

	input, err := repository.GetTrack()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, input)

}

func (repository *FavRepository) GetTrack() (input []Favourite, err error) {

	err = repository.Db.Select(&input, `SELECT track_id,user_id,fav_track_index from Favourite_tracks `)
	if err != nil {
		fmt.Println("error on display")
		return

	}

	return
}

// Get Track by ID

func (repository *FavRepository) FavTrackById(c *gin.Context) {

	track := Favourite{}

	id := c.Param("id")

	err := repository.Db.Get(&track, `SELECT track_id,user_id,fav_track_index  from Favourite_tracks WHERE fav_track_id=?`, id)

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

// DELETE Track
func (repository *FavRepository) Delete(c *gin.Context) {
	input := Favourite{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message cannot bind the STRUCT ": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`DELETE From Favourite_tracks  WHERE  fav_track_id=?`, input.FavTrackId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message cannot DELETE ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Track Removed from Favourite Successfully"})
}
