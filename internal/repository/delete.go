package repository

import (
	"context"
)

func (d *db) DeleteModel(ctx context.Context, uuid string) error {
	const query = `
		delete from orders
		where order_uid = $1;
	`
	_, err := d.pool.Exec(ctx, query, uuid)
	return err
}

func (d *db) DeleteData(ctx context.Context, id int) error {
	const query = `
		delete from invalid_data
		where id = $1;
	`
	_, err := d.pool.Exec(ctx, query, id)
	return err
}
