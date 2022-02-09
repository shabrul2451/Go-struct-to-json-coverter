package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserMetadata struct {
	CompanyId string `json:"company_id" bson:"company_id"`
}

type UserResourcePermission struct {
	Resources []ResourceWiseRoles `json:"resources" bson:"resources"`
}

type ResourceWiseRoles struct {
	Name  string `json:"name" bson:"name"`
	Roles []Role `json:"roles" bson:"roles"`
}

type Role struct {
	Name string `json:"name" bson:"name"`
}

type User struct {
	Metadata           UserMetadata           `json:"metadata" bson:"metadata"`
	ID                 string                 `json:"id" bson:"id"`
	FirstName          string                 `json:"first_name" bson:"first_name" `
	LastName           string                 `json:"last_name" bson:"last_name"`
	Email              string                 `json:"email" bson:"email" `
	Phone              string                 `json:"phone" bson:"phone" `
	Password           string                 `json:"password" bson:"password" `
	Status             string                 `json:"status" bson:"status"`
	CreatedDate        time.Time              `json:"created_date" bson:"created_date"`
	UpdatedDate        time.Time              `json:"updated_date" bson:"updated_date"`
	AuthType           string                 `json:"auth_type" bson:"auth_type"`
	ResourcePermission UserResourcePermission `json:"resource_permission" bson:"resource_permission"`
}

func main() {
	var user User
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
