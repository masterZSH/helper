# helper


```go
import "github.com/masterZSH/helper"
//  NilSliceToEmptySlice
helper.NilSliceToEmptySlice(yournilslice)

// slice to str  faster than string([]byte{})
helper.SliceByteToString([]byte[1,2,3])

//  str to slice  better that []byte("123")
helper.StringToSliceByte("123")


// is email 
helper.IsEmail("123@gmail.com")

// GetPhoneCipherText "13112312312" => "131****1234"
helper.GetPhoneCipherText("13112312312")

```
