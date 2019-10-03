package tablestore

import (
	//"fmt"
	//"strings"
	//"github.com/darshannevgi/fleetdb/kv_store"
	//"github.com/darshannevgi/fleetdb/tablestore"
)
/* 
Each values[0] represents byte[] data value of first column
if first column is country_code then  values[0] is byte representation of 'US' = byte[]{85,83}
*/
func TranslateToKV(columnSpecs []FleetDbColumnSpec, values [][]byte) []KVItem{
    var pKey []byte
    var cKey []byte
    for i, colSpec := range columnSpecs {
    	if colSpec.isPartition{
    		pKey = append(pKey, values[i]...)
    	}
    	if colSpec.isClustering{
    		cKey = append(cKey, values[i]...)
    	}
   }
    //Handle composite key and only pK and cK case
    //In case of composite key , Key is made in order of partition key specified in create table command
    //if PRIMARY KEY ((country_code, state_province, city) then key will be  US/NY/Buffalo
    
    kvItems := make([]KVItem, len(values))
    index := 0
    var prefix []byte
    prefix = pKey
    prefix = append(prefix,"/"...)    
    if len(cKey) > 0{
	    prefix = append(prefix,cKey...)
	    prefix = append(prefix,"/"...)
    }
    
    for i, colSpec := range columnSpecs {
    	if !colSpec.isPartition && !colSpec.isClustering{
    		key := append(prefix,colSpec.colname...)
    		val := values[i]
    		//fmt.Println("Key is =" + string(key))
    		//fmt.Println("Value is =" + string(val))
    		kvItems[index] =  KVItem{key, val}
    		index = index + 1
    	}
   }
	return kvItems;
}
