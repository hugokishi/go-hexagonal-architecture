package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Setup(cnn *gorm.DB) {
	if err := cnn.Exec(`
		DO $$ BEGIN
			CREATE TYPE product_type_enum AS ENUM ('snack', 'side-dish', 'drink', 'dessert');
		EXCEPTION
			WHEN duplicate_object THEN null;
		END $$;
	`,
	).AutoMigrate(); err != nil {
		logrus.Error("Error on exec auto migrate create product_type_enum", err)
	}

	if err := cnn.AutoMigrate(&Product{}); err != nil {
		logrus.Error("Error on auto migrate", err)
	}
}
