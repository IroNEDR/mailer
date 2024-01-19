run:
	go run main.go -from=$(FROMMAIL) -password=$(PASSWORD) -to=$(TOMAIL) -smtpHost=$(SMTPHOST) -smtpPort=$(SMTPPORT)