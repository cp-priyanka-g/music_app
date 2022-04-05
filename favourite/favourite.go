package favourite

import (
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
type FavouriteTrack struct {
	FavTrackId int    `json:"fav_track_id"`
	UserName   string `json:"name"`
	TrackUrl   string `json:"track_url"`
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
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Favourite_tracks (track_id,user_id,fav_track_index) VALUES (?,?,?)`, input.TrackId, input.UserID, input.FavTrackIndex)

	if err != nil {
		panic(err)
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

func (repository *FavRepository) GetTrack() (input []FavouriteTrack, err error) {

	err = repository.Db.Select(&input, `SELECT f.fav_track_id,u.name,t.track_url  from Favourite_tracks as f join Users as u on f.user_id=u.user_id 
	join Track as t on f.track_id= t.track_id`)
	if err != nil {

		return

	}

	return
}

// Get Track by ID

func (repository *FavRepository) FavTrackId(c *gin.Context) {

	track := FavouriteTrack{}

	id := c.Param("id")

	err := repository.Db.Get(&track, `SELECT f.fav_track_id,u.name,t.track_url  from Favourite_tracks as f join Users as u on f.user_id=u.user_id 
	join Track as t on f.track_id= t.track_id  WHERE f.fav_track_id=?`, id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, track)
}

// DELETE Track
func (repository *FavRepository) Delete(c *gin.Context) {
	input := Favourite{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`DELETE From Favourite_tracks  WHERE  fav_track_id=?`, input.FavTrackId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Track Removed from Favourite Successfully"})
}
