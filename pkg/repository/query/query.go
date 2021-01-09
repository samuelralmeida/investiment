package query

// Queries
const (
	InsertNota = `
		INSERT INTO public.nota_corretagems ("date", "receipt_id", "broker")
		VALUES ($1, $2, $3)
		RETURNING id;
	`
)
