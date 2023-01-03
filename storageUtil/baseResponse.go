package storageUtil

import()

type Response struct {
	Code int `json:"status"`
	Message string `json:"message"`
	Success bool `json:"success"`
	Data map[string]interface{} `json:"data"`
}