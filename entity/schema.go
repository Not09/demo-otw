package entity

import(
	"time"
	"gorm.io/gorm"
)

type Requirecustomer struct{
	gorm.Model

	ID_require	 string `gorm:"uniqueIndex"`
	Team	string
	Source  string
	FW		string
	Timerequested time.Time
	Product string
	QTY 	int
	Level	string
	ID_Ar	string
	Drive_sn string
	HSA_sn	string
	Faliure_mode string
	Failhead string
	Engowner string
	Sender	string
	R_remark string



}

type Labheadfa struct{
	Name_Recived string
	Timerecieved time.Time
	L_remark	string

}

type Engheadfa struct{
	Name_Approved string
	Faflow  string
	E_remark	string
}