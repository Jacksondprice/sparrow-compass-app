package notify

import "os"

var _ = os.Getenv

var TWILIO_ACCOUNT_SID = os.Getenv("TWILIO_ACCOUNT_SID")
var TWILIO_AUTH_TOKEN = os.Getenv("TWILIO_AUTH_TOKEN")
