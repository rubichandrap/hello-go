package models

import (
	"time"

	"gorm.io/gorm"
)

type ConfigFoundation struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null"`
	SiteSettings
	OauthStorageSettings
	MailSettings
	ThemeSettings
	SMSGatewaySettings
	PaymentGatewaySettings
	BankTransferSettings
	EWalletSettings
	CreditCardSettings
	DirectDebitSettings
	CreatedAt       time.Time
	UpdatedAt       time.Time
	EmailVerifiedAt time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type SiteSettings struct {
	SiteName        string `gorm:"size:255;not null"`
	SiteTitle       string `gorm:"size:255;not null"`
	SiteDescription string `gorm:"type:text;not null"`
	SiteKeyword     string `gorm:"not null"`
	SiteUrl         string `gorm:"unique;not null"`
	DashboardUrl    string `gorm:"unique;not null"`
	ApiUrl          string `gorm:"not null"`
	ApiDashboardUrl string `gorm:"not null"`
	SiteLanguage    string
	SiteTimezone    string
	FacebookUrl     string
	InstagramUrl    string
	PinterestUrl    string
	TwitterUrl      string
	YoutubeUrl      string
	UserAgent       string
	AllowUrl        string
	DisallowUrl     string
	ShouldIndex     string
	Status          bool `gorm:"default:0"`
}

type OauthStorageSettings struct {
	TwitterApiKey        string
	TwitterApiSecret     string
	GoogleClientId       string
	GoogleClientSecret   string
	GoogleMapKey         string
	GoogleAnalyticKey    string
	FacebookClientId     string
	FacebookClientSecret string
	StorageBucket        string `gorm:"default:local"`
}

type MailSettings struct {
	MailHost        string
	MailPort        string
	MailUsername    string
	MailPassword    string
	MailFromAddress string
	MailFromName    string
	MailPrimary     string
	Phone           string
}

type ThemeSettings struct {
	Logo  string `gorm:"default:logo.jpg"`
	Theme string `gorm:"default:default"`
}

type SMSGatewaySettings struct {
	SmsOption    string `gorm:"default:off"`
	TwilioAppId  string
	TwilioNumber string
	TwilioToken  string
}

type PaymentGatewaySettings struct {
	PaymentOption    string `gorm:"default:midtrans"`
	PaymentServerKey string
	PaymentClientKey string
	PaymentSecret    string
}

type BankTransferSettings struct {
	BcaVirtualAccount     string
	BniVirtualAccount     string
	PermataVirtualAccount string
	Echannel              string
	OtherVirtualAccount   string
}

type EWalletSettings struct {
	Gopay     string
	Shopeepay string
	Qris      string
}

type CreditCardSettings struct {
	CreditCard string
}

type DirectDebitSettings struct {
	BcaKlikbca    string
	BcaKlikpay    string
	BriEpay       string
	CimbClicks    string
	DanamonOnline string
	TelkomselCash string
}
