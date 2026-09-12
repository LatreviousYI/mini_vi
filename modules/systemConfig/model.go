package systemConfig

type SystemConfigModel struct {
	DeviceId            string `json:"device_id" toml:"device_id"`                       // 设备id
	CollectionFrequency string `json:"collection_frequency" toml:"collection_frequency"` // 采集频率
	ImgNameRules        string `json:"img_name_rules" toml:"img_name_rules"`             // 命名规则
	ResolutionRatio     string `json:"resolution_ratio" toml:"resolution_ratio"`         // 分辨率
	MachineId           string `json:"machine_id" toml:"machine_id"`                     // 机器id
	Ip                  string `json:"ip" toml:"-"`                                      // ip
}
