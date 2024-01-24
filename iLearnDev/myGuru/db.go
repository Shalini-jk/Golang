package main

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name   string
	Email  string  `gorm:"uniqueIndex"` // Unique index to prevent duplicate emails
	Skills []Skill `gorm:"many2many:user_skills"`
}

type Skill struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex"` // Unique index to prevent duplicate skill names
	MyGuruGitRepo string //git repo to configure mentor's content
	Users         []User `gorm:"many2many:user_skills"` // Many-to-many relationship
}

type UserSkill struct {
	gorm.Model

	UserID uint
	// User   User `gorm:"foreignKey:UserID"`

	SkillID uint
	// Skill   Skill `gorm:"foreignKey:SkillID"`

	GitRepo string
}

//skill name should be unique
func SkillAdd(db *gorm.DB, skillName, GitRepo string) error {
	skill := Skill{
		Name:          skillName,
		MyGuruGitRepo: GitRepo,
	}

	result := db.Create(&skill)

	return result.Error
}

func SkillFind(db *gorm.DB, skillName string) (*Skill, error) {
	skill := Skill{}

	result := db.First(&skill, "name = ?", skillName)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return &skill, nil
	}
}

func UserAdd(db *gorm.DB, userName, userEmail string) error {
	user := User{
		Name:  userName,
		Email: userEmail,
	}

	result := db.Create(&user)

	return result.Error
}

func UserFind(db *gorm.DB, userEmail string) (*User, error) {
	user := User{}

	result := db.First(&user, "email = ?", userEmail)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

//Add user and then update the user GitRepo too
func UserSkillUpdateGitRepo(db *gorm.DB, userEmail, skillName, GitRepo string) error {

	skill, err := SkillFind(db, skillName)
	if err != nil {
		return err
	}

	user, err := UserFind(db, userEmail)
	if err != nil {
		return err
	}

	//associate the skill and the user
	db.Model(user).Association("Skills").Append(skill)
	userSkill := UserSkill{}
	result := db.First(&userSkill, "user_id = ? AND skill_id = ?", user.ID, skill.ID)
	if result.Error != nil {
		return result.Error
	}

	userSkill.GitRepo = GitRepo
	db.Save(&userSkill)
	return nil
}

func initDB(logger logrus.FieldLogger) error {
	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database. Err: ", err)
		return err
	}
	// Auto-migrate schema (create tables if they don't exist)
	db.AutoMigrate(&User{}, &Skill{}, &UserSkill{})
	return nil
}

// func main() {
// 	var userskills []UserSkill
// 	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to database")
// 	}
// 	// Auto-migrate schema (create tables if they don't exist)
// 	db.AutoMigrate(&User{}, &Skill{}, &UserSkill{})

// 	err = SkillAdd(db, "Go", "http://github.com/myGuruGoRepo")
// 	if err != nil {
// 		fmt.Println("SkillAdd failed, Error:", err)
// 	} else {
// 		fmt.Println("SkillAdd passed, added skill :", "Go")
// 	}

// 	// err = SkillAdd(db, "Go", "test Repo")
// 	// if err != nil {
// 	// 	fmt.Println("SkillAdd failed, Error:", err)
// 	// } else {
// 	// 	fmt.Println("SkillAdd passed, added skill :", "Go")
// 	// }

// 	err = UserAdd(db, "Anurag", "anurag@ambersoft.in")
// 	if err != nil {
// 		fmt.Println("UserAdd failed, Error:", err)
// 	} else {
// 		fmt.Println("UserAdd passed, added skill :", "Anurag")
// 	}

// 	err = UserSkillUpdateGitRepo(db, "anurag@ambersoft.in", "Go", "http://github.com/anuragGoRepo")
// 	if err != nil {
// 		fmt.Println("UserSkillUpdateGitRepo failed, Error:", err)
// 	} else {
// 		fmt.Println("UserSkillUpdateGitRepo passed, added :", "http://github.com/anuragGoRepo")
// 	}

// 	return

// 	// Create users and skills
// 	user1 := User{Email: "johndoe@example.com"}
// 	user2 := User{Email: "janesmith@example.com"}
// 	skill1 := Skill{Name: "Go"}
// 	skill2 := Skill{Name: "Python"}
// 	skill3 := Skill{Name: "JavaScript"}

// 	// Create records
// 	db.Create(&user1)
// 	db.Create(&user2)
// 	db.Create(&skill1)
// 	db.Create(&skill2)
// 	db.Create(&skill3)

