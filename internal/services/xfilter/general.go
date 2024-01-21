package xfilter

type NoteFilterService struct {
	UserId string
	// 任意: なかったらDBに取りに行く（既に上位の関数で読み込まれてる場合に渡せばその情報を再利用してクエリを減らせる）
	CachedFollowees *[]string
	CachedMutees    *[]string
	CachedBlockers  *[]string
}
