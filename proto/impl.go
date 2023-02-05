package proto

type IService interface {
	CreateToken(CreateTokenRequest) (CreateTokenResponse, error)
}
