package views
const(
    AlertLvlError   = "danger"
    AlertLvlWarning = "warning"
    AlertLvlInfo    = "info"
    AlertLvlSuccess = "success"
    AlertMsgGeneric = "Something went wrong. Please try again, and contact us if the problem persists."
  
)
// Alert is used to render bootstrap alert message 
type Alert struct{
    Level string
    Message string
}
// Data is top level structure that views expect data to come in.
type Data struct{
    Alert *Alert
    Yield interface{}
} 