package mariadb

type Server struct {
	Server_no      int    `json:"serverNo,omitempty"`
	Server_name    string `json:"serverName,omitempty"`
	Group_name     string `json:"groupName,omitempty"`
	Group_id       int    `json:"groupID,omitempty"`
	Bmc_ip_addr    string `json:"bmcIpAddr,omitempty"`
	Bmc_ip_source  string `json:"bmcIpSource,omitempty"`
	Bmc_mac        string `json:"bmcMac,omitempty"`
	Bmc_version    string `json:"bmcVersion,omitempty"`
	Bmc_username   string `json:"bmcUsername,omitempty"`
	Bmc_password   string `json:"bmcPassword,omitempty"`
	Product_name   string `json:"productName,omitempty"`
	Board_part_num string `json:"boardPartNum,omitempty"`
	Smbios_uuid    string `json:"smbiosUuid,omitempty"`
	Bios_vendor    string `json:"biosVendor,omitempty"`
	Bios_version   string `json:"biosVersion,omitempty"`
	Os_type        string `json:"osType,omitempty"`
	Os_ip_addr     string `json:"osIpAddr,omitempty"`
	Os_version     string `json:"osVersion,omitempty"`
	Os_hostname    string `json:"osHostname,omitempty"`
	Os_arch        string `json:"osArch,omitempty"`
	Os_kernel      string `json:"osKernel,omitempty"`
	Status         string `json:"status,omitempty"`
	Start_time     string `json:"startTime,omitempty"`
	Create_date    string `json:"createDate,omitempty"`
	Critical_count int    `json:"criticalCount,omitempty"`
	Warning_count  int    `json:"warningCount,omitempty"`
	Alarm_status   int    `json:"alarmStatus,omitempty"`
}
