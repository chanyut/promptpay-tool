package dataobject

import (
	"fmt"
)

func Serialize(o DataObject) string {
	return fmt.Sprintf("%s%02d%s", o.ID(), o.Length(), o.Value())
}
