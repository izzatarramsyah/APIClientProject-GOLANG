package interfaces

import(
	"../models"
)
type Cache interface{

	Set(key string, value []models.Employee) (err error)
	Get(key string) (res []models.Employee, err error)
	
}