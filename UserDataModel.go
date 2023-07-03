package main

import "encoding/json"

func UnmarshalUserData(data []byte) (UserData, error) {
	var r UserData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UserData struct {
	Users []User `json:"users"`
}

type User struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	FavoriteFood string   `json:"favoriteFood"`
	Wishlist     []string `json:"wishlist"`
}
