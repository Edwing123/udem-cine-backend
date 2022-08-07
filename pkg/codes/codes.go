package codes

func NewCode(description string) Code {
	return Code(description)
}

type Code string

func (c Code) Error() string {
	return string(c)
}

var (
	AuthFailed   = NewCode("auth_failed")
	NotLoggedIn  = NewCode("not_logged_in")
	AccessDenied = NewCode("access_denied")
	AdminOnly    = NewCode("admin_only")

	ServerFailed = NewCode("database_server_error")
	Internal     = NewCode("internal_server_error")
	Client       = NewCode("client_error")

	NoRecords                    = NewCode("no_records")
	UserNameExists               = NewCode("username_exists")
	MovieTitleExists             = NewCode("movie_title_exists")
	RoomExists                   = NewCode("room_exists")
	ScheduleExists               = NewCode("schedule_exists")
	FunctionRoomScheduleConflict = NewCode("function_room_schedule_conflict")
	FunctionDependsOnMovie       = NewCode("function_depends_on_movie")
	FunctionDependsOnRoom        = NewCode("function_depends_on_room")
	FunctionDependsOnSchedule    = NewCode("function_depends_on_schedule")
)
