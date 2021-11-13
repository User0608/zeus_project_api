package repos

import (
	"github.com/User0608/zeus_project_api/errs"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	sign string
	gorm *gorm.DB
}

func NewUsuarioRepository(g *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		gorm: g,
		sign: "repos.UsuarioRepository",
	}
}
func (r *UsuarioRepository) Login(cu models.LogginRequest) (*models.Usuario, error) {
	usuario := &models.Usuario{}
	res := r.gorm.Raw("select * from fn_login('?','?')", cu.Username, cu.Password).Scan(usuario)
	if err := res.Error; err != nil {
		return nil, errs.WrapAndMessage(errs.Trc(r.sign, "Login"), res.Error, errs.ErrDatabaseRequest)
	} else {
		if res.RowsAffected > 0 {
			return nil, errs.Create(errs.Trc(r.sign, "Loggin"), errs.ErrUserOrPasswordInvalid)
		}
	}
	return usuario, nil
}

func (r *UsuarioRepository) Create(cu models.PostUsuario) (*models.Usuario, error) {
	var usuario = &models.Usuario{}
	ps := "select * from fn_create_user(?,?);"
	res := r.gorm.Raw(ps, cu.Username, cu.Password).Scan(usuario)
	if res.Error != nil {
		return nil, errs.DBBind(errs.Trc(r.sign, "Create"), res.Error)
	}
	return usuario, nil
}
