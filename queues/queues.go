package queues

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func GetClient() *azservicebus.Client {
	// namespace, ok := os.LookupEnv("AZURE_SERVICEBUS_HOSTNAME")
	// if !ok {
	// 	panic("AZURE_SERVICEBUS_HOSTNAME environment variable not found")
	// }

	// add env_vars to config.yml: https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63
	// https://www.loginradius.com/blog/engineering/environment-variables-in-golang/

	namespace := "mlsa3queues.servicebus.windows.net"
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azservicebus.NewClient(namespace, cred, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func SendMessage(message string) {
	client := GetClient()
	sender, err := client.NewSender("mlsa3onboardingqueue", nil)
	if err != nil {
		panic(err)
	}
	defer sender.Close(context.TODO())

	sbMessage := &azservicebus.Message{
		Body: []byte(message),
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		panic(err)
	}
}
