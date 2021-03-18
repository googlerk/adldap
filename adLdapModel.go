package adldap

type ReqAD struct {
	ReqUser string `json:"reqUser"`
	ReqPass string `json:"reqPass"`
	ReqDN   string `json:"reqDN"`
}

type Employee struct {
	Usr_id                 uint   `gorm:"column:usr_id; type:int(11); NOT NULL; AUTO_INCREMENT;" json:"usr_id"`
	AppGroupID             string `gorm:"column:AppGroupID; type:varchar(255); DEFAULT NULL;" json:"AppGroupID"`
	App_ID_Ref             string `gorm:"column:App_ID_Ref; type:text DEFAULT NULL;" json:"App_ID_Ref"`
	Benefit                string `gorm:"column:Benefit; type:varchar(255); DEFAULT NULL;" json:"Benefit"`
	Birth_Date             string `gorm:"column:Birth_Date; type:varchar(10); DEFAULT NULL;" json:"Birth_Date"`
	Business_Unit          string `gorm:"column:Business_Unit; type:varchar(200); DEFAULT NULL;" json:"Business_Unit"`
	Client_ip              string `gorm:"column:client_ip; type:varchar(25); DEFAULT NULL;" json:"client_ip"`
	Company_Contractor     string `gorm:"column:Company_Contractor; type:varchar(255); DEFAULT NULL;" json:"Company_Contractor"`
	Company                string `gorm:"column:Company; type:varchar(255); DEFAULT NULL;" json:"Company"`
	Contract_End_Date      string `gorm:"column:Contract_End_Date; type:varchar(255); DEFAULT NULL;" json:"Contract_End_Date"`
	Createdate             string `gorm:"column:createdate; type:datetime DEFAULT NULL;" json:"createdate"`
	DM_Organize            string `gorm:"column:DM_Organize; type:varchar(255); DEFAULT NULL;" json:"DM_Organize"`
	DM_Project             string `gorm:"column:DM_Project; type:varchar(255); DEFAULT NULL;" json:"DM_Project"`
	DS_Organize            string `gorm:"column:DS_Organize; type:varchar(255); DEFAULT NULL;" json:"DS_Organize"`
	DS_Project             string `gorm:"column:DS_Project; type:varchar(255); DEFAULT NULL;" json:"DS_Project"`
	Effective_date         string `gorm:"column:Effective_date; type:varchar(255); DEFAULT NULL;" json:"Effective_date"`
	Email                  string `gorm:"column:email; type:varchar(255); DEFAULT NULL;" json:"email"`
	Employee_ID            string `gorm:"column:Employee_ID; type:varchar(15); NOT NULL;" json:"Employee_ID"`
	Employee_Name          string `gorm:"column:Employee_Name; type:varchar(255); DEFAULT NULL;" json:"Employee_Name"`
	Employee_status        string `gorm:"column:Employee_status; type:varchar(255); DEFAULT NULL;" json:"Employee_status"`
	Exit_Interview_Date    string `gorm:"column:Exit_Interview_Date; type:varchar(255); DEFAULT NULL;" json:"Exit_Interview_Date"`
	Gender                 string `gorm:"column:gender; type:varchar(20); DEFAULT NULL;" json:"gender"`
	Generation             string `gorm:"column:Generation; type:varchar(100); DEFAULT NULL;" json:"Generation"`
	HR_Department_Division string `gorm:"column:HR_Department_Division; type:varchar(255); DEFAULT NULL;" json:"HR_Department_Division"`
	HR_Layer               string `gorm:"column:HR_Layer; type:varchar(255); DEFAULT NULL;" json:"HR_Layer"`
	HR_Position            string `gorm:"column:HR_Position; type:varchar(255); DEFAULT NULL;" json:"HR_Position"`
	Joining_Date           string `gorm:"column:Joining_Date; type:varchar(10); DEFAULT NULL;" json:"Joining_Date"`
	Last_Working_Date      string `gorm:"column:Last_Working_Date; type:varchar(255); DEFAULT NULL;" json:"Last_Working_Date"`
	Line_id                string `gorm:"column:line_id; type:varchar(222); DEFAULT NULL;" json:"line_id"`
	Local_Foreigner        string `gorm:"column:Local_Foreigner; type:varchar(255); DEFAULT NULL;" json:"Local_Foreigner"`
	MD                     string `gorm:"column:MD; type:varchar(255); DEFAULT NULL;" json:"MD"`
	Mobile                 string `gorm:"column:Mobile; type:varchar(255); DEFAULT NULL;" json:"Mobile"`
	Name_Thai              string `gorm:"column:Name_Thai; type:varchar(255); DEFAULT NULL;" json:"Name_Thai"`
	Nickname               string `gorm:"column:Nickname; type:varchar(255); DEFAULT NULL;" json:"Nickname"`
	Onboard_Date           string `gorm:"column:Onboard_Date; type:varchar(255); DEFAULT NULL;" json:"Onboard_Date"`
	Onboard_Month_Year     string `gorm:"column:Onboard_Month_Year; type:varchar(255); DEFAULT NULL;" json:"Onboard_Month_Year"`
	On_Off_Board           string `gorm:"column:On_Off_Board; type:varchar(255); DEFAULT NULL;" json:"On_Off_Board"`
	Organization           string `gorm:"column:Organization; type:varchar(100); DEFAULT NULL;" json:"Organization"`
	Password               string `gorm:"column:password; type:varchar(255); DEFAULT 'password';" json:"password"`
	People_Manager         string `gorm:"column:People_Manager; type:varchar(255); DEFAULT NULL;" json:"People_Manager"`
	Pictute                string `gorm:"column:pictute; type:varchar(1024); DEFAULT NULL;" json:"pictute"`
	Position               string `gorm:"column:Position; type:varchar(255); DEFAULT NULL;" json:"Position"`
	Profile                string `gorm:"column:profile; type:text DEFAULT NULL;" json:"profile"`
	Project_Role           string `gorm:"column:Project_Role; type:varchar(200); DEFAULT NULL;" json:"Project_Role"`
	Remarks                string `gorm:"column:Remarks; type:text DEFAULT NULL;" json:"Remarks"`
	Responsible_Party      string `gorm:"column:Responsible_Party; type:text DEFAULT NULL;" json:"Responsible_Party"`
	Sts_id                 string `gorm:"column:sts_id; type:int(11); DEFAULT 1;" json:"sts_id"`
	TerminateDate          string `gorm:"column:TerminateDate; type:varchar(100); DEFAULT NULL;" json:"TerminateDate"`
	Updatedate             string `gorm:"column:updatedate; type:datetime DEFAULT current_timestamp(); ON UPDATE current_timestamp();" json:"updatedate"`
	Userlevelid            string `gorm:"column:userlevelid; type:int(11); DEFAULT NULL;" json:"userlevelid"`
	Usr_id_ref             string `gorm:"column:usr_id_ref; type:varchar(32); DEFAULT NULL;" json:"usr_id_ref"`
	Workforce_Type         string `gorm:"column:Workforce_Type; type:varchar(255); DEFAULT NULL;" json:"Workforce_Type"`
}
type EmployeeAd struct {
	Ad_id                         uint   `grom:"primaryKey;autoIncrement;column:ad_id;" json:"ad_id"`
	AccountExpires                string `grom:"column:accountExpires" json:"accountExpires"`
	Admincount                    string `grom:"column:admincount" json:"admincount"`
	BadPasswordTime               string `grom:"column:badPasswordTime" json:"badPasswordTime"`
	BadPwdCount                   string `grom:"column:badPwdCount" json:"badPwdCount"`
	Cn                            string `grom:"column:cn" json:"cn"`
	CodePage                      string `grom:"column:codePage" json:"codePage"`
	Company                       string `grom:"column:company" json:"company"`
	CountryCode                   string `grom:"column:countryCode" json:"countryCode"`
	Createdate                    string `grom:"column:createdate" json:"createdate"`
	Department                    string `grom:"column:department" json:"department"`
	Description                   string `grom:"column:description" json:"description"`
	Directreports                 string `grom:"column:directreports" json:"directreports"`
	DisplayName                   string `grom:"column:displayName" json:"displayName"`
	DistinguishedName             string `grom:"column:distinguishedName" json:"distinguishedName"`
	Dn                            string `grom:"column:dn" json:"dn"`
	DSCorePropagationData         string `grom:"column:dSCorePropagationData" json:"dSCorePropagationData"`
	Employee_ID                   string `grom:"column:employee_ID" json:"employee_ID"`
	Employeeid                    string `grom:"column:employeeid" json:"employeeid"`
	Employeenumber                string `grom:"column:employeenumber" json:"employeenumber"`
	Gidnumber                     string `grom:"column:gidnumber" json:"gidnumber"`
	GivenName                     string `grom:"column:givenName" json:"givenName"`
	Homedirectory                 string `grom:"column:homedirectory" json:"homedirectory"`
	Homedrive                     string `grom:"column:homedrive" json:"homedrive"`
	Homephone                     string `grom:"column:homephone" json:"homephone"`
	Info                          string `grom:"column:info" json:"info"`
	InstanceType                  string `grom:"column:instanceType" json:"instanceType"`
	Ipphone                       string `grom:"column:ipphone" json:"ipphone"`
	Iscriticalsystemobject        string `grom:"column:iscriticalsystemobject" json:"iscriticalsystemobject"`
	Lastlogoff                    string `grom:"column:lastlogoff" json:"lastlogoff"`
	LastLogon                     string `grom:"column:lastLogon" json:"lastLogon"`
	LastLogonTimestamp            string `grom:"column:lastLogonTimestamp" json:"lastLogonTimestamp"`
	Lockouttime                   string `grom:"column:lockouttime" json:"lockouttime"`
	Loginshell                    string `grom:"column:loginshell" json:"loginshell"`
	LogonCount                    string `grom:"column:logonCount" json:"logonCount"`
	Logonhours                    string `grom:"column:logonhours" json:"logonhours"`
	Mail                          string `grom:"column:mail" json:"mail"`
	Mailnickname                  string `grom:"column:mailnickname" json:"mailnickname"`
	Manager                       string `grom:"column:manager" json:"manager"`
	Memberof                      string `grom:"column:memberof" json:"memberof"`
	Mobile                        string `grom:"column:mobile" json:"mobile"`
	Msdsconsistencyguid           string `grom:"column:ms-ds-consistencyguid" json:"ms-ds-consistencyguid"`
	Msdssupportedencryptiontypes  string `grom:"column:msds-supportedencryptiontypes" json:"msds-supportedencryptiontypes"`
	Msnpallowdialin               string `grom:"column:msnpallowdialin" json:"msnpallowdialin"`
	Msrtcsipdeploymentlocator     string `grom:"column:msrtcsip-deploymentlocator" json:"msrtcsip-deploymentlocator"`
	Msrtcsipfederationenabled     string `grom:"column:msrtcsip-federationenabled" json:"msrtcsip-federationenabled"`
	Msrtcsipinternetaccessenabled string `grom:"column:msrtcsip-internetaccessenabled" json:"msrtcsip-internetaccessenabled"`
	Msrtcsipoptionflags           string `grom:"column:msrtcsip-optionflags" json:"msrtcsip-optionflags"`
	Msrtcsipprimaryhomeserver     string `grom:"column:msrtcsip-primaryhomeserver" json:"msrtcsip-primaryhomeserver"`
	Msrtcsipprimaryuseraddress    string `grom:"column:msrtcsip-primaryuseraddress" json:"msrtcsip-primaryuseraddress"`
	Msrtcsipuserenabled           string `grom:"column:msrtcsip-userenabled" json:"msrtcsip-userenabled"`
	Mssfu30name                   string `grom:"column:mssfu30name" json:"mssfu30name"`
	Mssfu30nisdomain              string `grom:"column:mssfu30nisdomain" json:"mssfu30nisdomain"`
	Name                          string `grom:"column:name" json:"name"`
	ObjectCategory                string `grom:"column:objectCategory" json:"objectCategory"`
	Objectclass                   string `grom:"column:objectclass" json:"objectclass"`
	ObjectGUID                    string `grom:"column:objectGUID" json:"objectGUID"`
	ObjectSid                     string `grom:"column:objectSid" json:"objectSid"`
	Physicaldeliveryofficename    string `grom:"column:physicaldeliveryofficename" json:"physicaldeliveryofficename"`
	PrimaryGroupID                string `grom:"column:primaryGroupID" json:"primaryGroupID"`
	Protocolsettings              string `grom:"column:protocolsettings" json:"protocolsettings"`
	Proxyaddresses                string `grom:"column:proxyaddresses" json:"proxyaddresses"`
	PwdLastSet                    string `grom:"column:pwdLastSet" json:"pwdLastSet"`
	SAMAccountName                string `grom:"column:sAMAccountName" json:"sAMAccountName"`
	SAMAccountType                string `grom:"column:sAMAccountType" json:"sAMAccountType"`
	Scriptpath                    string `grom:"column:scriptpath" json:"scriptpath"`
	Serviceprincipalname          string `grom:"column:serviceprincipalname" json:"serviceprincipalname"`
	Showinadvancedviewonly        string `grom:"column:showinadvancedviewonly" json:"showinadvancedviewonly"`
	Sn                            string `grom:"column:sn" json:"sn"`
	Targetaddress                 string `grom:"column:targetaddress" json:"targetaddress"`
	Telephonenumber               string `grom:"column:telephonenumber" json:"telephonenumber"`
	Terminalserver                string `grom:"column:terminalserver" json:"terminalserver"`
	Title                         string `grom:"column:title" json:"title"`
	Uid                           string `grom:"column:uid" json:"uid"`
	Uidnumber                     string `grom:"column:uidnumber" json:"uidnumber"`
	Unixhomedirectory             string `grom:"column:unixhomedirectory" json:"unixhomedirectory"`
	Updatedate                    string `grom:"column:updatedate" json:"updatedate"`
	UserAccountControl            string `grom:"column:userAccountControl" json:"userAccountControl"`
	Userparameters                string `grom:"column:userparameters" json:"userparameters"`
	UserPrincipalName             string `grom:"column:userPrincipalName" json:"userPrincipalName"`
	USNChanged                    string `grom:"column:uSNChanged" json:"uSNChanged"`
	USNCreated                    string `grom:"column:uSNCreated" json:"uSNCreated"`
	WhenChanged                   string `grom:"column:whenChanged" json:"whenChanged"`
	WhenCreated                   string `grom:"column:whenCreated" json:"whenCreated"`
}
