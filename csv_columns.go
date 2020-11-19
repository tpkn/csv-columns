// CSV Tools, https://tpkn.me
package csvtools

import (
	"strings"
)

type results struct {
	name []string
	index []int
	mapped map[string]int
}

func FindColumns(columns_list []string, header_row []string, clean bool) ([]string, []int, map[string]int) {
	result := results{}
	result.mapped = make(map[string]int)
	
	if len(columns_list) == 0 {
		columns_list = header_row[0:]
	}
	
	for i := 0; i < len(columns_list); i++ {
		found_one := false
		target_name := strings.TrimSpace(columns_list[i])
		
		for index, row_column_name := range header_row {
			if target_name == row_column_name {
				result.index = append(result.index, index)
				result.mapped[target_name] = index
				
				found_one = true
				break
			}
		}
		
		if (found_one && clean) || !clean {
			result.name = append(result.name, target_name)
		}
		
		// Non-existent columns index = -1 
		if !found_one && !clean {
			result.index = append(result.index, -1)
			result.mapped[target_name] = -1
		}
	}
	
	return result.name, result.index, result.mapped
}

func IndexByName(key string, mapped_list map[string]int) int {
	if _, ok := mapped_list[key]; ok {
		return mapped_list[key]
	}else{
		return -1
	}
}
