package rate_model

type RateDB struct {
	ID      string  `db:"id"`
	Date    string  `db:"date"`
	Rate    float64 `db:"rate"`
	ValCode string  `db:"val_code"`
	Code    int     `db:"code"`
}

func (r *RateDB) GetID() string      { return r.ID }
func (r *RateDB) GetDate() string    { return r.Date }
func (r *RateDB) GetRate() float64   { return r.Rate }
func (r *RateDB) GetValCode() string { return r.ValCode }
func (r *RateDB) GetCode() int       { return r.Code }