// 	db.Find(&userskills)
// 	fmt.Println("1: Userskill is: ", userskills)
// 	for ind, userskill := range userskills {
// 		fmt.Println("ind : ", ind, "userID : ", userskill.UserID, "SkillID : ", userskill.SkillID, "Repo : ", userskill.GitRepo)
// 	}

// 	// Assign skills to users
// 	db.Model(&user1).Association("Skills").Append(&skill1, &skill3)

// 	db.Find(&userskills)
// 	fmt.Println("2: Userskill is: ", userskills)
// 	for ind, userskill := range userskills {
// 		fmt.Println("ind : ", ind, "userID : ", userskill.UserID, "SkillID : ", userskill.SkillID, "Repo : ", userskill.GitRepo)
// 	}

// 	db.Model(&user2).Association("Skills").Append(&skill2, &skill3)

// 	db.Find(&userskills)
// 	fmt.Println("3: Userskill is: ", userskills)
// 	for ind, userskill := range userskills {
// 		fmt.Println("ind : ", ind, "userID : ", userskill.UserID, "SkillID : ", userskill.SkillID, "Repo : ", userskill.GitRepo)
// 	}

// 	userSkill := UserSkill{}
// 	db.Where("user_id = ? AND skill_id = ? ", user1.ID, skill1.ID).First(&userSkill)

// 	// userSkill := UserSkill{
// 	// 	UserID: user1.ID,
// 	// 	// User:    user1,
// 	// 	SkillID: skill1.ID,
// 	// 	// Skill:   skill1,
// 	// 	GitRepo: "https://github.com/user1skill1repo",
// 	// }
// 	userSkill.GitRepo = "https://github.com/user1skill1repo"
// 	db.Save(&userSkill)

// 	// db.Find(&userskills)
// 	// fmt.Println("4: Userskill is: ", userskills)
// 	// for ind, userskill := range userskills {
// 	// 	fmt.Println("ind : ", ind, "userID : ", userskill.UserID, "SkillID : ", userskill.SkillID, "Repo : ", userskill.GitRepo)
// 	// }

// 	// userSkill = UserSkill{
// 	// 	UserID:  user1.ID,
// 	// 	SkillID: skill3.ID,
// 	// 	GitRepo: "https://github.com/user1skill3repo",
// 	// }
// 	// db.Create(&userSkill)

// 	// userSkill = UserSkill{
// 	// 	UserID:  user2.ID,
// 	// 	SkillID: skill2.ID,
// 	// 	GitRepo: "https://github.com/user2skill2repo",
// 	// }
// 	// db.Create(&userSkill)

// 	db.Find(&userskills)
// 	fmt.Println("5: Userskill is: ", userskills)
// 	for ind, userskill := range userskills {
// 		fmt.Println("ind : ", ind, "userID : ", userskill.UserID, "SkillID : ", userskill.SkillID, "Repo : ", userskill.GitRepo)
// 	}

// 	db.Where("user_id = ? AND skill_id = ? ", user1.ID, skill1.ID).Find(&userskills)

// 	fmt.Println("6: Userskill is: ", userskills)
// 	for ind, userskill := range userskills {
// 		fmt.Println("ind : ", ind, "userID : ", userskill.UserID, "SkillID : ", userskill.SkillID, "Repo : ", userskill.GitRepo)
// 	}

// 	fmt.Println("Userskill after search is: ", userskills)

// 	// List all users with their skills
// 	// users := []User{}
// 	// db.Preload("Skills").Find(&users)
// 	// for _, user := range users {
// 	// 	fmt.Println("User:", user.Email)
// 	// 	// user.Skills = []Skill{}

// 	// 	// db.Model(&user).Association("Skills").Find(&user.Skills)

// 	// 	// for _, skill := range user.Skills {
// 	// 	// 	fmt.Println(skill.Name)
// 	// 	// }
// 	// 	fmt.Println("Skills:", user.Skills)
// 	// }

// 	// skills := []Skill{}
// 	// db.Preload("Users").Find(&skills)
// 	// for _, skill := range skills {
// 	// 	fmt.Println("Skill:", skill.Name)
// 	// 	for _, user := range skill.Users {
// 	// 		fmt.Println(user.Email)
// 	// 	}
// 	// 	// fmt.Println("Skills:", user.Skills)
// 	// }

// 	// Update a user's email
// 	// db.Model(&user1).Update("Email", "johndoe@gmail.com")

// 	// // Delete a skill
// 	// db.Delete(&skill1)
// }
