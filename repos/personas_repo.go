package repos

import (
	"fmt"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"gorm.io/gorm"
)

type PersonaRepository struct {
	sign string
	conn *gorm.DB
}

func NewPersonaRepository(g *gorm.DB) *PersonaRepository {
	return &PersonaRepository{
		conn: g,
		sign: "repos.PersonaRepository",
	}
}

func (r *PersonaRepository) FindInfoBase() (*models.InformacionBaseEntity, error) {
	estados := []models.EstadoCivil{}
	niveles := []models.NivelEstudio{}
	if err := r.conn.Find(&estados).Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.FindInfoBase: %w", err))
	}
	if err := r.conn.Find(&niveles).Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.FindInfoBase: %w", err))
	}
	return &models.InformacionBaseEntity{EstadosCiviles: estados, NivelesEstudio: niveles}, nil
}
func (r *PersonaRepository) createOrUpdateEntity(tx *gorm.DB, entity *models.Entity) error {
	row := fmt.Sprintf("select count(*) from entity where dni = '%s'", entity.Dni)
	num := 0
	if err := tx.Raw(row).Row().Scan(&num); err != nil {
		return err
	}
	if num == 0 {
		return tx.Create(entity).Error
	}
	return tx.Where("dni = ?", entity.Dni).Updates(entity).Error
}
func (r *PersonaRepository) createOrUpdatePrimerJefe(tx *gorm.DB, pjefe *models.PrimerJefe) error {
	row := fmt.Sprintf("select count(*) from primer_jefe where dni = '%s'", pjefe.Dni)
	num := 0
	if err := tx.Raw(row).Row().Scan(&num); err != nil {
		return err
	}
	if err := tx.Exec("update primer_jefe set state= false where state=true").Error; err != nil {
		return err
	}
	pjefe.State = true
	if num == 0 {
		return tx.Create(pjefe).Error
	}
	return tx.Where("dni = ?", pjefe.Dni).Updates(pjefe).Error
}
func (r *PersonaRepository) RegistrarPrimerJefe(j *models.PrimerJefeEntity) (*models.PrimerJefeEntity, error) {
	return j, r.conn.Transaction(func(tx *gorm.DB) error {
		if err := r.createOrUpdateEntity(tx, &j.Entity); err != nil {
			return errores.NewInternalDBf(fmt.Errorf("repos.RegistrarPrimerJefe: %w", err))
		}
		if err := r.createOrUpdatePrimerJefe(tx, j.GetPrimerJefeObjetct()); err != nil {
			return errores.NewInternalDBf(fmt.Errorf("repos.RegistrarPrimerJefe: %w", err))
		}
		return nil
	})
}
func (r *PersonaRepository) LoadPrimerJefe() (*models.PrimerJefeEntity, error) {
	raw := "select e.* from entity e inner join primer_jefe pj on e.dni = pj.dni where pj.state = true"
	jefe := &models.PrimerJefeEntity{}
	if err := r.conn.Raw(raw).Scan(jefe).Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.LoadPrimerJefe: %w", err))
	}
	return jefe, nil
}

///////////////====================////// Segundo jefe
func (r *PersonaRepository) createOrUpdateSegundoJefe(tx *gorm.DB, pjefe *models.SegundoJefe) error {
	row := fmt.Sprintf("select count(*) from segundo_jefe where dni = '%s'", pjefe.Dni)
	num := 0
	if err := tx.Raw(row).Row().Scan(&num); err != nil {
		return err
	}
	if err := tx.Exec("update segundo_jefe set state=false where state=true").Error; err != nil {
		return err
	}
	pjefe.State = true
	if num == 0 {
		return tx.Create(pjefe).Error
	}
	return tx.Where("dni = ?", pjefe.Dni).Updates(pjefe).Error
}

func (r *PersonaRepository) RegistrarSegundoJefe(j *models.SegundoJefeEntity) (*models.SegundoJefeEntity, error) {
	return j, r.conn.Transaction(func(tx *gorm.DB) error {
		if err := r.createOrUpdateEntity(tx, &j.Entity); err != nil {
			return errores.NewInternalDBf(fmt.Errorf("repos.RegistrarSegundoJefe: %w", err))
		}
		if err := r.createOrUpdateSegundoJefe(tx, j.GetSegundoJefeObjetct()); err != nil {
			return errores.NewInternalDBf(fmt.Errorf("repos.RegistrarSegundoJefe: %w", err))
		}
		return nil
	})
}
func (r *PersonaRepository) LoadSegundoJefe() (*models.SegundoJefeEntity, error) {
	raw := "select e.* from entity e inner join segundo_jefe pj on e.dni = pj.dni where pj.state = true"
	jefe := &models.SegundoJefeEntity{}
	if err := r.conn.Raw(raw).Scan(jefe).Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.LoadSegundoJefe: %w", err))
	}
	return jefe, nil
}

