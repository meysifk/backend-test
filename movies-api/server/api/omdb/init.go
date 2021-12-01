package omdb

type Omdb struct {
	Credential string
	Host       string
}

func InitOmdb(credential, host string) *Omdb {
	return &Omdb{
		Credential: credential,
		Host:       host,
	}
}
