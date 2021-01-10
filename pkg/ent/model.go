package ent

import "apps/investimento/pkg/model"

func (n *Nota) ToModel() (*model.Nota, error) {
	nota := &model.Nota{
		ID:        n.ID,
		Date:      n.Date,
		ReceiptID: n.ReceiptID,
		Broker:    n.Broker,
	}

	ativos, err := n.Edges.MovimentacoesOrErr()
	if err != nil {
		return nil, err
	}

	for _, ativo := range ativos {

		a := model.Ativo{
			ID:          ativo.ID,
			Mercado:     ativo.Mercado,
			CV:          ativo.CV,
			TipoMercado: ativo.TipoMercado,
			Titulo:      ativo.Titulo,
			Qtde:        ativo.Qtde,
			Valor:       ativo.Valor,
			DC:          ativo.DC,
		}

		nota.Ativos = append(nota.Ativos, &a)
	}

	cblc, err := n.Edges.CblcsOrErr()
	if err != nil {
		return nil, err
	}

	c := model.Cblc{
		ID:             cblc.ID,
		TaxaLiquidacao: cblc.TaxaLiquidacao,
		TaxaRegistro:   cblc.TaxaRegistro,
	}

	nota.Cblc = &c

	bolsa, err := n.Edges.BolsasOrErr()
	if err != nil {
		return nil, err
	}

	b := model.Bolsa{
		ID:              bolsa.ID,
		Emolumentos:     bolsa.Emolumentos,
		TaxaAna:         bolsa.TaxaAna,
		TaxaTermoOpcoes: bolsa.TaxaTermoOpcoes,
	}

	nota.Bolsa = &b

	despesa, err := n.Edges.DespesasOrErr()
	if err != nil {
		return nil, err
	}

	d := model.Despesa{
		ID:         despesa.ID,
		Corretagem: despesa.Corretagem,
		Iss:        despesa.Iss,
		Irrf:       despesa.Irrf,
		Outros:     despesa.Outros,
	}

	nota.Despesa = &d

	return nota, nil
}
