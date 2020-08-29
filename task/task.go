package task
import (
	"net/http" 
	"github.com/gin-gonic/gin"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

type Customer struct {
    ID     int    `json:"id"` 
	Name  string `json:"name"`
	Email  string `json:"email"`
    Status string `json:"status"`
}

var db *sql.DB
func init() {
    var err error
    db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }
}


func CreateDetailHandler(c *gin.Context) {  //// POST/customers
    t := Customer{}
    if err := c.ShouldBindJSON(&t); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
 
    row := db.QueryRow("INSERT INTO customer (name, email, status) values ($1, $2, $3)  RETURNING id", t.Name,t.Email, t.Status)
 
    err := row.Scan(&t.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
 
    c.JSON(http.StatusCreated, t)
}

func GetCustomerByIdHandler(c *gin.Context) { ///GET BY ID
	id := c.Param("id")
    stmt, err := db.Prepare("SELECT id, name, email, status FROM customer where id=$1")
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    row := stmt.QueryRow(id)
    t := &Customer{}
    err = row.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, t)
}

func GetAllCustomersHandler(c *gin.Context) { //// GET ALL GET/customers
    status := c.Query("status")
    stmt, err := db.Prepare("SELECT id, name, email, status FROM customer")
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    rows, err := stmt.Query()
    if err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    customer := []Customer{}
    for rows.Next() {
        t := Customer{}
        err := rows.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
        if err != nil {
            c.JSON(http.StatusInternalServerError, err)
            return
        }
        customer = append(customer, t)
    }
    tt := []Customer{}
    for _, item := range customer {
        if status != "" {
            if item.Status == status {
                tt = append(tt, item)
            }
        } else { 
            tt = append(tt, item)
        }
    }
	c.JSON(http.StatusOK, tt)

}

func UpdateDetailCustomerHandler(c *gin.Context) { //PUT
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customer where id=$1")
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	row := stmt.QueryRow(id)
	t := &Customer{}
	err = row.Scan(&t.ID, &t.Name, &t.Email, &t.Status)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if err := c.ShouldBindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stmt, err = db.Prepare("UPDATE customer SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if _, err := stmt.Exec(id, t.Name, t.Email, t.Status); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, t)

}

func DelCustomerHandler(c *gin.Context){ //DEL
	id :=c.Param("id")
	stmt,err := db.Prepare("DELETE FROM customer WHERE id =$1")
	if err != nil {
		log.Fatal("can't prepare delete statement",err)
	}

	if _,err := stmt.Exec(id); err != nil {
		log.Fatal("can't execute delete statement",err)
	}
	c.JSON(http.StatusOK,gin.H{"message": "customer deleted"})

}

