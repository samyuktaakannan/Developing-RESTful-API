package main
import 
(
    "database/sql"
    //"encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    //"github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

func main() {
    router := gin.Default()
    router.GET("/movies", getMovies)
    router.POST("/movies", createMovie)
    router.DELETE("/movies/:movieid", deleteMovie)
    router.DELETE("/movies", deleteAllMovies)
    router.Run("localhost:8080")
}

const 
(
    DB_USER     = "postgres"
    DB_PASSWORD = "shreya123"
    DB_NAME     = "proj2"
)


func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    if err != nil {
        log.Fatal(err)
    }
    return db
}

type Movie struct 
{
    //id int `json:"ID"`
    MovieID   int `json:"movieid"`
    MovieName string `json:"moviename"`
}

type JsonResponse struct
{
    Type    string `json:"type"`
    Data    []Movie `json:"data"`
    Message string `json:"message"`
}



func handleMessage(c *gin.Context) {
    message := c.PostForm("message")
    c.JSON(http.StatusOK, gin.H{"message": message})
    fmt.Println(message)
    c.String(http.StatusOK, message)
}

/*func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}*/

func getMovies(c *gin.Context) {
    db, err := sql.Open("postgres", "user=postgres password=shreya123 dbname=proj2 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    rows, err := db.Query("SELECT * FROM moviestest")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var movies []Movie
    for rows.Next() {
        //var ID int
        var movieID int
        var movieName string

        err := rows.Scan(&movieID, &movieName)
        if err != nil {
            log.Fatal(err)
        }
        movies = append(movies,  Movie{MovieID: movieID, MovieName: movieName})
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }

    c.JSON(http.StatusOK, movies)
}

func createMovie(c *gin.Context) {
    var movie Movie
    if err := c.ShouldBindJSON(&movie); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Open a connection to the PostgreSQL database
    db, err := sql.Open("postgres", "user=postgres password=shreya123 dbname=proj2 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    query := "INSERT INTO moviestest (movieID,movieName) VALUES ($1, $2)"
    _,err = db.Exec(query,movie.MovieID,movie.MovieName)
    if err != nil {
        log.Fatal(err)
    } 

    c.Status(http.StatusCreated)
}

func deleteMovie(c *gin.Context) {
  // Extract the movie title from the URL parameter
  movieid := c.Param("movieid")

  // Open a database connection
  db, err := sql.Open("postgres", "user=postgres password=shreya123 dbname=proj2 sslmode=disable")
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  defer db.Close()

  // Delete the movie with the specified title from the "movies" table
  query := "DELETE FROM moviestest WHERE movieID= $1"
  result, err := db.Exec(query, movieid)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Check if any rows were affected by the delete operation
  count, err := result.RowsAffected()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  if count == 0 {
    c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
    return
  }

  c.Status(http.StatusNoContent)
}


func deleteAllMovies(c *gin.Context) {
    // Get database connection from the Gin context
    db, err := sql.Open("postgres", "user=postgres password=shreya123 dbname=proj2 sslmode=disable")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    defer db.Close()
        
    // Execute the DELETE statement
    result, err := db.Exec("DELETE FROM moviestest")
    if err != nil {
        // Handle the error
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // Get the number of rows deleted from the result object
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        // Handle the error
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // Return the number of rows deleted to the client
    c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d rows deleted", rowsAffected)})
}