package jupiter

// application
type ApplicationData struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	AData     AData  `json:"data"`
}

type AData struct {
	Request ApplicationRequest `json:"request"`
	Servers []Server           `json:"servers"`
}

type Server struct {
	ID              int    `json:"id"`
	ServerNo        string `json:"server_no"`
	Creator         string `json:"creator"`
	VMUUid          string `json:"vm_uuid"`
	Status          string `json:"status"`
	VMName          string `json:"vm_name"`
	ProjectGlobalID string `json:"project_global_id"`
	CloudName       string `json:"cloud_name"`
}

type ApplicationRequest struct {
	ProjectGlobalID string `json:"project_global_id"`
	CloudName       string `json:"cloud_name"`
	FlavorID        string `json:"flavor_id"`
	ImageID         string `json:"image_id"`
	VPCID           string `json:"vpc_id"`
	NetworkID       string `json:"network_id"`
}

// volume
type VolumeData struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	VData     VData  `json:"data"`
}

type VData struct {
	Detail VDetail `json:"Detail"`
}

type VDetail struct {
	Name            string `json:"name"`
	ProjectGlobalID string `json:"project_global_id"`
	CloudName       string `json:"cloud_name"`
	Description     string `json:"description"`
	Size            int    `json:"size"`
	UserId          string `json:"user_id"`
	VolumeFeature   string `json:"volume_feature"`
	VolumeType      string `json:"volume_type"`
	Vmuuid          string `json:"vm_uuid"`
}

// vpc
type VPCCreate struct {
	OrderMsg OrderMsg `json:"order_msg"`
}

type OrderMsg struct {
	RequestId string `json:"request_id"`
}

type VPCApplicationData struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Success   bool    `json:"success"`
	Timestamp int     `json:"timestamp"`
	VPCData   VPCData `json:"data"`
}

type VPCData struct {
	VPCform VPCform `json:"form"`
}

type VPCform struct {
	Vpctype      string       `json:"type"`
	Owner        []string     `json:"owner"`
	VpcName      string       `json:"vpc_name"`
	Description  string       `json:"description"`
	PlatformInfo PlatformInfo `json:"platform_info"`
}

type PlatformInfo struct {
	ProjectUUID  string `json:"project_uuid"`
	PlatformName string `json:"platform_name"`
}

//subnet
type SubnetApplicationData struct {
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	Success    bool       `json:"success"`
	Timestamp  int        `json:"timestamp"`
	SubnetData SubnetData `json:"data"`
}

type SubnetData struct {
	Subnetform Subnetform `json:"form"`
}

type Subnetform struct {
	Subnettype   string       `json:"type"`
	Owner        []string     `json:"owner"`
	VpcName      string       `json:"vpc_name"`
	Description  string       `json:"description"`
	PlatformInfo PlatformInfo `json:"platform_info"`
}

//outnet
type VPCApplicationData struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Success   bool    `json:"success"`
	Timestamp int     `json:"timestamp"`
	VPCData   VPCData `json:"data"`
}

type VPCData struct {
	VPCform VPCform `json:"form"`
}

type VPCform struct {
	Vpctype      string       `json:"type"`
	Owner        []string     `json:"owner"`
	VpcName      string       `json:"vpc_name"`
	Description  string       `json:"description"`
	PlatformInfo PlatformInfo `json:"platform_info"`
}

//dns
type VPCApplicationData struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Success   bool    `json:"success"`
	Timestamp int     `json:"timestamp"`
	VPCData   VPCData `json:"data"`
}

type VPCData struct {
	VPCform VPCform `json:"form"`
}

type VPCform struct {
	Vpctype      string       `json:"type"`
	Owner        []string     `json:"owner"`
	VpcName      string       `json:"vpc_name"`
	Description  string       `json:"description"`
	PlatformInfo PlatformInfo `json:"platform_info"`
}

//vpn
type VPCApplicationData struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Success   bool    `json:"success"`
	Timestamp int     `json:"timestamp"`
	VPCData   VPCData `json:"data"`
}

type VPCData struct {
	VPCform VPCform `json:"form"`
}

type VPCform struct {
	Vpctype      string       `json:"type"`
	Owner        []string     `json:"owner"`
	VpcName      string       `json:"vpc_name"`
	Description  string       `json:"description"`
	PlatformInfo PlatformInfo `json:"platform_info"`
}
