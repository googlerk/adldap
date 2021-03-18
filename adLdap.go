package adldap

import (
	"crypto/tls"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-ldap/ldap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/googlerk/adcrypt"
	"github.com/googlerk/adenv"
)

var DBsql *sql.DB
var DBgorm *gorm.DB

const dbEmpAd = "employee_ad"
const dbEmp = "employee"

// func main() {
// 	// var sttEmpAd = new(EmployeeAd)
// 	sttEmpAd, err := GetInfoFromAD("piyatida.c")
// 	log.Printf("\n %#v \n", sttEmpAd)
// 	log.Printf("\n %#v \n", err)
// }

func GetEmpByLoginName(loginname string) (sttEmpAd *EmployeeAd, err error) {
	var sttReq = new(ReqAD)
	sttEnv, err := adenv.GetAdEnv("../.env")
	sttReq.ReqUser = loginname
	sttEmpAd, err = adBindingGetInfo(sttReq, sttEnv)
	// fmt.Printf("%#v", sttEmpAd)
	// fmt.Printf("%#v", err)
	return sttEmpAd, err
}

func AddEmpAd(sttEmpAd *EmployeeAd) (boo bool, err error) {
	err = ConnectDB()
	if err != nil {
		return false, err
	}
	defer DBsql.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(sttEmpAd)
	if err != nil {
		return false, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	delete(jsonMap, "ad_id")

	resGORM := DBgorm.Table(dbEmpAd).Create(jsonMap)
	if resGORM.Error != nil {
		return false, resGORM.Error
	}
	return true, resGORM.Error
}

func UpdateEmpAd(sttEmpAd *EmployeeAd) (boo bool, err error) {
	err = ConnectDB()
	if err != nil {
		return false, err
	}
	defer DBsql.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(sttEmpAd)
	if err != nil {
		return false, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)

	resGORM := DBgorm.Table(dbEmpAd).Where("ad_id = ?", jsonMap["ad_id"]).Updates(jsonMap)
	if resGORM.Error != nil {
		return false, resGORM.Error
	}
	// fmt.Printf("\n Task GORM : Insert \n")
	return true, resGORM.Error
}

func InsertOnDupEmpAd(sttEmpAd *EmployeeAd) (boo bool, err error) {
	err = ConnectDB()
	if err != nil {
		return false, err
	}
	defer DBsql.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(sttEmpAd)
	if err != nil {
		return false, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	delete(jsonMap, "ad_id")

	tmpsAMAccountName := fmt.Sprintf("%s", jsonMap["sAMAccountName"])

	var QryEmployeeAd EmployeeAd
	resGORM := DBgorm.Table(dbEmpAd).Where("sAMAccountName = ?", tmpsAMAccountName).Last(&QryEmployeeAd)
	if resGORM.RowsAffected > 0 {
		resGORM = DBgorm.Table(dbEmpAd).Where("ad_id = ?", QryEmployeeAd.Ad_id).Updates(jsonMap)
		// fmt.Printf("\n Task GORM : Update \n")
		if resGORM.Error != nil {
			return false, resGORM.Error
		}
		return true, resGORM.Error
	}

	resGORM = DBgorm.Table(dbEmpAd).Create(jsonMap)
	if resGORM.Error != nil {
		return false, resGORM.Error
	}
	// fmt.Printf("\n Task GORM : Insert \n")
	return true, resGORM.Error
}

func AddEmp(sttEmp *Employee) (boo bool, err error) {
	err = ConnectDB()
	if err != nil {
		return false, err
	}
	defer DBsql.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(sttEmp)
	if err != nil {
		return false, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	delete(jsonMap, "usr_id")

	resGORM := DBgorm.Table(dbEmp).Create(jsonMap)
	if resGORM.Error != nil {
		return false, resGORM.Error
	}
	return true, resGORM.Error
}

func UpdateEmp(sttEmp *Employee) (boo bool, err error) {
	err = ConnectDB()
	if err != nil {
		return false, err
	}
	defer DBsql.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(sttEmp)
	if err != nil {
		return false, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)

	resGORM := DBgorm.Table(dbEmp).Where("usr_id = ?", jsonMap["usr_id"]).Updates(jsonMap)
	if resGORM.Error != nil {
		return false, resGORM.Error
	}
	// fmt.Printf("\n Task GORM : Insert \n")
	return true, resGORM.Error
}

func InsertOnDupEmp(sttEmp *Employee) (boo bool, err error) {
	err = ConnectDB()
	if err != nil {
		return false, err
	}
	defer DBsql.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(sttEmp)
	if err != nil {
		return false, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	delete(jsonMap, "usr_id")

	tmpWhere := fmt.Sprintf("%s", jsonMap["Employee_ID"])

	var QryEmployee Employee
	resGORM := DBgorm.Table(dbEmp).Where("Employee_ID = ?", tmpWhere).Last(&QryEmployee)
	if resGORM.RowsAffected > 0 {
		resGORM = DBgorm.Table(dbEmp).Where("usr_id = ?", QryEmployee.usr_id).Updates(jsonMap)
		// fmt.Printf("\n Task GORM : Update \n")
		if resGORM.Error != nil {
			return false, resGORM.Error
		}
		return true, resGORM.Error
	}

	resGORM = DBgorm.Table(dbEmp).Create(jsonMap)
	if resGORM.Error != nil {
		return false, resGORM.Error
	}
	// fmt.Printf("\n Task GORM : Insert \n")
	return true, resGORM.Error
}

func SelectEmpByID(EmpID string) (sttEmp *Employee, err error) {
	err = ConnectDB()
	if err != nil {
		return sttEmp, err
	}
	defer DBsql.Close()
	var QryEmployee Employee
	resGORM := DBgorm.Table(dbEmp).Where("Employee_ID = ?", EmpID).Last(&QryEmployee)
	if resGORM.Error != nil {
		return sttEmp, resGORM.Error
	}
	return sttEmp, resGORM.Error
}

func SelectEmpAdByID(EmpID string) (sttEmpAd *EmployeeAd, err error) {
	err = ConnectDB()
	if err != nil {
		return sttEmpAd, err
	}
	defer DBsql.Close()
	var QryEmployeeAd EmployeeAd
	resGORM := DBgorm.Table(dbEmpAd).Where("Employee_ID = ?", EmpID).Last(&QryEmployeeAd)
	if resGORM.Error != nil {
		return sttEmpAd, resGORM.Error
	}
	return sttEmpAd, resGORM.Error
}

func UserAuthenPassDN(dn string, pass string) (bool, error) {
	var sttReq = new(ReqAD)
	sttEnv, err := adenv.GetAdEnv("")
	sttReq.ReqDN = dn
	sttReq.ReqPass = pass
	// var sttEmployeeAd = new(EmployeeAd)
	auth, err := adBindingAuthenDnPass(sttEnv, sttReq)
	return auth, err
}

func SyncTableEmployee(sttEmpAd *EmployeeAd) (rowIns int64, err error) {
	err = ConnectDB()
	if err != nil {
		return 0, err
	}
	defer DBsql.Close()

	var sttEmp = new(Employee)
	resGORM := DBgorm.Table(dbEmp).Where("Employee_ID = ?", sttEmpAd.Employee_ID).Last(&sttEmp)
	if resGORM.RowsAffected == 0 {
		sttEmp.Employee_ID = sttEmpAd.Employee_ID
		sttEmp.Employee_Name = sttEmpAd.DisplayName
		sttEmp.Organization = sttEmpAd.Company
		sttEmp.Email = sttEmpAd.Mail
		sttEmp.Mobile = sttEmpAd.Mobile
		sttEmp.Sts_id = "2"
		resGORM = DBgorm.Table(dbEmp).Create(sttEmp)
		if resGORM.Error != nil {
			return resGORM.RowsAffected, resGORM.Error
		}
		// fmt.Printf("\n\n Employee : \n %#v \n\n", sttEmp)
	}
	if sttEmpAd.Msrtcsipuserenabled != "TRUE" {
		sttEmp.Sts_id = "2"
	}
	resGORM = DBgorm.Table(dbEmp).Where("Employee_ID = ?", sttEmp.Employee_ID).Updates(sttEmp)
	if resGORM.Error != nil {
		return resGORM.RowsAffected, resGORM.Error
	}
	return resGORM.RowsAffected, resGORM.Error
}

func adBindingAuthenDnPass(sttEnv *adenv.AdEnv, sttReq *ReqAD) (bool, error) {
	ad, err := ldap.DialTLS(sttEnv.DBProtocal, sttEnv.BindHost+":"+sttEnv.BindPort, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return false, err
	}
	// fmt.Println("ConnectLDAP : Pass")
	err = ad.Bind(sttReq.ReqDN, sttReq.ReqPass)
	if err != nil {
		return false, err
	}
	// fmt.Println("BindDN : Pass")
	return true, err
}

func ConnectDB() (err error) {
	sttEnv, err := adenv.GetAdEnv("")
	if err != nil {
		return err
	}
	// log.Printf("\n connectDB : %#v \n", err)
	db, err := sql.Open(sttEnv.DBDriver, concatConfigDB(sttEnv))
	if err != nil {
		// log.Printf("\n %#v \n", err)
		return err
	}
	DBsql = db

	dbg, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		// log.Printf("\n %#v \n", err)
		return err
	}
	DBgorm = dbg
	return nil
}

func GetInfoFromAD(s string) (sttEmpAd *EmployeeAd, err error) {
	var sttReq = new(ReqAD)
	sttEnv, err := adenv.GetAdEnv("")
	sttReq.ReqUser = s
	sttEmpAd, err = adBindingGetInfo(sttReq, sttEnv)
	// fmt.Printf("%#v", sttEmpAd)
	// fmt.Printf("%#v", err)
	return sttEmpAd, err
}

func GetDNtableEmpAd(user string) (string, error) {
	err := ConnectDB()
	defer DBsql.Close()
	if err != nil {
		return "", err
	}
	var QryEmployeeAd EmployeeAd
	resGORM := DBgorm.Limit(1).Table(dbEmpAd).Find(&QryEmployeeAd, "sAMAccountName = ?", user)
	return string(QryEmployeeAd.Dn), resGORM.Error
}

func InsertEmpAd(sttEmpAd *EmployeeAd) (rowIns int64, err error) {
	err = ConnectDB()
	if err != nil {
		return 0, err
	}
	defer DBsql.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(sttEmpAd)
	if err != nil {
		return 0, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	delete(jsonMap, "ad_id")

	tmpsAMAccountName := fmt.Sprintf("%s", jsonMap["sAMAccountName"])
	// fmt.Printf("\n\n %s \n", tmpsAMAccountName)

	var QryEmployeeAd EmployeeAd
	resGORM := DBgorm.Table(dbEmpAd).Where("sAMAccountName = ?", tmpsAMAccountName).Last(&QryEmployeeAd)
	if resGORM.RowsAffected > 0 {
		resGORM = DBgorm.Table(dbEmpAd).Where("ad_id = ?", QryEmployeeAd.Ad_id).Updates(jsonMap)
		// fmt.Printf("\n Task GORM : Update \n")
		return resGORM.RowsAffected, resGORM.Error
	}

	resGORM = DBgorm.Table(dbEmpAd).Create(jsonMap)
	if resGORM.Error != nil {
		return resGORM.RowsAffected, resGORM.Error
	}
	// fmt.Printf("\n Task GORM : Insert \n")
	return resGORM.RowsAffected, resGORM.Error
}

func adBindingGetInfo(sttReq *ReqAD, sttEnv *adenv.AdEnv) (sttEmpAd *EmployeeAd, err error) {
	ad, err := ldap.DialTLS("tcp", sttEnv.BindHost+":"+sttEnv.BindPort, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return sttEmpAd, err
	}
	err = ad.Bind(sttEnv.BindDN, string(adcrypt.Decrypt(sttEnv.BindPass, sttEnv.ValKey)))
	if err != nil {
		return sttEmpAd, err
	}
	searchReq := ldap.NewSearchRequest(sttEnv.BindBaseDN,
		ldap.ScopeWholeSubtree, 0, 0, 0, false,
		"(&(sAMAccountName="+sttReq.ReqUser+")(objectClass=user)(objectCategory=person))",
		[]string{"*"},
		[]ldap.Control{})
	sr, err := ad.Search(searchReq)
	if err != nil {
		log.Fatal(err)
	}

	txtJSON := ""
	for _, entry := range sr.Entries {
		for _, attr := range entry.Attributes {
			if txtJSON != "" {
				txtJSON += ", "
			}
			switch {
			case swList((strings.ToLower(attr.Name))):
				s, _ := base64.StdEncoding.DecodeString(attr.Values[0])
				txtJSON += `"` + string(attr.Name) + `"` + " : " + `"` + jsonEscape(string(s)) + `"` + ""
			default:
				txtJSON += `"` + string(attr.Name) + `"` + " : " + `"` + jsonEscape(attr.Values[0]) + `"` + ""
			}
		}
	}
	txtJSON = `{` + txtJSON + `}`
	err = json.Unmarshal([]byte(txtJSON), &sttEmpAd)
	sttEmpAd = validateEmpAd(sttEnv, sttEmpAd)

	return sttEmpAd, err
}

func validateEmpAd(sttEnv *adenv.AdEnv, sttEmpAd *EmployeeAd) *EmployeeAd {
	sttEmpAd.Createdate = sttEmpAd.WhenCreated
	sttEmpAd.Updatedate = sttEmpAd.WhenChanged
	if sttEmpAd.Dn == "" {
		sttEmpAd.Dn = sttEmpAd.DistinguishedName
	}
	switch strings.ToLower(sttEnv.ConstEnv) {
	case "production":
		sttEmpAd.Employee_ID = sttEmpAd.Employeenumber
	case "uat":
		sttEmpAd.Employee_ID = sttEmpAd.Description
	}
	return sttEmpAd
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	s := string(b)
	return s[1 : len(s)-1]
}

func concatConfigDB(sttEnv *adenv.AdEnv) string {
	res := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sttEnv.DBUser,
		sttEnv.DBPass,
		sttEnv.DBProtocal,
		sttEnv.DBHost,
		sttEnv.DBPort,
		sttEnv.DBName,
	)
	return res
}

func swList(s string) bool {
	// List ชื่อ Field ของ Table employee_ad ที่ไม่นำเข้า Insert
	a := []string{""}
	for _, i := range a {
		if i == s {
			return true
		}
	}
	return false
}
