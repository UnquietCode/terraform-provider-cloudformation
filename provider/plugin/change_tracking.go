package plugin


type ChangeType = string

// TODO namespace these
const (
  Unchanged ChangeType = "unchanged"
  Changed ChangeType = "changed"
  Maybe ChangeType = "maybe"
  // Added ChangeType = "added"
  Updated ChangeType = "updated"
  Deleted ChangeType = "deleted"
)


func removeFromChangesArray(item interface{}, removeArray interface{}) []interface{} {
  var newRemoved []interface{} = []interface{}{}
  
  for _, v := range removeArray.([]interface{}) {
    if v != item {
      newRemoved = append(newRemoved, v)
    }
  }
  
  return newRemoved
}

func removeFromStringArray(item string, removeArray []string) []string {
  var newRemoved []string = []string{}
  
  for _, v := range removeArray {
    if v != item {
      newRemoved = append(newRemoved, v)
    }
  }
  
  return newRemoved
}