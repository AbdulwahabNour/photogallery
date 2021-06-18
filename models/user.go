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

type User struct{
    gorm.Model
    Email    string `gorm:"uniqueIndex;not null"`
    Name  string  
    Password string `gorm:"-"`
    PasswordHash string `gorm:"not null"`
    Remember string `gorm:"-"`
    RememberHash string `gorm:"not null;uniqueIndex"`
}

func NewUserService(psqlInfo string)(*UserService, error){

     ug , err:= newUserGorm(psqlInfo)
     if err != nil{
         return nil, err
     }
     hmac := hash.NewHMAC(hmacSecretkey)
    return &UserService{ 
                         UserDB:  &userValidator{UserDB:ug,
                                                 hmac: hmac},
                         }, nil
}
//UserDB is used to interact with the users database
type UserDB interface{
    //Methods for querying for single users
    ByID(id uint)(*User, error)
    ByEmail(email string) (*User, error)
    ByRememberToken(token string)(* User, error)

    //Methods for altering users
    Create(user *User) error
    Update(user *User)error
    Delete(id uint) error

    // Used to close DB connection
    Close() error
    
    // Migration 
    AutoMigrate()error
    DestructiveReset() error
   
}

type UserService struct{
    UserDB
}

type userValFunc func(*User) error

func runuserValFunc(user *User, fns ...userValFunc) error{
    for _, f := range fns{
        err :=  f(user)
        if err != nil{
            return err
        }
    }
    return nil
}
type userValidator struct{
    UserDB
    hmac hash.HMAC 
}
 
//   ByRememberToken Function Take RememberToken and hash it 
//   and return that user.
 
func (u *userValidator) ByRememberToken(RememberToken string)(* User, error){
    hashedToken := u.hmac.Hash(RememberToken)
    return u.UserDB.ByRememberToken(hashedToken)
}
 
// Create Function create user by  
// adding pepper to user password  
// generate password from them 
// create remember hash
// then send User to Create
// on the subsequent UserDB 
 
func (u *userValidator)Create(user *User) error{
    err :=  runuserValFunc(user,  u.bcryptPassword)
    if err != nil{
        return err
    }

    if user.Remember == ""{
        token, err := rand.RememberToken()
        if err != nil{
            return err
        }
        user.Remember = token
    }
 
     user.RememberHash = u.hmac.Hash(user.Remember)

     return u.UserDB.Create(user)
}
// bcryptPassword will  hash auser's password with
// predefined pepper and bycrpt if the password field
// is not the empty string
func (u *userValidator) bcryptPassword(user *User) error{

    if user.Password == ""{
        return nil
    }

    password := []byte(user.Password + userPassPepper)
    hashbyte, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil{
        return err
    } 
    user.PasswordHash = string(hashbyte)
    user.Password = ""
   
    return  nil
}

 
// Update function will hash user remember token if it is provided
// then send User to Update on the subsequent UserDB
func (u *userValidator) Update(user *User) error{
    err :=  runuserValFunc(user,  u.bcryptPassword)
    if err != nil{
        return err
    }
    
    if user.Remember != ""{
        user.RememberHash = u.hmac.Hash(user.Remember)
   }
    return u.UserDB.Update(user)
}

func(u * userValidator)Delete(id uint) error{
    if id <=  0 {
     return ErrInvalidID
    }
    return u.UserDB.Delete(id)
}
 
 
func  newUserGorm(psqlInfo string) (*userGorm, error){
    db, err:= gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
    if err != nil{
        return nil, err
    }
  return &userGorm{db: db}, nil
}



var _ UserDB = &userGorm{}

type userGorm struct{
    db *gorm.DB
}
 
//   ByID will look up the user by id
//   if user is found we will return nil error
//   if user not found will return nil and ErrNotFound 
//   if there is another error, we will return  an error with more information
 
 
func (u *userGorm) ByID(id uint)(*User, error){
   
    var user User
    db := u.db.Where("id= ?",id)
    
    err := first(db, &user)
    if err == nil{
        return  &user, err
    }
   return nil, err
}
 
//   ByRememberToken Function Take Hashtoken  
//   and look up the user by hashedToken  
//   if Hashtoken is found will return the user  
//   if not found will return nil for user and error
 
func (u *userGorm) ByRememberToken(Hashtoken string)(* User, error){
    var user User
    db := u.db.Where("remember_hash = ?", Hashtoken)
    err := first(db, &user)
    if err != nil{
        return nil, err
    }
    return  &user, err
}
 
//   ByEmail look up auser with given email
//   and return that user
//   If user found will return the user and nil
//   If user not found will return nil for user and ErrNotFound
//   If there is another  error, we will return an error with more
//   information about  what went wrong 
 
func (u *userGorm)ByEmail(email string) (*User, error){
    var user User
    db := u.db.Where("email = ?", email)
    err := first(db ,&user)
    if err == nil{
        return  &user, err
    }
   return nil, err
    
}
//Create New user 
func (u *userGorm)Create(user *User) error{
     return u.db.Create(user).Error
}

 

//update User
func (u *userGorm) Update(user *User)error{
    return u.db.Save(&user).Error
}
//Delete User
func(u * userGorm)Delete(id uint) error{
    user := User{Model:gorm.Model{ID: id}}
    return u.db.Delete(&user).Error
}
//Migrate user table
func (u *userGorm)AutoMigrate()error {
    return u.db.AutoMigrate(&User{})
}

//Drops the user table  
func (u *userGorm)DestructiveReset() error{
    return u.db.Migrator().DropTable(&User{})
}


//close UserService db connection 
func (u *userGorm) Close() error{
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

