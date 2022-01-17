package repos

import (
	"fmt"

	"github.com/User0608/zeus_project_api/dberrs"
	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	sign string
	conn *gorm.DB
}

func NewUsuarioRepository(g *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		conn: g,
		sign: "repos.UsuarioRepository",
	}
}
func (r *UsuarioRepository) Login(cu models.LogginRequest) (*models.Usuario, error) {
	usuario := &models.Usuario{}
	res := r.conn.Raw("select * from fn_login(?,?)", cu.Username, cu.Password).Scan(usuario)
	if err := res.Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.Login: %w", err))
	} else {
		if res.RowsAffected == 0 {
			return nil, errores.NewForbiddenf(fmt.Errorf("repos.Login: usuario:`%s` intento de login", cu.Username), errores.ErrUserOrPasswordInvalid)
		}
	}
	return usuario, nil
}

func (r *UsuarioRepository) Create(cu models.PostUsuario) (*models.Usuario, error) {
	var usuario = &models.Usuario{}
	ps := "select * from fn_create_user(?,?);"
	res := r.conn.Raw(ps, cu.Username, cu.Password).Scan(usuario)
	if res.Error != nil {
		return nil, dberrs.DBBind(res.Error, "repos.Usuario.Create")
	}
	return usuario, nil
}
func (r *UsuarioRepository) Listar() ([]models.Usuario, error) {
	usuarios := []models.Usuario{}
	if err := r.conn.Omit("Password").Order("created_at desc").Find(&usuarios).Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.Usuario.Listar: %w", err))
	}
	return usuarios, nil
}
func (r *UsuarioRepository) FreeUsers() ([]models.Usuario, error) {
	usuarios := []models.Usuario{}
	if err := r.conn.Omit("Password").Find(&usuarios, "owner_entity=?", "").Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.Usuario.FreeUsers: %w", err))
	}
	r.conn.Clauses()
	return usuarios, nil
}
func (r *UsuarioRepository) Delete(username string) error {
	query := "delete from usuario where username=? and owner_entity = ''"
	res := r.conn.Exec(query, username)
	if err := res.Error; err != nil {
		return errores.NewInternalDBf(fmt.Errorf("repos.Usuario.Delete: %w", err))
	}
	if res.RowsAffected == 0 {
		const message = "Verifique que la cuenta de usuario `%s` no este asociado con ninguna persona"
		return errores.NewBadRequestf(nil, message, username)
	}
	return nil
}
func (r *UsuarioRepository) Update() (*models.Usuario, error) {
	return nil, nil
}
func (r *UsuarioRepository) SetOwner(username, owner string) (*models.Usuario, error) {
	return nil, nil

}
