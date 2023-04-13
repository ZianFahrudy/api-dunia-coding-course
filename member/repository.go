package member

type Repository interface {
	Save(member Member) (Member, error)
}
