package adldap

import (
	"crypto/md5"
	"crypto/tls"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

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
	sttEnv, err := adenv.GetAdEnv("")
	sttReq.ReqUser = loginname
	sttEmpAd, err = adBindingGetInfo(sttReq, sttEnv)
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
		resGORM = DBgorm.Table(dbEmp).Where("usr_id = ?", QryEmployee.Usr_id).Updates(jsonMap)
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

func SelectEmpByID(s string) (sttEmp *Employee, err error) {
	err = ConnectDB()
	if err != nil {
		return sttEmp, err
	}
	defer DBsql.Close()
	resGORM := DBgorm.Table(dbEmp).Where("Employee_ID = ?", s).Last(&sttEmp)
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
	resGORM := DBgorm.Table(dbEmpAd).Where("Employee_ID = ?", EmpID).Last(&sttEmpAd)
	if resGORM.Error != nil {
		return sttEmpAd, resGORM.Error
	}
	return sttEmpAd, resGORM.Error
}

func Authen(user string, pass string) (bool, error) {
	var sttReq = new(ReqAD)
	sttEnv, err := adenv.GetAdEnv("")
	sttReq.ReqUser = user
	sttReq.ReqPass = pass
	sttEmpAd, err := adBindingGetDn(sttReq, sttEnv)
	if err != nil {
		return false, err
	}
	sttReq.ReqDN = sttEmpAd.Dn
	// fmt.Printf("\n\n %#v \n\n", sttEmpAd)
	auth, err := adBindingAuthenDnPass(sttEnv, sttReq)
	return auth, err
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
		[]string{"AccountExpires", "Admincount", "BadPasswordTime", "BadPwdCount", "Cn", "CodePage", "Company", "CountryCode", "Createdate", "Department", "Description", "Directreports", "DisplayName", "DistinguishedName", "Dn", "DSCorePropagationData", "Employee_ID", "Employeeid", "Employeenumber", "Gidnumber", "GivenName", "Homedirectory", "Homedrive", "Homephone", "Info", "InstanceType", "Ipphone", "Iscriticalsystemobject", "Lastlogoff", "LastLogon", "LastLogonTimestamp", "Lockouttime", "Loginshell", "LogonCount", "Logonhours", "Mail", "Mailnickname", "Manager", "Memberof", "Mobile", "Msdsconsistencyguid", "Msdssupportedencryptiontypes", "Msnpallowdialin", "Msrtcsipdeploymentlocator", "Msrtcsipfederationenabled", "Msrtcsipinternetaccessenabled", "Msrtcsipoptionflags", "Msrtcsipprimaryhomeserver", "Msrtcsipprimaryuseraddress", "Msrtcsipuserenabled", "Mssfu30name", "Mssfu30nisdomain", "Name", "ObjectCategory", "Objectclass", "ObjectGUID", "ObjectSid", "Physicaldeliveryofficename", "PrimaryGroupID", "Protocolsettings", "Proxyaddresses", "PwdLastSet", "SAMAccountName", "SAMAccountType", "Scriptpath", "Serviceprincipalname", "Showinadvancedviewonly", "Sn", "Targetaddress", "Telephonenumber", "Terminalserver", "Title", "Uid", "Uidnumber", "Unixhomedirectory", "Updatedate", "UserAccountControl", "Userparameters", "UserPrincipalName", "USNChanged", "USNCreated", "WhenChanged", "WhenCreated"},
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

func adBindingGetDn(sttReq *ReqAD, sttEnv *adenv.AdEnv) (sttEmpAd *EmployeeAd, err error) {
	ad, err := ldap.DialTLS("tcp", sttEnv.BindHost+":"+sttEnv.BindPort, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return sttEmpAd, err
	}
	err = ad.Bind(sttEnv.BindDN, string(adcrypt.Decrypt(sttEnv.BindPass, sttEnv.ValKey)))
	if err != nil {
		return sttEmpAd, err
	}
	fieldFind := "sAMAccountName"
	switch strings.ToLower(sttEnv.ConstEnv) {
	case "uat":
		fieldFind = "Cn"
	}
	searchReq := ldap.NewSearchRequest(sttEnv.BindBaseDN,
		ldap.ScopeWholeSubtree, 0, 0, 0, false,
		"(&("+fieldFind+"="+sttReq.ReqUser+")(objectClass=user)(objectCategory=person))",
		[]string{"AccountExpires", "Admincount", "BadPasswordTime", "BadPwdCount", "Cn", "CodePage", "Company", "CountryCode", "Createdate", "Department", "Description", "Directreports", "DisplayName", "DistinguishedName", "Dn", "DSCorePropagationData", "Employee_ID", "Employeeid", "Employeenumber", "Gidnumber", "GivenName", "Homedirectory", "Homedrive", "Homephone", "Info", "InstanceType", "Ipphone", "Iscriticalsystemobject", "Lastlogoff", "LastLogon", "LastLogonTimestamp", "Lockouttime", "Loginshell", "LogonCount", "Logonhours", "Mail", "Mailnickname", "Manager", "Memberof", "Mobile", "Msdsconsistencyguid", "Msdssupportedencryptiontypes", "Msnpallowdialin", "Msrtcsipdeploymentlocator", "Msrtcsipfederationenabled", "Msrtcsipinternetaccessenabled", "Msrtcsipoptionflags", "Msrtcsipprimaryhomeserver", "Msrtcsipprimaryuseraddress", "Msrtcsipuserenabled", "Mssfu30name", "Mssfu30nisdomain", "Name", "ObjectCategory", "Objectclass", "ObjectGUID", "ObjectSid", "Physicaldeliveryofficename", "PrimaryGroupID", "Protocolsettings", "Proxyaddresses", "PwdLastSet", "SAMAccountName", "SAMAccountType", "Scriptpath", "Serviceprincipalname", "Showinadvancedviewonly", "Sn", "Targetaddress", "Telephonenumber", "Terminalserver", "Title", "Uid", "Uidnumber", "Unixhomedirectory", "Updatedate", "UserAccountControl", "Userparameters", "UserPrincipalName", "USNChanged", "USNCreated", "WhenChanged", "WhenCreated"},
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

	if sttEmpAd.Dn == "" {
		sttEmpAd.Dn = sttEmpAd.DistinguishedName
	}
	switch strings.ToLower(sttEnv.ConstEnv) {
	case "production":
		sttEmpAd.Employee_ID = sttEmpAd.Employeenumber
	case "uat":
		sttEmpAd.Employee_ID = sttEmpAd.Description
	case "local":
		sttEmpAd.Employee_ID = sttEmpAd.Employeenumber
	}

	// fmt.Printf("\n\n validateEmpAd %#v \n", sttEmpAd)
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

func SyncProcessFilterAD(s string) (boo bool, err error) {
	err = ConnectDB()
	if err != nil {
		return false, err
	}
	defer DBsql.Close()

	sr, err := LdapSearchFromFilter(s) //search AD ด้วย parameter s
	if err != nil {
		fmt.Println("LdapSearchFromFilter")
		return false, err
	}

	err = UnMarshalAD2EmpAD(sr)
	if err != nil {
		fmt.Println("UnMarshalAD2EmpAD")
		return false, err
	}

	boo, err = BatchUnableUser() //เปลี่ยน sts_id = 2 เมื่อหาใน AD Server ไม่พบ
	if err != nil {
		fmt.Println("BatchUnableUser")
		return false, err
	}

	return true, err
}

func LdapSearchFromFilter(s string) (sr *ldap.SearchResult, err error) {
	sttEnv, err := adenv.GetAdEnv("")
	ad, err := ldap.DialTLS("tcp", sttEnv.BindHost+":"+sttEnv.BindPort, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return sr, err
	}
	err = ad.Bind(sttEnv.BindDN, string(adcrypt.Decrypt(sttEnv.BindPass, sttEnv.ValKey)))
	if err != nil {
		return sr, err
	}
	searchReq := ldap.NewSearchRequest(sttEnv.BindBaseDN,
		ldap.ScopeWholeSubtree, 0, 0, 0, false,
		"(&"+s+"(objectClass=user)(objectCategory=person))",
		[]string{"AccountExpires", "Admincount", "BadPasswordTime", "BadPwdCount", "Cn", "CodePage", "Company", "CountryCode", "Createdate", "Department", "Description", "Directreports", "DisplayName", "DistinguishedName", "Dn", "DSCorePropagationData", "Employee_ID", "Employeeid", "Employeenumber", "Gidnumber", "GivenName", "Homedirectory", "Homedrive", "Homephone", "Info", "InstanceType", "Ipphone", "Iscriticalsystemobject", "Lastlogoff", "LastLogon", "LastLogonTimestamp", "Lockouttime", "Loginshell", "LogonCount", "Logonhours", "Mail", "Mailnickname", "Manager", "Memberof", "Mobile", "Msdsconsistencyguid", "Msdssupportedencryptiontypes", "Msnpallowdialin", "Msrtcsipdeploymentlocator", "Msrtcsipfederationenabled", "Msrtcsipinternetaccessenabled", "Msrtcsipoptionflags", "Msrtcsipprimaryhomeserver", "Msrtcsipprimaryuseraddress", "Msrtcsipuserenabled", "Mssfu30name", "Mssfu30nisdomain", "Name", "ObjectCategory", "Objectclass", "ObjectGUID", "ObjectSid", "Physicaldeliveryofficename", "PrimaryGroupID", "Protocolsettings", "Proxyaddresses", "PwdLastSet", "SAMAccountName", "SAMAccountType", "Scriptpath", "Serviceprincipalname", "Showinadvancedviewonly", "Sn", "Targetaddress", "Telephonenumber", "Terminalserver", "Title", "Uid", "Uidnumber", "Unixhomedirectory", "Updatedate", "UserAccountControl", "Userparameters", "UserPrincipalName", "USNChanged", "USNCreated", "WhenChanged", "WhenCreated"},
		[]ldap.Control{})
	sr, err = ad.Search(searchReq)
	if err != nil {
		log.Fatal(err)
	}
	return sr, err
}

func UnMarshalAD2EmpAD(sr *ldap.SearchResult) (err error) {
	// fmt.Println("Got", len(sr.Entries), "search results") //debug.fmt
	countSr := 0
	var SttResErr = new(ResErrHandle)
	var sttEmpAd = new(EmployeeAd)
	sttEnv, err := adenv.GetAdEnv("")
	if err != nil {
		fmt.Printf("%#v \n", "GetAdEnv")
		return err
	}

	for _, entry := range sr.Entries {
		txtJSON := ""
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
		txtJSON = `{` + txtJSON + `}`
		// fmt.Printf("\n %s \n", txtJSON)
		clearstt(&sttEmpAd)

		err = json.Unmarshal([]byte(txtJSON), &sttEmpAd)
		if err != nil {
			SttResErr.ResErr += "UnMarshalAD2EmpAD \n"
		}
		sttEmpAd = validateEmpAd(sttEnv, sttEmpAd)
		countSr = countSr + 1
		// fmt.Printf("( %v ) -> Employee_ID=%s SAMAccountName=%s Dn=%s \n", countSr, sttEmpAd.Employee_ID, sttEmpAd.SAMAccountName, sttEmpAd.Dn) //debug.fmt

		jsonMap, err := mapJsonEmpAd(sttEmpAd)
		if err != nil {
			SttResErr.ResErr += "mapJsonEmpAd :: Employee_ID : " + sttEmpAd.Employee_ID + " SAMAccountName = " + sttEmpAd.SAMAccountName + " :: Dn = " + sttEmpAd.Dn + "\n"
			goto endLoop
			return err
		}

		jsonMap, err = syncAD2EmpAD(jsonMap)
		if err != nil {
			SttResErr.ResErr += "syncAD2EmpAD :: Employee_ID : " + jsonMap["employee_ID"].(string) + " SAMAccountName = " + jsonMap["sAMAccountName"].(string) + " :: Dn = " + jsonMap["dn"].(string) + "\n"
			goto endLoop
			return err
		}
		// fmt.Printf("\n ------------------------------------------ \n jsonMap : %#v \n", jsonMap)

		jsonMap, err = syncEmpAd2Employee(jsonMap)
		if err != nil {
			SttResErr.ResErr += "syncEmpAd2Employee :: Employee_ID : " + jsonMap["employee_ID"].(string) + " SAMAccountName = " + jsonMap["sAMAccountName"].(string) + " :: Dn = " + jsonMap["dn"].(string) + "\n"
			SttResErr.ResErr += fmt.Sprintf("%s", err) + "\n"
			goto endLoop
			return err
		}
	endLoop:
	}
	// fmt.Printf("\n ===== \n SttResErr.ResErr :  %s \n", SttResErr.ResErr)
	return err
}

func syncAD2EmpAD(mapEmpAd map[string]interface{}) (map[string]interface{}, error) {
	err := errors.New("")
	// fmt.Printf("\n ------------------------------------------ \n mapEmpAd : %#v \n", mapEmpAd)
	delete(mapEmpAd, "ad_id")
	delete(mapEmpAd, "updatedate")

	var QryEmployeeAd []map[string]interface{}
	resGORM := DBgorm.Table(dbEmpAd).Find(&QryEmployeeAd, "sAMAccountName = ?", mapEmpAd["sAMAccountName"])
	// fmt.Printf("------------------------------------------------- \n QryEmployeeAd %#v \n", QryEmployeeAd)
	// fmt.Printf("\n\n sAMAccountName : %s resGORM.RowsAffected : %#v \n\n", mapEmpAd["sAMAccountName"], resGORM.RowsAffected)

	if resGORM.RowsAffected == 0 {
		// fmt.Printf("\n\n resGORM.RowsAffected : %#v \n\n", resGORM.RowsAffected)
		mapEmpAd["createdate"] = time.Now()
		resGORM = DBgorm.Table(dbEmpAd).Create(mapEmpAd)
		if resGORM.Error != nil {
			err = errors.New(fmt.Sprintf("Insert %s from struct EmployeeAd \n %s \n", dbEmpAd, resGORM.Error))
			return mapEmpAd, err
		}
		return mapEmpAd, resGORM.Error
	}

	// mapEmpAd["whenChanged"] = "20210320112255.0Z"
	t1 := convDateTime(mapEmpAd["whenChanged"].(string))
	t2 := convDateTime(QryEmployeeAd[0]["whenChanged"].(string))
	mapEmpAd["createdate"] = QryEmployeeAd[0]["createdate"]
	// fmt.Printf("\n ------------------------------------------ \n mapEmpAd : %#v \n", mapEmpAd)

	if t1 > t2 {
		resGORM = DBgorm.Table(dbEmpAd).Where("ad_id = ?", QryEmployeeAd[0]["ad_id"]).Updates(mapEmpAd)
		if resGORM.Error != nil {
			err = errors.New(fmt.Sprintf("Update %s Where ad_id = %s  \n %s \n", dbEmpAd, QryEmployeeAd[0]["ad_id"], resGORM.Error))
			return mapEmpAd, err
		}
	}

	return mapEmpAd, resGORM.Error
}

func BatchUnableUser() (boo bool, err error) {
	var QryEmployee []map[string]interface{}
	jsonMap := make(map[string]interface{})
	resGORM := DBgorm.Table(dbEmp).Find(&QryEmployee,
		"Sts_id=1")
	if resGORM.Error != nil {
		return false, err
	}
	for _, rec := range QryEmployee {
		var jsonData []byte
		jsonData, err = json.Marshal(rec)
		if err != nil {
			return false, err
		}
		err = json.Unmarshal([]byte(jsonData), &jsonMap)

		tmpSubStr := subStrbyEmail(rec["email"].(string))
		sttEmpAd, err := GetEmpByLoginName(tmpSubStr)
		if err != nil {
			return false, err
		}
		if sttEmpAd.Employee_ID == "" && sttEmpAd.SAMAccountName == "" && sttEmpAd.Mail == "" { //if struct empty
			jsonMap["sts_id"] = "2"
			resGORM = DBgorm.Table(dbEmp).Where("usr_id = ?", jsonMap["usr_id"]).Updates(jsonMap)
		}
		clearstt(&sttEmpAd)
	}

	return false, err
}

func subStrbyEmail(email string) (restr string) {
	runes := []rune(email)
	restr = ""
	for _, char := range runes {
		// fmt.Println(char)
		if char == '@' {
			return restr
		}
		restr += string(char)
	}
	return restr
}

func syncEmpAd2Employee(mapEmpAd map[string]interface{}) (map[string]interface{}, error) {
	chkUpdate := md5.New()
	chkExist := md5.New()
	err := errors.New("")
	var QryEmployee = new(EmpAdToEmployee)
	resGORM := DBgorm.Table(dbEmp).Find(&QryEmployee, "Employee_ID = ? AND Employee_Name = ? ", mapEmpAd["employee_ID"], mapEmpAd["displayName"])
	// fmt.Printf("Employee_Name : %s resGORM.RowsAffected : %#v \n", mapEmpAd["displayName"], resGORM.RowsAffected)
	// fmt.Printf("------------------------------------------------- \n QryEmployee %#v \n", QryEmployee)

	if resGORM.RowsAffected == 0 {
		mapEmpAd["createdate"] = time.Now()
		mapEmp, err := mapJsonEmp(mapEmpAd)
		if err != nil {

			return mapEmpAd, err
		}
		fmt.Println(mapEmpAd["whenCreated"].(string) + " --> " + convDateTimeFixFormat(mapEmpAd["whenCreated"].(string), "y-m-d"))
		mapEmp.Email = mapEmpAd["mail"].(string)
		mapEmp.Onboard_Date = convDateTimeFixFormat(mapEmpAd["whenCreated"].(string), "y-m-d")
		mapEmp.Onboard_Month_Year = convDateTimeFixFormat(mapEmpAd["whenCreated"].(string), "m-y")
		mapEmp.Joining_Date = convDateTimeFixFormat(mapEmpAd["whenCreated"].(string), "y-m-d")
		mapEmp.Sts_id = "2"
		resGORM = DBgorm.Table(dbEmp).Create(mapEmp)
		if resGORM.Error != nil {
			err = errors.New(fmt.Sprintf("Insert %s from struct EmpAdToEmployee \n %s \n", dbEmp, resGORM.Error))
			return mapEmpAd, err
		}
		resMap, _ := mapJsonEmployee(mapEmp)
		return resMap, resGORM.Error
	}
	// fmt.Println("Found Duplicate, Validate Data for Update table Employee")
	mapEmp, err := mapJsonEmp(mapEmpAd)
	if err != nil {
		err = errors.New(fmt.Sprintf("syncEmpAd2Employee func mapJsonEmp(mapEmpAd) : \n error : Usr_id= %s , Employee_ID= %s , Employee_Name= %s \n %#v \n", mapEmp.Usr_id, mapEmp.Employee_ID, mapEmp.Employee_Name, mapEmp))
		return mapEmpAd, err
	}
	strUpdate := mapEmp.Business_Unit + mapEmp.Email + mapEmp.Employee_ID + mapEmp.Employee_Name + mapEmp.Organization + mapEmp.Onboard_Date + mapEmp.Onboard_Month_Year + mapEmp.Mobile + mapEmp.Joining_Date + mapEmp.Mobile
	strExist := QryEmployee.Business_Unit + QryEmployee.Email + QryEmployee.Employee_ID + QryEmployee.Employee_Name + QryEmployee.Organization + QryEmployee.Onboard_Date + QryEmployee.Onboard_Month_Year + QryEmployee.Mobile + QryEmployee.Joining_Date + QryEmployee.Mobile
	io.WriteString(chkUpdate, strUpdate)
	io.WriteString(chkExist, strExist)
	strMD5Update := fmt.Sprintf("%x", chkUpdate.Sum(nil))
	strMD5Exist := fmt.Sprintf("%x", chkExist.Sum(nil))
	// fmt.Printf("\n %s \n %s \n %s \n %s \n", strUpdate, strExist, strMD5Update, strMD5Exist)
	mapEmp.Createdate = QryEmployee.Createdate
	if strMD5Update != strMD5Exist {
		resGORM = DBgorm.Table(dbEmp).Where("usr_id = ?", string(QryEmployee.Usr_id)).Updates(mapEmp)
		if resGORM.Error != nil {
			// fmt.Printf("\n %#v \n", mapEmp)
			err = errors.New(fmt.Sprintf("Update %s from struct EmpAdToEmployee \n error : Usr_id= %s , Employee_ID= %s , Employee_Name= %s \n %#v \n", QryEmployee.Usr_id, QryEmployee.Employee_ID, QryEmployee.Employee_Name, mapEmp))
			return mapEmpAd, err
		}
	}
	return mapEmpAd, resGORM.Error
}

func mapJsonEmpAd(stt *EmployeeAd) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	var jsonData []byte
	jsonData, err := json.Marshal(stt)
	if err != nil {
		return jsonMap, err
	}
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	return jsonMap, err
}
func mapJsonEmployee(stt *EmpAdToEmployee) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	var jsonData []byte
	jsonData, err := json.Marshal(stt)
	if err != nil {
		return jsonMap, err
	}
	err = json.Unmarshal([]byte(jsonData), &jsonMap)
	return jsonMap, err
}

func mapJsonEmp(mapEmpAd map[string]interface{}) (sttEmpAdToEmp *EmpAdToEmployee, err error) {
	// fmt.Printf("\n mapJsonEmp Map: \n %#v \n", mapEmpAd)
	var jsonData []byte
	jsonData, err = json.Marshal(mapEmpAd)
	if err != nil {
		return sttEmpAdToEmp, err
	}
	// fmt.Printf("\n mapJsonEmp jsonData: \n %s \n", jsonData)
	err = json.Unmarshal([]byte(jsonData), &sttEmpAdToEmp)
	// fmt.Printf("\n mapJsonEmp sttEmpAdToEmp: \n %#v \n", sttEmpAdToEmp)
	sttEmpAdToEmp.Email = mapEmpAd["mail"].(string)
	sttEmpAdToEmp.Onboard_Date = convDateTimeFixFormat(mapEmpAd["whenCreated"].(string), "y-m-d")
	sttEmpAdToEmp.Onboard_Month_Year = convDateTimeFixFormat(mapEmpAd["whenCreated"].(string), "m-y")
	sttEmpAdToEmp.Joining_Date = convDateTimeFixFormat(mapEmpAd["whenCreated"].(string), "y-m-d")
	return sttEmpAdToEmp, err
}

func findTimeFormat(s string) string {
	val := strings.ToLower(s)
	if strings.Index(val, "z") >= 0 {
		return "z"
	}
	if strings.Index(val, "-") >= 0 {
		return "-"
	}
	return "0"
}

func reFormatDate(format, y, m, d, h, i, s string) string {
	switch format {
	case "y-m-d h:i:s":
		return y + "-" + m + "-" + d + " " + h + ":" + i + ":" + s
	case "y-m-d":
		return y + "-" + m + "-" + d
	case "ymdhis":
		return y + m + d + h + i + s
	case "ymd":
		return y + m + d
	case "m-y":
		return m + "-" + y
	default:
		return y + m + d + h + i + s
	}
}

func convDateTimeFixFormat(s string, f string) string {
	value := s
	switch findTimeFormat(s) {
	case "z":
		//00000000011111111
		//12345678901234567
		//yyyymmddhhiiss.sZ
		runes := []rune(value)
		if len(runes) == 0 {
			return ""
		}
		yyyy := runes[0:4]
		mm := runes[4:6]
		dd := runes[6:8]
		hh := runes[8:10]
		ii := runes[10:12]
		ss := runes[12:14]
		cb := reFormatDate(f, string(yyyy), string(mm), string(dd), string(hh), string(ii), string(ss))
		return cb
	case "-":
		//0000000001111111111
		//1234-67-90 23:56:89
		//yyyy-mm-dd hh:ii:ss
		runes := []rune(value)
		if len(runes) == 0 {
			return ""
		}
		yyyy := runes[0:4]
		mm := runes[5:7]
		dd := runes[8:10]
		hh := runes[11:13]
		ii := runes[14:16]
		ss := runes[17:19]
		cb := reFormatDate(f, string(yyyy), string(mm), string(dd), string(hh), string(ii), string(ss))
		return cb
	}

	return ""
}

func convDateTime(s string) (i int) {
	i = 0
	value := s
	switch findTimeFormat(s) {
	case "z":
		runes := []rune(value)
		if len(runes) == 0 {
			return 0
		}
		yyyy := runes[0:4]
		mm := runes[4:6]
		dd := runes[6:8]
		hh := runes[8:10]
		ii := runes[10:12]
		ss := runes[12:14]
		cb := string(yyyy) + string(mm) + string(dd) + string(hh) + string(ii) + string(ss)
		i, err := strconv.Atoi(cb)
		if err != nil {
			fmt.Println(err)
		}
		return i
	case "-":
		runes := []rune(value)
		if len(runes) == 0 {
			return 0
		}
		yyyy := runes[0:4]
		mm := runes[5:7]
		dd := runes[8:10]
		cb := string(yyyy) + string(mm) + string(dd) + "00" + "00" + "00"
		i, err := strconv.Atoi(cb)
		if err != nil {
			fmt.Println(err)
		}
		return i
	}

	return i
}

func clearstt(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func RunAdSync(timeFilter int) (boo bool, err error) {
	// chars := []string{"a*", "b*", "c*", "d*", "e*", "f*", "g*", "h*", "i*", "j*", "k*", "l*", "m*", "n*", "o*", "p*", "q*", "r*", "s*", "t*", "u*", "v*", "w*", "x*", "y*", "z*", "0*", "1*", "2*", "3*", "4*", "5*", "6*", "7*", "8*", "9*"}
	// chars := []string{"0*", "1*", "2*", "3*", "4*", "5*", "6*", "7*", "8*", "9*"}
	chars := []string{"somchai*"}
	// fmt.Printf("--> loop count : %v \n", ii)
	t := time.Now()
	td := t.Add(time.Duration(-timeFilter) * time.Minute)
	timeThen := fmt.Sprintf("%d%02d%02d%02d%02d%02d.0Z",
		td.Year(), td.Month(), td.Day(),
		td.Hour(), td.Minute(), td.Second())
	for _, char := range chars {
		// fmt.Printf("Start Run SyncProcessFilterAD(\"%s\" \n", char)
		param := "(sAMAccountName=" + char + ")(whenChanged>=" + timeThen + ")"
		if timeFilter < 0 {
			param = "(sAMAccountName=" + char + ")"
		}

		fmt.Println(param)
		boo, err = FilterAdSync(param)
		if err != nil {
			txterr := fmt.Sprintf("%s", err.(net.Error))
			err = errors.New(txterr)
			if strings.Index(txterr, "No connection") >= 0 {
				//เคส login DB server ไม่ได้
				return false, err
			}
		}
	}
	return true, err
}

func FilterAdSync(char string) (boo bool, err error) {
	boo, err = SyncProcessFilterAD(char)
	// fmt.Printf("\n boo : %#v \n", boo)
	// fmt.Printf(" err : %#v \n\n", err)
	return
}
