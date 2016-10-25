package models_test

import (
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/apiseedprojects/go/bootstrap"
	"github.com/apiseedprojects/go/models"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
)

func setupFixtures(gdb *gorm.DB) []*models.User {
	userModels := []*models.User{
		&models.User{Username: "alfa", Password: "@lfa"},
		&models.User{Username: "beta", Password: "b3t4"},
		&models.User{Username: "gamma", Password: "g4mm@"},
	}
	gdb.Unscoped().Delete(models.User{})
	if gdb.Error != nil {
		panic(fmt.Errorf("error wiping users table: %s", gdb.Error.Error()))
	}
	for _, user := range userModels {
		gdb.Create(user)
		if gdb.Error != nil {
			panic(fmt.Errorf("unable to create users fixtures: %s", gdb.Error.Error()))
		}
	}
	return userModels
}

func getGormDB() *gorm.DB {
	logrusCurrentLevel := logrus.GetLevel()
	logrus.SetLevel(logrus.PanicLevel)
	_, config := bootstrap.ReadConfig("../local/config.json", false)
	logrus.SetLevel(logrusCurrentLevel)
	gdb, err := gorm.Open("mysql", config.DBConnectionString)
	if err != nil {
		panic(err)
	}
	return gdb
}

func TestUserModel(t *testing.T) {
	gdb := getGormDB()
	um := models.NewUsersModel(gdb)

	gdb.LogMode(false)

	Convey("Given an existing list of users ", t, func() {
		_ = setupFixtures(gdb)
		Convey("When I request a list of users", func() {
			userList, err := um.List()
			Convey("Then I get a list of users", func() {
				So(err, ShouldBeNil)
				So(len(*userList), ShouldEqual, 3)
			})
		})
	})
	Convey("Given an existing list of users ", t, func() {
		pk := setupFixtures(gdb)
		Convey("When I request a single existing User", func() {
			user, err := um.Get(pk[0].ID)
			Convey("Then I should get it", func() {
				So(err, ShouldBeNil)
				So(user.ID, ShouldEqual, pk[0].ID)
				So(user.Username, ShouldEqual, pk[0].Username)
				So(user.Password, ShouldEqual, pk[0].Password)
			})
		})
		Convey("When I request a single non-existing User", func() {
			user, err := um.Get(0)
			Convey("Then I should get an error about it", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "record not found")
				So(user, ShouldBeNil)
			})
		})
	})
	Convey("Given an existing list of users ", t, func() {
		pk := setupFixtures(gdb)
		Convey("When I create a new, unique user name", func() {
			inUser := models.User{
				Username: "testusername",
				Password: "testpassword",
			}
			outUser, err := um.Create(&inUser)
			Convey("Then I should not get errors and be able to retrieve it", func() {
				So(err, ShouldBeNil)
				So(outUser.ID, ShouldBeGreaterThan, 0)
				So(outUser.Username, ShouldEqual, "testusername")
				So(outUser.Password, ShouldEqual, "testpassword")

				fetchedUser := &models.User{}
				err := um.GDB.Find(fetchedUser, outUser.ID).Error
				So(err, ShouldBeNil)
				So(fetchedUser.Username, ShouldEqual, "testusername")
				So(fetchedUser.Password, ShouldEqual, "testpassword")
			})
		})
		Convey("When I create a new, duplicate username", func() {
			inUser := models.User{
				Username: pk[0].Username,
				Password: "testpassword",
			}
			outUser, err := um.Create(&inUser)
			Convey("Then I should get an error about duplicity of keys", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "duplicate")
				So(outUser, ShouldBeNil)
			})
		})
	})
	Convey("Given an existing list of users ", t, func() {
		pk := setupFixtures(gdb)
		Convey("When I update a user", func() {
			u := &models.User{
				Username: "seconduser",
				Password: "secondpass",
			}
			Convey("it should be correctly updated", func() {
				So(pk[0].Username, ShouldEqual, "alfa")
				So(pk[0].Password, ShouldEqual, "@lfa")
				uu, err := um.Update(pk[0].ID, u)
				So(err, ShouldBeNil)
				So(uu.Username, ShouldEqual, "seconduser")
				So(uu.Password, ShouldEqual, "secondpass")
				fu := &models.User{}
				err = um.GDB.First(fu, pk[0].ID).Error
				So(err, ShouldBeNil)
				So(fu.Username, ShouldEqual, "seconduser")
				So(fu.Password, ShouldEqual, "secondpass")
			})
		})
		Convey("When I try to update a non existing user", func() {
			u := &models.User{}
			_, err := um.Update(0, u)
			Convey("then I should get a record not found error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "record not found")
			})
		})
	})
	Convey("Given an existing list of users ", t, func() {
		pk := setupFixtures(gdb)
		Convey("When I delete an existing user", func() {
			Convey("Then the user should be deleted", func() {
				oid := pk[0].ID
				ou := &models.User{}
				err := um.GDB.First(ou, oid).Error
				So(err, ShouldBeNil)
				So(ou.Username, ShouldEqual, "alfa")
				So(ou.Password, ShouldEqual, "@lfa")
				err = um.Delete(oid)
				So(err, ShouldBeNil)
				xu := &models.User{}
				err = um.GDB.First(xu, oid).Error
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "record not found")
			})
		})
		Convey("When I delete a non-existing user", func() {
			Convey("Then we should get a record not found error", func() {
				err := um.Delete(65535)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "record not found")
			})
		})
	})
}
