package main

var SCRIPTS = map[string]string{
	"balance assets":     "b assets -X $",
	"register":           "r --tail 10",
	"balance this month": "b -b \"this month\"",
}

const QUERIES_FILE = "queries.txt"
const HTPASSWD_FILE = ".htpasswd"
const DEFAULT_JOURNAL = "ledger.txt"
const ARCHETYPES_DIR = "archetypes"
