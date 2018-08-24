package superType

type Maper interface {
	Merge(arr map[string]interface{}) interface{}
	Search(arr map[string]interface{}) string
	//Column(arr map[string]interface{}) []interface{}	// 获取
	Keys(arr map[string]interface{}) []string
	Values(arr map[string]interface{}) []interface{}

	KeyExists(arr string) bool

	Map(arr func())
	Reuce(arr func())
}
type Map map[string]interface{}
type MapSlice []map[string]interface{}

func (m Map) KeyExists(field string) bool {
	dict := (map[string]interface{})(m)
	if _,ok := dict[field];ok {
		return true
	}
	return false
}

func (m Map) HasKey(field string) bool {
	return m.KeyExists(field)
}

func (ms MapSlice) Column(field string) []interface{} {
	var tmp []interface{}
	arr := ([]map[string]interface{})(ms)
	if len(arr)>0 {
		for _,item := range arr{
			if Map(item).KeyExists(field) {
				tmp = append(tmp, item[field])
			}
		}
	}
	return tmp
}

func (m Map) Keys() []string {
	var tmp []string
	arr := (map[string]interface{})(m)
	if len(arr)>0 {
		for item,_ := range arr{
			tmp = append(tmp, item)
		}
	}
	return tmp
}

func (m Map) Values() []interface{} {
	var tmp []interface{}
	arr := (map[string]interface{})(m)
	if len(arr)>0 {
		for _,item := range arr{
			tmp = append(tmp, item)
		}
	}
	return tmp
}

