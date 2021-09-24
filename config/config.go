package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/frain-dev/convoy/config/algo"
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
	SqsQueueProvider        QueueProvider           = "sqs"
	DefaultStrategyProvider StrategyProvider        = "default"
	DefaultSignatureHeader  SignatureHeaderProvider = "X-Convoy-Signature"
)

type QueueConfiguration struct {
	Type  QueueProvider `json:"type"`
	Redis struct {
		DSN string `json:"dsn"`
	} `json:"redis"`
	Sqs struct {
		DSN               string `json:"dsn"`
		Profile           string `json:"profile"`
		DelaySeconds      uint16 `json:"delaySeconds"`
		MaxMessages       uint8  `json:"maxMessages"`
		VisibilityTimeout uint16 `json:"visibilityTimeout"`
	} `json:"sqs"`
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
	Hash   string                  `json:"hash"`
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

	// This enables us deploy to Heroku where the $PORT is provided
	// dynamically.
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		c.Server = struct {
			HTTP struct {
				Port int `json:"port"`
			} `json:"http"`
		}{
			HTTP: struct {
				Port int `json:"port"`
			}{
				Port: port,
			},
		}
	}

	if signatureHeader := os.Getenv("CONVOY_SIGNATURE_HEADER"); signatureHeader != "" {
		c.Signature.Header = SignatureHeaderProvider(signatureHeader)
	}

	if signatureHash := os.Getenv("CONVOY_SIGNATURE_HASH"); signatureHash != "" {
		c.Signature.Hash = signatureHash
	}
	err = ensureSignature(c.Signature)
	if err != nil {
		return err
	}

	if apiUsername := os.Getenv("CONVOY_API_USERNAME"); apiUsername != "" {
		var apiPassword string
		if apiPassword = os.Getenv("CONVOY_API_PASSWORD"); apiPassword == "" {
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
		if uiPassword = os.Getenv("CONVOY_UI_PASSWORD"); uiPassword == "" {
			return errors.New("Failed to retrieve uiPassword")
		}

		if jwtKey = os.Getenv("CONVOY_JWT_KEY"); jwtKey == "" {
			return errors.New("Failed to retrieve jwtKey")
		}

		if jwtExpiryString = os.Getenv("CONVOY_JWT_EXPIRY"); jwtExpiryString == "" {
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

	if retryStrategy := os.Getenv("CONVOY_RETRY_STRATEGY"); retryStrategy != "" {

		intervalSeconds, err := retrieveIntfromEnv("CONVOY_INTERVAL_SECONDS")
		if err != nil {
			return err
		}

		retryLimit, err := retrieveIntfromEnv("CONVOY_RETRY_LIMIT")
		if err != nil {
			return err
		}

		c.Strategy = StrategyConfiguration{
			Type: StrategyProvider(retryStrategy),
			Default: struct {
				IntervalSeconds uint64 `json:"intervalSeconds"`
				RetryLimit      uint64 `json:"retryLimit"`
			}{
				IntervalSeconds: intervalSeconds,
				RetryLimit:      retryLimit,
			},
		}

	}

	c.UIAuthorizedUsers = parseAuthorizedUsers(c.UIAuth)

	cfgSingleton.Store(c)
	return nil
}

func ensureSignature(signature SignatureConfiguration) error {
	_, ok := algo.M[signature.Hash]
	if !ok {
		return fmt.Errorf("invalid hash algorithm - '%s', must be one of %s", signature.Hash, reflect.ValueOf(algo.M).MapKeys())
	}
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

func retrieveIntfromEnv(config string) (uint64, error) {
	value, err := strconv.Atoi(os.Getenv(config))
	if err != nil {
		return 0, errors.New("Failed to parse - " + config)
	}

	if value == 0 {
		return 0, errors.New("Invalid - " + config)
	}

	return uint64(value), nil
}

// Get fetches the application configuration. LoadConfig must have been called
// previously for this to work.
// Use this when you need to get access to the config object at runtime
func Get() (Configuration, error) {
	c, ok := cfgSingleton.Load().(*Configuration)
	if !ok {
		return Configuration{}, errors.New("call Load before this function")
	}

	return *c, nil
}
