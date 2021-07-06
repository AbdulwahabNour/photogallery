package models

import "strings"
const ( 
    // ErrNotFound is returned when a resource not found
    ErrNotFound modelError = "Models: Resource not found"
    // ErrInvalid is returned when an invalid id  is provided
    ErrInvalidID modelError = "Models: ID  provided was invalid"
    // ErrInvalidPassword is returned  if hash don't matched 
    ErrInvalidPassword modelError = "Models: Incorrect password"
    ErrEmailRequired modelError = "Models: Email address is required"
    ErrEmailInvalid modelError = "Models: Email address is not valid "
    
    // ErrEmailTaken is returned when an update or create is attempted
    // with an email address that is already in use.
    ErrEmailTaken modelError = "Models: Email address is already taken"

    // ErrPasswordTooShort is returned  when an update or create is attempted
    // with a user password that is less than 8 characters
    ErrPasswordTooShort modelError = "Models: Password is too short must contain at least 8 characters"

    // ErrPasswordTooLong is returned when an update or create is attempted
    // with a user password that is more than 100 characters
    ErrPasswordTooLong modelError = "Models: Password is too long must contain at least 100 characters"
    // ErrPasswordRequired is returned when password is empty
    ErrPasswordRequired modelError = "Models:  Password is  required"
    ErrRememberTooShort modelError = "Models: Remember token is too short"
    // ErrRememberHashRequired is returned when Rememberhash is empty
    ErrRememberHashRequired modelError = "Rememberhash is required"
     // ErrNameRequired is returned when user name is empty
    ErrNameRequired modelError = "Models: Name is required "

    // ErrNameTooShort is returned  when an update or create is attempted
    // with a user name that is less than 8 characters
    ErrNameTooShort modelError = "Models: Name is too short must contain at least 8 characters"

     // ErrNameTooLong is returned when an update or create is attempted
    // with a user name that is more than 100 characters
    ErrNameTooLong modelError = "Models: Name is too long must contain at least 100 characters"
    
    
)
 
//                 -> string
// modelError ->   -> PublicError     
//                 -> error
type modelError  string

func (e modelError)Error() string{
    return string(e)
}
func (e modelError) Public() string{
    return strings.Replace(string(e), "Models: ", "", 1) 
}
