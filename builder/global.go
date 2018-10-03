package builder

// The non-derterministic ordering of map is kinda annoying. But it is not critical
// to the overall logic, avoiding worrying about it too much
var NameToBlock = map[string]string{}
