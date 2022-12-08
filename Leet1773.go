func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    matches := 0
    for i := 0; i < len(items); i++ {
        item := items[i]
 
        if ruleKey == "type" {
            if ruleValue == item[0] {
                matches = matches + 1
            }
        } else if ruleKey == "color" {
            if ruleValue == item[1] {
                matches = matches + 1
            }
        } else if ruleKey == "name" {
            if ruleValue == item[2] {
                matches = matches + 1
            }
        }
    }
    return matches
}
