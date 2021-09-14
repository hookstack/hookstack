package config

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

var cfgSingleton atomic.Value

type DatabaseConfiguration struct {
	Dsn string `json:"dsn"`
}

type Configuration struct {
	Auth              AuthConfiguration   `json:"auth"`
	UIAuth            UIAuthConfiguration `json:"ui"`
	UIAuthorizedUsers map[string]string
	Database          DatabaseConfiguration `json:"database"`
	Queue             QueueConfiguration    `json:"queue"`
	Server            struct {
		HTTP struct {
			Port int `json:"port"`
		} `json:"http"`
	}
	Strategy  StrategyConfiguration  `json:"strategy"`
	Signature SignatureConfiguration `json:"signature"`
}

type AuthProvider string
type QueueProvider string
type StrategyProvider string
type SignatureHeaderProvider string

const (
	NoAuthProvider          AuthProvider            = "none"
	BasicAuthProvider       AuthProvider            = "basic"
	RedisQueueProvider      QueueProvider           = "redis"
	DefaultStrategyProvider StrategyProvider        = "default"
	DefaultSignatureHeader  SignatureHeaderProvider = "X-Convoy-Signature"
)

type QueueConfiguration struct {
	Type  QueueProvider `json:"type"`
	Redis struct {
		DSN string `json:"dsn"`
	} `json:"redis"`
}

type AuthConfiguration struct {
	Type  AuthProvider `json:"type"`
	Basic Basic        `json:"basic"`
}

type UIAuthConfiguration struct {
	Type                  AuthProvider  `json:"type"`
	Basic                 []Basic       `json:"basic"`
	JwtKey                string        `json:"jwtKey"`
	JwtTokenExpirySeconds time.Duration `json:"jwtTokenExpirySeconds"`
}

type Basic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type StrategyConfiguration struct {
	Type    StrategyProvider `json:"type"`
	Default struct {
		IntervalSeconds uint64 `json:"intervalSeconds"`
		RetryLimit      uint64 `json:"retryLimit"`
	} `json:"default"`
}

type SignatureConfiguration struct {
	Header SignatureHeaderProvider `json:"header"`
}

func LoadConfig(p string) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}

	defer f.Close()

	c := new(Configuration)

	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return err
	}

	if mongoDsn := os.Getenv("CONVOY_MONGO_DSN"); mongoDsn != "" {
		c.Database = DatabaseConfiguration{Dsn: mongoDsn}
	}

	if queueDsn := os.Getenv("CONVOY_REDIS_DSN"); queueDsn != "" {
		c.Queue = QueueConfiguration{
			Type: "redis",
			Redis: struct {
				DSN string `json:"dsn"`
			}{
				DSN: queueDsn,
			},
		}
	}

	if signatureHeader := os.Getenv("CONVOY_SIGNATURE_HEADER"); signatureHeader != "" {
		c.Signature.Header = SignatureHeaderProvider(signatureHeader)
	}

	if apiUsername := os.Getenv("CONVOY_API_USERNAME"); apiUsername != "" {
		var apiPassword string
		if apiPassword = os.Getenv("CONVOY_API_PASSWORD"); apiPassword != "" {
			return errors.New("Failed to retrieve apiPassword")
		}

		c.Auth = AuthConfiguration{
			Type:  "basic",
			Basic: Basic{apiUsername, apiPassword},
		}
	}

	if uiUsername := os.Getenv("CONVOY_UI_USERNAME"); uiUsername != "" {
		var uiPassword, jwtKey, jwtExpiryString string
		var jwtExpiry time.Duration
		if uiPassword = os.Getenv("CONVOY_UI_PASSWORD"); uiPassword != "" {
			return errors.New("Failed to retrieve uiPassword")
		}

		if jwtKey = os.Getenv("CONVOY_JWT_KEY"); jwtKey != "" {
			return errors.New("Failed to retrieve jwtKey")
		}

		if jwtExpiryString = os.Getenv("CONVOY_JWT_EXPIRY"); jwtExpiryString != "" {
			return errors.New("Failed to retrieve jwtExpiry")
		}

		jwtExpiryInt, err := strconv.Atoi(jwtExpiryString)
		if err != nil {
			return errors.New("Failed to parse jwtExpiry")
		}

		jwtExpiry = time.Duration(jwtExpiryInt) * time.Second

		basicCredentials := Basic{uiUsername, uiPassword}
		c.UIAuth = UIAuthConfiguration{
			Type: "basic",
			Basic: []Basic{
				basicCredentials,
			},
			JwtKey:                jwtKey,
			JwtTokenExpirySeconds: jwtExpiry,
		}
	}

	c.UIAuthorizedUsers = parseAuthorizedUsers(c.UIAuth)

	cfgSingleton.Store(c)
	return nil
}

func parseAuthorizedUsers(auth UIAuthConfiguration) map[string]string {
	users := auth.Basic
	usersMap := make(map[string]string)
	for i := 0; i < len(users); i++ {
		usersMap[users[i].Username] = users[i].Password
	}
	return usersMap
}

// Get fetches the application configuration. LoadFromFile must have been called
// previously for this to work.
// Use this when you need to get access to the config object at runtime
func Get() (Configuration, error) {
	c, ok := cfgSingleton.Load().(*Configuration)
	if !ok {
		return Configuration{}, errors.New("call Load before this function")
	}

	return *c, nil
}
