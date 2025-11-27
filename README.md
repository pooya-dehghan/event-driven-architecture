ill add a user system management to engage from email or telegram , 

i want to every user be able to set a price for gold index (xauusd) and be alerted when the price hits a certain point ,

ill use microservice architecture containing few services like crawler and user service.

i want to use queue for publishing telegram and email messaging and also be able to track the message queues.



### Database Design

# User :
    _ID: uint
    _Name: string
    _PhoneNumber: string


# CurrencyRequest : 
    _ID: uint
    _Price: uint
    _UserID: uint
    _Currency: Enum(Gold, Oil, Silver)
    

# Notifier : 
    _ID: uint
    _Platform: Enum(Email, Telegram)
    _SendDate: Date
    _UserID: uint
    


