package core

import "github.com/MagicalTux/goro/core/tokenizer"

func compileTernaryOp(v Runnable, c compileCtx) (Runnable, error) {
	// v contains the first part, we already have read the ? too
	i, err := c.NextItem()
	if err != nil {
		return nil, err
	}

	var yes, no Runnable
	l := MakeLoc(i.Loc())

	if i.IsSingle(':') {
		yes = v
	} else if i.IsSingle('?') {
		yes = v
		v, _ = spawnOperator(tokenizer.T_IS_NOT_IDENTICAL, v, &runZVal{nil, l}, l)
	} else {
		yes, err = compileExpr(i, c)
		if err != nil {
			return nil, err
		}

		i, err = c.NextItem()
		if err != nil {
			return nil, err
		}

		if !i.IsSingle(':') {
			return nil, i.Unexpected()
		}
	}

	// check no
	no, err = compileExpr(nil, c)
	if err != nil {
		return nil, err
	}

	return &runnableIf{cond: v, yes: yes, no: no, l: l}, nil
}
