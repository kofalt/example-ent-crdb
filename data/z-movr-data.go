package data

/*
	Selected data from "cockroach demo" do demonstrate bulk insert
*/

var MovrUsers = []map[string]string{
	{
		"id":          "ae147ae1-47ae-4800-8000-000000000022",
		"city":        "amsterdam",
		"name":        "Tyler Dalton",
		"address":     "88194 Angela Gardens Suite 94",
		"credit_card": "4443538758",
	},
	{
		"id":          "bd70a3d7-0a3d-4000-8000-000000000025",
		"city":        "amsterdam",
		"name":        "David Stanton",
		"address":     "80015 Mark Views Suite 96",
		"credit_card": "3471210499",
	},
	{
		"id":          "c28f5c28-f5c2-4000-8000-000000000026",
		"city":        "amsterdam",
		"name":        "Maria Weber",
		"address":     "14729 Karen Radial",
		"credit_card": "5844236997",
	},
	{
		"id":          "2e147ae1-47ae-4400-8000-000000000009",
		"city":        "boston",
		"name":        "Cindy Medina",
		"address":     "31118 Allen Gateway Apt. 60",
		"credit_card": "6464362441",
	},
}

var MovrVehicles = []map[string]string{
	{
		"id":               "aaaaaaaa-aaaa-4800-8000-00000000000a",
		"city":             "amsterdam",
		"type":             "scooter",
		"owner_id":         "c28f5c28-f5c2-4000-8000-000000000026",
		"creation_time":    "2019-01-02T09:04:05.000Z",
		"status":           "in_use",
		"current_location": "62609 Stephanie Route",
		"ext":              "{\"color\": \"red\"}",
	},
	{
		"id":               "bbbbbbbb-bbbb-4800-8000-00000000000b",
		"city":             "amsterdam",
		"type":             "scooter",
		"owner_id":         "bd70a3d7-0a3d-4000-8000-000000000025",
		"creation_time":    "2019-01-02T09:04:05.000Z",
		"status":           "available",
		"current_location": "57637 Mitchell Shoals Suite 59",
		"ext":              "{\"color\": \"blue\"}",
	},
	{
		"id":               "22222222-2222-4200-8000-000000000002",
		"city":             "boston",
		"type":             "scooter",
		"owner_id":         "2e147ae1-47ae-4400-8000-000000000009",
		"creation_time":    "2019-01-02T09:04:05.000Z",
		"status":           "in_use",
		"current_location": "19659 Christina Ville",
		"ext":              "{\"color\": \"blue\"}",
	},
}

var MovrRides = []map[string]string{
	{
		"id":            "ab020c49-ba5e-4800-8000-00000000014e",
		"city":          "amsterdam",
		"vehicle_city":  "amsterdam",
		"rider_id":      "c28f5c28-f5c2-4000-8000-000000000026",
		"vehicle_id":    "aaaaaaaa-aaaa-4800-8000-00000000000a",
		"start_address": "1905 Christopher Locks Apt. 77",
		"end_address":   "66037 Belinda Plaza Apt. 93",
		"start_time":    "2018-12-13T09:04:05.000Z",
		"end_time":      "2018-12-14T14:04:05.000Z",
		"revenue":       "77.00",
	},
	{
		"id":            "ac083126-e978-4800-8000-000000000150",
		"city":          "amsterdam",
		"vehicle_city":  "amsterdam",
		"rider_id":      "c28f5c28-f5c2-4000-8000-000000000026",
		"vehicle_id":    "aaaaaaaa-aaaa-4800-8000-00000000000a",
		"start_address": "50217 Victoria Fields Apt. 44",
		"end_address":   "56217 Wilson Spring",
		"start_time":    "2018-12-07T09:04:05.000Z",
		"end_time":      "2018-12-07T16:04:05.000Z",
		"revenue":       "9.00",
	},
	{
		"id":            "ac8b4395-8106-4800-8000-000000000151",
		"city":          "amsterdam",
		"vehicle_city":  "amsterdam",
		"rider_id":      "ae147ae1-47ae-4800-8000-000000000022",
		"vehicle_id":    "bbbbbbbb-bbbb-4800-8000-00000000000b",
		"start_address": "34704 Stewart Ports Suite 56",
		"end_address":   "53889 Frank Lake Apt. 49",
		"start_time":    "2018-12-22T09:04:05.000Z",
		"end_time":      "2018-12-22T22:04:05.000Z",
		"revenue":       "27.00",
	},
	{
		"id":            "ad0e5604-1893-4800-8000-000000000152",
		"city":          "amsterdam",
		"vehicle_city":  "amsterdam",
		"rider_id":      "ae147ae1-47ae-4800-8000-000000000022",
		"vehicle_id":    "aaaaaaaa-aaaa-4800-8000-00000000000a",
		"start_address": "10806 Kevin Spur",
		"end_address":   "15744 Valerie Squares",
		"start_time":    "2018-12-08T09:04:05.000Z",
		"end_time":      "2018-12-09T04:04:05.000Z",
		"revenue":       "20.00",
	},
	{
		"id":            "ae147ae1-47ae-4800-8000-000000000154",
		"city":          "amsterdam",
		"vehicle_city":  "amsterdam",
		"rider_id":      "bd70a3d7-0a3d-4000-8000-000000000025",
		"vehicle_id":    "aaaaaaaa-aaaa-4800-8000-00000000000a",
		"start_address": "63503 Lisa Summit Suite 28",
		"end_address":   "26800 Brown Station",
		"start_time":    "2018-12-25T09:04:05.000Z",
		"end_time":      "2018-12-27T04:04:05.000Z",
		"revenue":       "0.00",
	},
}
