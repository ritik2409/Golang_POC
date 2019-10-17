package aws

import "github.com/aws/aws-sdk-go/aws/session"

func create() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	aws_access_key_id = ""
	aws_secret_access_key = ""
	
}
