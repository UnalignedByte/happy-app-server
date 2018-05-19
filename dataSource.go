package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "time"
import "math"
import "io/ioutil"
import "strings"

const dbUsername = "unalignedbyte"
const dbUrl = "mysql.unalignedbyte.com"
const dbName = "happyapp_unalignedbyte_com"
const dbPasswordFileName = "../private/happyapp/db_password.txt"

const tableHappiness = "happiness"
const happinessPercentage = "percentage"
const happinessDate = "date"

func connect() *sql.DB {
    connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUsername, getPassword(), dbUrl, dbName)
    db, err := sql.Open("mysql", connectString)
    checkError(err)
    return db
}

func getPassword() string {
    password, err := ioutil.ReadFile(dbPasswordFileName)
    checkError(err)
    passwordString := strings.Replace(string(password), "\n", "", -1)
    return passwordString
}

func AddHappiness(percentage int, date time.Time) {
    db := connect()
    defer db.Close()

    if percentage > 100 {
        percentage = 100
    } else if percentage < 0 {
        percentage = 0
    }

    query := fmt.Sprintf("insert into %s (%s, %s) values(\"%d\", \"%v\")", tableHappiness,
        happinessPercentage, happinessDate,
        percentage, date.UTC())

    _, err := db.Exec(query)
    checkError(err)
}

func GetOverallHappiness() int {
    db := connect()
    defer db.Close()

    var overallHappiness float64
    query := fmt.Sprintf("select avg(%s) from %s", happinessPercentage, tableHappiness)
    err := db.QueryRow(query).Scan(&overallHappiness)
    checkError(err)

    return int(math.Round(overallHappiness))
}
