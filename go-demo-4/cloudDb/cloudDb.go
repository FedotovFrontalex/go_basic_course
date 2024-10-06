package cloudDb

import ()

type CloudDb struct {
	Url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		Url: url,
	}
}

func (db *CloudDb) Write(content []byte) {
}

func (db *CloudDb) Read() ([]byte, error) {
	return nil, nil
}
