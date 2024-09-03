package args

import (
	"errors"
	"flag"
	"os"
	"strings"
)

func Parse() (*Args, error) {
	secret := os.Getenv("API_SECRET")

	urlFlag := flag.String("url", "REQUIRED", "Google Sheet URL")
	sinkFlag := flag.String("sink", "Warehouse", "Fivetran sink name")
	keyFlag := flag.String("key", "OPTIONAL", "Fivetran API key. Can be set with the 'API_KEY' environment variable")

	flag.Parse()

	key := os.Getenv("API_KEY")
	if strings.TrimSpace(*keyFlag) != "OPTIONAL" {
		key = *keyFlag
	}

	return &Args{
		APISecret: secret,
		URLPart:   *urlFlag,
		APIKey:    key,
		Sink:      *sinkFlag,
	}, nil
}

func Validate(args *Args) (bool, error) {
	if args == nil {
		return false, errors.New("args cannot be nil")
	}
	if strings.TrimSpace(args.APISecret) == "" {
		return false, errors.New("'API_SECRET' environment variable is required.\n")
	}
	if strings.TrimSpace(args.APIKey) == "" {
		return false, errors.New("'API_KEY' environment variable or the -key flag is required.\n")
	}
	if strings.TrimSpace(args.URLPart) == "REQUIRED" {
		return false, errors.New("The -url flag is required.\n")
	}
	return true, nil
}

func PrintUsage() {
	flag.PrintDefaults()
}
