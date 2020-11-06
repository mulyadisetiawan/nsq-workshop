package benefit

type publisher interface {
	Publish(topic string, msg []byte) error
}

type Repository struct {
	publisher publisher
}

func NewRepository(publisher publisher) *Repository {
	return &Repository{
		publisher,
	}
}

func (r *Repository) Publish(topic string, msg []byte) error {
	return r.publisher.Publish(topic, msg)
}
