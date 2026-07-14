//T的方法集合不包括 *T
	// //var r11 Pointerer =T{}
	// te := T{}
	// te.PointerMethod()
	// //res11:=r11.(T)
	// //fmt.Print(res11)

	// //*T可以转换为valuer类型，说明方法集合包括 T
	// //*T可以转换为 Pointerer类型，说明方法集包括 *T

	// //*T的方法集合里面有*T
	// var r2 Pointerer = &T{}
	// res2 := r2.(*T)
	// fmt.Println(res2)

	// //*T的方法集里面有 T
	// var r21 Valuer = &T{}
	// res21 := r21.(*T)
	// fmt.Println(res21)

	// //Qp的方法集合 内嵌*T
	// //Qp的方法集包括 T
	// var r31 Valuer = Qp{}
	// res31 := r31.(Qp)
	// fmt.Println(res31)
	// //Qp的方法集合包括 *T
	// var r32 Pointerer = Qp{}
	// res32 := r32.(Qp)
	// fmt.Println(res32)
	// //*Qp的方法集合包括 T
	// var r33 Valuer = &Qp{}
	// res33 := r33.(*Qp)
	// fmt.Println(res33)
	// //*Qp的方法集合包括 T
	// var r34 Pointerer = &Qp{}
	// res34 := r34.(*Qp)
	// fmt.Println(res34)

	// //Qv的方法集包括T
	// var r41 Valuer = Qv{}
	// res41 := r41.(Qv)
	// fmt.Println(res41)

	// //Qv的方法集不包括*T
	// //var r42 Pointerer =Qv{}

	// //*QV的方法集包括T
	// var r43 Valuer = &Qv{}
	// res43 := r43.(*Qv)
	// fmt.Println(res43)

	// //*QV的方法集包括*T
	// var r44 Pointerer = &Qv{}
	// res44 := r44.(*Qv)
	// fmt.Println(res44)
