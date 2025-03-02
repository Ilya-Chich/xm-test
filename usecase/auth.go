package usecase

type AuthUC struct {
}

func NewAuth() *AuthUC {
	return &AuthUC{}
}

func (c *AuthUC) Login() error {
	return nil
}
