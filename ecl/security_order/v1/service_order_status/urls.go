package service_order_status

import "github.com/nttcom/eclcloud"

func getURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/ScreenEventFGSOrderProgressRate")
}
