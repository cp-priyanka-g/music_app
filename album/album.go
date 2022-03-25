package album

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Album struct {
	Id          int    `json:"album_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	IsPublished int    `json:"is_published "`
	CreatedAt   string `json:"created_at "`
	UpdatedAt   string `json:"updated_at "`
	ArtistId    int    `json:"artist_id"`
}

type AlbumTrack struct {
	Id      int   `json:"id"`
	AlbumId int   `json:"album_id"`
	TrackId []int `json:"track_id"`
}

type AlbumRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *AlbumRepository {
	return &AlbumRepository{Db: db}
}

func (repository *AlbumRepository) Create(c *gin.Context) {
	input := Album{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`INSERT INTO Album(name,description,image_url,is_published,artist_id) VALUES (?,?,?,?,?)`, input.Name, input.Description, input.ImageUrl, input.IsPublished, input.ArtistId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Album Created Successfully"})
}

// Select Album
func (repository *AlbumRepository) Read(c *gin.Context) {

	input, err := repository.GetAlbum()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, input)

}

func (repository *AlbumRepository) GetAlbum() (input []Album, err error) {

	err = repository.Db.Select(&input, `SELECT name,description,image_url,is_published from Album`)
	if err != nil {
		panic(err)
	}
	return
}

//UPDATE Album
func (repository *AlbumRepository) Update(c *gin.Context) {
	input := Album{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`UPDATE Album SET name=?,description=?,image_url=? ,is_published=? WHERE id=?`, input.Name, input.Description, input.ImageUrl, input.IsPublished, input.Id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Album Updated Successfully"})
}

// DELETE Album
func (repository *AlbumRepository) Delete(c *gin.Context) {
	input := Album{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`DELETE From Album   WHERE id=?`, input.Id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Album  Removed Successfully"})
}

// Add/remove Track from Album

func (repository *AlbumRepository) Add(c *gin.Context) {
	input := AlbumTrack{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	query := `INSERT INTO AlbumTrack(album_id,track_id) VALUES`
	var inserts []string
	var params []int

	for _, v := range input.TrackId {
		inserts = append(inserts, "(?, ?)")
		params = append(params, input.AlbumId, v)
	}

	fmt.Println("Params:", params)
	s, _ := json.Marshal(params)
	queryparam := string(s)
	fmt.Println("Params after json marshal", queryparam)

	queryVals := strings.Join(inserts, ",")
	query = query + queryVals + queryparam
	fmt.Println("Query is:", query, queryparam)

	data := repository.Db.MustExec(query)

	fmt.Println(data)

	return
}

func (repository *AlbumRepository) Remove(c *gin.Context) {

	input := AlbumTrack{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	_, err = repository.Db.Exec(`DELETE From AlbumTrack WHERE album_id=? and track_id=? `, input.AlbumId, input.TrackId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Track  Removed from album  Successfully"})
}