//////
func (r *PersonaRepository) ResumenPrimerSegunJefe() (*models.PrimerJefeEntity, *models.SegundoJefeEntity, error) {
	primerJefe := &models.PrimerJefeEntity{}
	segundoJefe := &models.SegundoJefeEntity{}
	if err := r.conn.Limit(1).Find(primerJefe, "state=?", true).Error; err != nil {
		return nil, nil, errores.NewInternalDBf(fmt.Errorf("repos.ResumenPrimerSegunJefe: %w", err))
	}
	if err := r.conn.Table("entity").Limit(1).Find(primerJefe, "dni=?", primerJefe.Dni).Error; err != nil {
		return nil, nil, errores.NewInternalDBf(fmt.Errorf("repos.ResumenPrimerSegunJefe: %w", err))
	}

	if err := r.conn.Limit(1).Find(segundoJefe, "state=?", true).Error; err != nil {
		return nil, nil, errores.NewInternalDBf(fmt.Errorf("repos.ResumenPrimerSegunJefe: %w", err))
	}
	if err := r.conn.Table("entity").Limit(1).Find(segundoJefe, "dni=?", segundoJefe.Dni).Error; err != nil {
		return nil, nil, errores.NewInternalDBf(fmt.Errorf("repos.ResumenPrimerSegunJefe: %w", err))
	}
	return primerJefe, segundoJefe, nil
}

////////////===================///////////////////Jefe Instruccion
func (r *PersonaRepository) createOrUpdateJefeInstruccion(tx *gorm.DB, pjefe *models.JefeInstruccion) error {
	row := fmt.Sprintf("select count(*) from jefe_instruccion where dni = '%s'", pjefe.Dni)
	num := 0
	if err := tx.Raw(row).Row().Scan(&num); err != nil {
		return err
	}
	if err := tx.Exec("update jefe_instruccion set state=false where state=true").Error; err != nil {
		return err
	}
	pjefe.State = true
	if num == 0 {
		return tx.Create(pjefe).Error
	}
	return tx.Where("dni = ?", pjefe.Dni).Updates(pjefe).Error
}
func (r *PersonaRepository) RegistrarJefeInstruccion(j *models.JefeInstruccionEntity) (*models.JefeInstruccionEntity, error) {
	return j, r.conn.Transaction(func(tx *gorm.DB) error {
		if err := r.createOrUpdateEntity(tx, &j.Entity); err != nil {
			return errores.NewInternalDBf(fmt.Errorf("repos.RegistrarJefeInstruccion: %w", err))
		}
		if err := r.createOrUpdateJefeInstruccion(tx, j.GetJefeInstructor()); err != nil {
			return errores.NewInternalDBf(fmt.Errorf("repos.RegistrarJefeInstruccion: %w", err))
		}
		return nil
	})
}
func (r *PersonaRepository) LoadJefeInstruccion() (*models.JefeInstruccionEntity, error) {
	jefe := &models.JefeInstruccionEntity{}
	res := r.conn.Table("jefe_instruccion").Limit(1).Find(jefe, "state=?", true)
	if err := res.Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.LoadJefeInstruccion: %w", err))
	}
	if res.RowsAffected == 0 {
		return nil, errores.NewNotFoundf(nil, errores.ErrRecordNotFaund)
	}
	if err := r.conn.Table("entity").Limit(1).Find(jefe, "dni=?", jefe.Dni).Error; err != nil {
		return nil, errores.NewInternalDBf(fmt.Errorf("repos.LoadJefeInstruccion: %w", err))
	}
	return jefe, nil
}

/////////////////////////// Instructores CRUD
func (r *PersonaRepository) RegistrarInstructor(i *models.InstructorEntity) error {
	return r.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("entity").Create(&i.Entity).Error; err != nil {
			return errores.NewInternalDBf(err)
		}
		instructor := i.GetInstructor()
		if err := tx.Table("instructor").Create(instructor).Error; err != nil {
			return errores.NewInternalDBf(err)
		}
		i.SetInstructor(instructor)
		return nil
	})
}
func (r *PersonaRepository) UpdateInstructor(i *models.InstructorEntity) error {
	return r.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("entity").Where("dni=?", i.Dni).Updates(&i.Entity).Error; err != nil {
			return errores.NewInternalDBf(err)
		}
		instructor := i.GetInstructor()
		if err := tx.Table("instructor").Where("dni=?", i.Dni).Updates(instructor).Error; err != nil {
			return errores.NewInternalDBf(err)
		}
		i.SetInstructor(instructor)
		return nil
	})
}

func (r *PersonaRepository) FindInstructorByDNI(dni string) (*models.InstructorEntity, error) {
	instructor := &models.InstructorEntity{}
	result := r.conn.Table("vw_instructor").Limit(1).Find(instructor, "dni=?", dni)
	if err := result.Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	if result.RowsAffected == 0 {
		return nil, errores.NewNotFoundf(nil, "no se encontro registro con dni `%s`", dni)
	}
	return instructor, nil
}
func (r *PersonaRepository) FindAllInstructors() ([]*models.InstructorEntity, error) {
	instructores := []*models.InstructorEntity{}
	if err := r.conn.Table("vw_instructor").Where("state=?", true).Find(&instructores).Error; err != nil {
		return nil, errores.NewInternalDBf(err)
	}
	return instructores, nil
}
