package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lenslocked.com/hash"
	"lenslocked.com/rand"
)
 
 
var ( 
    //ErrNotFound is returned when a resource not found
    ErrNotFound = errors.New("Models: resource not found")
    // ErrInvalid is returned when an invalid id  is provided
    ErrInvalidID = errors.New("Models: ID  provided was invalid")
    // ErrInvalidPassword is if hash don't matched 
    ErrInvalidPassword = errors.New("Models: Incorrect password")
)

const (
    userPassPepper = "photogallery-App"
    hmacSecretkey  = "This-Must-Key-For-HMAC" 
)
func NewUserService(psqlInfo string)(*UserService, error){

     db, err:= gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
     if err != nil{
         return nil, err
     }
     hmac := hash.NewHMAC(hmacSecretkey)
    return &UserService{db:db, hmac: hmac}, nil
}

type UserService struct{
    db *gorm.DB
    hmac hash.HMAC 
}
/*
** ByID will look up the user by id
** if user is found we will return nil error
** if user not found will return nil and ErrNotFound 
** if there is another error, we will return  an error with more information
**
*/
func (u *UserService) ByID(id uint)(*User, error){
   
    var user User
    db := u.db.Where("id= ?",id)
    
    err := first(db, &user)
    if err == nil{
        return  &user, err
    }
   return nil, err
}
/*
** ByHash Function will look up the user by RememberHash
** if RememberHash is found will return the user for this RememberHash
** if not found will return nil for user and error
*/
func (u *UserService) ByHash(token string)(* User, error){
    var user User
    
    hashedToken := u.hmac.Hash(token)
    db := u.db.Where("remember_hash = ?", hashedToken)
    err := first(db, &user)
    if err != nil{
        return nil, err
    }
    return  &user, err
  
}
/*
** ByEmail look up auser with given email
** and return that user
** If user found will return the user and nil
** If user not found will return nil for user and ErrNotFound
** If there is another  error, we will return an error with more
** information about  what went wrong 
*/
func (u *UserService)ByEmail(email string) (*User, error){
    var user User
    db := u.db.Where("email = ?", email)
    err := first(db ,&user)
    if err == nil{
        return  &user, err
    }
   return nil, err
    
}
//Create New user 
func (u *UserService)Create(user *User) error{
    hashbyte, err := generatePassword([]byte(user.Password + userPassPepper))
    if err != nil{
        return err
    }
    user.PasswordHash = string(hashbyte)
    user.Password = ""
    if user.Rememer == ""{
        token, err := rand.RememberToken()
        if err != nil{
            return err
        }
        user.Rememer = token
    }
 
     user.RememberHash = u.hmac.Hash(user.Rememer)
    return u.db.Create(user).Error
}

//Get First user 
func (u *UserService)First(user *User) error{
    return u.db.First(user).Error
}

//Get Last user 
func (u *UserService)Last(user *User) error{
    return u.db.Last(user).Error
}

//update User
func (u *UserService)Update(user *User)error{
    if user.Rememer != ""{
        user.RememberHash = u.hmac.Hash(user.Rememer)
   }
    return u.db.Save(&user).Error
}
//Delete User
func(u * UserService)Delete(id uint) error{
    var user User 
    user.ID = id
    return u.db.Delete(&user).Error
}
//Drops the user table and rebuilds it
func (u *UserService)DestructiveReset(){
    u.db.Migrator().DropTable(&User{})
    u.db.AutoMigrate(&User{})
}


//close UserService db connection 
func (u *UserService) Close() error{
    con,_ := u.db.DB()
    return con.Close()
    
}

func (u *UserService)Authenticate(email string, password string)(*User, error){
      foundUser, err := u.ByEmail(email)
      if err != nil{
          return nil, err
      }
      err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password + userPassPepper))
       if err != nil{
            switch err{
                case bcrypt.ErrMismatchedHashAndPassword:
                    return nil, ErrInvalidPassword

                default:
                    return nil, err
            }
       }
    
      return foundUser, nil
}


 
func first(db *gorm.DB, data  interface{}) error{
    err := db.First(data).Error
    if err == gorm.ErrRecordNotFound{
        return  ErrNotFound
    }
    return  err
}
func generatePassword(password []byte)([]byte,error){
    hashbyte, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil{
       return nil, err
    }
    return hashbyte, nil
}

type User struct{
    gorm.Model
    Email    string `gorm:"uniqueIndex;not null"`
    Name  string  
    Password string `gorm:"-"`
    PasswordHash string `gorm:"not null"`
    Rememer string `gorm:"-"`
    RememberHash string `gorm:"not null;uniqueIndex"`
}