package user

func New() *User {
	return new(User)
}

type User struct {
	Username string
	Password string
	APIToken string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Initials string `json:"initials"`
	Timezone struct {
		Kind      string `json:"kind"`
		Offset    string `json:"offset"`
		OlsonName string `json:"olson_name"`
	} `json:"time_zone"`
}

func (u *User) Login(name, pass string) {
	u.Username = name
	u.Password = pass
}

// ****
// API response:
// {
//   "projects": [
//     {
//       "project_id": 1027438,
//       "project_name": "SantaMonicaBoocamp1",
//       "id": 4001730,
//       "kind": "membership_summary",
//       "project_color": "ffc100",
//       "role": "owner"
//     }
//   ],
//   "username": "gobootcamp",
//   "id": 1266284,
//   "time_zone": {
//     "kind": "time_zone",
//     "olson_name": "America/Los_Angeles",
//     "offset": "-08:00"
//   },
//   "has_google_identity": false,
//   "name": "Matt Aimonetti (GoBootcamp)",
//   "kind": "me",
//   "api_token": "3ab7cfb10fcf49ee0f825380627c63c4",
//   "email": "mattaimonetti+gobootcamp@gmail.com",
//   "initials": "MA("
// }