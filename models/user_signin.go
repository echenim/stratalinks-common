package models

import "time"

type UserSignIn struct {
	Id                        string     `json:"id"`
	UserId                    string     `json:"user_id"`
	UserName                  string     `json:"user"`
	PasswordHash              string     `json:"password"`
	TwoFactorEnabled          bool       `json:"two_factor_enabled"`
	Lockoutaenabled           bool       `json:"lockout"`
	LockOutEnd                *time.Time `json:"lockout_end"`
	LastSuccessfuleSignInDate *time.Time `json:"last_successfule_sign_in_date"`
	LastFailedSignInDate      *time.Time `json:"last_failed_sign_data"`
}
