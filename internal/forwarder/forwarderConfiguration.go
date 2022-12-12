package forwarder

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

type ForwarderConfiguration struct {
  EhubNamespaceConnectionString   string
  EventHubName                    string
  BlobStorageConnectionString     string
  BlobContainerName               string
  WebSinkURL                      string

  VerboseLevel                    int

  PostMaxRetries                  int
  PostRetryIncrementalDelay       int
}

func InitConfig() *ForwarderConfiguration {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  fmt.Println("InitConfig")
  
  config := &ForwarderConfiguration{}

  config.EhubNamespaceConnectionString = os.Getenv("EHUBNAMESPACECONNECTIONSTRING")
  if config.EhubNamespaceConnectionString == "" {
    log.Fatal("EhubNamespaceConnectionString is not set")
  }

  config.EventHubName = os.Getenv("EVENTHUBNAME")
  if config.EventHubName == "" {
    log.Fatal("EventHubName is not set")
  }

  config.BlobStorageConnectionString = os.Getenv("BLOBSTORAGECONNECTIONSTRING")
  if config.BlobStorageConnectionString == "" {
    log.Fatal("BlobStorageConnectionString is not set")
  }

  config.BlobContainerName = os.Getenv("BLOBCONTAINERNAME")
  if config.BlobContainerName == "" {
    log.Fatal("BlobContainerName is not set")
  }

  config.WebSinkURL = os.Getenv("WEBSINKURL")
  if config.WebsinkURL == "" {
    log.Fatal("WebSinkURL is not set")
  }

  verboseLevel := os.Getenv("VERBOSELEVEL")
  if verboseLevel != "" {
    config.VerboseLevel, err = strconv.Atoi(verboseLevel)
    if err != nil {
      log.Fatal("VerboseLevel is not set")
    }
  }

  config.PostMaxRetries = 10
  config.PostRetryIncrementalDelay = 1000

}

