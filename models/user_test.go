package models

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm/logger"
)

func testingUserService() (*UserService, error){

    const (
        host = "localhost"
        port = 5433
        user = "postgres"
        password = "admin"
        dbname = "lenlocked_test"
    )

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    us, err := NewUserService(psqlInfo)
    us.db.Config.Logger = logger.Default.LogMode(logger.Silent )
    if err != nil {
        return nil, err
    }
    us.DestructiveReset()
    return us, nil   
}
func TestCreateUser(t *testing.T){
    us, err := testingUserService()
    if err != nil{
        t.Fatal(err)
    }
    user := User{
        Name: "Ahmed",
        Email: "ahmedTest@yahoo.com",
    }
    err = us.Create(&user)
    if err != nil{
        t.Fatal(err)
    }
 
    CreateTime := time.Duration(5*time.Second)
    
    if time.Since(user.CreatedAt) > CreateTime{
        t.Errorf("Expected CreatedAt to be recent. Received %s\n", user.CreatedAt)
    }
    if time.Since(user.UpdatedAt) > CreateTime{
 
        t.Errorf("Expected CreatedAt to be recent. Received %s\n", user.UpdatedAt)
    }
    
    
}