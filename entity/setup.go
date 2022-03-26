package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("demo-otw.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(
		&Requirecustomer{},
		&Labheadfa{},
		&Engheadfa{},
	)

	db = database

	tracking1 := Requirecustomer{
		ID_require:    "630",
		Team:          "DRIVE PE",
		Source:        "DRIVE",
		FW:            "FW2236",
		Timerequested: time.Now(),
		Product:       "LongsPeak",
		QTY:           1,
		Level:         "HSA",
		ID_Ar:         "AR1858110",
		Drive_sn:      "ZVT0Qy4G",
		HSA_sn:        "BFATQCZGB",
		Faliure_mode:  "spiral scratch",
		Failhead:      "19",
		Engowner:      "Thunchanok/4854",
		Sender:        "Saowarod/6435",
		R_remark:      "",
	}
	db.Model(&Requirecustomer{}).Create(&tracking1)

	tracking2 := Requirecustomer{
		ID_require:    "631",
		Team:          "ASE",
		Source:        "Reliability",
		FW:            "FW2236",
		Timerequested: time.Now(),
		Product:       "LongsPeak",
		QTY:           1,
		Level:         "HSA",
		ID_Ar:         "AR1858920",
		Drive_sn:      "ZVT0SCA6",
		HSA_sn:        "DEL3QA7ST",
		Faliure_mode:  "H.Deg.rd",
		Failhead:      "all hd",
		Engowner:      "Suchart/3192",
		Sender:        "Phitchayanin / 3297",
		R_remark:      "",
	}
	db.Model(&Requirecustomer{}).Create(&tracking2)

	tracking3 := Requirecustomer{
		ID_require:    "28",
		Team:          "HME",
		Source:        "Drive ",
		FW:            "FW2237",
		Timerequested: time.Now(),
		Product:       "LongsPeakBP",
		QTY:           1,
		Level:         "HSA",
		ID_Ar:         "AR1859719",
		Drive_sn:      "ZX2007RT",
		HSA_sn:        "BEMYQDGS0",
		Faliure_mode:  "EC10591",
		Failhead:      "4",
		Engowner:      "Penpathu/6119",
		Sender:        "Waraporn/4518",
		R_remark:      "",
	}
	db.Model(&Requirecustomer{}).Create(&tracking3)

	tracking4 := Requirecustomer{
		ID_require:    "843",
		Team:          "ASE",
		Source:        "Customer ",
		FW:            "FW2238",
		Timerequested: time.Now(),
		Product:       "Rosewood",
		QTY:           1,
		Level:         "HSA",
		ID_Ar:         "AR1861588",
		Drive_sn:      "ZDZLSTZJ",
		HSA_sn:        "TFBKDBG8Z",
		Faliure_mode:  "High PES - Other",
		Failhead:      "3",
		Engowner:      "Jakkaphan C/4810",
		Sender:        "Phitchayanin / 3297",
		R_remark:      "",
	}
	db.Model(&Requirecustomer{}).Create(&tracking4)

	tracking5 := Requirecustomer{
		ID_require:    "5",
		Team:          "ASE",
		Source:        "Reliability ",
		FW:            "FW2238",
		Timerequested: time.Now(),
		Product:       "EagleRayBP",
		QTY:           1,
		Level:         "HSA",
		ID_Ar:         "AR1861304 ",
		Drive_sn:      "ZMQ01NPC",
		HSA_sn:        "DEK6Q1TSP",
		Faliure_mode:  "H.Deg.WR",
		Failhead:      "3",
		Engowner:      "Jakkaphan C/4810",
		Sender:        "Phitchayanin / 3297",
		R_remark:      "",
	}
	db.Model(&Requirecustomer{}).Create(&tracking5)

	recived1 := Labheadfa{
		Name_Recived: "KANCHANAPHORN",
		Timerecieved: time.Now(),
		L_remark:     "",
	}
	db.Model(&Labheadfa{}).Create(&recived1)

	approved1 := Engheadfa{
		Name_Approved: "PAKORN",
		Faflow:        "Full flow",
		E_remark:      "",
	}
	db.Model(&Engheadfa{}).Create(&approved1)
}