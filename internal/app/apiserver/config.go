package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	FcmKey      string `toml:"fcm_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":3052",
		LogLevel:    "debug",
		DatabaseURL: "account:PN-!UPVm@tcp(178.159.45.196)/user1100741_account",
		//DatabaseURL: "host=178.159.45.196 dbname=user1100741_account sslmode=disable password=PN-!UPVm user=account",
		//DatabaseURL: "host=localhost port=3053 dbname=mitsodb sslmode=require password=mitsodbpass user=dbuser",
		FcmKey:      "AAAAbgOPVb4:APA91bGTZhVFprUUAGEAOVnzEV_H34OOL_GXHGruuz2supgUUR9pryzDGvj_70OS4iTkJgOU15SdDjL4P7jsXI-8znwlJPEoB2Na_DIhSZI-WqXfaY1BOEnT5xfkNrwxTWid2zC9t3AQ",
	}

}
