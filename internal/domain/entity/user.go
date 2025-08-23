package entity

import (
    "errors"
    "time"
)

type User struct {
    ID        string    `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Follow struct {
    ID         string    `json:"id" db:"id"`
    FollowerID string    `json:"follower_id" db:"follower_id"`
    FolloweeID string    `json:"followee_id" db:"followee_id"`
    CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

func (u *User) Validate() error {
    if u.Name == "" {
        return errors.New("name cannot be empty")
    }
    return nil
}